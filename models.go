package fyersgosdk

import (
	"log"
	"net/http"
)

type Client struct {
	clientId     string
	authToken    string
	accessToken  string
	appId        string
	appSecret    string
	redirectUrl  string
	refreshToken string
	pin          string
	retryCount   int //1-5

	httpClient HTTPClient
}

// FyersModel is the API client for profile, orders, positions, data, and alerts.
// Use Client only for GetLoginURL and GenerateAccessToken; use FyersModel for all other API calls.
type FyersModel struct {
	appId       string
	accessToken string
	httpClient  HTTPClient
}

type ClientOptions struct {
	Debug      bool
	Logger     *log.Logger
	HTTPClient *http.Client
}

// Common API response structure
type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	S       string `json:"s"`
}

type AccessTokenResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	S            string `json:"s"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// User

type Profile struct {
	APIResponse
	Name          string `json:"name"`
	Image         string `json:"image"`
	DisplayName   string `json:"display_name"`
	EmailId       string `json:"email_id"`
	PAN           string `json:"PAN"`
	FyId          string `json:"fy_id"`
	PinChangeDate string `json:"pin_change_date"`
	MobileNumber  string `json:"mobile_number"`
	Totp          bool   `json:"totp"`
	PwdChangeDate string `json:"pwd_change_date"`
	PwdToExpire   int    `json:"pwd_to_expire"`
	DdpiEnabled   bool   `json:"ddpi_enabled"`
	MtfEnabled    bool   `json:"mtf_enabled"`
}

type Funds struct {
	APIResponse
	FundLimit []FundLimit `json:"fund_limit"`
}

type FundLimit struct {
	ID              int     `json:"id"`
	Title           string  `json:"title"`
	EquityAmount    float64 `json:"equityAmount"`
	CommodityAmount float64 `json:"commodityAmount"`
}

type Holdings struct {
	APIResponse
	Holdings []Holding `json:"holdings"`
	Overall  Overall   `json:"overall"`
}

type Overall struct {
	CountTotal        int     `json:"count_total"`
	TotalInvestment   float64 `json:"total_investment"`
	TotalCurrentValue float64 `json:"total_current_value"`
	TotalPl           float64 `json:"total_pl"`
	PnlPerc           float64 `json:"pnl_perc"`
}

type Holding struct {
	HoldingType             string  `json:"holdingType"`
	Quantity                int     `json:"quantity"`
	CostPrice               float64 `json:"costPrice"`
	MarketVal               float64 `json:"marketVal"`
	RemainingQuantity       int     `json:"remainingQuantity"`
	Pl                      float64 `json:"pl"`
	Ltp                     float64 `json:"ltp"`
	Id                      int     `json:"id"`
	FyToken                 string  `json:"fyToken"`
	Exchange                int     `json:"exchange"`
	Symbol                  string  `json:"symbol"`
	Segment                 int     `json:"segment"`
	Isin                    string  `json:"isin"`
	QtyT1                   int     `json:"qty_t1"`
	RemainingPledgeQuantity int     `json:"remainingPledgeQuantity"`
	CollateralQuantity      int     `json:"collateralQuantity"`
}

// Transaction Info

type OrderBook struct {
	APIResponse
	OrderBook []OrderBookItem `json:"orderBook"`
}

type OrderBookItem struct {
	ClientId          string  `json:"clientId"`
	Id                string  `json:"id"`
	ExchOrdId         string  `json:"exchOrdId"`
	Qty               int     `json:"qty"`
	RemainingQuantity int     `json:"remainingQuantity"`
	FilledQty         int     `json:"filledQty"`
	DiscloseQty       int     `json:"discloseQty"`
	LimitPrice        float64 `json:"limitPrice"`
	StopPrice         float64 `json:"stopPrice"`
	TradedPrice       float64 `json:"tradedPrice"`
	Type              int     `json:"type"`
	FyToken           string  `json:"fyToken"`
	Exchange          int     `json:"exchange"`
	Segment           int     `json:"segment"`
	Symbol            string  `json:"symbol"`
	Instrument        int     `json:"instrument"`
	Message           string  `json:"message"`
	OfflineOrder      bool    `json:"offlineOrder"`
	OrderDateTime     string  `json:"orderDateTime"`
	OrderValidity     string  `json:"orderValidity"`
	Pan               string  `json:"pan"`
	ProductType       string  `json:"productType"`
	Side              int     `json:"side"`
	Status            int     `json:"status"`
	Source            string  `json:"source"`
	ExSym             string  `json:"ex_sym"`
	Description       string  `json:"description"`
	Ch                float64 `json:"ch"`
	Chp               float64 `json:"chp"`
	Lp                float64 `json:"lp"`
	SlNo              int     `json:"slNo"`
	DqQtyRem          int     `json:"dqQtyRem"`
	OrderNumStatus    string  `json:"orderNumStatus"`
	DisclosedQty      int     `json:"disclosedQty"`
	OrderTag          string  `json:"orderTag"`
}

type Position struct {
	APIResponse
	NetPositions []NetPosition   `json:"netPositions"`
	Overall      OverallPosition `json:"overall"`
}

type NetPosition struct {
	NetQty           int     `json:"netQty"`
	Qty              int     `json:"qty"`
	AvgPrice         float64 `json:"avgPrice"`
	NetAvg           float64 `json:"netAvg"`
	Side             int     `json:"side"`
	ProductType      string  `json:"productType"`
	RealizedProfit   float64 `json:"realized_profit"`
	UnrealizedProfit float64 `json:"unrealized_profit"`
	Pl               float64 `json:"pl"`
	Ltp              float64 `json:"ltp"`
	BuyQty           int     `json:"buyQty"`
	BuyAvg           float64 `json:"buyAvg"`
	BuyVal           float64 `json:"buyVal"`
	SellQty          int     `json:"sellQty"`
	SellAvg          float64 `json:"sellAvg"`
	SellVal          float64 `json:"sellVal"`
	SlNo             int     `json:"slNo"`
	FyToken          string  `json:"fyToken"`
	CrossCurrency    string  `json:"crossCurrency"`
	RbiRefRate       float64 `json:"rbiRefRate"`
	QtyMultiCom      float64 `json:"qtyMulti_com"`
	Segment          int     `json:"segment"`
	Symbol           string  `json:"symbol"`
	Id               string  `json:"id"`
	CFBuyQty         int     `json:"cfBuyQty"`
	CFSellQty        int     `json:"cfSellQty"`
	DayBuyQty        int     `json:"dayBuyQty"`
	DaySellQty       int     `json:"daySellQty"`
	Exchange         int     `json:"exchange"`
}

type OverallPosition struct {
	CountTotal   int     `json:"count_total"`
	CountOpen    int     `json:"count_open"`
	PlTotal      float64 `json:"pl_total"`
	PlRealized   float64 `json:"pl_realized"`
	PlUnrealized float64 `json:"pl_unrealized"`
}

type TradeBook struct {
	APIResponse
	TradeBook []TradeBookItem `json:"tradeBook"`
}

type TradeBookItem struct {
	ClientId        string  `json:"clientId"`
	OrderDateTime   string  `json:"orderDateTime"`
	OrderNumber     string  `json:"orderNumber"`
	ExchangeOrderNo string  `json:"exchangeOrderNo"`
	Exchange        int     `json:"exchange"`
	Side            int     `json:"side"`
	Segment         int     `json:"segment"`
	OrderType       int     `json:"orderType"`
	FyToken         string  `json:"fyToken"`
	ProductType     string  `json:"productType"`
	TradedQty       int     `json:"tradedQty"`
	TradePrice      float64 `json:"tradePrice"`
	TradeValue      float64 `json:"tradeValue"`
	TradeNumber     string  `json:"tradeNumber"`
	Row             int     `json:"row"`
	Symbol          string  `json:"symbol"`
	OrderTag        string  `json:"orderTag"`
}

// Order Placement

type OrderRequest struct {
	Symbol       string  `json:"symbol"`
	Qty          int     `json:"qty"`
	Type         int     `json:"type"`
	Side         int     `json:"side"`
	ProductType  string  `json:"productType"`
	LimitPrice   float64 `json:"limitPrice"`
	StopPrice    float64 `json:"stopPrice"`
	Validity     string  `json:"validity"`
	DisclosedQty int     `json:"disclosedQty"`
	OfflineOrder bool    `json:"offlineOrder"`
	StopLoss     float64 `json:"stopLoss"`
	TakeProfit   float64 `json:"takeProfit"`
	OrderTag     string  `json:"orderTag"`
}

type OrderResponse struct {
	APIResponse
	Id string `json:"id"`
}

type MultiLegOrderRequest struct {
	OrderTag     string `json:"orderTag"`
	ProductType  string `json:"productType"`
	OfflineOrder bool   `json:"offlineOrder"`
	OrderType    string `json:"orderType"`
	Validity     string `json:"validity"`
	Legs         Leg    `json:"legs"`
}

type Leg struct {
	Leg1 LegBody `json:"leg1"`
	Leg2 LegBody `json:"leg2"`
	Leg3 LegBody `json:"leg3"`
}

type LegBody struct {
	Symbol     string  `json:"symbol"`
	Qty        int     `json:"qty"`
	Side       int     `json:"side"`
	Type       int     `json:"type"`
	LimitPrice float64 `json:"limitPrice"`
}

// GTT Order

type GTTOrderRequest struct {
	Side        int       `json:"side"`
	Symbol      string    `json:"symbol"`
	ProductType string    `json:"productType"`
	OrderInfo   OrderInfo `json:"orderInfo"`
}

type OrderInfo struct {
	Leg1 Leg1 `json:"leg1"`
	Leg2 Leg2 `json:"leg2,omitempty"`
}

type Leg1 struct {
	Price        float64 `json:"price"`
	TriggerPrice float64 `json:"triggerPrice"`
	Qty          int     `json:"qty"`
}
type Leg2 struct {
	Price        float64 `json:"price"`
	TriggerPrice float64 `json:"triggerPrice"`
	Qty          int     `json:"qty"`
}

type GTTOrderResponse struct {
	APIResponse
	OrderBook []OrderBookItem `json:"orderBook"`
}

type GTTOrderBookItem struct {
	ClientId        string  `json:"clientId"`
	Exchange        int     `json:"exchange"`
	FyToken         string  `json:"fy_token"`
	IdFyers         string  `json:"id_fyers"`
	Id              string  `json:"id"`
	Instrument      int     `json:"instrument"`
	LotSize         int     `json:"lot_size"`
	Multiplier      int     `json:"multiplier"`
	OrdStatus       int     `json:"ord_status"`
	Precision       int     `json:"precision"`
	PriceLimit      float64 `json:"price_limit"`
	Price2Limit     float64 `json:"price2_limit"`
	PriceTrigger    float64 `json:"price_trigger"`
	Price2Trigger   float64 `json:"price2_trigger"`
	ProductType     string  `json:"product_type"`
	Qty             int     `json:"qty"`
	Qty2            int     `json:"qty2"`
	ReportType      string  `json:"report_type"`
	Segment         int     `json:"segment"`
	Symbol          string  `json:"symbol"`
	SymbolDesc      string  `json:"symbol_desc"`
	SymbolExch      string  `json:"symbol_exch"`
	TickSize        float64 `json:"tick_size"`
	TranSide        int     `json:"tran_side"`
	GttOcoInd       int     `json:"gtt_oco_ind"`
	CreateTime      string  `json:"create_time"`
	CreateTimeEpoch int     `json:"create_time_epoch"`
	OmsMsg          string  `json:"oms_msg"`
	LtpCh           float64 `json:"ltp_ch"`
	LtpChp          float64 `json:"ltp_chp"`
	Ltp             float64 `json:"ltp"`
}

type ModifyGTTOrderRequest struct {
	Id        string    `json:"id"`
	OrderInfo OrderInfo `json:"orderInfo"`
}

type CancelGTTOrderRequest struct {
	Id string `json:"id"`
}

// Smart Order

// CreateSmartOrderLimitRequest is the request body for create smart order limit.
type CreateSmartOrderLimitRequest struct {
	Symbol       string   `json:"symbol"`
	Side         int      `json:"side"`
	Qty          int      `json:"qty"`
	ProductType  string   `json:"productType"`
	LimitPrice   float64  `json:"limitPrice"`
	StopPrice    *float64 `json:"stopPrice,omitempty"`
	OrderType    int      `json:"orderType"`
	EndTime      int64    `json:"endTime"`
	Hpr          float64  `json:"hpr"`
	Lpr          float64  `json:"lpr"`
	Mpp          float64  `json:"mpp"`
	OnExp        int      `json:"onExp"`
	Type         *int     `json:"type,omitempty"`
	Validity     *string  `json:"validity,omitempty"`
	DisclosedQty *int     `json:"disclosedQty,omitempty"`
	OfflineOrder *bool    `json:"offlineOrder,omitempty"`
}

// CreateSmartOrderStepRequest is the request body for create smart order step.
type CreateSmartOrderStepRequest struct {
	Symbol       string   `json:"symbol"`
	Side         int      `json:"side"`
	Qty          int      `json:"qty"`
	ProductType  string   `json:"productType"`
	InitQty      int      `json:"initQty"`
	Avgqty       int      `json:"avgqty"`
	Avgdiff      float64  `json:"avgdiff"`
	Direction    int      `json:"direction"`
	LimitPrice   float64  `json:"limitPrice"`
	OrderType    int      `json:"orderType"`
	StartTime    int64    `json:"startTime"`
	EndTime      int64    `json:"endTime"`
	Hpr          float64  `json:"hpr"`
	Lpr          float64  `json:"lpr"`
	Mpp          float64  `json:"mpp"`
	Type         *int     `json:"type,omitempty"`
	StopPrice    *float64 `json:"stopPrice,omitempty"`
	Validity     *string  `json:"validity,omitempty"`
	DisclosedQty *int     `json:"disclosedQty,omitempty"`
	OfflineOrder *bool    `json:"offlineOrder,omitempty"`
}

// CreateSmartOrderSIPRequest is the request body for create smart order SIP.
type CreateSmartOrderSIPRequest struct {
	Symbol       string   `json:"symbol"`
	Side         int      `json:"side"`
	Amount       float64  `json:"amount"`
	ProductType  string   `json:"productType"`
	Freq         int      `json:"freq"`
	SipDay       int      `json:"sip_day"`
	ImdStart     bool     `json:"imd_start"`
	EndTime      int64    `json:"endTime"`
	Hpr          float64  `json:"hpr"`
	Lpr          float64  `json:"lpr"`
	StepUpFreq   int      `json:"step_up_freq"`
	StepUpAmount float64  `json:"step_up_amount"`
	Qty          *int     `json:"qty,omitempty"`
	Type         *int     `json:"type,omitempty"`
	LimitPrice   *float64 `json:"limitPrice,omitempty"`
	StopPrice    *float64 `json:"stopPrice,omitempty"`
	Validity     *string  `json:"validity,omitempty"`
	DisclosedQty *int     `json:"disclosedQty,omitempty"`
	OfflineOrder *bool    `json:"offlineOrder,omitempty"`
}

// CreateSmartOrderTrailRequest is the request body for create smart order trail.
type CreateSmartOrderTrailRequest struct {
	Symbol       string   `json:"symbol"`
	Side         int      `json:"side"`
	Qty          int      `json:"qty"`
	ProductType  string   `json:"productType"`
	StopPrice    float64  `json:"stopPrice"`
	JumpDiff     float64  `json:"jump_diff"`
	OrderType    int      `json:"orderType"`
	Mpp          float64  `json:"mpp"`
	Type         *int     `json:"type,omitempty"`
	LimitPrice   *float64 `json:"limitPrice,omitempty"`
	Validity     *string  `json:"validity,omitempty"`
	DisclosedQty *int     `json:"disclosedQty,omitempty"`
	OfflineOrder *bool    `json:"offlineOrder,omitempty"`
}

// ModifySmartOrderRequest is the request body for modify smart order.
// Use Qty, LimitPrice, EndTime, StopPrice, Hpr, Lpr for Limit/Step/Trail flows.
// Use Amount, SipDay, Hpr, Lpr for SIP flows.
type ModifySmartOrderRequest struct {
	FlowId      string  `json:"flowId"`
	Qty         int     `json:"qty,omitempty"`
	LimitPrice  float64 `json:"limitPrice,omitempty"`
	EndTime     int64   `json:"endTime,omitempty"`
	StopPrice   float64 `json:"stopPrice,omitempty"`
	Hpr         float64 `json:"hpr,omitempty"`
	Lpr         float64 `json:"lpr,omitempty"`
	Amount      float64 `json:"amount,omitempty"`  // SIP flow
	SipDay      int     `json:"sip_day,omitempty"` // SIP flow
	ProductType string  `json:"productType,omitempty"`
}

// FlowIdRequest is the request body for cancel/pause/resume smart order and activate/deactivate smart exit trigger.
type FlowIdRequest struct {
	FlowId string `json:"flowId"`
}

// GetSmartOrderBookFilter is optional filter for smart order book.
type GetSmartOrderBookFilter struct {
	Exchange    []string `json:"exchange,omitempty"`
	Side        []int    `json:"side,omitempty"`
	Flowtype    []int    `json:"flowtype,omitempty"`
	Product     []string `json:"product,omitempty"`
	MessageType []int    `json:"messageType,omitempty"`
	Search      string   `json:"search,omitempty"`
	SortBy      string   `json:"sort_by,omitempty"`
	OrdBy       int      `json:"ord_by,omitempty"`
	PageNo      int      `json:"page_no,omitempty"`
	PageSize    int      `json:"page_size,omitempty"`
}

// CreateSmartExitTriggerRequest is the request body for create smart exit trigger.
type CreateSmartExitTriggerRequest struct {
	Name       string  `json:"name"`
	Type       int     `json:"type"` // 1: Only Alert, 2: Exit with Alert, 3: Exit with Alert + Wait for Recovery
	ProfitRate float64 `json:"profitRate"`
	LossRate   float64 `json:"lossRate"`           // negative value
	WaitTime   *int    `json:"waitTime,omitempty"` // required for type 3, in minutes
}

// GetSmartExitTriggerFilter is optional filter for get smart exit trigger.
type GetSmartExitTriggerFilter struct {
	FlowId string `json:"flowId,omitempty"`
}

// UpdateSmartExitTriggerRequest is the request body for update smart exit trigger.
type UpdateSmartExitTriggerRequest struct {
	FlowId       string   `json:"flowId"`
	ProfitRate   *float64 `json:"profitRate,omitempty"`
	LossRate     *float64 `json:"lossRate,omitempty"`
	Type         *int     `json:"type,omitempty"`
	Name         *string  `json:"name,omitempty"`
	WaitTime     *int     `json:"waitTime,omitempty"`
	TriggerPrice *float64 `json:"triggerPrice,omitempty"`
	StopLoss     *float64 `json:"stopLoss,omitempty"`
	TakeProfit   *float64 `json:"takeProfit,omitempty"`
}

// Trade Operations

type ModifyOrderRequest struct {
	Id         string  `json:"id"`
	Qty        int     `json:"qty"`
	Type       int     `json:"type"`
	Side       int     `json:"side"`
	LimitPrice float64 `json:"limitPrice"`
}

// ModifyMultiOrderItem is one item in the PATCH /multi-order/sync body (array of these).
type ModifyMultiOrderItem struct {
	Id         int64   `json:"id"`
	Type       int     `json:"type"`
	LimitPrice float64 `json:"limitPrice"`
	Qty        int     `json:"qty"`
}

type CancelOrderRequest struct {
	Id string `json:"id"`
}

type ConvertPositionRequest struct {
	Symbol       string `json:"symbol"`
	PositionSide int    `json:"positionSide"`
	ConvertQty   int    `json:"convertQty"`
	ConvertFrom  string `json:"convertFrom"`
	ConvertTo    string `json:"convertTo"`
	Overnight    int    `json:"overnight"`
}

type ConvertPositionResponse struct {
	OrderResponse
	PositionDetails int `json:"positionDetails"`
}

type ExitPositionByProductTypeRequest struct {
	Segment     []int    `json:"segment,omitempty"`
	Side        []int    `json:"side,omitempty"`
	ProductType []string `json:"productType,omitempty"`
}

type CancelPendingOrdersRequest struct {
	Id                  string `json:"id,omitempty"`
	PendingOrdersCancel int    `json:"pending_orders_cancel"`
}

// Broker Config

type BrokerConfig struct {
	APIResponse
	MarketStatus []MarketStatus `json:"marketStatus"`
}

type MarketStatus struct {
	Exchange   int    `json:"exchange"`
	MarketType string `json:"market_type"`
	Segment    int    `json:"segment"`
	Status     string `json:"status"`
}

// Data
type HistoryRequest struct {
	Symbol     string `json:"symbol"`
	Resolution string `json:"resolution"`
	DateFormat string `json:"date_format"`
	RangeFrom  string `json:"range_from"`
	RangeTo    string `json:"range_to"`
	ContFlag   string `json:"cont_flag,omitempty"`
}

type HistoryResponse struct {
	APIResponse
	Candles [][]interface{} `json:"candles"`
}

type StockQuotesResponse struct {
	APIResponse
	Data []StockQuote `json:"d"`
}

type StockQuote struct {
	N string      `json:"n"`
	S string      `json:"s"`
	V QuoteValues `json:"v"`
}

type QuoteValues struct {
	Ch             float64 `json:"ch"`
	Chp            float64 `json:"chp"`
	Lp             float64 `json:"lp"`
	Spread         float64 `json:"spread"`
	Ask            float64 `json:"ask"`
	Bid            float64 `json:"bid"`
	OpenPrice      float64 `json:"open_price"`
	HighPrice      float64 `json:"high_price"`
	LowPrice       float64 `json:"low_price"`
	PrevClosePrice float64 `json:"prev_close_price"`
	Atp            float64 `json:"atp"`
	Volume         int     `json:"volume"`
	ShortName      string  `json:"short_name"`
	Exchange       string  `json:"exchange"`
	Description    string  `json:"description"`
	OriginalName   string  `json:"original_name"`
	Symbol         string  `json:"symbol"`
	FyToken        string  `json:"fyToken"`
	Tt             string  `json:"tt"`
}

type MarketDepthRequest struct {
	Symbol string `json:"symbol"`
	OHLCV  string `json:"ohlcv_flag,omitempty"`
}

type MarketDepthResponse struct {
	APIResponse
	Data map[string]MarketDepthSymbol `json:"d"`
}

type MarketDepthSymbol struct {
	TotalBuyQty  int          `json:"totalbuyqty"`
	TotalSellQty int          `json:"totalsellqty"`
	Ask          []DepthLevel `json:"ask"`
	Bids         []DepthLevel `json:"bids"`
	O            float64      `json:"o"`         // Open price
	H            float64      `json:"h"`         // High price
	L            float64      `json:"l"`         // Low price
	C            float64      `json:"c"`         // Close price
	Chp          float64      `json:"chp"`       // Change percentage
	TickSize     float64      `json:"tick_Size"` // Tick size
	Ch           float64      `json:"ch"`        // Change
	Ltq          int          `json:"ltq"`       // Last traded quantity
	Ltt          int64        `json:"ltt"`       // Last traded time
	Ltp          float64      `json:"ltp"`       // Last traded price
	V            int          `json:"v"`         // Volume
	Atp          float64      `json:"atp"`       // Average traded price
	LowerCkt     float64      `json:"lower_ckt"` // Lower circuit
	UpperCkt     float64      `json:"upper_ckt"` // Upper circuit
	Expiry       string       `json:"expiry"`    // Expiry date
	Oi           int          `json:"oi"`        // Open interest
	OiFlag       bool         `json:"oiflag"`    // Open interest flag
	Pdoi         int          `json:"pdoi"`      // Previous day open interest
	OiPercent    int          `json:"oipercent"` // Open interest percentage
}

type DepthLevel struct {
	Price  float64 `json:"price"`
	Volume int     `json:"volume"`
	Ord    int     `json:"ord"`
}

type OptionChainRequest struct {
	Symbol      string `json:"symbol"`
	StrikeCount int    `json:"strikecount,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
}

