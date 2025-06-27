package fyersgosdk

import (
	"encoding/json"
	"fmt"
	fyersws "fyers-go-sdk/websocket"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

// Store the latest data messages for retrieval
var latestSymbolUpdateStruct SymbolUpdateResponse
var latestDepthUpdateStruct DepthUpdateResponse
var latestIndexUpdateStruct IndexUpdateResponse
var latestDataMessageJSON string

// Global callback for real-time data
var globalDataCallback DataSocketCallback

// SymbolUpdateResponse struct for SymbolUpdate data type
type SymbolUpdateResponse struct {
	AskPrice       float64 `json:"ask_price"`
	AskSize        int     `json:"ask_size"`
	AvgTradePrice  float64 `json:"avg_trade_price"`
	BidPrice       float64 `json:"bid_price"`
	BidSize        int     `json:"bid_size"`
	Ch             float64 `json:"ch"`
	Chp            float64 `json:"chp"`
	ExchFeedTime   int64   `json:"exch_feed_time"`
	HighPrice      float64 `json:"high_price"`
	LastTradedQty  int     `json:"last_traded_qty"`
	LastTradedTime int64   `json:"last_traded_time"`
	LowPrice       float64 `json:"low_price"`
	LowerCkt       float64 `json:"lower_ckt"`
	Ltp            float64 `json:"ltp"`
	OpenPrice      float64 `json:"open_price"`
	PrevClosePrice float64 `json:"prev_close_price"`
	Symbol         string  `json:"symbol"`
	TotBuyQty      int     `json:"tot_buy_qty"`
	TotSellQty     int     `json:"tot_sell_qty"`
	Type           string  `json:"type"`
	UpperCkt       float64 `json:"upper_ckt"`
	VolTradedToday int     `json:"vol_traded_today"`
}

// DepthUpdateResponse struct for DepthUpdate data type
type DepthUpdateResponse struct {
	Symbol    string      `json:"symbol"`
	Type      string      `json:"type"`
	Bids      [][]float64 `json:"bids"`
	Asks      [][]float64 `json:"asks"`
	Timestamp int64       `json:"timestamp"`
}

// IndexUpdateResponse struct for IndexUpdate data type
type IndexUpdateResponse struct {
	Symbol        string  `json:"symbol"`
	Type          string  `json:"type"`
	IndexValue    float64 `json:"index_value"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"change_percent"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Open          float64 `json:"open"`
	PrevClose     float64 `json:"prev_close"`
	Volume        int64   `json:"volume"`
	Timestamp     int64   `json:"timestamp"`
}

// DataSocketCallback is the callback function type for real-time data
type DataSocketCallback func(interface{}, string)

// DataSocket connects to the Fyers Data WebSocket and streams data via callback
func DataSocket(fyClient *Client, webSocketRequest DataSocketRequest, callback DataSocketCallback) error {
	// Set global callback for onDataMessage to use
	globalDataCallback = callback

	// Replace with your actual access token
	accessTokenStr := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)

	// Create a FyersDataSocket instance
	dataSocket := fyersws.NewFyersDataSocket(
		accessTokenStr, // Access token in the format "appid:accesstoken"
		"",             // Log path - leave empty to auto-create logs in the current directory
		false,          // Lite mode disabled. Set to true if you want a lite response
		false,          // Save response in a log file instead of printing it
		true,           // Enable auto-reconnection to WebSocket on disconnection
		onDataConnect,  // Callback function to subscribe to data upon connection
		onDataClose,    // Callback function to handle WebSocket connection close events
		onDataError,    // Callback function to handle WebSocket errors
		onDataMessage,  // Callback function to handle incoming messages from the WebSocket
	)

	// Establish a connection to the Fyers Data WebSocket
	err := dataSocket.Connect()
	if err != nil {
		return fmt.Errorf("failed to connect to Data Socket: %v", err)
	}

	// Subscribe to symbols
	dataSocket.Subscribe(webSocketRequest.Symbols, webSocketRequest.DataType)

	// Set up signal handling to keep the connection alive
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for interrupt signal
	<-sigChan
	fmt.Println("\nReceived interrupt signal, closing connection...")

	// Close the connection
	dataSocket.CloseConnection()
	fmt.Println("Data Socket connection closed")

	return nil
}

// DataSocketSimple connects to the Fyers Data WebSocket without callback (JSON only)
func DataSocketSimple(fyClient *Client, webSocketRequest DataSocketRequest) error {
	// Set global callback to nil to skip callback processing
	globalDataCallback = nil

	// Replace with your actual access token
	accessTokenStr := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)

	// Create a FyersDataSocket instance
	dataSocket := fyersws.NewFyersDataSocket(
		accessTokenStr, // Access token in the format "appid:accesstoken"
		"",             // Log path - leave empty to auto-create logs in the current directory
		false,          // Lite mode disabled. Set to true if you want a lite response
		false,          // Save response in a log file instead of printing it
		true,           // Enable auto-reconnection to WebSocket on disconnection
		onDataConnect,  // Callback function to subscribe to data upon connection
		onDataClose,    // Callback function to handle WebSocket connection close events
		onDataError,    // Callback function to handle WebSocket errors
		func(message map[string]interface{}) {
			onDataMessageSimple(message, webSocketRequest.Fields)
		}, // Custom message handler for simple mode
	)

	// Establish a connection to the Fyers Data WebSocket
	err := dataSocket.Connect()
	if err != nil {
		return fmt.Errorf("failed to connect to Data Socket: %v", err)
	}

	// Subscribe to symbols
	dataSocket.Subscribe(webSocketRequest.Symbols, webSocketRequest.DataType)

	// Set up signal handling to keep the connection alive
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for interrupt signal
	<-sigChan
	fmt.Println("\nReceived interrupt signal, closing connection...")

	// Close the connection
	dataSocket.CloseConnection()
	fmt.Println("Data Socket connection closed")

	return nil
}

// onDataMessageSimple handles messages for simple mode with field filtering but no callback
func onDataMessageSimple(message map[string]interface{}, requestedFields []string) {
	jsonData, _ := json.Marshal(message)
	jsonString := string(jsonData)
	latestDataMessageJSON = jsonString

	// Determine data type and unmarshal to appropriate struct
	if dataType, exists := message["type"]; exists {
		switch dataType {
		case "sf": // Symbol feed
			var resp SymbolUpdateResponse
			if err := json.Unmarshal(jsonData, &resp); err == nil {
				latestSymbolUpdateStruct = resp

				// Apply field filtering and print filtered struct
				filteredStruct := filterStructFields(resp, requestedFields)
				fmt.Println("Response: ", jsonString)
				fmt.Printf("Filtered Struct: %+v\n", filteredStruct)
			}
		case "dp": // Depth feed
			var resp DepthUpdateResponse
			if err := json.Unmarshal(jsonData, &resp); err == nil {
				latestDepthUpdateStruct = resp

				// Apply field filtering and print filtered struct
				filteredStruct := filterStructFields(resp, requestedFields)
				fmt.Println("Response: ", jsonString)
				fmt.Printf("Filtered Struct: %+v\n", filteredStruct)
			}
		case "if": // Index feed
			var resp IndexUpdateResponse
			if err := json.Unmarshal(jsonData, &resp); err == nil {
				latestIndexUpdateStruct = resp

				// Apply field filtering and print filtered struct
				filteredStruct := filterStructFields(resp, requestedFields)
				fmt.Println("Response: ", jsonString)
				fmt.Printf("Filtered Struct: %+v\n", filteredStruct)
			}
		default:
			// For other message types (auth, subscription, etc.), just print JSON
			fmt.Println("Response: ", jsonString)
		}
	} else {
		fmt.Println("Response: ", jsonString)
	}
}

// Data Socket callback functions
func onDataMessage(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	jsonString := string(jsonData)
	latestDataMessageJSON = jsonString

	// If no callback is provided, only print JSON and skip struct processing
	if globalDataCallback == nil {
		fmt.Println("Response: ", jsonString)
		return
	}

	// Determine data type and unmarshal to appropriate struct
	if dataType, exists := message["type"]; exists {
		switch dataType {
		case "sf": // Symbol feed
			var resp SymbolUpdateResponse
			if err := json.Unmarshal(jsonData, &resp); err == nil {
				latestSymbolUpdateStruct = resp

				// Call global callback with filtered data
				filteredStruct := filterStructFields(resp, []string{"Ch", "Ltp", "Symbol"}) // Default fields
				globalDataCallback(filteredStruct, jsonString)

				fmt.Println("Response: ", jsonString)
			}
		case "dp": // Depth feed
			var resp DepthUpdateResponse
			if err := json.Unmarshal(jsonData, &resp); err == nil {
				latestDepthUpdateStruct = resp

				// Call global callback with filtered data
				filteredStruct := filterStructFields(resp, []string{"Symbol", "Bids", "Asks"}) // Default fields
				globalDataCallback(filteredStruct, jsonString)

				fmt.Println("Response: ", jsonString)
			}
		case "if": // Index feed
			var resp IndexUpdateResponse
			if err := json.Unmarshal(jsonData, &resp); err == nil {
				latestIndexUpdateStruct = resp

				// Call global callback with filtered data
				filteredStruct := filterStructFields(resp, []string{"IndexValue", "Change", "ChangePercent"}) // Default fields
				globalDataCallback(filteredStruct, jsonString)

				fmt.Println("Response: ", jsonString)
			}
		default:
			// For other message types (auth, subscription, etc.), pass the original message
			globalDataCallback(message, jsonString)
		}
	} else {
		globalDataCallback(message, jsonString)
		fmt.Println("Response: ", jsonString)
	}
}

// filterStructFields creates a copy of the struct with only requested fields populated
func filterStructFields(structData interface{}, requestedFields []string) interface{} {
	if len(requestedFields) == 0 {
		// If no fields specified, return all fields
		return structData
	}

	// Create a copy of the original struct
	originalValue := reflect.ValueOf(structData)
	if originalValue.Kind() == reflect.Ptr {
		originalValue = originalValue.Elem()
	}

	// Create a new instance of the same type
	newValue := reflect.New(originalValue.Type()).Elem()

	// Copy all fields from original to new
	for i := 0; i < originalValue.NumField(); i++ {
		newValue.Field(i).Set(originalValue.Field(i))
	}

	// Create a map of requested field names for quick lookup
	requestedMap := make(map[string]bool)
	for _, field := range requestedFields {
		requestedMap[field] = true
	}

	// Zero out fields that are not requested
	for i := 0; i < newValue.NumField(); i++ {
		field := newValue.Type().Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}

		// If this field is not in the requested fields, zero it out
		if !requestedMap[jsonTag] && !requestedMap[field.Name] {
			newValue.Field(i).Set(reflect.Zero(field.Type))
		}
	}

	return newValue.Interface()
}

// Data Socket callback functions
func onDataError(message map[string]interface{}) {
	if invalid, exists := message["invalid_symbols"]; exists {
		resp := map[string]interface{}{
			"type":            "sub",
			"code":            -300,
			"message":         "Please provide a valid symbol",
			"s":               "error",
			"invalid_symbols": invalid,
		}
		jsonResp, _ := json.Marshal(resp)
		fmt.Printf("Error: %s\n", jsonResp)
		return
	}
	jsonData, _ := json.Marshal(message)
	fmt.Printf("Error: %s\n", string(jsonData))
}

func onDataClose(message map[string]interface{}) {
	fmt.Printf("Connection closed: %v\n", message)
}

func onDataConnect() {
	// fmt.Println("Data Socket - Connection established, subscribing to symbols...")
}

// Order Socket Example
func OrderSocket(fyClient *Client, orderSocketRequest OrderSocketRequest) (map[string]interface{}, error) {
	// Replace with your actual access token
	accessTokenStr := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	// accessToken := "Z0G0WQQT6T-101:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCb1dpTC1kYlZrbXZGVmJwQk85RVBwWnpOMEdIVFBxY09zNXEwOTRjamZQd3RKSU9IMDJMd3pLdFF0ZDA5X2RIaHF1SUEtUUFvTWpXT1dldk1kVi03R0RRdjIzckxoYzRsbFh6c1hTeTg5Vzk5ZWNJbz0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiIyZDVjZGZiMmZmMzU5NDg2YWFmNGQyOTViZWM0YjIzMTFlYzVmZTU0NDc1Mjc5MGUzZGZiMmFhNSIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiTiIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzUwODExNDAwLCJpYXQiOjE3NTA3Mzc2NjIsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc1MDczNzY2Miwic3ViIjoiYWNjZXNzX3Rva2VuIn0.QLPwwLxeXNuYEgRldhIBGGeZ4IaXXr9ogYqmZFRGgh0"

	// Create a FyersOrderSocket instance
	orderSocket := fyersws.NewFyersOrderSocket(
		accessTokenStr,   // Access token in the format "appid:accesstoken"
		false,            // Write to file - set to true if you want to save responses to a log file
		"",               // Log path - leave empty to auto-create logs in the current directory
		onOrderTrades,    // Callback function to handle trade events
		onOrderPositions, // Callback function to handle position events
		onOrderUpdates,   // Callback function to handle order events
		onOrderGeneral,   // Callback function to handle general events
		onOrderError,     // Callback function to handle WebSocket errors
		onOrderConnect,   // Callback function called when WebSocket connection is established
		onOrderClose,     // Callback function to handle WebSocket connection close events
		true,             // Enable auto-reconnection to WebSocket on disconnection
		5,                // Maximum number of reconnection attempts
	)

	// Establish a connection to the Fyers Order WebSocket
	err := orderSocket.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Order Socket: %v", err)
	}

	// fmt.Println("Order Socket connected successfully!")
	// fmt.Println("Receiving real-time order updates...")
	// fmt.Println("Press Ctrl+C to stop")

	// Subscribe to order updates
	// fmt.Printf("Subscribing to subscriptions: %v\n", orderSocketRequest.TradeOperations)
	for _, subscription := range orderSocketRequest.TradeOperations {
		orderSocket.Subscribe(subscription)
	}

	// Set up signal handling to keep the connection alive
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for interrupt signal
	<-sigChan
	fmt.Println("\nReceived interrupt signal, closing connection...")

	// Close the connection
	orderSocket.CloseConnection()
	fmt.Println("Order Socket connection closed")

	// Return connection status and subscription info
	return map[string]interface{}{
		"status":        "connected",
		"subscriptions": orderSocketRequest.TradeOperations,
		"message":       "Order Socket connected and subscribed successfully",
	}, nil
}

// Order Socket callback functions
func onOrderTrades(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	fmt.Println("Trade Response: ", string(jsonData))
}

func onOrderPositions(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	fmt.Println("Position Response: ", string(jsonData))
}

func onOrderUpdates(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	fmt.Println("Order Response: ", string(jsonData))
}

func onOrderGeneral(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	fmt.Println("General: ", string(jsonData))
}

func onOrderError(message map[string]interface{}) {
	if invalid, exists := message["invalid_symbols"]; exists {
		resp := map[string]interface{}{
			"type":            "sub",
			"code":            -300,
			"message":         "Please provide a valid symbol",
			"s":               "error",
			"invalid_symbols": invalid,
		}
		jsonResp, _ := json.Marshal(resp)
		fmt.Println("Error: ", jsonResp)
		return
	}
	jsonData, _ := json.Marshal(message)
	fmt.Println("Error: ", string(jsonData))
}

func onOrderClose(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	fmt.Println("Response: ", string(jsonData))
}

func onOrderConnect() {
	// fmt.Println("Order Socket - Connection established, subscribing to order updates...")

	// Subscribe to order updates
	// fmt.Println("Subscribing to OnOrders, OnTrades, OnPositions")
	// Note: This function is called during connection, so we can't access the orderSocket variable here
	// The subscription is handled in the main function
}
