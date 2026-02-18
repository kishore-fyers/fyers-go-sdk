package main

// import (
// 	"fmt"
// 	fyersgosdk "fyers-go-sdk"
// )

// Get Auth Code URL
// func main() {
// 	appId := "M0R4WW1PYU-100"
// 	appSecret := "XKCP7PAISD"
// 	redirectUrl := "https://trade.fyers.in/api-login/redirect-uri/index.html"

// 	fyClient := fyersgosdk.SetClientData(appId, appSecret, redirectUrl)
// 	fmt.Println(fyClient.GetLoginURL())
// }

// Generate Access Token
// func main() {
// 	appId := "M0R4WW1PYU-100"
// 	appSecret := "XKCP7PAISD"
// 	redirectUrl := "https://trade.fyers.in/api-login/redirect-uri/index.html"
// 	authToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOiJNMFI0V1cxUFlVIiwidXVpZCI6IjRkYWU5MjQ0NmY4MDRlMWM5Y2RhNjE5NmU0MmY0MjE0IiwiaXBBZGRyIjoiIiwibm9uY2UiOiIiLCJzY29wZSI6IiIsImRpc3BsYXlfbmFtZSI6IllLMDQzOTEiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImF1ZCI6IltcImQ6MVwiLFwiZDoyXCIsXCJ4OjBcIixcIng6MVwiLFwieDoyXCJdIiwiZXhwIjoxNzcxNDIyMDMwLCJpYXQiOjE3NzEzOTIwMzAsImlzcyI6ImFwaS5sb2dpbi5meWVycy5pbiIsIm5iZiI6MTc3MTM5MjAzMCwic3ViIjoiYXV0aF9jb2RlIn0.rnEMaa8MigGEs_LSwEGoc-y0UbqjVRIwahvVccssMwU"

// 		fyClient := fyersgosdk.SetClientData(appId, appSecret, redirectUrl)
// 		response, err := fyClient.GenerateAccessToken(authToken, fyClient)
// 		if err != nil {
// 			fmt.Printf("Error generating access token: %v\n", err)
// 		}
// 		fmt.Println("access token:", response)
// 	}

// Generate Access Token From Refresh Token
// func main() {
// 	appId := "M0R4WW1PYU-100"
// 	appSecret := "XKCP7PAISD"
// 	redirectUrl := "https://trade.fyers.in/api-login/redirect-uri/index.html"
// 	refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiZDoxIiwiZDoyIiwieDowIiwieDoxIiwieDoyIl0sImF0X2hhc2giOiJnQUFBQUFCcGxVNE1xZUpQUFdYNWhWQzcyNkZQVTVXTHpjcFRXMjQ3OGFDdGJMemwybE42TmkzUEZpU0xKMHVDRldPVE9Fc3JIbjVlbWxVdVNiQ2F2UXlySTh0LXozeFdaWFo4MFRXZWFxb0JKeVdtbFFfYVNacz0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMCwiZXhwIjoxNzcyNjcwNjAwLCJpYXQiOjE3NzEzOTI1MjQsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5MjUyNCwic3ViIjoicmVmcmVzaF90b2tlbiJ9.ogZNRYM6lWQ4RRpVeOMuzpwmbAK9MLhPB89UBaFtxCY"
// 	pin := "0000"

// 	fyClient := fyersgosdk.SetClientData(appId, appSecret, redirectUrl)
// 	response, err := fyClient.GenerateAccessTokenFromRefreshToken(refreshToken, pin, fyClient)
// 	if err != nil {
// 		fmt.Printf("Error generating access token from refresh token: %v\n", err)
// 	}
// 	fmt.Println("access token:", response)
// }

// Get Profile
// func main() {
// 	appId := "M0R4WW1PYU-100"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiZDoxIiwiZDoyIiwieDowIiwieDoxIiwieDoyIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXTm5QcEV5WG9xZDJTcFZoQkFLcVhRaTRBa1BtX1IzWkpWZnVwRnNxQlFyNVQxbnd5NC00eW9IOFJwVW5DV21xOVlCb3FrTVRIWDZrdUlhOHN1TzUzdVhZSVJ0VHZMdTVvU0h4YVE5cG1iUllaUT0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMCwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTc5OTEsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5Nzk5MSwic3ViIjoiYWNjZXNzX3Rva2VuIn0.TKbd3hc9vRH-OwotVg8FzvxodA4-7MjIL80p7MBzfZI"

