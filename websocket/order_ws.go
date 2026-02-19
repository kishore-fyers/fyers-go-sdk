package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// FyersOrderSocket represents the order WebSocket client
type FyersOrderSocket struct {
	accessToken          string
	logPath              string
	wsObject             *websocket.Conn
	wsRun                bool
	pingThread           *time.Ticker
	writeToFile          bool
	backgroundFlag       bool
	reconnectDelay       int
	onTrades             func(map[string]interface{})
	onPosition           func(map[string]interface{})
	restartFlag          bool
	onOrder              func(map[string]interface{})
	onGeneral            func(map[string]interface{})
	onError              func(map[string]interface{})
	onOpen               func()
	maxReconnectAttempts int
	reconnectAttempts    int
	onClose              func(map[string]interface{})
	runningThread        bool
	url                  string
	positionMapper       map[string]interface{}
	orderMapper          map[string]interface{}
	tradeMapper          map[string]interface{}
	orderLogger          *FyersLogger
	websocketTask        bool
	socketType           map[string]interface{}
	mu                   sync.Mutex
	connected            bool
	stopChan             chan bool
}

// NewFyersOrderSocket creates a new FyersOrderSocket instance
func NewFyersOrderSocket(
	accessToken string,
	writeToFile bool,
	logPath string,
	onTrades func(map[string]interface{}),
	onPositions func(map[string]interface{}),
	onOrders func(map[string]interface{}),
	onGeneral func(map[string]interface{}),
	onError func(map[string]interface{}),
	onConnect func(),
	onClose func(map[string]interface{}),
	reconnect bool,
	reconnectRetry int,
) *FyersOrderSocket {

	maxReconnectAttempts := 50
	if reconnectRetry < maxReconnectAttempts {
		maxReconnectAttempts = reconnectRetry
	}

	// Load map.json
	mapData, err := loadMapJSON()
	if err != nil {
		fmt.Printf("Failed to load map.json: %v\n", err)
	}

	positionMapper := make(map[string]interface{})
	orderMapper := make(map[string]interface{})
	tradeMapper := make(map[string]interface{})

	if mapData != nil {
		if posMap, exists := mapData["position_mapper"]; exists {
			positionMapper = posMap.(map[string]interface{})
		}
		if ordMap, exists := mapData["order_mapper"]; exists {
			orderMapper = ordMap.(map[string]interface{})
		}
		if trdMap, exists := mapData["trade_mapper"]; exists {
			tradeMapper = trdMap.(map[string]interface{})
		}
	}

	// Setup logger
	var loggerPath string
	if logPath != "" {
		loggerPath = logPath // Use the logPath as is (it's already a directory)
	} else {
		loggerPath = "" // Empty string means no logging to file
	}

	orderLogger := NewFyersLogger("FyersOrderSocket", "DEBUG", 2, loggerPath)

	socketType := map[string]interface{}{
		"OnOrders":    "orders",
		"OnTrades":    "trades",
		"OnPositions": "positions",
		"OnGeneral":   []string{"edis", "pricealerts", "login"},
	}

	return &FyersOrderSocket{
		accessToken:          accessToken,
		logPath:              logPath,
		wsObject:             nil,
		wsRun:                false,
		pingThread:           nil,
		writeToFile:          writeToFile,
		backgroundFlag:       false,
		reconnectDelay:       0,
		onTrades:             onTrades,
		onPosition:           onPositions,
		restartFlag:          reconnect,
		onOrder:              onOrders,
		onGeneral:            onGeneral,
		onError:              onError,
		onOpen:               onConnect,
		maxReconnectAttempts: maxReconnectAttempts,
		reconnectAttempts:    0,
		onClose:              onClose,
		runningThread:        false,
		url:                  "wss://socket.fyers.in/trade/v3",
		positionMapper:       positionMapper,
		orderMapper:          orderMapper,
		tradeMapper:          tradeMapper,
		orderLogger:          orderLogger,
		websocketTask:        false,
		socketType:           socketType,
		connected:            false,
		stopChan:             make(chan bool),
	}
}

// loadMapJSON loads the map.json file
func loadMapJSON() (map[string]interface{}, error) {
	// Get the directory of the current file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get current file path")
	}

	dir := filepath.Dir(filename)
	mapPath := filepath.Join(dir, "map.json")

	data, err := ioutil.ReadFile(mapPath)
	if err != nil {
		return nil, err
	}

	var mapData map[string]interface{}
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		return nil, err
	}

	return mapData, nil
}

