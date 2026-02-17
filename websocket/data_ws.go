package fyersgosdk

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// FyersDataSocket represents the main WebSocket client
type FyersDataSocket struct {
	url                  string
	accessToken          string
	hsmToken             string
	logPath              string
	lite                 bool
	maxRetry             int
	source               string
	channelNum           int
	channels             []int
	runningChannels      map[int]bool
	dataType             string
	OnMessage            func(map[string]interface{})
	OnError              func(map[string]interface{})
	OnOpen               func()
	OnClose              func(map[string]interface{})
	updateTick           bool
	ackCount             int
	wsRun                *websocket.Conn
	writeToFile          bool
	backgroundFlag       bool
	updateCount          int
	liteResp             map[string]interface{}
	channelSymbol        []string
	symbolDict           map[string]string
	scripsCount          map[string]int
	scripsPerChannel     map[int][]string
	restartFlag          bool
	reconnectAttempts    int
	reconnectDelay       int
	maxReconnectAttempts int
	mu                   sync.Mutex
	connected            bool
	messageQueue         chan []byte
	stopChan             chan bool
	stopOnce             sync.Once

	// State management for continuous updates (like Python)
	scripsSym     map[uint16]string                 // topic_id -> topic_name for scrips
	indexSym      map[uint16]string                 // topic_id -> topic_name for index
	dpSym         map[uint16]string                 // topic_id -> topic_name for depth
	resp          map[string]map[string]interface{} // topic_name -> response data
	fieldMappings map[string][]string               // Add this line

	// lastLtpScrips/lastLtpIndex: dedup lite ticks (emit only when LTP changes, match Python). Key = topic name, value = last raw int32 LTP.
	lastLtpScrips map[string]int32
	lastLtpIndex  map[string]int32

	// Subscription cache: on reconnect we re-subscribe without calling the symbol-token API (avoids "unexpected EOF" / subscription failed when API is flaky).
	cachedSubscribeSymbols  []string
	cachedSubscribeDataType string
	cachedSubscribeDataDict map[string]string
}

// readDeadlineDuration is the max time to wait for a message. Prevents blocking forever when network is lost (no internet).
// Fixed duration (e.g. 30s) so timeout triggers sooner regardless of maxReconnectAttempts.
const readDeadlineDuration = 30 * time.Second

// NewFyersDataSocket creates a new FyersDataSocket instance.
// reconnectRetry sets max reconnect attempts; default 5. Same rule as Python: cap at 50, use min(50, reconnectRetry).
func NewFyersDataSocket(
	accessToken string,
	logPath string,
	liteMode bool,
	writeToFile bool,
	reconnect bool,
	reconnectRetry int,
	onConnect func(),
	onClose func(map[string]interface{}),
	onError func(map[string]interface{}),
	onMessage func(map[string]interface{}),
) *FyersDataSocket {
	// Load field mappings from map.json
	fieldMappings, err := loadFieldMappingsOnce()
	if err != nil {
		fmt.Printf("Failed to load field mappings: %v\n", err)
		return nil
	}
	if reconnectRetry <= 0 {
		reconnectRetry = 5
	}
	maxReconnectAttempts := 50
	if reconnectRetry < maxReconnectAttempts {
		maxReconnectAttempts = reconnectRetry
	}
	// fmt.Printf("[FyersDataSocket RECONNECT] init: reconnect=%v, maxReconnectAttempts=%d\n", reconnect, maxReconnectAttempts)
	return &FyersDataSocket{
		url:                  "wss://socket.fyers.in/hsm/v1-5/prod",
		accessToken:          accessToken,
		hsmToken:             "",
		logPath:              logPath,
		lite:                 liteMode,
		maxRetry:             reconnectRetry,
		source:               "GoSDK-1.0.0",
		channelNum:           11,
		channels:             []int{},
		runningChannels:      make(map[int]bool),
		dataType:             "",
		OnMessage:            onMessage,
		OnError:              onError,
		OnOpen:               onConnect,
		OnClose:              onClose,
		updateTick:           false,
		ackCount:             0,
		wsRun:                nil,
		writeToFile:          writeToFile,
		backgroundFlag:       false,
		updateCount:          0,
		liteResp:             make(map[string]interface{}),
		channelSymbol:        []string{},
		symbolDict:           make(map[string]string),
		scripsCount:          make(map[string]int),
		scripsPerChannel:     make(map[int][]string),
		restartFlag:          reconnect,
		reconnectAttempts:    0,
		reconnectDelay:       0,
		maxReconnectAttempts: maxReconnectAttempts,
		connected:            false,
		messageQueue:         make(chan []byte, 1000),
		stopChan:             make(chan bool),
		scripsSym:            make(map[uint16]string),
		indexSym:             make(map[uint16]string),
		dpSym:                make(map[uint16]string),
		resp:                 make(map[string]map[string]interface{}),
		fieldMappings:        fieldMappings,
		lastLtpScrips:        make(map[string]int32),
		lastLtpIndex:         make(map[string]int32),
	}
}

// Helper function to load field mappings once
func loadFieldMappingsOnce() (map[string][]string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get current file path")
	}
	dir := filepath.Dir(filename)
	mapPath := filepath.Join(dir, "map.json")
	data, err := os.ReadFile(mapPath)
	if err != nil {
		return nil, err
	}
	var mapData map[string]interface{}
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		return nil, err
	}
	fieldMappings := make(map[string][]string)
	if dataVal, ok := mapData["data_val"].([]interface{}); ok {
		fieldMappings["data_val"] = make([]string, len(dataVal))
		for i, v := range dataVal {
			fieldMappings["data_val"][i] = v.(string)
		}
	}
	if indexVal, ok := mapData["index_val"].([]interface{}); ok {
		fieldMappings["index_val"] = make([]string, len(indexVal))
		for i, v := range indexVal {
			fieldMappings["index_val"][i] = v.(string)
		}
	}
	if depthVal, ok := mapData["depthvalue"].([]interface{}); ok {
		fieldMappings["depthvalue"] = make([]string, len(depthVal))
		for i, v := range depthVal {
			fieldMappings["depthvalue"][i] = v.(string)
		}
	}
	return fieldMappings, nil
}

var (
	marshalMappings     map[string][]string
	marshalMappingsOnce sync.Once
)

func getMappingsForMarshal() map[string][]string {
	marshalMappingsOnce.Do(func() {
		marshalMappings, _ = loadFieldMappingsOnce()
	})
	return marshalMappings
}

