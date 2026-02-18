package fyersgosdk

const (
	//Live
	BaseURL      = "https://api-t1.fyers.in/api/v3"
	BaseDataURL  = "https://api-t1.fyers.in/data"
	Websocket    = "wss://socket.fyers.in/trade/v3"
	HSMWebsocket = "socket.fyers.in/hsm/v1-5/prod"

	ValidateAuthCodeURL     = BaseURL + "/validate-authcode"
	GenerateAuthCodeURL     = BaseURL + "/generate-authcode?"
	ValidateRefreshTokenURL = BaseURL + "/validate-refresh-token"
	LogoutURL               = BaseURL + "/logout"
	SymbolTokenURL          = BaseDataURL + "/symbol-token"

	ProfileURL             = BaseURL + "/profile"
	FundURL                = BaseURL + "/funds"
	HoldingsURL            = BaseURL + "/holdings"
	PositionURL            = BaseURL + "/positions"
	TradeBookURL           = BaseURL + "/tradebook"
	TradeBookByTagURL      = BaseURL + "/tradebook?order_tag="
	SingleOrderActionURL   = BaseURL + "/orders/sync"
	MultipleOrderActionURL = BaseURL + "/multi-order/sync"
	MultiLegOrderURL       = BaseURL + "/multileg/orders/sync"
	GTTOrderURL            = BaseURL + "/gtt/orders/sync"
	GTTOrderBookURL        = BaseURL + "/gtt/orders"
	OrdersByTagURL         = BaseURL + "/orders??order_tag="
	OrderByIdURL           = BaseURL + "/orders?id="
	OrderBookURL           = BaseURL + "/orders"
	OrderCheckMarginURL    = BaseURL + "/multiorder/margin"
	MarketDepthURL         = BaseDataURL + "/depth?symbol="
	MarketStatusURL        = BaseDataURL + "/marketStatus"
	StockHistoryURL        = BaseDataURL + "/history?"
	StockQuotesURL         = BaseDataURL + "/quotes?symbols="
	OptionChainURl         = BaseDataURL + "/options-chain-v3?"
	AlertsURL              = BaseURL + "/price-alert"
	ToggleAlertURL         = BaseURL + "/toggle-alert"
)
