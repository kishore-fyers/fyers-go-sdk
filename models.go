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
	Leg1 Leg1  `json:"leg1"`
	Leg2 *Leg2 `json:"leg2,omitempty"`
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
// Required: symbol, side, qty, productType, limitPrice, endTime, orderType, onExp.
// Optional: stopPrice (required when orderType=4), hpr, lpr, mpp (default 0); mpp valid 0–3 or -1.
// productType: "CNC" | "MARGIN" | "INTRADAY" | "MTF". side: 1=Buy, -1=Sell. orderType: 1=Limit, 4=Stop-Limit. onExp: 1=Cancel, 2=Market.
type CreateSmartOrderLimitRequest struct {
	Symbol       string   `json:"symbol"`                 // Required. e.g. "NSE:SBIN-EQ"
	Side         int      `json:"side"`                   // Required. 1=Buy, -1=Sell
	Qty          int      `json:"qty"`                    // Required. Min 1, multiple of lot size
	ProductType  string   `json:"productType"`            // Required. CNC, MARGIN, INTRADAY, MTF
	LimitPrice   float64  `json:"limitPrice"`             // Required. Min 0.01
	EndTime      int64    `json:"endTime"`                // Required. Unix timestamp (epoch)
	OrderType    int      `json:"orderType"`              // Required. 1=Limit, 4=Stop-Limit
	OnExp        int      `json:"onExp"`                  // Required. 1=Cancel, 2=Market
	StopPrice    *float64 `json:"stopPrice,omitempty"`    // Optional. Required when orderType=4. Default 0
	Hpr          *float64 `json:"hpr,omitempty"`          // Optional. 0=no upper limit. Order executes only below this price
	Lpr          *float64 `json:"lpr,omitempty"`          // Optional. 0=no lower limit. Order executes only above this price
	Mpp          *float64 `json:"mpp,omitempty"`          // Optional. 0=no market protection. Valid: 0–3 or -1 (disabled)
	Type         *int     `json:"type,omitempty"`
	Validity     *string  `json:"validity,omitempty"`
	DisclosedQty *int     `json:"disclosedQty,omitempty"`
	OfflineOrder *bool    `json:"offlineOrder,omitempty"`
}

// CreateSmartOrderStepRequest is the request body for create smart order step.
// Required: symbol, side, qty, productType, orderType, avgqty, avgdiff, direction, startTime, endTime.
// Conditional: limitPrice (required if orderType=1). Optional: initQty (default 0), hpr, lpr, mpp (default 0; mpp valid 0–3 or -1).
// productType: "CNC"|"MARGIN"|"INTRADAY"|"MTF". side: 1=Buy, -1=Sell. orderType: 1=Limit, 2=Market. direction: 1=avg on price decrease, -1=avg on price increase.
// qty must be >= initQty + avgqty. avgdiff Min: 0.01.
type CreateSmartOrderStepRequest struct {
	Symbol       string   `json:"symbol"`                // Required. e.g. "NSE:SBIN-EQ"
	Side         int      `json:"side"`                  // Required. 1=Buy, -1=Sell
	Qty          int      `json:"qty"`                   // Required. Total qty; must be >= initQty + avgqty
	ProductType  string   `json:"productType"`           // Required. CNC, MARGIN, INTRADAY, MTF
	OrderType    int      `json:"orderType"`             // Required. 1=Limit, 2=Market
	Avgqty       int      `json:"avgqty"`                // Required. Qty at each step (Min: 1)
	Avgdiff      float64  `json:"avgdiff"`               // Required. Price diff between steps (Min: 0.01)
	Direction    int      `json:"direction"`             // Required. 1=avg on decrease, -1=avg on increase
	StartTime    int64    `json:"startTime"`             // Required. Unix timestamp (epoch)
	EndTime      int64    `json:"endTime"`               // Required. Unix timestamp (epoch)
	LimitPrice   *float64 `json:"limitPrice,omitempty"`  // Conditional. Required if orderType=1. Default 0
	InitQty      *int     `json:"initQty,omitempty"`     // Optional. Qty to place immediately. Default 0 (no initial order)
	Hpr          *float64 `json:"hpr,omitempty"`         // Optional. 0=no upper limit. Execute only below this price
	Lpr          *float64 `json:"lpr,omitempty"`         // Optional. 0=no lower limit. Execute only above this price
	Mpp          *float64 `json:"mpp,omitempty"`         // Optional. 0=no market protection. Valid: 0–3 or -1
	Type         *int     `json:"type,omitempty"`
	StopPrice    *float64 `json:"stopPrice,omitempty"`
	Validity     *string  `json:"validity,omitempty"`
	DisclosedQty *int     `json:"disclosedQty,omitempty"`
	OfflineOrder *bool    `json:"offlineOrder,omitempty"`
}