// buildOrderedKeys returns keys in data_val / index_val / depth order, then ch/chp, then any remainder.
func buildOrderedKeys(resp map[string]interface{}, mappings map[string][]string) []string {
	if mappings == nil {
		// fallback: return keys in map iteration order (undefined)
		keys := make([]string, 0, len(resp))
		for k := range resp {
			keys = append(keys, k)
		}
		return keys
	}
	seen := make(map[string]bool)
	ordered := make([]string, 0, len(resp))

	dataType, _ := resp["type"].(string)
	var fieldOrder []string
	switch dataType {
	case "sf", "scrips":
		fieldOrder = mappings["data_val"]
	case "if":
		fieldOrder = mappings["index_val"]
	case "dp":
		fieldOrder = mappings["depthvalue"]
	}

	for _, key := range fieldOrder {
		if _, exists := resp[key]; exists {
			ordered = append(ordered, key)
			seen[key] = true
		}
	}
	for _, key := range []string{"ch", "chp"} {
		if _, exists := resp[key]; exists && !seen[key] {
			ordered = append(ordered, key)
			seen[key] = true
		}
	}
	// Remainder: use fixed order so lite mode (ltp, symbol, type) is deterministic: ltp already from fieldOrder, then symbol, type.
	for _, key := range []string{"symbol", "type"} {
		if _, exists := resp[key]; exists && !seen[key] {
			ordered = append(ordered, key)
			seen[key] = true
		}
	}
	// Any other keys in sorted order for determinism (Go map iteration is random).
	var remainder []string
	for key := range resp {
		if !seen[key] {
			remainder = append(remainder, key)
		}
	}
	sort.Strings(remainder)
	ordered = append(ordered, remainder...)
	return ordered
}

// MarshalDataResponseInOrder marshals the data response map to JSON with keys in the same order
// as the Fyers data_val / index_val / depth feed: ltp, vol_traded_today, ..., type, symbol, ch, chp.
// Use this when you need output order to match the Python SDK or a fixed field order.
func MarshalDataResponseInOrder(resp map[string]interface{}) ([]byte, error) {
	if resp == nil {
		return []byte("null"), nil
	}
	mappings := getMappingsForMarshal()
	orderedKeys := buildOrderedKeys(resp, mappings)
	var buf strings.Builder
	buf.Grow(512)
	buf.WriteByte('{')
	for i, key := range orderedKeys {
		if i > 0 {
			buf.WriteByte(',')
		}
		keyBytes, _ := json.Marshal(key)
		buf.Write(keyBytes)
		buf.WriteByte(':')
		valBytes, err := json.Marshal(resp[key])
		if err != nil {
			return nil, err
		}
		buf.Write(valBytes)
	}
	buf.WriteByte('}')
	return []byte(buf.String()), nil
}

// AccessTokenToHSMToken converts access token to HSM token by decoding JWT
func (f *FyersDataSocket) AccessTokenToHSMToken() bool {
	// Check if access token is in the correct format
	if !strings.Contains(f.accessToken, ":") {
		fmt.Printf("Access token format error: expected format 'APPID:TOKEN', got: %s\n", f.accessToken)
		if f.OnError != nil {
			f.OnError(map[string]interface{}{
				"type":    AUTH_TYPE,
				"code":    AUTH_ERROR_CODE,
				"message": INVALID_TOKEN,
				"s":       ERROR,
			})
		}
		return false
	}

	// Extract the token part (after the colon)
	parts := strings.Split(f.accessToken, ":")
	if len(parts) != 2 {
		fmt.Printf("Access token format error: expected exactly one colon, got: %s\n", f.accessToken)
		if f.OnError != nil {
			f.OnError(map[string]interface{}{
				"type":    AUTH_TYPE,
				"code":    AUTH_ERROR_CODE,
				"message": INVALID_TOKEN,
				"s":       ERROR,
			})
		}
		return false
	}

	// Use only the token part for JWT decoding
	tokenPart := parts[1]

	// Split the JWT token into parts
	tokenParts := strings.Split(tokenPart, ".")
	if len(tokenParts) != 3 {
		fmt.Printf("JWT token format error: expected 3 parts, got %d\n", len(tokenParts))
		if f.OnError != nil {
			f.OnError(map[string]interface{}{
				"type":    AUTH_TYPE,
				"code":    AUTH_ERROR_CODE,
				"message": INVALID_TOKEN,
				"s":       ERROR,
			})
		}
		return false
	}

	// Decode the payload (second part)
	payloadB64 := tokenParts[1]

	// Add padding if needed
	if len(payloadB64)%4 != 0 {
		payloadB64 += strings.Repeat("=", 4-len(payloadB64)%4)
	}

	// Decode base64
	payloadBytes, err := base64.URLEncoding.DecodeString(payloadB64)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		if f.OnError != nil {
			f.OnError(map[string]interface{}{
				"type":    AUTH_TYPE,
				"code":    AUTH_ERROR_CODE,
				"message": INVALID_TOKEN,
				"s":       ERROR,
			})
		}
		return false
	}

	// Parse JSON payload
	var payload map[string]interface{}
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
		if f.OnError != nil {
			f.OnError(map[string]interface{}{
				"type":    AUTH_TYPE,
				"code":    AUTH_ERROR_CODE,
				"message": INVALID_TOKEN,
				"s":       ERROR,
			})
		}
		return false
	}

	// Check if token is expired
	if exp, exists := payload["exp"]; exists {
		expTime := int64(exp.(float64))
		currentTime := time.Now().Unix()
		if expTime-currentTime < 0 {
			if f.OnError != nil {
				f.OnError(map[string]interface{}{
					"type":    AUTH_TYPE,
					"code":    TOKEN_EXPIRED,
					"message": TOKEN_EXPIRED_MSG,
					"s":       ERROR,
				})
			}
			fmt.Printf("Token expired: exp=%d, current=%d\n", expTime, currentTime)
			return false
		}
	}

	// Extract hsm_key
	if hsmKey, exists := payload["hsm_key"]; exists {
		f.hsmToken = hsmKey.(string)
		// Do not print here: auth success is reported via OnMessage when server responds after Connect()
		return true
	}

	fmt.Printf("hsm_key not found in token payload\n")
	if f.OnError != nil {
		f.OnError(map[string]interface{}{
			"type":    AUTH_TYPE,
			"code":    AUTH_ERROR_CODE,
			"message": INVALID_TOKEN,
			"s":       ERROR,
		})
	}
	return false
}

// Connect establishes WebSocket connection
func (f *FyersDataSocket) Connect() error {
	if !f.AccessTokenToHSMToken() {
		return fmt.Errorf("failed to get HSM token")
	}

	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(f.url, nil)
	if err != nil {
		return err
	}

	f.mu.Lock()
	_ = f.reconnectAttempts > 0 // wasReconnect (debug commented)
	f.wsRun = conn
	f.connected = true
	f.reconnectAttempts = 0
	f.reconnectDelay = 0
	f.mu.Unlock()

	// if wasReconnect {
	// 	fmt.Printf("[FyersDataSocket RECONNECT] on_open: reconnection successful, reset attempt counter and delay\n")
	// } else {
	// 	fmt.Printf("[FyersDataSocket RECONNECT] on_open: first connection established\n")
	// }

	// Start message processing goroutine
	go f.processMessageQueue()

	// Send authentication message
	authMsg := f.createAuthMessage()
	err = f.wsRun.WriteMessage(websocket.BinaryMessage, authMsg)
	if err != nil {
		return err
	}

	// Send mode message (lite or full)
	var modeMsg []byte
	if f.lite {
		modeMsg = f.createLiteModeMessage()
	} else {
		modeMsg = f.createFullModeMessage()
	}
	err = f.wsRun.WriteMessage(websocket.BinaryMessage, modeMsg)
	if err != nil {
		return err
	}

	// Start reading messages
	go f.readMessages()

	// Call onOpen callback
	if f.OnOpen != nil {
		f.OnOpen()
	}

	return nil
}