type OptionChainResponse struct {
	APIResponse
	Data OptionChainData `json:"data"`
}

type OptionChainData struct {
	CallOi       int                `json:"callOi"`
	ExpiryData   []ExpiryData       `json:"expiryData"`
	IndiavixData IndiavixData       `json:"indiavixData"`
	OptionsChain []OptionsChainItem `json:"optionsChain"`
	Message      string             `json:"message"`
	S            string             `json:"s"`
}

type ExpiryData struct {
	Date   string `json:"date"`
	Expiry string `json:"expiry"`
}

type IndiavixData struct {
	Ask         float64 `json:"ask"`
	Bid         float64 `json:"bid"`
	Description string  `json:"description"`
	ExSymbol    string  `json:"ex_symbol"`
	Exchange    string  `json:"exchange"`
	FyToken     string  `json:"fyToken"`
	Ltp         float64 `json:"ltp"`
	LtpCh       float64 `json:"ltpch"`
	LtpChp      float64 `json:"ltpchp"`
	OptionType  string  `json:"option_type"`
	StrikePrice float64 `json:"strike_price"`
	Symbol      string  `json:"symbol"`
	Volume      int     `json:"volume"`
}

type OptionsChainItem struct {
	Ask         float64 `json:"ask"`
	Bid         float64 `json:"bid"`
	Description string  `json:"description"`
	ExSymbol    string  `json:"ex_symbol"`
	Exchange    string  `json:"exchange"`
	FyToken     string  `json:"fyToken"`
	Ltp         float64 `json:"ltp"`
	LtpCh       float64 `json:"ltpch"`
	LtpChp      float64 `json:"ltpchp"`
	FP          float64 `json:"fp,omitempty"`
	FPCh        float64 `json:"fpch,omitempty"`
	FPChp       float64 `json:"fpchp,omitempty"`
	Oi          int     `json:"oi,omitempty"`
	OiCh        int     `json:"oich,omitempty"`
	OiChp       float64 `json:"oichp,omitempty"`
	OptionType  string  `json:"option_type"`
	PrevOi      int     `json:"prev_oi,omitempty"`
	StrikePrice float64 `json:"strike_price,omitempty"`
	Symbol      string  `json:"symbol,omitempty"`
	Volume      int     `json:"volume,omitempty"`
}

// Web Socket
type DataSocketRequest struct {
	Symbols  []string
	DataType string
	LiteMode bool
}

type OrderSocketRequest struct {
	TradeOperations []string
}

// Alerts

type AlertsResponse struct {
	APIResponse
	Data map[string]AlertItem `json:"data"`
}

type AlertItem struct {
	FyToken string       `json:"fyToken"`
	Symbol  string       `json:"symbol"`
	Alert   AlertDetails `json:"alert"`
}

type AlertDetails struct {
	ComparisonType string  `json:"comparisonType"`
	Condition      string  `json:"condition"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Value          float64 `json:"value"`
	TriggeredAt    string  `json:"triggeredAt"`
	CreatedAt      string  `json:"createdAt"`
	Status         int     `json:"status"`
}

type AlertRequest struct {
	Symbol         string  `json:"symbol"`
	Agent          string  `json:"agent"`
	AlertType      int     `json:"alert-type"`
	ComparisonType string  `json:"comparisonType"`
	Condition      string  `json:"condition"`
	Value          float64 `json:"value"`
	Name           string  `json:"name"`
	AlertId        string  `json:"alertId,omitempty"`
}

type Alert struct {
	AlertId string `json:"alert_id"`
	Symbol  string `json:"symbol"`
}