// parsePositionData parses position data from a message
func (f *FyersOrderSocket) parsePositionData(msg map[string]interface{}) map[string]interface{} {
	positionData := make(map[string]interface{})

	if positions, exists := msg["positions"]; exists {
		if posMap, ok := positions.(map[string]interface{}); ok {
			// Use the position mapper to convert field names like Python version
			for rawKey, mappedKey := range f.positionMapper {
				if val, exists := posMap[rawKey]; exists {
					positionData[mappedKey.(string)] = val
				}
			}
		}
	}

	return map[string]interface{}{
		"s":         msg["s"],
		"positions": positionData,
	}
}

// parseTradeData parses trade data from a message
func (f *FyersOrderSocket) parseTradeData(msg map[string]interface{}) map[string]interface{} {
	tradeData := make(map[string]interface{})

	if trades, exists := msg["trades"]; exists {
		if trdMap, ok := trades.(map[string]interface{}); ok {
			// Use the trade mapper to convert field names like Python version
			for rawKey, mappedKey := range f.tradeMapper {
				if val, exists := trdMap[rawKey]; exists {
					tradeData[mappedKey.(string)] = val
				}
			}
		}
	}

	return map[string]interface{}{
		"s":      msg["s"],
		"trades": tradeData,
	}
}

// parseOrderData parses order data from a message
func (f *FyersOrderSocket) parseOrderData(msg map[string]interface{}) map[string]interface{} {
	orderData := make(map[string]interface{})

	if orders, exists := msg["orders"]; exists {
		if ordMap, ok := orders.(map[string]interface{}); ok {
			// Use the order mapper to convert field names like Python version
			for rawKey, mappedKey := range f.orderMapper {
				if val, exists := ordMap[rawKey]; exists {
					orderData[mappedKey.(string)] = val
				}
			}

			// Add orderNumStatus like Python version
			if id, exists := ordMap["id"]; exists {
				if orgOrdStatus, exists := ordMap["org_ord_status"]; exists {
					orderData["orderNumStatus"] = fmt.Sprintf("%v:%v", id, orgOrdStatus)
				}
			}
		}
	}

	return map[string]interface{}{
		"s":      msg["s"],
		"orders": orderData,
	}
}

// OnTrades handles trade events
func (f *FyersOrderSocket) OnTrades(message map[string]interface{}) {
	if f.onTrades != nil {
		f.onTrades(message)
	} else {
		fmt.Printf("Trade : %v\n", message)
	}
}

// OnPositions handles position events
func (f *FyersOrderSocket) OnPositions(message map[string]interface{}) {
	if f.onPosition != nil {
		f.onPosition(message)
	} else {
		fmt.Printf("Position : %v\n", message)
	}
}

// OnOrder handles order events
func (f *FyersOrderSocket) OnOrder(message map[string]interface{}) {
	if f.onOrder != nil {
		f.onOrder(message)
	} else {
		fmt.Printf("Order : %v\n", message)
	}
}

// OnGeneral handles general events
func (f *FyersOrderSocket) OnGeneral(message map[string]interface{}) {
	if f.onGeneral != nil {
		f.onGeneral(message)
	} else {
		fmt.Printf("General : %v\n", message)
	}
}

// OnError handles error events
func (f *FyersOrderSocket) OnError(message interface{}) {
	if f.onError != nil {
		f.onError(map[string]interface{}{"error": message})
	} else {
		fmt.Printf("Error : %v\n", message)
	}
}

// handleMessage processes incoming messages
func (f *FyersOrderSocket) handleMessage(message []byte) {
	// Convert message to string first
	messageStr := string(message)

	// Handle pong messages like Python version
	if messageStr == "pong" {
		return
	}

	var data map[string]interface{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		f.OnError("Failed to parse message")
		return
	}

	// Use field-based parsing like Python version
	if _, exists := data["orders"]; exists {
		parsedData := f.parseOrderData(data)
		f.OnOrder(parsedData)
	} else if _, exists := data["positions"]; exists {
		parsedData := f.parsePositionData(data)
		f.OnPositions(parsedData)
	} else if _, exists := data["trades"]; exists {
		parsedData := f.parseTradeData(data)
		f.OnTrades(parsedData)
	} else {
		f.OnGeneral(data)
	}
}

// readMessages reads messages from WebSocket
func (f *FyersOrderSocket) readMessages() {
	for {
		select {
		case <-f.stopChan:
			return
		default:
			if f.wsObject == nil {
				return
			}

			_, message, err := f.wsObject.ReadMessage()
			if err != nil {
				// Check if this is a normal close error
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					f.OnError(err.Error())
				} else {
					// Normal close, just log it
					f.orderLogger.Info("WebSocket connection closed normally")
				}
				return
			}

			f.handleMessage(message)
		}
	}
}

