package fyersgosdk

import (
	"encoding/json"
	"fmt"
	fyersws "fyers-go-sdk/websocket"
	"os"
	"os/signal"
	"syscall"
)

// Data Socket Example
func DataSocket(fyModel *FyersModel, webSocketRequest DataSocketRequest) (map[string]interface{}, error) {
	// Replace with your actual access token
	accessTokenStr := fmt.Sprintf("%s:%s", fyModel.appId, fyModel.accessToken)
	// accessToken := "Z0G0WQQT6T-101:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCb1dpTC1kYlZrbXZGVmJwQk85RVBwWnpOMEdIVFBxY09zNXEwOTRjamZQd3RKSU9IMDJMd3pLdFF0ZDA5X2RIaHF1SUEtUUFvTWpXT1dldk1kVi03R0RRdjIzckxoYzRsbFh6c1hTeTg5Vzk5ZWNJbz0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiIyZDVjZGZiMmZmMzU5NDg2YWFmNGQyOTViZWM0YjIzMTFlYzVmZTU0NDc1Mjc5MGUzZGZiMmFhNSIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiTiIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzUwODExNDAwLCJpYXQiOjE3NTA3Mzc2NjIsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc1MDczNzY2Miwic3ViIjoiYWNjZXNzX3Rva2VuIn0.QLPwwLxeXNuYEgRldhIBGGeZ4IaXXr9ogYqmZFRGgh0"

	var dataSocket *fyersws.FyersDataSocket
	// onDataConnect is called on every connection (including after reconnect). Re-subscribe here so feed data resumes after reconnect.
	onDataConnect := func() {
		dataSocket.Subscribe(webSocketRequest.Symbols, webSocketRequest.DataType)
	}

	// Create a FyersDataSocket instance
	dataSocket = fyersws.NewFyersDataSocket(
		accessTokenStr,            // Access token in the format "appid:accesstoken"
		"",                        // Log path - leave empty to auto-create logs in the current directory
		webSocketRequest.LiteMode, // Lite mode disabled. Set to true if you want a lite response
		false,                     // Save response in a log file instead of printing it
		true,                      // Enable auto-reconnection to WebSocket on disconnection
		50,                        // reconnectRetry: max reconnect attempts (same as Python default; cap 50)
		onDataConnect,             // Callback: subscribe on every connect (first + after reconnect)
		onDataClose,               // Callback function to handle WebSocket connection close events
		onDataError,               // Callback function to handle WebSocket errors
		onDataMessage,             // Callback function to handle incoming messages from the WebSocket
	)

	// Establish a connection to the Fyers Data WebSocket
	err := dataSocket.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Data Socket: %v", err)
	}

	// Set up signal handling to keep the connection alive
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for interrupt signal
	<-sigChan
	fmt.Println("\nReceived interrupt signal, closing connection...")

	// Close the connection
	dataSocket.CloseConnection()
	fmt.Println("Data Socket connection closed")

	// Return connection status and subscription info
	return map[string]interface{}{
		"status":       "connected",
		"symbols":      webSocketRequest.Symbols,
		"subscription": webSocketRequest.DataType,
		"message":      "Data Socket connected and subscribed successfully",
	}, nil
}

// Data Socket callback functions
func onDataMessage(message map[string]interface{}) {
	// Marshal in data_val order (ltp, vol_traded_today, ..., type, symbol, ch, chp) to match Python SDK
	jsonData, _ := fyersws.MarshalDataResponseInOrder(message)
	fmt.Printf("Response: %s\n", string(jsonData))
}

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
	fmt.Printf("Trade Response: %s\n", string(jsonData))
}

func onOrderPositions(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	fmt.Printf("Position Response: %s\n", string(jsonData))
}

func onOrderUpdates(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	fmt.Printf("Order Response: %s\n", string(jsonData))
}

func onOrderGeneral(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	fmt.Printf("General : %s\n", string(jsonData))
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
		fmt.Printf("Error: %s\n", jsonResp)
		return
	}
	jsonData, _ := json.Marshal(message)
	fmt.Printf("Error: %s\n", string(jsonData))
}

func onOrderClose(message map[string]interface{}) {
	jsonData, _ := json.Marshal(message)
	fmt.Printf("Response: %s\n", string(jsonData))
}

func onOrderConnect() {
	// fmt.Println("Order Socket - Connection established, subscribing to order updates...")

	// Subscribe to order updates
	// fmt.Println("Subscribing to OnOrders, OnTrades, OnPositions")
	// Note: This function is called during connection, so we can't access the orderSocket variable here
	// The subscription is handled in the main function
}