// createAuthMessage creates the authentication message in binary format
func (f *FyersDataSocket) createAuthMessage() []byte {
	// Calculate buffer size: 18 + len(hsm_token) + len(source)
	bufferSize := 18 + len(f.hsmToken) + len(f.source)

	// Create byte buffer
	buffer := make([]byte, 0, bufferSize)

	// Pack data length (buffer_size - 2)
	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, uint16(bufferSize-2))
	buffer = append(buffer, lengthBytes...)

	// Set ReqType (1 for auth)
	buffer = append(buffer, 1)

	// Set FieldCount (4 fields)
	buffer = append(buffer, 4)

	// Field-1: AuthToken
	buffer = append(buffer, 1) // field ID
	tokenLenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(tokenLenBytes, uint16(len(f.hsmToken)))
	buffer = append(buffer, tokenLenBytes...)
	buffer = append(buffer, []byte(f.hsmToken)...)

	// Field-2: Mode (1 byte)
	buffer = append(buffer, 2)    // field ID
	buffer = append(buffer, 0, 1) // length = 1
	buffer = append(buffer, '1')  // mode = '1'

	// Field-3: Version (1 byte)
	buffer = append(buffer, 3)    // field ID
	buffer = append(buffer, 0, 1) // length = 1
	buffer = append(buffer, 1)    // version = 1

	// Field-4: Source
	buffer = append(buffer, 4) // field ID
	sourceLenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(sourceLenBytes, uint16(len(f.source)))
	buffer = append(buffer, sourceLenBytes...)
	buffer = append(buffer, []byte(f.source)...)

	return buffer
}

// readMessages reads messages from WebSocket. On connection error, runs reconnect logic (same as Python __on_close) when restartFlag is true.
// A read deadline is set so that when the network is lost we get a timeout error and can reconnect instead of blocking forever.
func (f *FyersDataSocket) readMessages() {
	for {
		select {
		case <-f.stopChan:
			return
		default:
		}

		f.mu.Lock()
		conn := f.wsRun
		f.mu.Unlock()
		if conn == nil {
			return
		}

		conn.SetReadDeadline(time.Now().Add(readDeadlineDuration))
		_, message, err := conn.ReadMessage()
		if err != nil {
			f.mu.Lock()
			f.wsRun = nil
			f.connected = false
			restartFlag := f.restartFlag
			maxAttempts := f.maxReconnectAttempts
			_ = f.reconnectAttempts // attempts (debug commented)
			f.mu.Unlock()

			// if err != nil && (strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "deadline")) {
			// 	fmt.Printf("[FyersDataSocket RECONNECT] on_close: read deadline/timeout (no data) -> triggering reconnect: %v\n", err)
			// }
			// fmt.Printf("[FyersDataSocket RECONNECT] on_close: error=%v, restartFlag=%v, attempts=%d/%d\n", err, restartFlag, attempts, maxAttempts)

			if f.OnError != nil {
				errMsg := err.Error()
				if strings.Contains(errMsg, "timeout") || strings.Contains(errMsg, "deadline") {
					errMsg = "Connection timed out"
				}
				f.OnError(map[string]interface{}{"error": errMsg})
			}

			if !restartFlag {
				// fmt.Printf("[FyersDataSocket RECONNECT] reconnection disabled (restartFlag=false); invoking OnClose\n")
				if f.OnClose != nil {
					f.OnClose(map[string]interface{}{
						"code":    SUCCESS_CODE,
						"message": CONNECTION_CLOSED,
						"s":       SUCCESS,
					})
				}
				return
			}

			// fmt.Printf("[FyersDataSocket RECONNECT] entering runReconnectLoop (maxAttempts=%d)\n", maxAttempts)
			f.runReconnectLoop(maxAttempts)
			return
		}

		f.handleMessage(message)
	}
}

// runReconnectLoop runs the same logic as Python __on_close reconnect path: backoff, clear state, Connect(); loops until success or max attempts. Does not close stopChan.
func (f *FyersDataSocket) runReconnectLoop(maxAttempts int) {
	for {
		select {
		case <-f.stopChan:
			// fmt.Printf("[FyersDataSocket RECONNECT] runReconnectLoop: stopChan received, exiting\n")
			return
		default:
		}

		f.mu.Lock()
		attempts := f.reconnectAttempts
		f.mu.Unlock()

		if attempts >= maxAttempts {
			// fmt.Printf("[FyersDataSocket RECONNECT] giving up: max attempts reached (%d), invoking OnClose\n", maxAttempts)
			if f.OnClose != nil {
				f.OnClose(map[string]interface{}{
					"code":    SUCCESS_CODE,
					"message": MAX_RECONNECT_ATTEMPTS_REACHED,
					"s":       ERROR,
				})
			}
			return
		}

		attemptNum := attempts + 1
		_ = attemptNum
		// fmt.Printf("[FyersDataSocket RECONNECT] will retry: attempt %d of %d\n", attemptNum, maxAttempts)
		// fmt.Printf("Attempting reconnect %d of %d...\n", attemptNum, maxAttempts)

		f.mu.Lock()
		delay := f.reconnectDelay
		if attempts%5 == 0 {
			f.reconnectDelay += 5
			delay = f.reconnectDelay
		}
		f.mu.Unlock()

		if delay > 0 {
			// fmt.Printf("[FyersDataSocket RECONNECT] backoff: Reconnection backoff: waiting %ds before reconnect (delay increases by 5s every 5 attempts).\n", delay)
			select {
			case <-f.stopChan:
				// fmt.Printf("[FyersDataSocket RECONNECT] backoff interrupted by stopChan\n")
				return
			case <-time.After(time.Duration(delay) * time.Second):
			}
		}

		f.mu.Lock()
		f.reconnectAttempts++
		f.scripsPerChannel[f.channelNum] = nil
		// Do not clear symbolDict on reconnect (like Python: symbol_token is kept). OnOpen will re-subscribe and we'll repopulate from cache if needed.
		f.mu.Unlock()

		// fmt.Printf("[FyersDataSocket RECONNECT] clearing state and calling Connect() for attempt %d\n", attemptNum)
		err := f.Connect()
		if err == nil {
			// fmt.Printf("[FyersDataSocket RECONNECT] Connect() succeeded, returning from runReconnectLoop\n")
			return
		}
		// Retry silently (like Python): do not call OnError for each failed attempt; only the initial "Connection timed out" was already reported
	}
}