// CreateSmartOrderSIPRequest is the request body for create smart order SIP (Equity only).
// Required: symbol, productType, freq, sip_day, and at least one of qty or amount.
// Conditional: sip_time (required if freq=1 Daily). Unix timestamp within market hours.
// Optional: side, imd_start, endTime, hpr, lpr, step_up_freq (3|5), step_up_qty, step_up_amount, exp_qty.
// productType: "CNC"|"MTF". freq: 1=Daily, 2, 3, 6. sip_day: 1–28.
type CreateSmartOrderSIPRequest struct {
	Symbol       string   `json:"symbol"`                 // Required. Equity only, e.g. "NSE:SBIN-EQ"
	ProductType  string   `json:"productType"`            // Required. CNC, MTF
	Freq         int      `json:"freq"`                    // Required. 1=Daily, 2, 3, 6
	SipDay       int      `json:"sip_day"`                 // Required. Day of month (1–28)
	Qty          *int     `json:"qty,omitempty"`           // At least one of Qty or Amount required. Max 999999
	Amount       *float64 `json:"amount,omitempty"`       // At least one of Qty or Amount required
	SipTime      *int64   `json:"sip_time,omitempty"`      // Conditional. Required if freq=1. Unix timestamp (market hours)
	ImdStart     *bool    `json:"imd_start,omitempty"`     // Optional. true=start now, false=wait for schedule
	EndTime      *int64   `json:"endTime,omitempty"`       // Optional. 0=no end. Unix timestamp when SIP ends
	Hpr          *float64 `json:"hpr,omitempty"`           // Optional. Skips SIP if price above this
	Lpr          *float64 `json:"lpr,omitempty"`           // Optional. Skips SIP if price below this
	StepUpFreq   *int     `json:"step_up_freq,omitempty"`  // Optional. 3 or 5. 0=no step-up
	StepUpQty    *int     `json:"step_up_qty,omitempty"`   // Optional. Qty increase per step-up (Max 999999)
	StepUpAmount *float64 `json:"step_up_amount,omitempty"` // Optional. Amount increase per step-up
	ExpQty       *int     `json:"exp_qty,omitempty"`       // Optional. Qty for expiry/final SIP order (Max 999999)
	Side         *int     `json:"side,omitempty"`         // Optional. 1=Buy, -1=Sell (not in spec; kept for compatibility)
	Type         *int     `json:"type,omitempty"`
	LimitPrice   *float64 `json:"limitPrice,omitempty"`
	StopPrice    *float64 `json:"stopPrice,omitempty"`
	Validity     *string  `json:"validity,omitempty"`
	DisclosedQty *int     `json:"disclosedQty,omitempty"`
	OfflineOrder *bool    `json:"offlineOrder,omitempty"`
}

// CreateSmartOrderTrailRequest is the request body for create smart order trail.
// Required: symbol, side, qty, productType, orderType, stopPrice, jump_diff.
// Optional: limitPrice (required if orderType=1), target_price, mpp (default 0; valid 0–3 or -1).
// productType: "CNC"|"MARGIN"|"INTRADAY"|"MTF". side: 1=Buy, -1=Sell. orderType: 1=Limit, 2=Market. jump_diff Min: 0.2.
type CreateSmartOrderTrailRequest struct {
	Symbol       string   `json:"symbol"`                // Required. e.g. "NSE:SBIN-EQ"
	Side         int      `json:"side"`                  // Required. 1=Buy, -1=Sell
	Qty          int      `json:"qty"`                   // Required. Min 1, multiple of lot size
	ProductType  string   `json:"productType"`           // Required. CNC, MARGIN, INTRADAY, MTF
	OrderType    int      `json:"orderType"`             // Required. 1=Limit, 2=Market
	StopPrice    float64  `json:"stopPrice"`             // Required. Initial stop/trigger price, > 0
	JumpDiff     float64  `json:"jump_diff"`             // Required. Jump price — stop trails by this (Min: 0.2)
	LimitPrice   *float64 `json:"limitPrice,omitempty"`   // Optional. Required if orderType=1. Default 0 (market)
	TargetPrice  *float64 `json:"target_price,omitempty"` // Optional. Default 0 (no target). If set, must be > current LTP
	Mpp          *float64 `json:"mpp,omitempty"`         // Optional. 0=no market protection. Valid: 0–3 or -1
	Type         *int     `json:"type,omitempty"`
	Validity     *string  `json:"validity,omitempty"`
	DisclosedQty *int     `json:"disclosedQty,omitempty"`
	OfflineOrder *bool    `json:"offlineOrder,omitempty"`
}