// 		fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 		response, err := fyModel.GetProfile()
// 		if err != nil {
// 			fmt.Printf("Error getting profile: %v\n", err)
// 			return
// 		}
// 		fmt.Println("profile:", response)
// 	}

// Get Funds
// func main() {
// 	appId := "M0R4WW1PYU-100"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiZDoxIiwiZDoyIiwieDowIiwieDoxIiwieDoyIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXTm5QcEV5WG9xZDJTcFZoQkFLcVhRaTRBa1BtX1IzWkpWZnVwRnNxQlFyNVQxbnd5NC00eW9IOFJwVW5DV21xOVlCb3FrTVRIWDZrdUlhOHN1TzUzdVhZSVJ0VHZMdTVvU0h4YVE5cG1iUllaUT0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMCwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTc5OTEsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5Nzk5MSwic3ViIjoiYWNjZXNzX3Rva2VuIn0.TKbd3hc9vRH-OwotVg8FzvxodA4-7MjIL80p7MBzfZI"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.GetFunds()
// 	if err != nil {
// 		fmt.Printf("Error getting funds: %v\n", err)
// 		return
// 	}
// 	fmt.Println("funds:", response)
// }

// Get Holdings
// func main() {
// 	appId := "M0R4WW1PYU-100"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiZDoxIiwiZDoyIiwieDowIiwieDoxIiwieDoyIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXTm5QcEV5WG9xZDJTcFZoQkFLcVhRaTRBa1BtX1IzWkpWZnVwRnNxQlFyNVQxbnd5NC00eW9IOFJwVW5DV21xOVlCb3FrTVRIWDZrdUlhOHN1TzUzdVhZSVJ0VHZMdTVvU0h4YVE5cG1iUllaUT0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMCwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTc5OTEsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5Nzk5MSwic3ViIjoiYWNjZXNzX3Rva2VuIn0.TKbd3hc9vRH-OwotVg8FzvxodA4-7MjIL80p7MBzfZI"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	holdingsResp, err := fyModel.GetHoldings()
// 	if err != nil {
// 		fmt.Printf("Error getting holdings: %v\n", err)
// 		return
// 	}
// 	fmt.Println("holdings:", holdingsResp)
// }

// Logout
// func main() {
// 	appId := "M0R4WW1PYU-100"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiZDoxIiwiZDoyIiwieDowIiwieDoxIiwieDoyIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXTm5QcEV5WG9xZDJTcFZoQkFLcVhRaTRBa1BtX1IzWkpWZnVwRnNxQlFyNVQxbnd5NC00eW9IOFJwVW5DV21xOVlCb3FrTVRIWDZrdUlhOHN1TzUzdVhZSVJ0VHZMdTVvU0h4YVE5cG1iUllaUT0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMCwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTc5OTEsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5Nzk5MSwic3ViIjoiYWNjZXNzX3Rva2VuIn0.TKbd3hc9vRH-OwotVg8FzvxodA4-7MjIL80p7MBzfZI"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.Logout()
// 	if err != nil {
// 		fmt.Printf("Error logging out: %v", err)
// 	} else {
// 		fmt.Println("logout: ", response)
// 	}
// }

// All Trades
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.GetTradeBook()
// 	if err != nil {
// 		fmt.Printf("Error getting trade book: %v", err)
// 	} else {
// 		fmt.Println("trade book: ", response)
// 	}
// }

// Trade Book by Tag
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.GetTradeBookByTag("2:Exit")
// 	if err != nil {
// 		fmt.Printf("Error getting trade book by tag: %v", err)
// 	} else {
// 		fmt.Println("trade book by tag: ", response)
// 	}
// }

// Get Order Book
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.GetOrderBook()
// 	if err != nil {
// 		fmt.Printf("Error getting order book: %v", err)
// 	} else {
// 		fmt.Println("order book: ", response)
// 	}
// }

// Get Order Book by Tag
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.GetOrderBookByTag("1:Ordertag")
// 	if err != nil {
// 		fmt.Printf("Error getting order book by tag: %v", err)
// 	} else {
// 		fmt.Println("order book by tag: ", response)
// 	}
// }

// Get Positions
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.GetPositions()
// 	if err != nil {
// 		fmt.Printf("Error getting positions: %v", err)
// 	} else {
// 		fmt.Println("positions: ", response)
// 	}
// }