// handleMessage processes incoming messages
func (f *FyersDataSocket) handleMessage(message []byte) {
	// Check if message is too short
	if len(message) < 3 {
		if f.OnError != nil {
			f.OnError(map[string]interface{}{"error": "Message too short"})
		}
		return
	}

	// Parse the response type like Python implementation
	// struct.unpack("!HB", data[:3]) - 2 bytes for length, 1 byte for response type
	_, respType := binary.BigEndian.Uint16(message[:2]), message[2]

	switch respType {
	case 1: // Authentication response
		f.handleAuthResponse(message)
	case 4: // Subscription response
		f.handleSubscribeResponse(message)
	case 5: // Unsubscription response
		f.handleUnsubscribeResponse(message)
	case 6: // Data Feed Response
		f.handleDataFeedResponse(message)
	case 7, 8: // Resume/Pause response
		f.handleResumePauseResponse(message, int(respType))
	case 12: // Full Mode Data Response
		f.handleLiteFullModeResponse(message)
	default:
		// Unknown response type - ignore silently like Python
	}
}

// handleAuthResponse handles authentication response
func (f *FyersDataSocket) handleAuthResponse(data []byte) {
	if len(data) < 4 {
		return
	}

	offset := 4
	if offset+1 >= len(data) {
		return
	}
	offset += 1

	if offset+2 >= len(data) {
		return
	}
	fieldLength := binary.BigEndian.Uint16(data[offset : offset+2])
	offset += 2

	if offset+int(fieldLength) > len(data) {
		return
	}
	stringVal := string(data[offset : offset+int(fieldLength)])
	offset += int(fieldLength)

	if stringVal == "K" {
		if f.OnMessage != nil {
			f.OnMessage(map[string]interface{}{
				"type":    AUTH_TYPE,
				"code":    SUCCESS_CODE,
				"message": AUTH_SUCCESS,
				"s":       SUCCESS,
			})
		}
	} else {
		if f.OnError != nil {
			f.OnError(map[string]interface{}{
				"type":    AUTH_TYPE,
				"code":    AUTH_ERROR_CODE,
				"message": AUTH_FAIL,
				"s":       ERROR,
			})
		}
	}

	// Parse ack count
	if offset+1+2+4 <= len(data) {
		offset += 1
		offset += 2
		f.ackCount = int(binary.BigEndian.Uint32(data[offset : offset+4]))
	}
}

// handleSubscribeResponse handles subscription response
func (f *FyersDataSocket) handleSubscribeResponse(data []byte) {
	if len(data) < 5 {
		return
	}

	offset := 5
	if offset+2 >= len(data) {
		return
	}
	fieldLength := binary.BigEndian.Uint16(data[offset : offset+2])
	offset += 2

	if offset+int(fieldLength) > len(data) {
		return
	}
	stringVal := string(data[offset : offset+int(fieldLength)])

	if stringVal == "K" {
		if f.OnMessage != nil {
			f.OnMessage(map[string]interface{}{
				"type":    SUBS_TYPE,
				"code":    SUCCESS_CODE,
				"message": SUBSCRIBE_SUCCESS,
				"s":       SUCCESS,
			})
		}
	} else {
		if f.OnError != nil {
			f.OnError(map[string]interface{}{
				"type":    SUBS_TYPE,
				"code":    SUBS_ERROR_CODE,
				"message": SUBSCRIBE_FAIL,
				"s":       ERROR,
			})
		}
	}
}

// handleUnsubscribeResponse handles unsubscription response
func (f *FyersDataSocket) handleUnsubscribeResponse(data []byte) {
	if len(data) < 5 {
		return
	}

	offset := 5
	if offset+2 >= len(data) {
		return
	}
	fieldLength := binary.BigEndian.Uint16(data[offset : offset+2])
	offset += 2

	if offset+int(fieldLength) > len(data) {
		return
	}
	stringVal := string(data[offset : offset+int(fieldLength)])

	if stringVal == "K" {
		if f.OnMessage != nil {
			f.OnMessage(map[string]interface{}{
				"type":    UNSUBS_TYPE,
				"code":    SUCCESS_CODE,
				"message": UNSUBSCRIBE_SUCCESS,
				"s":       SUCCESS,
			})
		}
	} else {
		if f.OnError != nil {
			f.OnError(map[string]interface{}{
				"type":    UNSUBS_TYPE,
				"code":    UNSUBS_ERROR_CODE,
				"message": UNSUBSCRIBE_FAIL,
				"s":       ERROR,
			})
		}
	}
}

// handleDataFeedResponse handles data feed response
func (f *FyersDataSocket) handleDataFeedResponse(data []byte) {
	fieldMappings := f.fieldMappings
	if fieldMappings == nil {
		if f.OnError != nil {
			f.OnError(map[string]interface{}{"error": "Field mappings not loaded"})
		}
		return
	}

	// Parse message number and scrip count
	if len(data) < 9 {
		return
	}

	// Handle acknowledgment if needed
	if f.ackCount > 0 {
		f.updateCount++
		messageNum := binary.BigEndian.Uint32(data[3:7])
		if f.updateCount == f.ackCount {
			ackMsg := f.createAcknowledgmentMessage(int(messageNum))
			f.messageQueue <- ackMsg
			f.updateCount = 0
		}
	}

	scripCount := binary.BigEndian.Uint16(data[7:9])
	offset := 9

	for i := 0; i < int(scripCount); i++ {
		if offset >= len(data) {
			break
		}

		dataType := data[offset]
		offset++

		if dataType == 83 { // Snapshot datafeed
			offset = f.handleSnapshotData(data, offset, fieldMappings)
		} else if dataType == 85 { // Full mode datafeed
			offset = f.handleFullModeData(data, offset, fieldMappings)
		} else if dataType == 76 { // Lite mode datafeed
			offset = f.handleLiteModeData(data, offset, fieldMappings)
		}
	}
}

