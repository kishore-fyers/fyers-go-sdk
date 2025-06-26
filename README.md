# Fyers Go SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/your-username/fyersgosdk)](https://goreportcard.com/report/github.com/your-username/fyersgosdk)
[![Go Version](https://img.shields.io/github/go-mod/go-version/your-username/fyersgosdk)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Fyers API](https://img.shields.io/badge/API-Fyers%20v3-blue.svg)](https://myapi.fyers.in/)

A comprehensive Go SDK for the [Fyers API](https://myapi.fyers.in/) that provides seamless integration with Fyers trading platform. This SDK enables you to build trading applications, algorithmic trading systems, and market data analysis tools using Go.

## 🚀 Features

- **Complete API Coverage**: Full implementation of Fyers API v3 endpoints
- **Authentication**: OAuth2-based authentication with automatic token management
- **Real-time Data**: WebSocket streaming for live market data and order updates
- **Trading Operations**: Place, modify, and cancel orders with support for various order types
- **Portfolio Management**: Access holdings, positions, and fund information
- **Market Data**: Historical data, real-time quotes, market depth, and options chain
- **Error Handling**: Comprehensive error handling with detailed error messages
- **Type Safety**: Strongly typed Go structs for all API responses
- **Examples**: Extensive examples for all major functionality
- **Testing**: Comprehensive test suite with mock responses

## 📋 Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Authentication](#authentication)
- [API Reference](#api-reference)
- [Examples](#examples)
- [WebSocket Streaming](#websocket-streaming)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)

## 📦 Installation

```bash
go get github.com/your-username/fyersgosdk
```

## 🚀 Quick Start

### Basic Setup

```go
package main

import (
    "fmt"
    "log"
    fyersgosdk "github.com/your-username/fyersgosdk"
)

func main() {
    // Initialize client with your credentials
    client := fyersgosdk.SetClientData(
        "YOUR_CLIENT_ID",
        "YOUR_APP_ID", 
        "YOUR_APP_SECRET",
        "YOUR_REDIRECT_URL",
    )

    // Generate login URL
    loginURL := client.GetLoginURL()
    fmt.Println("Login URL:", loginURL)

    // After user authorization, get auth code and generate access token
    authCode := "AUTH_CODE_FROM_REDIRECT"
    response, accessTokenResp, err := client.GenerateAccessToken(authCode, client)
    if err != nil {
        log.Fatal("Error generating access token:", err)
    }

    // Set access token for subsequent API calls
    client.SetAccessToken(accessTokenResp.AccessToken)

    // Get user profile
    _, profile, err := client.GetProfile(client)
    if err != nil {
        log.Fatal("Error getting profile:", err)
    }

    fmt.Printf("Welcome, %s!\n", profile.Name)
}
```

### Get Market Data

```go
// Get real-time quotes
symbols := "NSE:NIFTY50-INDEX,NSE:BANKNIFTY-INDEX"
_, quotes, err := client.GetQuotes(symbols)
if err != nil {
    log.Fatal("Error getting quotes:", err)
}

for _, quote := range quotes {
    fmt.Printf("%s: %.2f\n", quote.Symbol, quote.Ltp)
}
```

### Place an Order

```go
order := fyersgosdk.OrderRequest{
    Symbol:        "NSE:RELIANCE-EQ",
    Qty:           100,
    Side:          1, // Buy
    Type:          2, // Market
    ProductType:   "INTRADAY",
    Validity:      "DAY",
    DisclosedQty:  0,
    OfflineOrder:  "False",
    StopLoss:      0,
    TakeProfit:    0,
}

_, orderResp, err := client.PlaceOrder(order)
if err != nil {
    log.Fatal("Error placing order:", err)
}

fmt.Printf("Order placed successfully! Order ID: %s\n", orderResp.Id)
```

## 🔐 Authentication

The Fyers Go SDK uses OAuth2 authentication. Here's the complete flow:

### 1. App Registration
First, register your application on the [Fyers Developer Portal](https://myapi.fyers.in/):

1. Create a new app
2. Get your `App ID` and `App Secret`
3. Set your redirect URL
4. Note your `Client ID`

### 2. Authentication Flow

```go
// Step 1: Initialize client
client := fyersgosdk.SetClientData(clientId, appId, appSecret, redirectUrl)

// Step 2: Generate login URL
loginURL := client.GetLoginURL()
fmt.Println("Please visit:", loginURL)

// Step 3: User authorizes and you get auth code from redirect
// Step 4: Exchange auth code for access token
response, accessTokenResp, err := client.GenerateAccessToken(authCode, client)
if err != nil {
    log.Fatal(err)
}

// Step 5: Use access token for API calls
client.SetAccessToken(accessTokenResp.AccessToken)
```

### 3. Token Management

```go
// Set refresh token for automatic token renewal
client.SetRefreshToken(accessTokenResp.RefreshToken)

// The SDK automatically handles token refresh when needed
```

## 📚 API Reference

### User Operations

#### Get Profile
```go
_, profile, err := client.GetProfile(client)
```

#### Get Funds
```go
_, funds, err := client.GetFunds(client)
```

#### Get Holdings
```go
_, holdings, err := client.GetHoldings(client)
```

### Trading Operations

#### Place Order
```go
order := fyersgosdk.OrderRequest{
    Symbol:      "NSE:RELIANCE-EQ",
    Qty:         100,
    Side:        1, // 1=Buy, 2=Sell
    Type:        2, // 1=Limit, 2=Market, 3=Stop, 4=StopLimit
    ProductType: "INTRADAY", // INTRADAY, CNC, MARGIN
    Validity:    "DAY", // DAY, IOC, TTL
}

_, orderResp, err := client.PlaceOrder(order)
```

#### Modify Order
```go
modifyReq := fyersgosdk.ModifyOrderRequest{
    Id:     "ORDER_ID",
    Qty:    150,
    Price:  2500.50,
    Type:   1, // Limit order
}

_, modifyResp, err := client.ModifyOrder(modifyReq)
```

#### Cancel Order
```go
_, cancelResp, err := client.CancelOrder("ORDER_ID")
```

### Market Data

#### Get Quotes
```go
symbols := "NSE:NIFTY50-INDEX,NSE:BANKNIFTY-INDEX"
_, quotes, err := client.GetQuotes(symbols)
```

#### Get Historical Data
```go
historyReq := fyersgosdk.HistoryRequest{
    Symbol:     "NSE:NIFTY50-INDEX",
    Resolution: "1D", // 1, 2, 3, 5, 10, 15, 20, 30, 60, 120, 1D, 1W, 1M
    DateFormat: 1,
    RangeFrom:  "2024-01-01",
    RangeTo:    "2024-01-31",
    ContFlag:   "1",
}

_, history, err := client.GetHistory(historyReq)
```

#### Get Market Depth
```go
_, depth, err := client.GetMarketDepth("NSE:RELIANCE-EQ")
```

### Portfolio Management

#### Get Positions
```go
_, positions, err := client.GetPositions(client)
```

#### Get Trade Book
```go
_, trades, err := client.GetTradeBook(client)
```

#### Get Orders
```go
_, orders, err := client.GetOrders(client)
```

## 🔌 WebSocket Streaming

The SDK provides real-time data streaming capabilities:

### General Stream (Orders, Trades, Positions)

```go
// Initialize general stream
stream := fyersgosdk.NewGeneralStream(client)

// Subscribe to order updates
stream.SubscribeOrders(func(order fyersgosdk.OrderUpdate) {
    fmt.Printf("Order Update: %+v\n", order)
})

// Subscribe to trade updates
stream.SubscribeTrades(func(trade fyersgosdk.TradeUpdate) {
    fmt.Printf("Trade Update: %+v\n", trade)
})

// Subscribe to position updates
stream.SubscribePositions(func(position fyersgosdk.PositionUpdate) {
    fmt.Printf("Position Update: %+v\n", position)
})

// Connect to stream
err := stream.Connect()
if err != nil {
    log.Fatal("Error connecting to stream:", err)
}

// Keep connection alive
select {}
```

### Market Data Stream

```go
// Initialize market data stream
marketStream := fyersgosdk.NewMarketDataStream(client)

// Subscribe to symbol updates
symbols := []string{"NSE:NIFTY50-INDEX", "NSE:BANKNIFTY-INDEX"}
marketStream.SubscribeSymbols(symbols, func(quote fyersgosdk.QuoteUpdate) {
    fmt.Printf("Quote Update: %s = %.2f\n", quote.Symbol, quote.Ltp)
})

// Connect to stream
err := marketStream.Connect()
if err != nil {
    log.Fatal("Error connecting to market stream:", err)
}
```

## 📁 Examples

The SDK includes comprehensive examples in the `examples/` directory:

### Basic Examples
- [Authentication](examples/fyers/basic/fyers.go) - Basic authentication flow
- [Profile Management](examples/user/profile/main.go) - Get user profile
- [Funds Information](examples/user/funds/main.go) - Get account funds
- [Holdings](examples/user/holdings/main.go) - Get portfolio holdings

### Trading Examples
- [Basic Order Placement](examples/orders/basic/main.go) - Place simple orders
- [Multi-Leg Orders](examples/orders/multileg/main.go) - Complex order types
- [GTT Orders](examples/orders/gtt/single/main.go) - Good Till Triggered orders
- [Order Modifications](examples/trading/modify/main.go) - Modify existing orders

### Data Examples
- [Historical Data](examples/data/history/main.go) - Get historical prices
- [Real-time Quotes](examples/data/quotes/main.go) - Live market data
- [Market Depth](examples/data/depth/main.go) - Order book data
- [Options Chain](examples/data/options/main.go) - Options data

### Streaming Examples
- [Order Updates](examples/stream/general/orders/main.go) - Real-time order updates
- [Trade Updates](examples/stream/general/trades/main.go) - Live trade information
- [Market Data](examples/stream/marketdata/symbols/main.go) - Live quotes

## ⚠️ Error Handling

The SDK provides comprehensive error handling:

```go
_, profile, err := client.GetProfile(client)
if err != nil {
    // Check for specific error types
    if fyersErr, ok := err.(*fyersgosdk.FyersError); ok {
        switch fyersErr.Code {
        case 401:
            fmt.Println("Authentication failed - check your tokens")
        case 429:
            fmt.Println("Rate limit exceeded - wait before retrying")
        case 500:
            fmt.Println("Server error - try again later")
        default:
            fmt.Printf("API Error %d: %s\n", fyersErr.Code, fyersErr.Message)
        }
    } else {
        fmt.Printf("Network error: %v\n", err)
    }
    return
}
```

## 🔧 Configuration

### Environment Variables

You can configure the SDK using environment variables:

```bash
export FYERS_CLIENT_ID="your_client_id"
export FYERS_APP_ID="your_app_id"
export FYERS_APP_SECRET="your_app_secret"
export FYERS_REDIRECT_URL="your_redirect_url"
```

### Custom HTTP Client

```go
// Create custom HTTP client with timeout
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

// Initialize client with custom HTTP client
client := fyersgosdk.SetClientData(clientId, appId, appSecret, redirectUrl)
client.SetHTTPClient(httpClient)
```

## 🧪 Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run specific test files:

```bash
go test ./auth_test.go
go test ./orders_test.go
```

## 📊 Rate Limits

The Fyers API has rate limits to ensure fair usage:

- **REST API**: 120 requests per minute per user
- **WebSocket**: 1 connection per user
- **Data API**: 100 requests per minute per user

The SDK automatically handles rate limiting and provides appropriate error messages.

## 🔒 Security

- Never commit your API credentials to version control
- Use environment variables for sensitive data
- Implement proper token storage and refresh mechanisms
- Validate all user inputs before sending to API
- Use HTTPS for all communications

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/fyersgosdk.git`
3. Create a feature branch: `git checkout -b feature/amazing-feature`
4. Make your changes and add tests
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -m 'Add amazing feature'`
7. Push to the branch: `git push origin feature/amazing-feature`
8. Open a Pull Request

### Code Style

- Follow Go conventions and [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for code formatting
- Add comments for exported functions and types
- Write tests for new functionality
- Update documentation for API changes

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

- **Documentation**: [Fyers API Documentation](https://myapi.fyers.in/)
- **Issues**: [GitHub Issues](https://github.com/your-username/fyersgosdk/issues)
- **Discussions**: [GitHub Discussions](https://github.com/your-username/fyersgosdk/discussions)
- **Email**: support@your-email.com

## 🙏 Acknowledgments

- [Fyers](https://fyers.in/) for providing the excellent trading API
- The Go community for the amazing ecosystem
- All contributors who have helped improve this SDK

## 📈 Changelog

See [CHANGELOG.md](CHANGELOG.md) for a detailed history of changes.

---

**Disclaimer**: This SDK is not officially affiliated with Fyers. Please refer to the [official Fyers API documentation](https://myapi.fyers.in/) for the most up-to-date information about the API.

**Trading Risk**: Trading in financial markets involves risk. This SDK is provided as-is without any warranty. Users are responsible for their trading decisions and should understand the risks involved.
