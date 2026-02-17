package main

import (
	"fmt"

	fyersgosdk "fyers-go-sdk"
)

const (
	clientId     string = ""
	authToken    string = ""
	accessToken  string = ""
	appId        string = "Z0G0WQQT6T-101"
	appSecret    string = "TJHJFWBP0Q"
	redirectUrl  string = "https://trade.fyers.in/api-login/redirect-uri/index.html"
	refreshToken string = ""
	pin          string = "1223"
)

func main() {
	// Client is used only for GetLoginURL and GenerateAccessToken
	fyClient := fyersgosdk.SetClientData(clientId, appId, appSecret, redirectUrl, pin)
	fmt.Println(fyClient.GetLoginURL())

	// Get access token (use auth code from login redirect)
	response, err := fyClient.GenerateAccessToken(authToken, fyClient)
	if err != nil {
		fmt.Printf("Error generating access token: %v\n", err)
	}
	fmt.Println("access token:", response)

	fyClient.SetRefreshToken(refreshToken)
	response, err = fyClient.GenerateAccessTokenFromRefreshToken(fyClient)
	if err != nil {
		fmt.Printf("Error generating access token from refresh token: %v\n", err)
	}
	fmt.Println("access token from refresh token:", response)

	// FyersModel is used for all API calls (profile, funds, orders, etc.)
	// fyModel := fyersgosdk.NewFyersModel(appId, accessToken)

	// Get profile (returns raw JSON string)
	// response, err := fyModel.GetProfile()
	// if err != nil {
	// 	fmt.Printf("Error getting profile: %v\n", err)
	// 	return
	// }
	// fmt.Println("profile:", response)

	// // Get funds
	// response, err = fyModel.GetFunds()
	// if err != nil {
	// 	fmt.Printf("Error getting funds: %v\n", err)
	// 	return
	// }
	// fmt.Println("funds:", response)

	// // Get holdings (returns raw JSON string)
	// holdingsResp, err := fyModel.GetHoldings()
	// if err != nil {
	// 	fmt.Printf("Error getting holdings: %v\n", err)
	// 	return
	// }
	// fmt.Println("holdings:", holdingsResp)

	// Logout
	// response, err := fyModel.Logout()
	// if err != nil {
	// 	fmt.Printf("Error logging out: %v", err)
	// } else {
	// 	fmt.Println("logout: ", response)
	// }

	// Get order book
	// response, err := fyModel.GetOrderBook()
	// if err != nil {
	// 	fmt.Printf("Error getting order book: %v", err)
	// } else {
	// 	fmt.Println("order book: ", response)
	// }

	// Get order book by tag
	// response, err := fyModel.GetOrderBookByTag("1:Ordertag")
	// if err != nil {
	// 	fmt.Printf("Error getting order book by tag: %v", err)
	// } else {
	// 	fmt.Println("order book by tag: ", response)
	// }

	// Get order by id
	// response, err := fyModel.GetOrderById("25070100257946")
	// if err != nil {
	// 	fmt.Printf("Error getting order by id: %v", err)
	// } else {
	// 	fmt.Println("order by id: ", response)
	// }

	// Get positions
	// response, err := fyModel.GetPositions()
	// if err != nil {
	// 	fmt.Printf("Error getting positions: %v", err)
	// } else {
	// 	fmt.Println("positions: ", response)
	// }

	// Get trade book
	// response, err := fyModel.GetTradeBook()
	// if err != nil {
	// 	fmt.Printf("Error getting trade book: %v", err)
	// } else {
	// 	fmt.Println("trade book: ", response)
	// }

	// Get trade book by tag
	// response, err := fyModel.GetTradeBookByTag("2:Exit")
	// if err != nil {
	// 	fmt.Printf("Error getting trade book by tag: %v", err)
	// } else {
	// 	fmt.Println("trade book by tag: ", response)
	// }

	// Single Order Action
	// response, err := fyModel.SingleOrderAction(fyersgosdk.OrderRequest{
	// 	Symbol:       "NSE:IDEA-EQ",
	// 	Qty:          1,
	// 	Type:         1,
	// 	Side:         1,
	// 	ProductType:  "CNC",
	// 	LimitPrice:   100,
	// 	StopPrice:    100,
	// 	Validity:     "DAY",
	// 	DisclosedQty: 1,
	// 	OfflineOrder: false,
	// 	StopLoss:     100,
	// 	TakeProfit:   100,
	// 	OrderTag:     "TESTEST",
	// })
	// if err != nil {
	// 	fmt.Printf("Error single order action: %v", err)
	// } else {
	// 	fmt.Println("single order action: ", response)
	// }

	// Multi Order Action
	// response, err := fyModel.MultiOrderAction([]fyersgosdk.OrderRequest{
	// 	{Symbol: "NSE:SBIN-EQ", Qty: 1, Type: 1, Side: 1, ProductType: "CNC"},
	// 	{Symbol: "NSE:ABB-EQ", Qty: 1, Type: 1, Side: 1, ProductType: "CNC"},
	// })
	// if err != nil {
	// 	fmt.Printf("Error multi order action: %v", err)
	// } else {
	// 	fmt.Println("multi order action: ", response)
	// }

	// Multi Leg Order Action
	// response, err := fyModel.MultiLegOrderAction([]fyersgosdk.MultiLegOrderRequest{
	// 	{
	// 		OrderTag:     "tag1",
	// 		ProductType:  "MARGIN",
	// 		OfflineOrder: false,
	// 		OrderType:    "3L",
	// 		Validity:     "IOC",
	// 		Legs: fyersgosdk.Leg{
	// 			Leg1: fyersgosdk.LegBody{
	// 				Symbol:     "NSE:SBIN25JULFUT",
	// 				Qty:        750,
	// 				Side:       1,
	// 				Type:       1,
	// 				LimitPrice: 800,
	// 			},
	// 			Leg2: fyersgosdk.LegBody{
	// 				Symbol:     "NSE:SBIN25AUGFUT",
	// 				Qty:        750,
	// 				Side:       1,
	// 				Type:       1,
	// 				LimitPrice: 800,
	// 			},
	// 			Leg3: fyersgosdk.LegBody{
	// 				Symbol:     "NSE:SBIN25SEPFUT",
	// 				Qty:        750,
	// 				Side:       1,
	// 				Type:       1,
	// 				LimitPrice: 3,
	// 			},
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Printf("Error multi leg order action: %v", err)
	// } else {
	// 	fmt.Println("multi leg order action: ", response)
	// }

	// GTT Order Action
	// response, err := fyModel.GTTSingleOrderAction(fyersgosdk.GTTOrderRequest{
	// 	Side:        1,
	// 	Symbol:      "NSE:SBIN-EQ",
	// 	ProductType: "CNC",
	// 	OrderInfo: fyersgosdk.OrderInfo{
	// 		Leg1: fyersgosdk.Leg1{
	// 			Price:        100,
	// 			TriggerPrice: 100,
	// 			Qty:          1,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Printf("Error GTT order action: %v", err)
	// } else {
	// 	fmt.Println("GTT order action: ", response)
	// }

	// GTT Multi Order Action
	// response, err := fyModel.GTTMultiOrderAction([]fyersgosdk.GTTOrderRequest{
	// 	{
	// 		Side:        1,
	// 		Symbol:     "NSE:SBIN-EQ",
	// 		ProductType: "CNC",
	// 		OrderInfo: fyersgosdk.OrderInfo{
	// 			Leg1: fyersgosdk.Leg1{Price: 10000, TriggerPrice: 10000, Qty: 1},
	// 			Leg2: fyersgosdk.Leg1{Price: 990, TriggerPrice: 990, Qty: 3},
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Printf("Error GTT multi order action: %v", err)
	// } else {
	// 	fmt.Println("GTT multi order action: ", response)
	// }

	// GTT Modify Order Action
	// response, err := fyModel.ModifyGTTOrder(fyersgosdk.ModifyGTTOrderRequest{
	// 	Id: "23030900015105",
	// 	OrderInfo: fyersgosdk.OrderInfo{
	// 		Leg1: fyersgosdk.Leg1{Price: 1010, TriggerPrice: 1010, Qty: 5},
	// 		Leg2: fyersgosdk.Leg1{Price: 1010, TriggerPrice: 1010, Qty: 5},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Printf("Error GTT modify order action: %v", err)
	// } else {
	// 	fmt.Println("GTT modify order action: ", response)
	// }

	// GTT Cancel Order Action
	// response, err := fyModel.CancelGTTOrder("23030900015105")
	// if err != nil {
	// 	fmt.Printf("Error GTT cancel order action: %v", err)
	// } else {
	// 	fmt.Println("GTT cancel order action: ", response)
	// }

	// GTT Order Book
	// response, err := fyModel.GetGTTOrderBook()
	// if err != nil {
	// 	fmt.Printf("Error getting GTT order book: %v", err)
	// } else {
	// 	fmt.Println("GTT order book: ", response)
	// }

	// Modify Order
	// response, err := fyModel.ModifyOrder(fyersgosdk.ModifyOrderRequest{
	// 	Id:         "23030900015105",
	// 	Qty:        10,
	// 	Type:       1,
	// 	Side:       1,
	// 	LimitPrice: 100,
	// })
	// if err != nil {
	// 	fmt.Printf("Error modify order action: %v", err)
	// } else {
	// 	fmt.Println("modify order action: ", response)
	// }

	// Get history
	// response, err := fyModel.GetHistory(fyersgosdk.HistoryRequest{
	// 	Symbol:     "NSE:SBIN-EQ",
	// 	Resolution: "30",
	// 	DateFormat: "1",
	// 	RangeFrom:  "2021-01-01",
	// 	RangeTo:    "2021-01-02",
	// })
	// if err != nil {
	// 	fmt.Printf("Error getting history: %v", err)
	// } else {
	// 	fmt.Println("history: ", response)
	// }

	// Get stock quotes
	// response, err := fyModel.GetStockQuotes("NSE:SBIN-EQ")
	// if err != nil {
	// 	fmt.Printf("Error getting stock quotes: %v", err)
	// } else {
	// 	fmt.Println("stock quotes: ", response)
	// }

	// Get market depth
	// response, err := fyModel.GetMarketDepth(fyersgosdk.MarketDepthRequest{
	// 	Symbol: "NSE:SBIN-EQ",
	// 	OHLCV:  "1",
	// })
	// if err != nil {
	// 	fmt.Printf("Error getting market depth: %v", err)
	// } else {
	// 	fmt.Println("market depth: ", response)
	// }

	// Get option chain
	// response, err := fyModel.GetOptionChain(fyersgosdk.OptionChainRequest{
	// 	Symbol:      "NSE:SBIN-EQ",
	// 	StrikeCount: 1,
	// })
	// if err != nil {
	// 	fmt.Printf("Error getting option chain: %v", err)
	// } else {
	// 	fmt.Println("option chain: ", response)
	// }

	// Get alerts (response is raw JSON; parse with json.Unmarshal if you need alert IDs etc.)
	// response, err := fyModel.GetAlerts()
	// if err != nil {
	// 	fmt.Printf("Error getting alerts: %v\n", err)
	// } else {
	// 	fmt.Println("Get Alerts Response: ", response)
	// }

	// Example alert ID (use an ID from GetAlerts response after parsing)
	// exampleAlertId := ""

	// Create alert
	// alertReq := fyersgosdk.AlertRequest{
	// 	Symbol:         "NSE:SBIN-EQ",
	// 	Name:           "NSE:SBIN-EQ alert",
	// 	Agent:          "fyers-api",
	// 	AlertType:      1,
	// 	ComparisonType: "LTP",
	// 	Condition:      "GT",
	// 	Value:          600.0,
	// }
	// response, err := fyModel.CreateAlert(alertReq)
	// if err != nil {
	// 	fmt.Printf("Error creating alert: %v\n", err)
	// } else {
	// 	fmt.Println("Create Alert Response: ", response)
	// }

	// Toggle alert
	// if exampleAlertId != "" {
	// 	response, err := fyModel.ToggleAlert(exampleAlertId)
	// 	if err != nil {
	// 		fmt.Printf("Error toggling alert: %v\n", err)
	// 	} else {
	// 		fmt.Println("Toggle Alert Response: ", response)
	// 	}
	// }

	// Delete alert
	// if exampleAlertId != "" {
	// 	response, err := fyModel.DeleteAlert(exampleAlertId)
	// 	if err != nil {
	// 		fmt.Printf("Error deleting alert: %v\n", err)
	// 	} else {
	// 		fmt.Println("Delete Alert Response: ", response)
	// 	}
	// }

	// Update alert
	// alertReq := fyersgosdk.AlertRequest{
	// 	Symbol:         "NSE:SBIN-EQ",
	// 	Name:           "NSE:NIFTY50-INDEX",
	// 	Agent:          "fyers-api",
	// 	AlertType:      1,
	// 	ComparisonType: "LTP",
	// 	Condition:      "GT",
	// 	Value:          25423.49,
	// }
	// response, err := fyModel.UpdateAlert("6137795", alertReq)
	// if err != nil {
	// 	fmt.Printf("Error updating alert: %v\n", err)
	// } else {
	// 	fmt.Println("Update Alert Response: ", response)
	// }

	// WEBSOCKET EXAMPLES (these may still use fyClient if the SDK expects Client for websocket auth)
	// Data Socket (Market Data WebSocket)
	// wsResponse, wsErr := fyersgosdk.DataSocket(fyModel, fyersgosdk.DataSocketRequest{
	// 	Symbols:  []string{"MCX:SILVER26MARFUT"},
	// 	DataType: "SymbolUpdate",
	// 	Mode:     false,
	// })
	// if wsErr != nil {
	// 	fmt.Printf("Data Socket Error: %v\n", wsErr)
	// } else {
	// 	fmt.Printf("Data Socket Response: %+v\n", wsResponse)
	// }

	// Order Socket (Order Updates WebSocket)
	// wsResponse2, wsErr := fyersgosdk.OrderSocket(fyClient, fyersgosdk.OrderSocketRequest{
	// 	TradeOperations: []string{"OnOrders", "OnTrades", "OnPositions"},
	// })
	// if wsErr != nil {
	// 	fmt.Printf("Order Socket Error: %v\n", wsErr)
	// } else {
	// 	fmt.Printf("Order Socket Response: %+v\n", wsResponse2)
	// }
}