// handleSnapshotData handles snapshot data feed
func (f *FyersDataSocket) handleSnapshotData(data []byte, offset int, fieldMappings map[string][]string) int {
	if offset+2 >= len(data) {
		return offset
	}

	topicID := binary.BigEndian.Uint16(data[offset : offset+2])
	offset += 2

	if offset+1 >= len(data) {
		return offset
	}
	topicNameLen := data[offset]
	offset++

	if offset+int(topicNameLen) > len(data) {
		return offset
	}
	topicName := string(data[offset : offset+int(topicNameLen)])
	offset += int(topicNameLen)

	// Determine data type based on topic name prefix and maintain mappings
	var dataType string
	var fieldNames []string

	if strings.HasPrefix(topicName, "dp") {
		dataType = "depth"
		fieldNames = fieldMappings["depthvalue"]
		f.dpSym[topicID] = topicName
		f.resp[topicName] = make(map[string]interface{})
	} else if strings.HasPrefix(topicName, "if") {
		dataType = "index"
		fieldNames = fieldMappings["index_val"]
		f.indexSym[topicID] = topicName
		f.resp[topicName] = make(map[string]interface{})
	} else if strings.HasPrefix(topicName, "sf") {
		dataType = "scrips"
		fieldNames = fieldMappings["data_val"]
		f.scripsSym[topicID] = topicName
		f.resp[topicName] = make(map[string]interface{})
	} else {
		return offset
	}

	// Parse field count
	if offset+1 >= len(data) {
		return offset
	}
	fieldCount := data[offset]
	offset++

	// Parse field values
	for i := 0; i < int(fieldCount); i++ {
		if offset+4 > len(data) {
			break
		}
		value := int32(binary.BigEndian.Uint32(data[offset : offset+4]))
		offset += 4

		if value != -2147483648 && i < len(fieldNames) {
			f.resp[topicName][fieldNames[i]] = value
		}
	}

	// Skip 2 bytes
	offset += 2

	// Parse multiplier and precision
	if offset+3 <= len(data) {
		multiplier := binary.BigEndian.Uint16(data[offset : offset+2])
		offset += 2
		precision := data[offset]
		offset++

		f.resp[topicName]["multiplier"] = multiplier
		f.resp[topicName]["precision"] = precision
	}

	// Parse exchange, exchange_token, symbol
	valNames := []string{"exchange", "exchange_token", "symbol"}
	for _, valName := range valNames {
		if offset+1 <= len(data) {
			stringLen := data[offset]
			offset++
			if offset+int(stringLen) <= len(data) {
				stringData := string(data[offset : offset+int(stringLen)])
				f.resp[topicName][valName] = stringData
				offset += int(stringLen)
			}
		}
	}

	f.resp[topicName]["type"] = dataType
	if symbol, exists := f.symbolDict[topicName]; exists {
		f.resp[topicName]["symbol"] = symbol
	}

	// Apply precision and multiplier calculations
	processedResponse := f.applyPrecisionAndMultiplier(f.resp[topicName], dataType)

	// Send decoded data to callback
	if f.OnMessage != nil {
		f.OnMessage(processedResponse)
	}

	// Seed lastLtp from snapshot so the first lite tick with the same value does not emit again (avoid duplicate response).
	if dataType == "scrips" {
		if ltp, ok := f.resp[topicName]["ltp"].(int32); ok {
			f.lastLtpScrips[topicName] = ltp
		}
	} else if dataType == "index" {
		if ltp, ok := f.resp[topicName]["ltp"].(int32); ok {
			f.lastLtpIndex[topicName] = ltp
		}
	}

	return offset
}

// fullModeTickRelevantScrips: only emit OnMessage when one of these fields changes (avoids spam from exch_feed_time / last_traded_time).
var fullModeTickRelevantScrips = map[string]bool{
	"ltp": true, "vol_traded_today": true, "bid_size": true, "ask_size": true,
	"bid_price": true, "ask_price": true, "last_traded_qty": true, "avg_trade_price": true,
	"low_price": true, "high_price": true, "open_price": true, "prev_close_price": true,
	"OI": true, "lower_ckt": true, "upper_ckt": true,
}

// fullModeTickRelevantIndex: only emit when one of these index fields changes.
var fullModeTickRelevantIndex = map[string]bool{
	"ltp": true, "prev_close_price": true, "high_price": true, "low_price": true, "open_price": true,
}

// handleFullModeData handles full mode data feed
func (f *FyersDataSocket) handleFullModeData(data []byte, offset int, fieldMappings map[string][]string) int {
	if offset+2 >= len(data) {
		return offset
	}

	topicID := binary.BigEndian.Uint16(data[offset : offset+2])
	offset += 2

	if offset+1 >= len(data) {
		return offset
	}
	fieldCount := data[offset]
	offset++

	// Track flags for different data types
	sfFlag, idxFlag, dpFlag := false, false, false
	f.updateTick = false

	// Process field values and check for changes
	for i := 0; i < int(fieldCount); i++ {
		if offset+4 > len(data) {
			break
		}
		value := int32(binary.BigEndian.Uint32(data[offset : offset+4]))
		offset += 4

		// Check if topic ID exists in our mappings
		if topicName, exists := f.scripsSym[topicID]; exists {
			if i < len(fieldMappings["data_val"]) {
				fieldName := fieldMappings["data_val"][i]
				if existingValue, hasValue := f.resp[topicName][fieldName]; hasValue {
					if existingValue != value && value != -2147483648 {
						f.resp[topicName][fieldName] = value
						if fullModeTickRelevantScrips[fieldName] {
							f.updateTick = true
						}
					}
				} else if value != -2147483648 {
					f.resp[topicName][fieldName] = value
					if fullModeTickRelevantScrips[fieldName] {
						f.updateTick = true
					}
				}
			}
			sfFlag = true
		} else if topicName, exists := f.indexSym[topicID]; exists {
			if i < len(fieldMappings["index_val"]) {
				fieldName := fieldMappings["index_val"][i]
				if existingValue, hasValue := f.resp[topicName][fieldName]; hasValue {
					if existingValue != value && value != -2147483648 {
						f.resp[topicName][fieldName] = value
						if fullModeTickRelevantIndex[fieldName] {
							f.updateTick = true
						}
					}
				} else if value != -2147483648 {
					f.resp[topicName][fieldName] = value
					if fullModeTickRelevantIndex[fieldName] {
						f.updateTick = true
					}
				}
			}
			idxFlag = true
		} else if topicName, exists := f.dpSym[topicID]; exists {
			if i < len(fieldMappings["depthvalue"]) {
				fieldName := fieldMappings["depthvalue"][i]
				if existingValue, hasValue := f.resp[topicName][fieldName]; hasValue {
					if existingValue != value && value != -2147483648 {
						f.resp[topicName][fieldName] = value
						f.updateTick = true
					}
				} else if value != -2147483648 {
					f.resp[topicName][fieldName] = value
					f.updateTick = true
				}
			}
			dpFlag = true
		}
	}

	// Send update only if there were changes
	if f.updateTick {
		if sfFlag {
			if topicName, exists := f.scripsSym[topicID]; exists {
				// Ensure precision and multiplier are preserved for full mode updates
				if _, hasPrecision := f.resp[topicName]["precision"]; !hasPrecision {
					// If precision is missing, set default values
					f.resp[topicName]["precision"] = uint8(2)
					f.resp[topicName]["multiplier"] = uint16(100)
				}
				processedResponse := f.applyPrecisionAndMultiplier(f.resp[topicName], "scrips")
				if f.OnMessage != nil {
					f.OnMessage(processedResponse)
				}
			}
		} else if idxFlag {
			if topicName, exists := f.indexSym[topicID]; exists {
				// Ensure precision and multiplier are preserved for full mode updates
				if _, hasPrecision := f.resp[topicName]["precision"]; !hasPrecision {
					// If precision is missing, set default values
					f.resp[topicName]["precision"] = uint8(2)
					f.resp[topicName]["multiplier"] = uint16(100)
				}
				processedResponse := f.applyPrecisionAndMultiplier(f.resp[topicName], "index")
				if f.OnMessage != nil {
					f.OnMessage(processedResponse)
				}
			}
		} else if dpFlag {
			if topicName, exists := f.dpSym[topicID]; exists {
				// Ensure precision and multiplier are preserved for full mode updates
				if _, hasPrecision := f.resp[topicName]["precision"]; !hasPrecision {
					// If precision is missing, set default values
					f.resp[topicName]["precision"] = uint8(2)
					f.resp[topicName]["multiplier"] = uint16(100)
				}
				processedResponse := f.applyPrecisionAndMultiplier(f.resp[topicName], "depth")
				if f.OnMessage != nil {
					f.OnMessage(processedResponse)
				}
			}
		}
	}

	return offset
}

