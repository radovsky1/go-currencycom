package go_currencycom

import "context"

type ExchangeInfoService struct {
	c *Client
}

type ExchangeFilter struct {
	FilterType string `json:"filterType"`
	MinPrice   string `json:"minPrice"`
	MaxPrice   string `json:"maxPrice"`
}

type RateLimit struct {
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	RateLimitType string `json:"rateLimitType"`
}

type ExchangeSymbolInfo struct {
	AssetType          string           `json:"assetType"`
	BaseAsset          string           `json:"baseAsset"`
	BaseAssetPrecision int              `json:"baseAssetPrecision"`
	Country            string           `json:"country"`
	ExchangeFee        float64          `json:"exchangeFee"`
	Filters            []ExchangeFilter `json:"filters"`
	Industry           string           `json:"industry"`
	LongRate           float64          `json:"longRate"`
	MakerFee           float64          `json:"makerFee"`
	MarketModes        []string         `json:"marketModes"`
	MarketType         string           `json:"marketType"`
	MaxSLGap           float64          `json:"maxSLGap"`
	MaxTPGap           float64          `json:"maxTPGap"`
	MinSLGap           float64          `json:"minSLGap"`
	MinTPGap           float64          `json:"minTPGap"`
	Name               string           `json:"name"`
	OrderTypes         []OrderType      `json:"orderTypes"`
	QuoteAsset         string           `json:"quoteAsset"`
	QuoteAssetID       string           `json:"quoteAssetId"`
	QuotePrecision     int              `json:"quotePrecision"`
	Sector             string           `json:"sector"`
	ShortRate          float64          `json:"shortRate"`
	Status             string           `json:"status"`
	SwapChargeInterval int64            `json:"swapChargeInterval"`
	Symbol             string           `json:"symbol"`
	TakerFee           float64          `json:"takerFee"`
	TickSize           float64          `json:"tickSize"`
	TickValue          float64          `json:"tickValue"`
	TradingFee         float64          `json:"tradingFee"`
	TradingHours       string           `json:"tradingHours"`
}

type ExchangeInfo struct {
	ExchangeFilters []ExchangeFilter     `json:"exchangeFilters"`
	RateLimits      []RateLimit          `json:"rateLimits"`
	ServerTime      int64                `json:"serverTime"`
	Symbols         []ExchangeSymbolInfo `json:"symbols"`
	Timezone        string               `json:"timezone"`
}

func (s *ExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeInfo, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v2/exchange_info",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ExchangeInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