// Single Order Placement
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.SingleOrderAction(fyersgosdk.OrderRequest{
// 		Symbol:       "NSE:IDEA-EQ",
// 		Qty:          1,
// 		Type:         1,
// 		Side:         1,
// 		ProductType:  "CNC",
// 		LimitPrice:   100,
// 		StopPrice:    0,
// 		Validity:     "DAY",
// 		DisclosedQty: 1,
// 		OfflineOrder: false,
// 		StopLoss:     0,
// 		TakeProfit:   0,
// 		OrderTag:     "TESTEST",
// 	})
// 	if err != nil {
// 		fmt.Printf("Error single order action: %v", err)
// 	} else {
// 		fmt.Println("single order action: ", response)
// 	}
// }

// Multi Order Placement
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.MultiOrderAction([]fyersgosdk.OrderRequest{
// 		{Symbol: "NSE:ITC-EQ", Qty: 1, Type: 1, Side: 1, ProductType: "CNC", LimitPrice: 165, StopPrice: 0, DisclosedQty: 0, Validity: "DAY", OfflineOrder: true, StopLoss: 0, TakeProfit: 0, OrderTag: "tag1"},
// 		{Symbol: "NSE:ITC-EQ", Qty: 1, Type: 1, Side: 1, ProductType: "CNC", LimitPrice: 165, StopPrice: 0, DisclosedQty: 0, Validity: "DAY", OfflineOrder: true, StopLoss: 0, TakeProfit: 0, OrderTag: "tag1"},
// 		{Symbol: "NSE:ITC-EQ", Qty: 1, Type: 1, Side: 1, ProductType: "CNC", LimitPrice: 165, StopPrice: 0, DisclosedQty: 0, Validity: "DAY", OfflineOrder: true, StopLoss: 0, TakeProfit: 0, OrderTag: "tag1"},
// 	})
// 	if err != nil {
// 		fmt.Printf("Error multi order action: %v", err)
// 	} else {
// 		fmt.Println("multi order action: ", response)
// 	}
// }

// MultiLeg Order
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 		fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 		response, err := fyModel.MultiLegOrderAction([]fyersgosdk.MultiLegOrderRequest{
// 			{
// 				OrderTag:     "tag1",
// 				ProductType:  "MARGIN",
// 				OfflineOrder: false,
// 				OrderType:    "3L",
// 				Validity:     "IOC",
// 				Legs: fyersgosdk.Leg{
// 					Leg1: fyersgosdk.LegBody{
// 						Symbol:     "NSE:SBIN26FEBFUT",
// 						Qty:        750,
// 						Side:       1,
// 						Type:       1,
// 						LimitPrice: 800,
// 					},
// 					Leg2: fyersgosdk.LegBody{
// 						Symbol:     "NSE:SBIN26FEBFUT",
// 						Qty:        750,
// 						Side:       1,
// 						Type:       1,
// 						LimitPrice: 800,
// 					},
// 					Leg3: fyersgosdk.LegBody{
// 						Symbol:     "NSE:ABB26FEBFUT",
// 						Qty:        750,
// 						Side:       1,
// 						Type:       1,
// 						LimitPrice: 3,
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			fmt.Printf("Error multi leg order action: %v", err)
// 		} else {
// 			fmt.Println("multi leg order action: ", response)
// 		}
// 	}

// GTT Single
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.GTTSingleOrderAction(fyersgosdk.GTTOrderRequest{
// 		Side:        1,
// 		Symbol:      "NSE:SBIN-EQ",
// 		ProductType: "CNC",
// 		OrderInfo: fyersgosdk.OrderInfo{
// 			Leg1: fyersgosdk.Leg1{
// 				Price:        100,
// 				TriggerPrice: 100,
// 				Qty:          1,
// 			},
// 		},
// 	})
// 	if err != nil {
// 		fmt.Printf("Error GTT order action: %v", err)
// 	} else {
// 		fmt.Println("GTT order action: ", response)
// 	}
// }

// GTT OCO
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.GTTMultiOrderAction([]fyersgosdk.GTTOrderRequest{
// 		{
// 			Side:        1,
// 			Symbol:      "NSE:SBIN-EQ",
// 			ProductType: "CNC",
// 			OrderInfo: fyersgosdk.OrderInfo{
// 				Leg1: fyersgosdk.Leg1{Price: 10000, TriggerPrice: 10000, Qty: 3},
// 				Leg2: fyersgosdk.Leg2{Price: 990, TriggerPrice: 990, Qty: 3},
// 				Leg3: fyersgosdk.Leg3{Price: 990, TriggerPrice: 990, Qty: 3},
// 			},
// 		},
// 	})
// 	if err != nil {
// 		fmt.Printf("Error GTT multi order action: %v", err)
// 	} else {
// 		fmt.Println("GTT multi order action: ", response)
// 	}
// }

// GTT Modify
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 		fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 		response, err := fyModel.ModifyGTTOrder([]fyersgosdk.ModifyGTTOrderRequest{
// 			{
// 				Id: "26021800001427",
// 				OrderInfo: fyersgosdk.OrderInfo{
// 					Leg1: fyersgosdk.Leg1{Price: 1750, TriggerPrice: 1749, Qty: 5},
// 					Leg2: fyersgosdk.Leg2{Price: 1700, TriggerPrice: 1701, Qty: 5},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			fmt.Printf("Error GTT modify order action: %v", err)
// 		} else {
// 			fmt.Println("GTT modify order action: ", response)
// 		}
// 	}

// GTT Cancel
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.CancelGTTOrder("25111300002007")
// 	if err != nil {
// 		fmt.Printf("Error GTT cancel order action: %v", err)
// 	} else {
// 		fmt.Println("GTT cancel order action: ", response)
// 	}
// }

// CancelGTT
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.CancelGTTOrder("25111300002007")
// 	if err != nil {
// 		fmt.Printf("Error GTT cancel order action: %v", err)
// 	} else {
// 		fmt.Println("GTT cancel order action: ", response)
// 	}
// }

// GTT Get Order Book
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.GetGTTOrderBook()
// 	if err != nil {
// 		fmt.Printf("Error getting GTT order book: %v", err)
// 	} else {
// 		fmt.Println("GTT order book: ", response)
// 	}
// }

// Modify Orders
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.ModifyOrder(fyersgosdk.ModifyOrderRequest{
// 		Id:         "23030900015105",
// 		Qty:        10,
// 		Type:       1,
// 		Side:       1,
// 		LimitPrice: 100,
// 	})
// 	if err != nil {
// 		fmt.Printf("Error modify order action: %v", err)
// 	} else {
// 		fmt.Println("modify order action: ", response)
// 	}
// }

// Modify Multi Orders
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.ModifyMutliOrder([]fyersgosdk.ModifyMultiOrderItem{
// 		{Id: 8102710298291, Type: 1, LimitPrice: 61049, Qty: 1},
// 		{Id: 8102710298292, Type: 1, LimitPrice: 61049, Qty: 1},
// 	})
// 	if err != nil {
// 		fmt.Printf("Error modify order action: %v", err)
// 	} else {
// 		fmt.Println("modify order action: ", response)
// 	}
// }

// Cancel Order
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.CancelOrder("23030900015105")
// 	if err != nil {
// 		fmt.Printf("Error cancel order action: %v", err)
// 	} else {
// 		fmt.Println("cancel order action: ", response)
// 	}
// }

// Multi Cancel Order
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	Id := []string{"808058117761", "808058117762"}

// 	response, err := fyModel.CancelMutliOrder(Id)
// 	if err != nil {
// 		fmt.Printf("Error cancel order action: %v", err)
// 	} else {
// 		fmt.Println("Mutli cancel order action: ", response)
// 	}
// }

// Exit Order
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	response, err := fyModel.ExitPosition()
// 	if err != nil {
// 		fmt.Printf("Error Exit : %v", err)
// 	} else {
// 		fmt.Println("Exit Position: ", response)
// 	}
// }

// Exit Position - By Id
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	Id := []string{"NSE:IDEA-EQ-INTRADAY", "NSE:EASEMYTRIP-EQ-INTRADAY"}
// 	response, err := fyModel.ExitPositionById(Id)
// 	if err != nil {
// 		fmt.Printf("Error Exit : %v", err)
// 	} else {
// 		fmt.Println("Exit Position: ", response)
// 	}
// }

// Exit Position by Tag
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)

// 	response, err := fyModel.ExitPositionByProductType(fyersgosdk.ExitPositionByProductTypeRequest{
// 		Segment:     []int{10},
// 		Side:        []int{1, -1},
// 		ProductType: []string{"INTRADAY", "CNC"},
// 	})
// 	if err != nil {
// 		fmt.Printf("Error Exit : %v", err)
// 	} else {
// 		fmt.Println("Exit Position: ", response)
// 	}
// }

// Pending Order Cancel
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)

// 	response, err := fyModel.CancelPendingOrders(fyersgosdk.CancelPendingOrdersRequest{PendingOrdersCancel: 1, Id: "NSE:SBIN-EQ-INTRADAY"})
// 	if err != nil {
// 		fmt.Printf("Error cancel pending orders: %v", err)
// 	} else {
// 		fmt.Println("Cancel pending orders: ", response)
// 	}
// }

// Convert Position
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)

// 	response, err := fyModel.ConvertPosition(fyersgosdk.ConvertPositionRequest{
// 		Symbol:       "MCX:SILVERMIC20NOVFUT-INTRADAY",
// 		PositionSide: 1,
// 		ConvertQty:   1,
// 		ConvertFrom:  "INTRADAY",
// 		ConvertTo:    "CNC",
// 		Overnight:    1,
// 	})
// 	if err != nil {
// 		fmt.Printf("Error convert position: %v", err)
// 	} else {
// 		fmt.Println("Convert position: ", response)
// 	}
// }

// Market Status
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)

// 	response, err := fyModel.GetMarketStatus()
// 	if err != nil {
// 		fmt.Printf("Error convert position: %v", err)
// 	} else {
// 		fmt.Println("Convert position: ", response)
// 	}
// }

// Quotes
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// 	symbol := []string{"NSE:SBIN-EQ"}
// 	response, err := fyModel.GetStockQuotes(symbol)
// 	if err != nil {
// 		fmt.Printf("Error getting quotes: %v", err)
// 	} else {
// 		fmt.Println("Quotes: ", response)
// 	}
// }

// Market depth
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)

// 	response, err := fyModel.GetMarketDepth(fyersgosdk.MarketDepthRequest{
// 		Symbol: "NSE:SBIN-EQ",
// 		OHLCV:  "1",
// 	})
// 	if err != nil {
// 		fmt.Printf("Error getting market depth: %v", err)
// 	} else {
// 		fmt.Println("Market depth: ", response)
// 	}
// }

// Option Chain
// func main() {
// 	appId := "Z0G0WQQT6T-101"
// 	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsieDowIiwieDoxIl0sImF0X2hhc2giOiJnQUFBQUFCcGxXc3F0Tmp6UlBHZFN2S2Z0dnFWS29DdmRNN1pOclJqV09kcXFjUEtxVmhrQTdZSkVCSm1GdTVTMGpENW56YmVaY2JfSXVpYy1pd2tQWm9XRzltd2p6Y2hiUzJ5Njk4ZFQtd204aWk4NTBIaE5kMD0iLCJkaXNwbGF5X25hbWUiOiIiLCJvbXMiOiJLMSIsImhzbV9rZXkiOiJkNWU1YWY5ZmM0NWMwMzZhY2FkZmE2M2ZhZDc1YzZhMmEwZjc3ZDRmMDFlMWJkMTNlMTc4YWI3YyIsImlzRGRwaUVuYWJsZWQiOiJZIiwiaXNNdGZFbmFibGVkIjoiWSIsImZ5X2lkIjoiWUswNDM5MSIsImFwcFR5cGUiOjEwMSwiZXhwIjoxNzcxNDYxMDAwLCJpYXQiOjE3NzEzOTk5NzgsImlzcyI6ImFwaS5meWVycy5pbiIsIm5iZiI6MTc3MTM5OTk3OCwic3ViIjoiYWNjZXNzX3Rva2VuIn0._MjzvF4-6aKO6gv_AZoAo3tQUhQtJQVJS_1KldVxJLE"

// 	fyModel := fyersgosdk.NewFyersModel(appId, accessToken)

// 	response, err := fyModel.GetOptionChain(fyersgosdk.OptionChainRequest{
// 		Symbol:      "NSE:TCS-EQ",
// 		StrikeCount: 1,
// 		Timestamp:   "", // optional
// 	})
// 	if err != nil {
// 		fmt.Printf("Error getting option chain: %v", err)
// 	} else {
// 		fmt.Println("Option chain: ", response)
// 	}
// }