// FloatSDK is a float64 that marshals to JSON as "x.0" when whole (to match Python SDK output).
type FloatSDK float64

func (f FloatSDK) MarshalJSON() ([]byte, error) {
	v := float64(f)
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return json.Marshal(v)
	}
	if v == math.Trunc(v) {
		return []byte(fmt.Sprintf("%.1f", v)), nil
	}
	return json.Marshal(v)
}

// asFloat64 returns the numeric value as float64 whether it was stored as int, float64, or FloatSDK.
func asFloat64(v interface{}) (float64, bool) {
	switch n := v.(type) {
	case int:
		return float64(n), true
	case int32:
		return float64(n), true
	case float64:
		return n, true
	case FloatSDK:
		return float64(n), true
	default:
		return 0, false
	}
}

// applyPrecisionAndMultiplier applies precision and multiplier calculations
func (f *FyersDataSocket) applyPrecisionAndMultiplier(response map[string]interface{}, dataType string) map[string]interface{} {
	precision, precisionOk := response["precision"].(uint8)
	multiplier, multiplierOk := response["multiplier"].(uint16)

	if !precisionOk || !multiplierOk {
		return response
	}

	precisionCalcValue := []string{
		"ltp", "bid_price", "ask_price", "avg_trade_price", "low_price",
		"high_price", "open_price", "prev_close_price",
	}

	newResponse := make(map[string]interface{})

	if f.lite {
		if ltp, exists := response["ltp"]; exists {
			if intValue, ok := ltp.(int32); ok {
				newResponse["ltp"] = FloatSDK(float64(intValue) / (math.Pow(10, float64(precision)) * float64(multiplier)))
			}
		}
		// Symbol with exchange prefix to match Python (e.g. MCX:SILVER26MARFUT)
		if symbol, exists := response["symbol"]; exists {
			symbolStr, _ := symbol.(string)
			if symbolStr != "" && !strings.Contains(symbolStr, ":") {
				if exchange, hasExchange := response["exchange"]; hasExchange {
					if exStr, ok := exchange.(string); ok && exStr != "" {
						symbolStr = exStr + ":" + symbolStr
					} else {
						symbolStr = "NSE:" + symbolStr
					}
				} else {
					symbolStr = "NSE:" + symbolStr
				}
			}
			newResponse["symbol"] = symbolStr
		}
		newResponse["type"] = "sf"
	} else {
		switch dataType {
		case "scrips":
			fieldMappings := f.fieldMappings
			if fieldMappings != nil {
				for _, fieldName := range fieldMappings["data_val"] {
					if value, exists := response[fieldName]; exists {
						if intValue, ok := value.(int32); ok {
							needsPrecision := false
							for _, precisionField := range precisionCalcValue {
								if fieldName == precisionField {
									needsPrecision = true
									break
								}
							}
							if needsPrecision && fieldName != "upper_ckt" && fieldName != "lower_ckt" {
								newResponse[fieldName] = FloatSDK(float64(intValue) / (math.Pow(10, float64(precision)) * float64(multiplier)))
							} else {
								newResponse[fieldName] = value
							}
						} else {
							newResponse[fieldName] = value
						}
					}
				}
			}
			newResponse["lower_ckt"] = 0
			newResponse["upper_ckt"] = 0
			delete(newResponse, "OI")
			delete(newResponse, "Yhigh")
			delete(newResponse, "Ylow")
			newResponse["type"] = "sf"
			if symbol, exists := newResponse["symbol"]; exists {
				symbolStr := symbol.(string)
				if !strings.Contains(symbolStr, ":") {
					symbolStr = "NSE:" + symbolStr
				}
				newResponse["symbol"] = symbolStr
			}
			if ltp, ltpOk := asFloat64(newResponse["ltp"]); ltpOk {
				if prevClose, prevCloseOk := asFloat64(newResponse["prev_close_price"]); prevCloseOk && prevClose != 0 {
					change := ltp - prevClose
					changePercent := (change / prevClose) * 100
					newResponse["ch"] = FloatSDK(math.Round(change*10000) / 10000)
					newResponse["chp"] = FloatSDK(math.Round(changePercent*10000) / 10000)
				}
			}
		case "index":
			fieldMappings := f.fieldMappings
			if fieldMappings != nil {
				for _, fieldName := range fieldMappings["index_val"] {
					if value, exists := response[fieldName]; exists {
						if intValue, ok := value.(int32); ok {
							needsPrecision := false
							for i, indexField := range fieldMappings["index_val"] {
								if fieldName == indexField && (i == 0 || i == 1 || i == 3 || i == 4 || i == 5) {
									needsPrecision = true
									break
								}
							}
							if needsPrecision {
								newResponse[fieldName] = FloatSDK(float64(intValue) / (math.Pow(10, float64(precision)) * float64(multiplier)))
							} else {
								newResponse[fieldName] = value
							}
						} else {
							newResponse[fieldName] = value
						}
					}
				}
			}
			newResponse["type"] = "if"
			if ltp, ltpOk := asFloat64(newResponse["ltp"]); ltpOk {
				if prevClose, prevCloseOk := asFloat64(newResponse["prev_close_price"]); prevCloseOk && prevClose != 0 {
					change := ltp - prevClose
					changePercent := (change / prevClose) * 100
					newResponse["ch"] = FloatSDK(math.Round(change*100) / 100)
					newResponse["chp"] = FloatSDK(math.Round(changePercent*100) / 100)
				}
			}
		case "depth":
			fieldMappings := f.fieldMappings
			if fieldMappings != nil {
				for i, fieldName := range fieldMappings["depthvalue"] {
					if value, exists := response[fieldName]; exists {
						if intValue, ok := value.(int32); ok {
							if i < 10 {
								newResponse[fieldName] = FloatSDK(float64(intValue) / (math.Pow(10, float64(precision)) * float64(multiplier)))
							} else {
								newResponse[fieldName] = value
							}
						} else {
							newResponse[fieldName] = value
						}
					}
				}
			}
			newResponse["type"] = "dp"
		}
	}
	return newResponse
}

// createAcknowledgmentMessage creates acknowledgment message
func (f *FyersDataSocket) createAcknowledgmentMessage(messageNumber int) []byte {
	buffer := make([]byte, 0, 10)

	// Length (will be calculated)
	buffer = append(buffer, 0, 0)

	// Request type (6 for acknowledgment)
	buffer = append(buffer, 6)

	// Message number (4 bytes, big endian)
	msgNumBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(msgNumBytes, uint32(messageNumber))
	buffer = append(buffer, msgNumBytes...)

	// Update length
	length := len(buffer) - 2
	binary.BigEndian.PutUint16(buffer[:2], uint16(length))

	return buffer
}

