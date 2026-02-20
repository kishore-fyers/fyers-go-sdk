<a href="https://fyers.in/"><img src="https://assets.fyers.in/images/logo.svg" align="right" /></a>

# Fyers Go SDK : fyers-api-v3 - v1.0.0

The official Fyers Go SDK for API-V3 Users [FYERS API](https://fyers.in/products/api/).

Fyers API is a set of REST-like APIs that provide integration with our in-house trading platform with which you can build your own customized trading applications.

## Documentation

- [Fyers API documentation](https://myapi.fyers.in/docsv3)

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

- [Compatible Go Versions](#compatible-go-versions)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Authentication](#authentication)
- [API Reference](#api-reference)
- [Examples](#examples)
- [WebSocket Streaming](#websocket-streaming)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)

## Compatible Go Versions

- **Minimum:** Go **1.18** (required by `go.mod`).
- **Compatible:** Go 1.18 and any later release (1.19, 1.20, 1.21, 1.22, 1.23, etc.).

Using Go 1.19 or newer is recommended for security and tooling support.

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
    // 1. Initialize client (auth only)
    fyClient := fyersgosdk.SetClientData("YOUR_APP_ID", "YOUR_APP_SECRET", "YOUR_REDIRECT_URL")
    fmt.Println("Login URL:", fyClient.GetLoginURL())

    // 2. After user authorizes, exchange auth code for access token
    authCode := "AUTH_CODE_FROM_REDIRECT"
    response, err := fyClient.GenerateAccessToken(authCode, fyClient)
    if err != nil {
        log.Fatal("Error generating access token:", err)
    }
    // Parse response JSON to get access_token; then create model for API calls

    // 3. Use FyersModel for all API calls (profile, orders, data, etc.)
    fyModel := fyersgosdk.NewFyersModel("YOUR_APP_ID", "ACCESS_TOKEN")
    profile, err := fyModel.GetProfile()
    if err != nil {
        log.Fatal("Error getting profile:", err)
    }
    fmt.Println("Profile:", profile)
}
```

### Get Market Data

```go
fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
// Quotes (up to 50 symbols)
quotes, err := fyModel.GetStockQuotes([]string{"NSE:SBIN-EQ", "NSE:NIFTY50-INDEX"})
// History
history, err := fyModel.GetHistory(fyersgosdk.HistoryRequest{
    Symbol: "NSE:SBIN-EQ", Resolution: "30", DateFormat: "1",
    RangeFrom: "2021-01-01", RangeTo: "2021-01-02", ContFlag: "",
})
```

### Place an Order

```go
fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
response, err := fyModel.SingleOrderAction(fyersgosdk.OrderRequest{
    Symbol: "NSE:IDEA-EQ", Qty: 1, Type: 1, Side: 1, ProductType: "CNC",
    LimitPrice: 100, Validity: "DAY", DisclosedQty: 0, OfflineOrder: false,
})
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
// Step 1: Initialize client (appId, appSecret, redirectUrl)
fyClient := fyersgosdk.SetClientData(appId, appSecret, redirectUrl)

// Step 2: Generate login URL
loginURL := fyClient.GetLoginURL()
fmt.Println("Please visit:", loginURL)

// Step 3: User authorizes; you get auth code from redirect URL
// Step 4: Exchange auth code for access token (returns JSON string)
response, err := fyClient.GenerateAccessToken(authCode, fyClient)

// Step 5: Or refresh token with PIN
response, err := fyClient.GenerateAccessTokenFromRefreshToken(refreshToken, pin, fyClient)

// Step 6: Create FyersModel for all API calls
fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
```

## 📚 API Reference

All API methods return `(string, error)`; the string is the raw JSON response. Use **Client** only for auth; use **FyersModel** for everything after login.

### Authentication (Client)

| Method | Description |
|--------|-------------|
| `SetClientData(appId, appSecret, redirectUrl string) *Client` | Create client for login/refresh. |
| `GetLoginURL() string` | URL for user to authorize and get auth code. |
| `GenerateAccessToken(authToken string, fyClient *Client) (string, error)` | Exchange auth code for access token. |
| `GenerateAccessTokenFromRefreshToken(refreshToken, pin string, fyClient *Client) (string, error)` | Get new access token from refresh token. |
| `SetAccessToken(accessToken string) *Client` | Set token on client (optional). |

### User & Profile (FyersModel)

| Method | Description |
|--------|-------------|
| `GetProfile() (string, error)` | User profile. |
| `GetFunds() (string, error)` | Account funds. |
| `GetHoldings() (string, error)` | Portfolio holdings. |
| `Logout() (string, error)` | Logout session. |

### Transaction Info (FyersModel)

| Method | Description |
|--------|-------------|
| `GetOrderBook() (string, error)` | All orders. |
| `GetOrderBookByTag(tag string) (string, error)` | Orders by tag. |
| `GetOrderById(id string) (string, error)` | Single order by ID. |
| `GetPositions() (string, error)` | Open positions. |
| `GetTradeBook() (string, error)` | Trade book. |
| `GetTradeBookByTag(tag string) (string, error)` | Trades by tag. |

### Orders (FyersModel)

| Method | Description |
|--------|-------------|
| `SingleOrderAction(orderRequest OrderRequest) (string, error)` | Place single order. |
| `MultiOrderAction(orderRequests []OrderRequest) (string, error)` | Place multiple orders. |
| `MultiLegOrderAction(orderRequests []MultiLegOrderRequest) (string, error)` | Place multi-leg orders. |
| `ModifyOrder(orderRequest ModifyOrderRequest) (string, error)` | Modify single order (PATCH). |
| `ModifyMutliOrder(requests []ModifyMultiOrderItem) (string, error)` | Modify multiple orders (PATCH). |
| `CancelOrder(Id string) (string, error)` | Cancel single order. |
| `CancelMutliOrder(orderIds []string) (string, error)` | Cancel multiple orders (DELETE). |

### GTT Orders (FyersModel)

| Method | Description |
|--------|-------------|
| `GTTSingleOrderAction(orderRequest GTTOrderRequest) (string, error)` | Place GTT order. |
| `GTTMultiOrderAction(orderRequests []GTTOrderRequest) (string, error)` | Place GTT (sends first only). |
| `ModifyGTTOrder(orderRequests []ModifyGTTOrderRequest) (string, error)` | Modify GTT order. |
| `CancelGTTOrder(orderId string) (string, error)` | Cancel GTT order. |
| `GetGTTOrderBook() (string, error)` | GTT order book. |

### Trade Operations / Positions (FyersModel)

| Method | Description |
|--------|-------------|
| `ExitPosition() (string, error)` | Exit all positions (DELETE with exit_all). |
| `ExitPositionById(orderId []string) (string, error)` | Exit by order IDs. |
| `ExitPositionByProductType(req ExitPositionByProductTypeRequest) (string, error)` | Exit by segment/side/productType. |
| `CancelPendingOrders(req CancelPendingOrdersRequest) (string, error)` | Cancel pending orders (optional Id for single symbol). |
| `ConvertPosition(req ConvertPositionRequest) (string, error)` | Convert position (e.g. INTRADAY to CNC). |

### Market Data (FyersModel)

| Method | Description |
|--------|-------------|
| `GetHistory(req HistoryRequest) (string, error)` | OHLCV history (symbol, resolution, range). |
| `GetStockQuotes(symbols []string) (string, error)` | Quotes for up to 50 symbols. |
| `GetMarketDepth(req MarketDepthRequest) (string, error)` | Market depth (symbol, ohlcv_flag). |
| `GetOptionChain(req OptionChainRequest) (string, error)` | Options chain. |

### Broker (FyersModel)

| Method | Description |
|--------|-------------|
| `GetMarketStatus() (string, error)` | Market status. |

### Alerts (FyersModel)

| Method | Description |
|--------|-------------|
| `GetAlerts() (string, error)` | List price alerts. |
| `CreateAlert(req AlertRequest) (string, error)` | Create alert. |
| `UpdateAlert(alertId string, req AlertRequest) (string, error)` | Update alert. |
| `DeleteAlert(alertId string) (string, error)` | Delete alert. |
| `ToggleAlert(alertId string) (string, error)` | Toggle alert. |

### WebSocket (package-level)

| Function | Description |
|----------|-------------|
| `DataSocket(fyModel *FyersModel, req DataSocketRequest) (map[string]interface{}, error)` | Real-time data stream (quotes, depth). |
| `OrderSocket(fyClient *Client, req OrderSocketRequest) (map[string]interface{}, error)` | Real-time orders, trades, positions. |

## 🔌 WebSocket Streaming

Use **DataSocket** for live market data (quotes, depth) and **OrderSocket** for orders, trades, and positions. Both run until interrupt (e.g. Ctrl+C).

```go
// Data WebSocket (symbols, lite/full mode)
fyModel := fyersgosdk.NewFyersModel(appId, accessToken)
result, err := fyersgosdk.DataSocket(fyModel, fyersgosdk.DataSocketRequest{
    Symbols: []string{"NSE:SBIN-EQ"},
    LiteMode: false,
    DataType: "symbol",
})

// Order WebSocket (trades, positions, orders)
fyClient := fyersgosdk.SetClientData(appId, appSecret, redirectUrl)
fyClient.SetAccessToken(accessToken)
result, err := fyersgosdk.OrderSocket(fyClient, fyersgosdk.OrderSocketRequest{
    TradeOperations: []string{"OnOrders", "OnTrades", "OnPositions"},
})
```

## 📁 Examples

All runnable examples live in **[examples/fyers/fyers.go](examples/fyers/fyers.go)**. Uncomment the `main()` you need and run from that directory. The file includes:

### Authentication
- **Get Auth Code URL** – `SetClientData`, `GetLoginURL`
- **Generate Access Token** – `GenerateAccessToken`
- **Generate Access Token From Refresh Token** – `GenerateAccessTokenFromRefreshToken`

### User & Profile
- **Get Profile** – `GetProfile`
- **Get Funds** – `GetFunds`
- **Get Holdings** – `GetHoldings`
- **Logout** – `Logout`

### Transaction Info
- **All Trades** – `GetTradeBook`
- **Trade Book by Tag** – `GetTradeBookByTag`
- **Get Order Book** – `GetOrderBook`
- **Get Order Book by Tag** – `GetOrderBookByTag`
- **Get Positions** – `GetPositions`

### Orders
- **Single Order Placement** – `SingleOrderAction`
- **Multi Order Placement** – `MultiOrderAction`
- **MultiLeg Order** – `MultiLegOrderAction`
- **Modify Orders** – `ModifyOrder`
- **Modify Multi Orders** – `ModifyMutliOrder`
- **Cancel Order** – `CancelOrder`
- **Multi Cancel Order** – `CancelMutliOrder`

### GTT
- **GTT Single** – `GTTSingleOrderAction`
- **GTT OCO** – `GTTMultiOrderAction`
- **GTT Modify** – `ModifyGTTOrder`
- **GTT Cancel / CancelGTT** – `CancelGTTOrder`
- **GTT Get Order Book** – `GetGTTOrderBook`

### Trade Operations / Positions
- **Exit Order** – `ExitPosition`
- **Exit Position By Id** – `ExitPositionById`
- **Exit Position by Tag** – `ExitPositionByProductType`
- **Pending Order Cancel** – `CancelPendingOrders`
- **Convert Position** – `ConvertPosition`

### Market Data & Broker
- **Market Status** – `GetMarketStatus`
- **Quotes** – `GetStockQuotes`
- **Market depth** – `GetMarketDepth`
- **Option Chain** – `GetOptionChain`

For **Get History**, use `GetHistory(HistoryRequest{...})` as in the API Reference above.

## ⚠️ Error Handling

API methods return `(string, error)`. Check `err` and parse the response string as JSON when needed.

```go
response, err := fyModel.GetProfile()
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
// response is JSON string; unmarshal for structured access
```

## 🔧 Configuration

Pass credentials into `SetClientData(appId, appSecret, redirectUrl)` and `NewFyersModel(appId, accessToken)`. Prefer environment variables or a config file; do not commit secrets.

## 🧪 Testing

All tests live in a single file under the **`test/`** folder. They cover every SDK function (auth, profile, orders, GTT, trade operations, market data, alerts) without requiring real credentials; invalid/empty tokens are used so the suite runs offline.

**Run all tests:**

```bash
go test ./test/
```

**Run with verbose output:**

```bash
go test ./test/ -v
```

**Run with coverage:**

```bash
go test ./test/ -cover
```

**Run with valid token (log all API responses):**  
Set `FYERS_APP_ID` and `FYERS_ACCESS_TOKEN`; then run the JSON-driven test. All case inputs and API responses are written to `test/test_output.json` and `test/report-YYYY-MM-DD.txt`. Do not commit the token; use env only.

```bash
FYERS_APP_ID=your_app_id FYERS_ACCESS_TOKEN=your_token go test ./test/ -v -run TestRunAllFromJSON
```

**Validation test with token:**  
Runs the case names listed in `test_cases.json` under `validationCases` (profile, funds, holdings, order book, positions, trade book, market status, history, quotes, depth, option chain, alerts by default) and fails if any response is not success (`s:"ok"` or `code: 200`). Add or remove names in `validationCases` to change which APIs are validated. Skipped if `FYERS_APP_ID` or `FYERS_ACCESS_TOKEN` is not set.

```bash
FYERS_APP_ID=your_app_id FYERS_ACCESS_TOKEN=your_token go test ./test/ -v -run TestRunValidationWithToken
```

**Optional integration test:** set `FYERS_APP_ID` and `FYERS_ACCESS_TOKEN` to run `TestGetProfile_WithToken` against the real API; otherwise it is skipped.

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
