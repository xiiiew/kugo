package kugo

import "github.com/shopspring/decimal"

// URI
const (
	UriSymbols        = "/api/v2/symbols"
	UriAccounts       = "/api/v1/accounts"
	UriSpotOrders     = "/api/v1/orders"
	UriSpotOrderFills = "/api/v1/fills"

	UriFutureAccount = "/api/v1/account-overview"
	UriFutureOrders  = "/api/v1/orders"
)

type BaseResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// SymbolsResponse Response of /api/v2/symbols
type SymbolsResponse struct {
	BaseResponse
	Data []SymbolsData `json:"data"`
}
type SymbolsData struct {
	Symbol          string          `json:"symbol"`
	Name            string          `json:"name"`
	BaseCurrency    string          `json:"baseCurrency"`
	QuoteCurrency   string          `json:"quoteCurrency"`
	FeeCurrency     string          `json:"feeCurrency"`
	Market          string          `json:"market"`
	BaseMinSize     decimal.Decimal `json:"baseMinSize"`
	QuoteMinSize    decimal.Decimal `json:"quoteMinSize"`
	BaseMaxSize     decimal.Decimal `json:"baseMaxSize"`
	QuoteMaxSize    decimal.Decimal `json:"quoteMaxSize"`
	BaseIncrement   decimal.Decimal `json:"baseIncrement"`
	QuoteIncrement  decimal.Decimal `json:"quoteIncrement"`
	PriceIncrement  decimal.Decimal `json:"priceIncrement"`
	PriceLimitRate  decimal.Decimal `json:"priceLimitRate"`
	MinFunds        decimal.Decimal `json:"minFunds"`
	IsMarginEnabled bool            `json:"isMarginEnabled"`
	EnableTrading   bool            `json:"enableTrading"`
}

// AccountsResponse Response of /api/v2/accounts
type AccountsResponse struct {
	BaseResponse
	Data []AccountsData `json:"data"`
}
type AccountsData struct {
	Id        string          `json:"id"`
	Currency  string          `json:"currency"`
	Type      string          `json:"type"`
	Balance   decimal.Decimal `json:"balance"`
	Available decimal.Decimal `json:"available"`
	Holds     decimal.Decimal `json:"holds"`
}

// SpotOrdersRequest Request of /api/v1/orders
type SpotOrdersRequest struct {
	ClientOid   string          `json:"clientOid,omitempty"`
	Side        string          `json:"side,omitempty"`   // buy or sell
	Symbol      string          `json:"symbol,omitempty"` // e.g. BTC-USDT
	Type        string          `json:"type,omitempty"`   // limit or market
	Remark      string          `json:"remark,omitempty"`
	Stp         string          `json:"stp,omitempty"`
	TradeType   string          `json:"tradeType,omitempty"`
	Price       decimal.Decimal `json:"price,omitempty"`
	Size        decimal.Decimal `json:"size,omitempty"`
	TimeInForce string          `json:"timeInForce,omitempty"` // GTC, GTT, IOC or FOK
	CancelAfter int64           `json:"cancelAfter,omitempty"`
	PostOnly    bool            `json:"postOnly,omitempty"`
	Hidden      bool            `json:"hidden,omitempty"`
	Iceberg     bool            `json:"iceberg,omitempty"`
	VisibleSize string          `json:"visibleSize,omitempty"`
	Funds       string          `json:"funds,omitempty"` // MARKET order only, It is required that you use one of the two parameters, size or funds.
}

// SpotOrderResponse Response of /api/v1/orders
type SpotOrderResponse struct {
	BaseResponse
	Data SpotOrderData `json:"data"`
}
type SpotOrderData struct {
	OrderId string `json:"orderId"`
}

// SpotOrderFillsRequest Request of /api/v1/fills
type SpotOrderFillsRequest struct {
	OrderId   string `json:"orderId,omitempty"` // If you specify orderId, other parameters can be ignored
	Symbol    string `json:"symbol,omitempty"`
	Side      string `json:"side,omitempty"`      // buy or sell
	Type      string `json:"type,omitempty"`      // limit, market, limit_stop or market_stop
	StartAt   int64  `json:"startAt,omitempty"`   // Start time (millisecond)
	EndAt     int64  `json:"endAt,omitempty"`     // End time (millisecond)
	TradeType string `json:"tradeType,omitempty"` // TRADE（Spot Trading）, MARGIN_TRADE (Margin Trading), TRADE as default.
}