// processMessageQueue processes queued messages
func (f *FyersDataSocket) processMessageQueue() {
	for {
		select {
		case msg := <-f.messageQueue:
			if f.wsRun != nil && f.connected {
				err := f.wsRun.WriteMessage(websocket.BinaryMessage, msg)
				if err != nil {
					if f.OnError != nil {
						f.OnError(map[string]interface{}{"error": err.Error()})
					}
				}
			}
		case <-f.stopChan:
			return
		}
	}
}

// createLiteModeMessage creates the lite mode message in binary format
func (f *FyersDataSocket) createLiteModeMessage() []byte {
	buffer := make([]byte, 0, 20)

	// Length (0 for now, will be calculated)
	buffer = append(buffer, 0, 0)

	// Request type (12 for mode change)
	buffer = append(buffer, 12)

	// Field count (2 fields)
	buffer = append(buffer, 2)

	// Field-1: Channel bits
	buffer = append(buffer, 1)               // field ID
	buffer = append(buffer, 0, 8)            // length = 8
	channelBits := uint64(1) << f.channelNum // Set bit for channel
	channelBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(channelBytes, channelBits)
	buffer = append(buffer, channelBytes...)

	// Field-2: Mode (76 for lite mode)
	buffer = append(buffer, 2)    // field ID
	buffer = append(buffer, 0, 1) // length = 1
	buffer = append(buffer, 76)   // lite mode

	// Update length
	length := len(buffer) - 2
	binary.BigEndian.PutUint16(buffer[:2], uint16(length))

	return buffer
}

// createFullModeMessage creates the full mode message in binary format
func (f *FyersDataSocket) createFullModeMessage() []byte {
	buffer := make([]byte, 0, 20)

	// Length (0 for now, will be calculated)
	buffer = append(buffer, 0, 0)

	// Request type (12 for mode change)
	buffer = append(buffer, 12)

	// Field count (2 fields)
	buffer = append(buffer, 2)

	// Field-1: Channel bits
	buffer = append(buffer, 1)               // field ID
	buffer = append(buffer, 0, 8)            // length = 8
	channelBits := uint64(1) << f.channelNum // Set bit for channel
	channelBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(channelBytes, channelBits)
	buffer = append(buffer, channelBytes...)

	// Field-2: Mode (70 for full mode)
	buffer = append(buffer, 2)    // field ID
	buffer = append(buffer, 0, 1) // length = 1
	buffer = append(buffer, 70)   // full mode

	// Update length
	length := len(buffer) - 2
	binary.BigEndian.PutUint16(buffer[:2], uint16(length))

	return buffer
}

// createSubscriptionMessage creates the subscription message in binary format
func (f *FyersDataSocket) createSubscriptionMessage(symbols map[string]string) []byte {
	// Create symbols data
	symbolsData := make([]byte, 0)
	symbolsData = append(symbolsData, 0, 0) // Placeholder for count

	symbolCount := 0
	for hsmSymbol := range symbols {
		symbolBytes := []byte(hsmSymbol)
		symbolsData = append(symbolsData, byte(len(symbolBytes)))
		symbolsData = append(symbolsData, symbolBytes...)
		symbolCount++
	}

	// Update symbol count
	binary.BigEndian.PutUint16(symbolsData[:2], uint16(symbolCount))

	// Calculate total length
	dataLen := 18 + len(symbolsData) + len(f.accessToken) + len(f.source)

	buffer := make([]byte, 0, dataLen)

	// Length
	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, uint16(dataLen))
	buffer = append(buffer, lengthBytes...)

	// Request type (4 for subscription)
	buffer = append(buffer, 4)

	// Field count (2 fields)
	buffer = append(buffer, 2)

	// Field-1: Symbols data
	buffer = append(buffer, 1) // field ID
	symbolsLenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(symbolsLenBytes, uint16(len(symbolsData)))
	buffer = append(buffer, symbolsLenBytes...)
	buffer = append(buffer, symbolsData...)

	// Field-2: Channel number
	buffer = append(buffer, 2)    // field ID
	buffer = append(buffer, 0, 1) // length = 1
	buffer = append(buffer, byte(f.channelNum))

	return buffer
}

// createUnsubscriptionMessage creates the unsubscription message in binary format
func (f *FyersDataSocket) createUnsubscriptionMessage(symbols map[string]string) []byte {
	// Create symbols data
	symbolsData := make([]byte, 0)
	symbolsData = append(symbolsData, 0, 0) // Placeholder for count

	symbolCount := 0
	for hsmSymbol := range symbols {
		symbolBytes := []byte(hsmSymbol)
		symbolsData = append(symbolsData, byte(len(symbolBytes)))
		symbolsData = append(symbolsData, symbolBytes...)
		symbolCount++
	}

	// Update symbol count
	binary.BigEndian.PutUint16(symbolsData[:2], uint16(symbolCount))

	// Calculate total length
	dataLen := 18 + len(symbolsData) + len(f.accessToken) + len(f.source)

	buffer := make([]byte, 0, dataLen)

	// Length
	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, uint16(dataLen))
	buffer = append(buffer, lengthBytes...)

	// Request type (5 for unsubscription)
	buffer = append(buffer, 5)

	// Field count (2 fields)
	buffer = append(buffer, 2)

	// Field-1: Symbols data
	buffer = append(buffer, 1) // field ID
	symbolsLenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(symbolsLenBytes, uint16(len(symbolsData)))
	buffer = append(buffer, symbolsLenBytes...)
	buffer = append(buffer, symbolsData...)

	// Field-2: Channel number
	buffer = append(buffer, 2)    // field ID
	buffer = append(buffer, 0, 1) // length = 1
	buffer = append(buffer, byte(f.channelNum))

	return buffer
}

// slicesEqual returns true if a and b have the same length and elements in the same order.
func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Subscribe subscribes to symbols. Uses cached HSM mapping on reconnect when the symbol-token API fails (e.g. unexpected EOF) so subscription still succeeds.
func (f *FyersDataSocket) Subscribe(symbols []string, dataType string) {
	f.dataType = dataType

	// Convert symbols to HSM tokens (may fail on reconnect if API is flaky)
	sc := newSymbolConversion(f.accessToken, dataType, f.logPath)
	dataDict, wrongSymbols, dpIndexFlag, errorMsg := sc.symbolToHSMToken(symbols)

	if errorMsg != "" || dataDict == nil {
		// API failed (e.g. unexpected EOF right after reconnect). Use cache if we have one for same symbols+dataType.
		f.mu.Lock()
		cached := f.cachedSubscribeDataDict != nil && f.cachedSubscribeDataType == dataType && slicesEqual(f.cachedSubscribeSymbols, symbols)
		dataDictToUse := f.cachedSubscribeDataDict
		f.mu.Unlock()
		if cached && dataDictToUse != nil {
			f.mu.Lock()
			for hsm, userSym := range dataDictToUse {
				f.symbolDict[hsm] = userSym
			}
			f.mu.Unlock()
			msg := f.createSubscriptionMessage(dataDictToUse)
			f.messageQueue <- msg
			return
		}
		if f.OnError != nil {
			f.OnError(map[string]interface{}{"error": errorMsg})
		}
		return
	}

	// API succeeded: update cache and symbolDict (topic -> user symbol, like Python symbol_token)
	f.mu.Lock()
	f.cachedSubscribeSymbols = make([]string, len(symbols))
	copy(f.cachedSubscribeSymbols, symbols)
	f.cachedSubscribeDataType = dataType
	f.cachedSubscribeDataDict = make(map[string]string, len(dataDict))
	for k, v := range dataDict {
		f.cachedSubscribeDataDict[k] = v
	}
	for hsm, userSym := range dataDict {
		f.symbolDict[hsm] = userSym
	}
	f.mu.Unlock()

	if len(wrongSymbols) > 0 {
		if f.OnError != nil {
			f.OnError(map[string]interface{}{"invalid_symbols": wrongSymbols})
		}
	}

	if dpIndexFlag {
		if f.OnError != nil {
			f.OnError(map[string]interface{}{"error": INDEX_DEPTH_ERROR_MESSAGE})
		}
	}

	// Create subscription message in binary format
	msg := f.createSubscriptionMessage(dataDict)
	f.messageQueue <- msg
}

