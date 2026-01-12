package main

import (
	"fmt"
	fyersgosdk "fyers-go-sdk"
)

const (
	clientId    string = "YK04391"
	authToken   string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOiJNMFI0V1cxUFlVIiwidXVpZCI6ImI3NzE4MmMyMTdhODQ4M2JiZjk3NzJhYTQ0Yjk1YjcxIiwiaXBBZGRyIjoiIiwibm9uY2UiOiIiLCJzY29wZSI6IiIsImRpc3BsYXlfbmFtZSI6IllLMDQzOTEiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiI4N2Q3YTBmZTNmMGVjZDdhNzAwNzAyOTY1YTNkZWU3ZmYwNDA0OTBmOTQ3NzI0MWU3YWY5ZTljYSIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiTiIsImF1ZCI6IltcImQ6MVwiLFwiZDoyXCIsXCJ4OjBcIixcIng6MVwiLFwieDoyXCJdIiwiZXhwIjoxNzUxMzgzMDQxLCJpYXQiOjE3NTEzNTMwNDEsImlzcyI6ImFwaS5sb2dpbi5meWVycy5pbiIsIm5iZiI6MTc1MTM1MzA0MSwic3ViIjoiYXV0aF9jb2RlIn0.9hSzBigEZ3mCr7FBNxcql1iqe2mZJWnkossL5z0nca8"
	accessToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcFpLYU1iRlJBTWVXZnVuTExRa2ttd2tFM3ZIdlJRbzdxUVBmTjBmRm5uVXZXRDRndVpnekw4X2NNSTZBRWtCUktPemhqSTVqQlBFYy05SERIUEZGUW1tMER6NE9YOUZCZmRwLXQ1MllJQmZfWXpqQT0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJjMDFlNDQ3ZDc1YzNjZDIwYTQ3YTQ0ZjA3NjJmNTM3Mjc5YjJmZWY4NGY4NzY4MTk4NTQwOTdkMiIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzY4MjY0MjAwLCJpYXQiOjE3NjgyMDM5MTYsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc2ODIwMzkxNiwic3ViIjoiYWNjZXNzX3Rva2VuIn0.GRxdraCQNX0vYBYlTQs3YkRUINit-rZvuJqbHAUSM3E"
	appId       string = "M0R4WW1PYU-100"
	appSecret   string = "XKCP7PAISD"
	redirectUrl string = "https://trade.fyers.in/api-login/redirect-uri/index.html"
)