// ping sends periodic ping messages
func (f *FyersOrderSocket) ping() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			f.mu.Lock()
			if f.wsObject != nil && f.connected {
				err := f.wsObject.WriteMessage(websocket.TextMessage, []byte("ping"))
				if err != nil {
					f.orderLogger.Error("Failed to send ping: " + err.Error())
					// Don't call OnError for ping failures as they're not critical
				}
			}
			f.mu.Unlock()
		case <-f.stopChan:
			return
		}
	}
}

// Connect establishes WebSocket connection
func (f *FyersOrderSocket) Connect() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	// Set up headers with authorization
	headers := http.Header{}
	headers.Set("authorization", f.accessToken)

	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(f.url, headers)
	if err != nil {
		return err
	}

	f.wsObject = conn
	f.connected = true
	f.wsRun = true

	// Display authentication success message
	authResponse := map[string]interface{}{
		"type":    AUTH_TYPE,
		"code":    200,
		"message": AUTH_SUCCESS,
		"s":       "ok",
	}
	jsonData, _ := json.Marshal(authResponse)
	fmt.Printf("Response: %s\n", string(jsonData))

	// Start reading messages
	go f.readMessages()

	// Start ping thread
	go f.ping()

	// Call onOpen callback
	if f.onOpen != nil {
		f.onOpen()
	}

	return nil
}

// socketTypeToSlist converts one or more socket type keys (e.g. "OnOrders", "OnTrades")
// into a single SLIST slice, matching Python's logic: comma-separated or multiple keys
// map to one list; list values (e.g. OnGeneral) are expanded.
func (f *FyersOrderSocket) socketTypeToSlist(dataTypes []string) []string {
	var dataTypeList []string
	for _, dataType := range dataTypes {
		if socketType, exists := f.socketType[dataType]; exists {
			if list, ok := socketType.([]string); ok {
				dataTypeList = append(dataTypeList, list...)
			} else if str, ok := socketType.(string); ok {
				dataTypeList = append(dataTypeList, str)
			}
		}
	}
	return dataTypeList
}

// Subscribe subscribes to a single data type.
func (f *FyersOrderSocket) Subscribe(dataType string) {
	f.SubscribeMultiple([]string{dataType})
}

// SubscribeMultiple subscribes to multiple data types in one SUB_ORD message (like Python:
// one message with SLIST e.g. ["orders","trades","positions"] so the server returns one
// "Successfully subscribed" instead of one per type).
func (f *FyersOrderSocket) SubscribeMultiple(dataTypes []string) {
	dataTypeList := f.socketTypeToSlist(dataTypes)
	if len(dataTypeList) == 0 {
		return
	}

	msg := map[string]interface{}{
		"T":     "SUB_ORD",
		"SLIST": dataTypeList,
		"SUB_T": 1,
	}

	jsonData, _ := json.Marshal(msg)
	if f.wsObject != nil && f.connected {
		err := f.wsObject.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			f.OnError(err.Error())
		}
	}
}

// Unsubscribe unsubscribes from a single data type.
func (f *FyersOrderSocket) Unsubscribe(dataType string) {
	f.UnsubscribeMultiple([]string{dataType})
}

// UnsubscribeMultiple unsubscribes from multiple data types in one SUB_ORD message.
func (f *FyersOrderSocket) UnsubscribeMultiple(dataTypes []string) {
	dataTypeList := f.socketTypeToSlist(dataTypes)
	if len(dataTypeList) == 0 {
		return
	}

	msg := map[string]interface{}{
		"T":     "SUB_ORD",
		"SLIST": dataTypeList,
		"SUB_T": -1,
	}

	jsonData, _ := json.Marshal(msg)
	if f.wsObject != nil && f.connected {
		err := f.wsObject.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			f.OnError(err.Error())
		}
	}
}

// KeepRunning keeps the WebSocket connection alive
func (f *FyersOrderSocket) KeepRunning() {
	select {}
}

// StopRunning stops the WebSocket connection
func (f *FyersOrderSocket) StopRunning() {
	f.CloseConnection()
}

// CloseConnection closes the WebSocket connection
func (f *FyersOrderSocket) CloseConnection() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.wsObject != nil {
		f.wsObject.Close()
		f.wsObject = nil
	}

	f.connected = false
	f.wsRun = false

	// Only close stopChan if it hasn't been closed already
	select {
	case <-f.stopChan:
		// Channel already closed, do nothing
	default:
		close(f.stopChan)
	}

	if f.onClose != nil {
		f.onClose(map[string]interface{}{"message": "Connection closed"})
	}
}

// IsConnected returns the connection status
func (f *FyersOrderSocket) IsConnected() bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.connected && f.wsObject != nil
}