// ModifySmartOrderRequest is the request body for modify smart order.
// Required: flowId. Optional fields depend on flowtype (4=Limit, 6=Trail, 3=Step, 7=SIP).
type ModifySmartOrderRequest struct {
	FlowId string `json:"flowId"` // Required. Unique identifier of the Smart Order to modify

	// Limit (flowtype 4): qty, limitPrice, stopPrice, endTime, hpr, lpr, mpp, onExp
	// Trail (flowtype 6): qty, limitPrice, stopPrice, jump_diff, target_price, unsetTargetPrice, mpp
	// Step (flowtype 3): qty, startTime, endTime, hpr, lpr, mpp, avgqty, avgdiff, initQty, limitPrice, direction
	// SIP (flowtype 7): qty, amount, hpr, lpr, sip_day, sip_time, step_up_amount, step_up_qty, exp_qty, exp_amount
	Qty              int     `json:"qty,omitempty"`
	LimitPrice       float64 `json:"limitPrice,omitempty"`
	StopPrice        float64 `json:"stopPrice,omitempty"`
	EndTime          int64   `json:"endTime,omitempty"`
	StartTime        int64   `json:"startTime,omitempty"`        // Step
	Hpr              float64 `json:"hpr,omitempty"`
	Lpr              float64 `json:"lpr,omitempty"`
	Mpp              float64 `json:"mpp,omitempty"`
	OnExp            int     `json:"onExp,omitempty"`           // Limit. 1=Cancel, 2=Market
	JumpDiff         float64 `json:"jump_diff,omitempty"`        // Trail
	TargetPrice      float64 `json:"target_price,omitempty"`    // Trail
	UnsetTargetPrice bool    `json:"unsetTargetPrice,omitempty"` // Trail. true to remove target_price
	Avgqty           int     `json:"avgqty,omitempty"`          // Step
	Avgdiff          float64 `json:"avgdiff,omitempty"`          // Step
	InitQty          int     `json:"initQty,omitempty"`          // Step (before order starts)
	Direction        int     `json:"direction,omitempty"`        // Step. 1=price drop, -1=price rise
	Amount           float64 `json:"amount,omitempty"`           // SIP
	SipDay           int     `json:"sip_day,omitempty"`         // SIP
	SipTime          int64   `json:"sip_time,omitempty"`         // SIP (daily/custom freq)
	StepUpAmount     float64 `json:"step_up_amount,omitempty"`   // SIP
	StepUpQty        int     `json:"step_up_qty,omitempty"`      // SIP
	ExpQty           int     `json:"exp_qty,omitempty"`          // SIP
	ExpAmount        float64 `json:"exp_amount,omitempty"`       // SIP
	ProductType      string  `json:"productType,omitempty"`
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

// CreateSmartExitTriggerRequest is the request body for create smart exit trigger (monitors overall position P&L).
// Required: name. At least one of profitRate or lossRate required.
// Optional: type (default 1). Conditional: waitTime required if type=3 (Min: 1, Max: 60 minutes).
// Type: 1=Only Alert (notification only), 2=Exit with Alert (immediate exit), 3=Exit with Alert + Wait for Recovery (waits waitTime then exits).
// profitRate/lossRate range: Min -1,00,00,000, Max 1,00,00,000.
type CreateSmartExitTriggerRequest struct {
	Name       string   `json:"name"`                 // Required. Unique name for the trigger
	Type       *int     `json:"type,omitempty"`       // Optional. Default 1. 1=Only Alert, 2=Exit+Alert, 3=Exit+Alert+Wait
	ProfitRate *float64 `json:"profitRate,omitempty"` // Conditional. At least one with lossRate. Book profit (positive) or minimize loss (negative)
	LossRate   *float64 `json:"lossRate,omitempty"`    // Conditional. At least one with profitRate. Max loss (negative) or min profit (positive)
	WaitTime   *int     `json:"waitTime,omitempty"`    // Conditional. Required if type=3. Wait time in minutes (1–60) before exiting
}

// GetSmartExitTriggerFilter is optional filter for get smart exit trigger.
type GetSmartExitTriggerFilter struct {
	FlowId string `json:"flowId,omitempty"`
}

// UpdateSmartExitTriggerRequest is the request body for modify smart exit trigger (PUT).
// Required: flowId. Optional: name, profitRate, lossRate, type, waitTime.
// Either profitRate or lossRate should be provided when updating targets.
// Type: 1=Only Alert, 2=Exit with Alert, 3=Exit+Alert+Wait (waitTime required if type=3). waitTime: 0–60 minutes.
// profitRate/lossRate range: Min -1,00,00,000, Max 1,00,00,000.
type UpdateSmartExitTriggerRequest struct {
	FlowId       string   `json:"flowId"`                 // Required. Unique identifier of the smart exit to update
	Name         *string  `json:"name,omitempty"`          // Optional. Unique name for the trigger
	ProfitRate   *float64 `json:"profitRate,omitempty"`    // Optional. Book profit (positive) or minimize loss (negative)
	LossRate     *float64 `json:"lossRate,omitempty"`      // Optional. Max loss (negative) or min profit (positive)
	Type         *int     `json:"type,omitempty"`          // Optional. Default 1. 1=Only Alert, 2=Exit+Alert, 3=Exit+Alert+Wait
	WaitTime     *int     `json:"waitTime,omitempty"`      // Optional. Required if type=3. Minutes (0–60). Default 0
	TriggerPrice *float64 `json:"triggerPrice,omitempty"`  // Optional (API may support)
	StopLoss     *float64 `json:"stopLoss,omitempty"`      // Optional (API may support)
	TakeProfit   *float64 `json:"takeProfit,omitempty"`    // Optional (API may support)
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