// SpotOrderFillsResponse Response of /api/v1/fills
type SpotOrderFillsResponse struct {
	BaseResponse
	Data SpotOrderFillsData `json:"data"`
}
type SpotOrderFillsData struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	TotalNum    int `json:"totalNum"`
	TotalPage   int `json:"totalPage"`
	Items       []struct {
		Symbol         string          `json:"symbol"`
		TradeId        string          `json:"tradeId"`
		OrderId        string          `json:"orderId"`
		CounterOrderId string          `json:"counterOrderId"`
		Side           string          `json:"side"`
		Liquidity      string          `json:"liquidity"`
		ForceTaker     bool            `json:"forceTaker"`
		Price          decimal.Decimal `json:"price"`
		Size           decimal.Decimal `json:"size"`
		Funds          decimal.Decimal `json:"funds"`
		Fee            decimal.Decimal `json:"fee"`
		FeeRate        decimal.Decimal `json:"feeRate"`
		FeeCurrency    string          `json:"feeCurrency"`
		Stop           string          `json:"stop"`
		Type           string          `json:"type"`
		CreatedAt      int64           `json:"createdAt"`
		TradeType      string          `json:"tradeType"`
	} `json:"items"`
}

// FutureAccountResponse Response of /api/v1/account-overview
type FutureAccountResponse struct {
	BaseResponse
	Data FutureAccountData `json:"data"`
}
type FutureAccountData struct {
	AccountEquity    decimal.Decimal `json:"accountEquity"` // Account equity = marginBalance + Unrealised PNL
	UnrealisedPNL    decimal.Decimal `json:"unrealisedPNL"` // Unrealised profit and loss
	MarginBalance    decimal.Decimal `json:"marginBalance"` // Margin balance = positionMargin + orderMargin + frozenFunds + availableBalance - unrealisedPNL
	PositionMargin   decimal.Decimal `json:"positionMargin"`
	OrderMargin      decimal.Decimal `json:"orderMargin"`
	FrozenFunds      decimal.Decimal `json:"frozenFunds"` // Frozen funds for withdrawal and out-transfer
	AvailableBalance decimal.Decimal `json:"availableBalance"`
	Currency         string          `json:"currency"`
}

// FutureOrderRequest Request of /api/v1/orders
type FutureOrderRequest struct {
	ClientOid     string          `json:"clientOid,omitempty"`
	Side          string          `json:"side,omitempty"`   // buy or sell
	Symbol        string          `json:"symbol,omitempty"` // e.g. BTC-USDT
	Type          string          `json:"type,omitempty"`   // limit or market
	Leverage      decimal.Decimal `json:"leverage,omitempty"`
	Remark        string          `json:"remark,omitempty"`
	Stop          string          `json:"stop,omitempty"`
	StopPriceType string          `json:"stopPriceType,omitempty"` // TP, IP or MP
	StopPrice     decimal.Decimal `json:"stopPrice,omitempty"`
	ReduceOnly    bool            `json:"reduceOnly,omitempty"`
	CloseOrder    bool            `json:"closeOrder,omitempty"`
	ForceHold     bool            `json:"forceHold,omitempty"`
	Price         decimal.Decimal `json:"price,omitempty"`
	Size          int             `json:"size,omitempty"`        // Cont
	TimeInForce   string          `json:"timeInForce,omitempty"` // GTC, GTT, IOC or FOK
	PostOnly      bool            `json:"postOnly,omitempty"`
	Hidden        bool            `json:"hidden,omitempty"`
	Iceberg       bool            `json:"iceberg,omitempty"`
	VisibleSize   string          `json:"visibleSize,omitempty"`
}

// FutureOrderResponse Response of /api/v1/orders
type FutureOrderResponse struct {
	BaseResponse
	Data FutureOrderData `json:"data"`
}
type FutureOrderData struct {
	OrderId string `json:"orderId"`
}