// Unsubscribe unsubscribes from symbols
func (f *FyersDataSocket) Unsubscribe(symbols []string, dataType string) {
	// Convert symbols to HSM tokens
	sc := newSymbolConversion(f.accessToken, dataType, f.logPath)
	dataDict, _, _, _ := sc.symbolToHSMToken(symbols)

	// Create unsubscription message in binary format
	msg := f.createUnsubscriptionMessage(dataDict)
	f.messageQueue <- msg
}

// KeepRunning keeps the WebSocket connection alive
func (f *FyersDataSocket) KeepRunning() {
	// Send periodic ping messages
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if f.wsRun != nil && f.connected {
				err := f.wsRun.WriteMessage(websocket.TextMessage, []byte("ping"))
				if err != nil {
					if f.OnError != nil {
						f.OnError(map[string]interface{}{"error": err.Error()})
					}
				}
			}
		case <-f.stopChan:
			return
		}
	}
}

// CloseConnection closes the WebSocket connection. Sets restartFlag false first so reconnect is not triggered (same as Python close_connection).
func (f *FyersDataSocket) CloseConnection() {
	// fmt.Printf("[FyersDataSocket RECONNECT] CloseConnection: setting restartFlag=false and closing socket\n")
	f.mu.Lock()
	f.restartFlag = false
	if f.wsRun != nil {
		f.wsRun.Close()
		f.wsRun = nil
	}
	f.connected = false
	f.mu.Unlock()

	f.stopOnce.Do(func() { close(f.stopChan) })

	if f.OnClose != nil {
		f.OnClose(map[string]interface{}{
			"code":    SUCCESS_CODE,
			"message": CONNECTION_CLOSED,
			"s":       SUCCESS,
		})
	}
}

// IsConnected returns the connection status
func (f *FyersDataSocket) IsConnected() bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.connected && f.wsRun != nil
}

// handleResumePauseResponse handles resume/pause response
func (f *FyersDataSocket) handleResumePauseResponse(data []byte, respType int) {
	// Implementation for resume/pause response
	if f.OnMessage != nil {
		messageType := "resume"
		if respType == 7 {
			messageType = "pause"
		}
		f.OnMessage(map[string]interface{}{
			"type": messageType,
			"data": data,
		})
	}
}

// handleLiteFullModeResponse handles lite/full mode response
func (f *FyersDataSocket) handleLiteFullModeResponse(data []byte) {
	if len(data) < 4 {
		return
	}

	offset := 3
	if offset+1 >= len(data) {
		return
	}
	fieldCount := data[offset]
	offset += 1

	if fieldCount >= 1 && offset+1+2 < len(data) {
		offset += 1 // field ID
		fieldLength := binary.BigEndian.Uint16(data[offset : offset+2])
		offset += 2

		if offset+int(fieldLength) <= len(data) {
			stringVal := string(data[offset : offset+int(fieldLength)])

			if stringVal == "K" {
				if f.OnMessage != nil {
					messageType := FULL_MODE_TYPE
					message := FULL_MODE
					if f.lite {
						messageType = LITE_MODE_TYPE
						message = LITE_MODE
					}
					f.OnMessage(map[string]interface{}{
						"type":    messageType,
						"code":    SUCCESS_CODE,
						"message": message,
						"s":       SUCCESS,
					})
				}
			} else {
				if f.OnError != nil {
					f.OnError(map[string]interface{}{
						"code":    MODE_ERROR_CODE,
						"message": MODE_CHANGE_ERROR,
						"s":       ERROR,
					})
				}
			}
		}
	}
}

// handleLiteModeData handles lite mode data feed
func (f *FyersDataSocket) handleLiteModeData(data []byte, offset int, fieldMappings map[string][]string) int {
	if offset+2 >= len(data) {
		return offset
	}

	topicID := binary.BigEndian.Uint16(data[offset : offset+2])
	offset += 2

	// Check if topic ID exists in our mappings
	if topicName, exists := f.scripsSym[topicID]; exists {
		if offset+4 <= len(data) {
			value := int32(binary.BigEndian.Uint32(data[offset : offset+4]))
			offset += 4

			// Dedup: emit only when LTP actually changed (same as Python: value != self.resp[...][data_val[0]]).
			if value != -2147483648 && len(fieldMappings["data_val"]) > 0 {
				last := f.lastLtpScrips[topicName]
				if last != value {
					f.lastLtpScrips[topicName] = value
					fieldName := fieldMappings["data_val"][0]
					if f.resp[topicName] == nil {
						f.resp[topicName] = make(map[string]interface{})
					}
					f.resp[topicName][fieldName] = value
					f.resp[topicName]["type"] = "sf"
					processedResponse := f.applyPrecisionAndMultiplier(f.resp[topicName], "scrips")
					if f.OnMessage != nil {
						f.OnMessage(processedResponse)
					}
				}
			}
		}
	} else if topicName, exists := f.indexSym[topicID]; exists {
		if offset+4 <= len(data) {
			value := int32(binary.BigEndian.Uint32(data[offset : offset+4]))
			offset += 4

			// Dedup: emit only when LTP actually changed (same as Python for index).
			if value != -2147483648 && len(fieldMappings["index_val"]) > 0 {
				last := f.lastLtpIndex[topicName]
				if last != value {
					f.lastLtpIndex[topicName] = value
					fieldName := fieldMappings["index_val"][0]
					if f.resp[topicName] == nil {
						f.resp[topicName] = make(map[string]interface{})
					}
					f.resp[topicName][fieldName] = value
					f.resp[topicName]["type"] = "if"
					processedResponse := f.applyPrecisionAndMultiplier(f.resp[topicName], "index")
					if f.OnMessage != nil {
						f.OnMessage(processedResponse)
					}
				}
			}
		}
	}

	return offset
}