func main() {
	// Create a new instance
	fyClient := fyersgosdk.SetClientData(clientId, appId, appSecret, redirectUrl)
	// fmt.Println(fyClient.GetLoginURL())

	// Get access token
	// response, accessToken, err := fyClient.GenerateAccessToken(authToken,fyClient)
	// if err != nil {
	// 	fmt.Printf("Error getting access token: %v", err)
	// }
	// fmt.Println("profile: ", response)
	// fmt.Println("profile: ", accessToken.AccessToken)

	// Set access token
	fyClient.SetAccessToken(accessToken)

	// Get profile
	// response, profile, err := fyClient.GetProfile(fyClient)
	// if err != nil {
	// 	fmt.Printf("Error getting profile: %v", err)
	// }
	// fmt.Println("profile: ", response)
	// fmt.Println("profile: ", profile.DdpiEnabled)

	// Get funds
	// response, funds, err := fyClient.GetFunds(fyClient)
	// if err != nil {
	// 	fmt.Printf("Error getting funds: %v", err)
	// }
	// fmt.Println("funds: ", response)
	// fmt.Println("fundsFundLimit : ", funds.FundLimit[0].Title)

	// Get holdings
	// response, holdings, err := fyClient.GetHoldings(fyClient)
	// if err != nil {
	// 	fmt.Printf("Error getting holdings: %v", err)
	// }
	// fmt.Println("holdings: ", response)
	// fmt.Println("holdings: ", holdings.Overall.CountTotal)

	// Logout
	// response, logout, err := fyClient.Logout(fyClient)
	// if err != nil {
	// 	fmt.Printf("Error logging out: %v", err)
	// }
	// fmt.Println("logout: ", response)
	// fmt.Println("logout: ", logout.Message)

	// Get order book
	// response, orderBook, err := fyClient.GetOrderBook(fyClient)
	// if err != nil {
	// 	fmt.Printf("Error getting order book: %v", err)
	// }
	// fmt.Println("order book: ", response)
	// fmt.Println("order book: ", orderBook.OrderBook[1].ExSym)

	// // Get order book by tag
	// response, orderBookByTag, err := fyClient.GetOrderBookByTag(fyClient, "1:Ordertag")
	// if err != nil {
	// 	fmt.Printf("Error getting order book by tag: %v", err)
	// }
	// fmt.Println("order book by tag: ", response)
	// fmt.Println("order book by tag: ", orderBookByTag.OrderBook[0])

	// Get order by id
	// response, orderById, err := fyClient.GetOrderById(fyClient, "25070100257946")										 //check again
	// if err != nil {
	// 	fmt.Printf("Error getting order by id: %v", err)
	// }
	// fmt.Println("order by id: ", response)
	// fmt.Println("order by id: ", orderById.OrderBook[0].Description)

	// // Get positions
	// response,positions, err := fyClient.GetPositions(fyClient)
	// if err != nil {
	// 	fmt.Printf("Error getting positions: %v", err)
	// }
	// fmt.Println("positions: ", response)
	// fmt.Println("positions: ", positions.APIResponse.Code)

	// // Get trade book
	// response, tradeBook, err := fyClient.GetTradeBook(fyClient)
	// if err != nil {
	// 	fmt.Printf("Error getting trade book: %v", err)
	// }
	// fmt.Println("trade book: ", response)
	// fmt.Println("trade book: ", tradeBook.TradeBook[0])

	// // Get trade book by tag
	// response, tradeBookByTag, err := fyClient.GetTradeBookByTag(fyClient, "2:Exit")
	// if err != nil {
	// 	fmt.Printf("Error getting trade book by tag: %v", err)
	// }
	// fmt.Println("trade book by tag: ", response)
	// fmt.Println("trade book by tag: ", tradeBookByTag.TradeBook[0].ClientId)

	// // Single Order Action
	// response, singleOrderAction, err := fyClient.SingleOrderAction(fyClient, fyersgosdk.OrderRequest{
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
	// 	fmt.Printf("Error getting single order action: %v", err)
	// }
	// fmt.Println("single order action: ", response)
	// fmt.Println("single order action: ", singleOrderAction.Message)

	// // Multi Order Action
	// response, multiOrderAction, err := fyClient.MultiOrderAction(fyClient, []fyersgosdk.OrderRequest{{
	// 	Symbol:      "NSE:SBIN-EQ",
	// 	Qty:         1,
	// 	Type:        1,
	// 	Side:        1,
	// 	ProductType: "CNC",
	// }, {
	// 	Symbol:      "NSE:ABB-EQ",
	// 	Qty:         1,
	// 	Type:        1,
	// 	Side:        1,
	// 	ProductType: "CNC",
	// }})
	// if err != nil {
	// 	fmt.Printf("Error getting multi order action: %v", err)
	// }
	// fmt.Println("multi order action: ", response)
	// fmt.Println("multi order action: ", multiOrderAction.Message)

	// // Multi Leg Order Action
	// response, multiLegOrderAction, err := fyClient.MultiLegOrderAction(fyClient, []fyersgosdk.MultiLegOrderRequest{
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
	// 	}})
	// if err != nil {
	// 	fmt.Printf("Error getting multi leg order action: %v", err)
	// }
	// fmt.Println("multi leg order action: ", response)
	// fmt.Println("multi leg order action: ", multiLegOrderAction.Message)

	// // GTT Order Action
	// response, gttOrderAction, err := fyClient.GTTSingleOrderAction(fyClient, fyersgosdk.GTTOrderRequest{
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
	// 	fmt.Printf("Error getting GTT order action: %v", err)
	// }
	// fmt.Println("GTT order action: ", response)
	// fmt.Println("GTT order action: ", gttOrderAction.Message)

	// GTT Multi Order Action
	// response, gttMultiOrderAction, err := fyClient.GTTMultiOrderAction(fyClient, []fyersgosdk.GTTOrderRequest{
	// 	{
	// 		Side: 1,
	// 		Symbol: "NSE:SBIN-EQ",
	// 		ProductType: "CNC",
	// 		OrderInfo: fyersgosdk.OrderInfo{
	// 			Leg1: fyersgosdk.Leg1{
	// 				Price: 10000,
	// 				TriggerPrice: 10000,
	// 				Qty: 1,
	// 			},
	// 			Leg2: fyersgosdk.Leg1{
	// 				Price: 990,
	// 				TriggerPrice: 990,
	// 				Qty: 3,
	// 			},
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Printf("Error getting GTT multi order action: %v", err)
	// }
	// fmt.Println("GTT multi order action: ", response)
	// fmt.Println("GTT multi order action: ", gttMultiOrderAction.Message)

	// GTT Modify Order Action
	// response, gttModifyOrderAction, err := fyClient.ModifyGTTOrder(fyClient, fyersgosdk.ModifyGTTOrderRequest{
	// 	Id: "23030900015105",
	// 	OrderInfo: fyersgosdk.OrderInfo{
	// 		Leg1: fyersgosdk.Leg1{
	// 			Price:        1010,
	// 			TriggerPrice: 1010,
	// 			Qty:          5,
	// 		},
	// 		Leg2: fyersgosdk.Leg1{
	// 			Price:        1010,
	// 			TriggerPrice: 1010,
	// 			Qty:          5,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Printf("Error getting GTT modify order action: %v", err)
	// }
	// fmt.Println("GTT modify order action: ", response)
	// fmt.Println("GTT modify order action: ", gttModifyOrderAction.Message)

	// // GTT Cancel Order Action
	// response, gttCancelOrderAction, err := fyClient.CancelGTTOrder(fyClient, "23030900015105")
	// if err != nil {
	// 	fmt.Printf("Error getting GTT cancel order action: %v", err)
	// }
	// fmt.Println("GTT cancel order action: ", response)
	// fmt.Println("GTT cancel order action: ", gttCancelOrderAction.Message)

	// GTT Order Book
	// response, gttOrderBook, err := fyClient.GetGTTOrderBook(fyClient)
	// if err != nil {
	// 	fmt.Printf("Error getting GTT order book: %v", err)
	// }
	// fmt.Println("GTT order book: ", response)
	// fmt.Println("GTT order book: ", gttOrderBook.ClientId)

	// Modify Orders
	// response, modifyOrderAction, err := fyClient.ModifyOrder(fyClient, fyersgosdk.ModifyOrderRequest{
	// 	Id: "23030900015105",
	// 	Qty: 10,
	// 	Type: 1,
	// 	Side: 1,
	// 	LimitPrice: 100,
	// })
	// if err != nil {
	// 	fmt.Printf("Error getting modify order action: %v", err)
	// }
	// fmt.Println("modify order action: ", response)
	// fmt.Println("modify order action: ", modifyOrderAction.Message)

	// Get history
	// response, history, err := fyClient.GetHistory(fyClient, fyersgosdk.HistoryRequest{
	// 	Symbol:     "NSE:SBIN-EQ",
	// 	Resolution: "30",
	// 	DateFormat: "1",
	// 	RangeFrom:  "2021-01-01",
	// 	RangeTo:    "2021-01-02",
	// })
	// if err != nil {
	// 	fmt.Printf("Error getting history: %v", err)
	// }
	// fmt.Println("history: ", response)
	// fmt.Println("history: ", history.Candles)

	// // Get stock quotes
	// response, stockQuotes , err := fyClient.GetStockQuotes(fyClient, "NSE:SBIN-EQ")
	// if err != nil {
	// 	fmt.Printf("Error getting stock quotes: %v", err)
	// }
	// fmt.Println("stock quotes: ", response)
	// fmt.Println("stock quotes: ", stockQuotes.Data[0].V.Symbol)

	// // Get market depth
	// response, marketDepth, err := fyClient.GetMarketDepth(fyClient, fyersgosdk.MarketDepthRequest{
	// 	Symbol: "NSE:SBIN-EQ",
	// 	OHLCV:  "1",
	// })
	// if err != nil {
	// 	fmt.Printf("Error getting market depth: %v", err)
	// }
	// fmt.Println("market depth: ", response)
	// fmt.Println("market depth: ", marketDepth.Data["NSE:SBIN-EQ"].TotalBuyQty)

	// // Get option chain
	// response, optionChain, err := fyClient.GetOptionChain(fyClient, fyersgosdk.OptionChainRequest{
	// 	Symbol:      "NSE:SBIN-EQ",
	// 	StrikeCount: 1,
	// })
	// if err != nil {
	// 	fmt.Printf("Error getting option chain: %v", err)
	// }
	// fmt.Println("option chain: ", response)
	// fmt.Println("option chain: ", optionChain.Data.OptionsChain)

	// Get alerts
	// response, priceAlert, err := fyClient.GetAlerts(fyClient)
	// if err != nil {
	// 	fmt.Printf("Error getting alerts: %v\n", err)
	// } else {
	// 	fmt.Println("Get Alerts Response: ", response)
	// 	if len(priceAlert.Data) > 0 {
	// 		for id, item := range priceAlert.Data {
	// 			fmt.Printf("Alert ID: %s, Symbol: %s, Value: %f\n", id, item.Symbol, item.Alert.Value)
	// 			// We can use this alertId for other operations
	// 		}
	// 	}
	// }

	// // Example alert ID (found from GetAlerts)
	// exampleAlertId := ""
	// for id := range priceAlert.Data {
	// 	exampleAlertId = id
	// 	break
	// }
	// fmt.Println("Example Alert ID available for use: ", exampleAlertId)

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
	// response, createAlert, err := fyClient.CreateAlert(fyClient, alertReq)
	// if err != nil {
	// 	fmt.Printf("Error creating alert: %v\n", err)
	// } else {
	// 	fmt.Println("Create Alert Response: ", response)
	// }

	// Toggle alert
	// if exampleAlertId != "" {
	// 	response, toggleAlert, err := fyClient.ToggleAlert(fyClient, exampleAlertId)
	// 	if err != nil {
	// 		fmt.Printf("Error toggling alert: %v\n", err)
	// 	} else {
	// 		fmt.Println("Toggle Alert Response: ", response)
	// 	}
	// }

	// Delete alert
	// if exampleAlertId != "" {
	// 	response, deleteAlert, err := fyClient.DeleteAlert(fyClient, exampleAlertId)
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
	// response, updateAlert, err := fyClient.UpdateAlert(fyClient, "6137795", alertReq)
	// if err != nil {
	// 	fmt.Printf("Error updating alert: %v\n", err)
	// } else {
	// 	fmt.Println("Update Alert Response: ", response)
	// 	fmt.Println("Update Alert Message: ", updateAlert.Message)
	// }


	// WEBSOCKET EXAMPLES
	// Data Socket (Market Data WebSocket)
	// wsResponse, wsErr := fyersgosdk.DataSocket(fyClient, fyersgosdk.DataSocketRequest{
	// 	Symbols: []string{"NSE:NH-EQ"},
	// 	DataType:    "SymbolUpdate",
	// })
	// if wsErr != nil {
	// 	fmt.Printf("Data Socket Error: %v\n", wsErr)
	// } else {
	// 	fmt.Printf("Data Socket Response: %+v\n", wsResponse)
	// }

	// Order Socket (Order Updates WebSocket)
	// wsResponse2, wsErr := fyersgosdk.OrderSocket(fyClient,fyersgosdk.OrderSocketRequest{
	// 	TradeOperations: []string{"OnOrders", "OnTrades", "OnPositions"},
	// })
	// if wsErr != nil {
	// 	fmt.Printf("Order Socket Error: %v\n", wsErr)
	// } else {
	// 	fmt.Printf("Order Socket Response: %+v\n", wsResponse2)
	// }
}
