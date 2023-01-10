package go_currencycom

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type exchangeInfoServiceTestSuite struct {
	baseTestSuite
}

func TestExchangeInfoService(t *testing.T) {
	suite.Run(t, new(exchangeInfoServiceTestSuite))
}

func (s *exchangeInfoServiceTestSuite) TestGetExchangeInfo() {
	data := []byte(`{
		"timezone":"UTC",
		"serverTime":1628193845310,
		"rateLimits":[
		],
		"exchangeFilters":[
		],
		"symbols":[
			{
				"symbol":"EVK",
				"name":"Evonik",
				"status":"BREAK",
				"baseAsset":"EVK",
				"baseAssetPrecision":3,
				"quoteAsset":"EUR",
				"quoteAssetId":"EUR",
				"quotePrecision":3,
				"orderTypes":[
					"LIMIT",
					"MARKET"
				],
				"filters":[
					{
						"filterType":"LOT_SIZE",
						"minQty":"1",
						"maxQty":"27000",
						"stepSize":"1"
					},
					{
						"filterType":"MIN_NOTIONAL",
						"minNotional":"29"
					}
				],
				"marketModes":[
					"REGULAR"
				],
				"marketType":"SPOT",
				"country":"DE",
				"sector":"Basic Materials",
				"industry":"Diversified Chemicals",
				"tradingHours":"UTC; Mon 07:02 - 15:30; Tue 07:02 - 15:30; Wed 07:02 - 15:30; Thu 07:02 - 15:30; Fri 07:02 - 15:30",
				"tickSize":0.005,
				"tickValue":0.14475,
				"exchangeFee":0.05
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	exchangeInfo, err := s.client.NewExchangeInfoService().Do(newContext())
	r := s.r()
	r.NoError(err)
	r.NotNil(exchangeInfo)
	e := &ExchangeInfo{
		Timezone:        "UTC",
		ServerTime:      1628193845310,
		RateLimits:      []RateLimit{},
		ExchangeFilters: []ExchangeFilter{},
		Symbols: []ExchangeSymbolInfo{
			{
				Symbol:             "EVK",
				Name:               "Evonik",
				Status:             "BREAK",
				BaseAsset:          "EVK",
				BaseAssetPrecision: 3,
				QuoteAsset:         "EUR",
				QuoteAssetID:       "EUR",
				QuotePrecision:     3,
				OrderTypes:         []OrderType{OrderTypeLimit, OrderTypeMarket},
				Filters: []ExchangeSymbolFilter{
					{
						FilterType: "LOT_SIZE",
						MinQty:     "1",
						MaxQty:     "27000",
						StepSize:   "1",
					},
					{
						FilterType:  "MIN_NOTIONAL",
						MinNotional: "29",
					},
				},
				MarketModes:  []string{"REGULAR"},
				MarketType:   "SPOT",
				Country:      "DE",
				Sector:       "Basic Materials",
				Industry:     "Diversified Chemicals",
				TradingHours: "UTC; Mon 07:02 - 15:30; Tue 07:02 - 15:30; Wed 07:02 - 15:30; Thu 07:02 - 15:30; Fri 07:02 - 15:30",
				TickSize:     0.005,
				TickValue:    0.14475,
				ExchangeFee:  0.05,
			},
		},
	}
	s.assertExchangeInfoEqual(e, exchangeInfo)
}

func (s *exchangeInfoServiceTestSuite) assertExchangeInfoEqual(e, a *ExchangeInfo) {
	r := s.r()
	r.Equal(e.Timezone, a.Timezone, "Timezone")
	r.Equal(e.ServerTime, a.ServerTime, "ServerTime")
	r.Len(a.RateLimits, len(e.RateLimits), "RateLimits")
	for i := range e.RateLimits {
		s.assertRateLimitEqual(&e.RateLimits[i], &a.RateLimits[i])
	}
	r.Len(a.ExchangeFilters, len(e.ExchangeFilters), "ExchangeFilters")
	for i := range e.ExchangeFilters {
		s.assertExchangeFilterEqual(&e.ExchangeFilters[i], &a.ExchangeFilters[i])
	}
	r.Len(a.Symbols, len(e.Symbols), "Symbols")
	for i := range e.Symbols {
		s.assertSymbolEqual(&e.Symbols[i], &a.Symbols[i])
	}
}

func (s *exchangeInfoServiceTestSuite) assertRateLimitEqual(e, a *RateLimit) {
	r := s.r()
	r.Equal(e.RateLimitType, a.RateLimitType, "RateLimitType")
	r.Equal(e.Interval, a.Interval, "Interval")
	r.Equal(e.IntervalNum, a.IntervalNum, "IntervalNum")
	r.Equal(e.Limit, a.Limit, "Limit")
}

func (s *exchangeInfoServiceTestSuite) assertExchangeFilterEqual(e, a *ExchangeFilter) {
	r := s.r()
	r.Equal(e.FilterType, a.FilterType, "FilterType")
	r.Equal(e.MinPrice, a.MinPrice, "MinPrice")
	r.Equal(e.MaxPrice, a.MaxPrice, "MaxPrice")
}

func (s *exchangeInfoServiceTestSuite) assertExchangeSymbolFilterEqual(e, a *ExchangeSymbolFilter) {
	r := s.r()
	r.Equal(e.FilterType, a.FilterType, "FilterType")
	r.Equal(e.MinQty, a.MinQty, "MinQty")
	r.Equal(e.MaxQty, a.MaxQty, "MaxQty")
	r.Equal(e.StepSize, a.StepSize, "StepSize")
}

func (s *exchangeInfoServiceTestSuite) assertSymbolEqual(e, a *ExchangeSymbolInfo) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Name, a.Name, "Name")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.BaseAsset, a.BaseAsset, "BaseAsset")
	r.Equal(e.BaseAssetPrecision, a.BaseAssetPrecision, "BaseAssetPrecision")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.QuotePrecision, a.QuotePrecision, "QuotePrecision")
	r.Len(a.OrderTypes, len(e.OrderTypes), "OrderTypes")
	for i := range e.OrderTypes {
		r.Equal(e.OrderTypes[i], a.OrderTypes[i], "OrderTypes[%d]", i)
	}
	r.Len(a.Filters, len(e.Filters), "Filters")
	for i := range e.Filters {
		s.assertExchangeSymbolFilterEqual(&e.Filters[i], &a.Filters[i])
	}
	r.Len(a.MarketModes, len(e.MarketModes), "MarketModes")
	for i := range e.MarketModes {
		r.Equal(e.MarketModes[i], a.MarketModes[i], "MarketModes[%d]", i)
	}
	r.Equal(e.MarketType, a.MarketType, "MarketType")
	r.Equal(e.Country, a.Country, "Country")
	r.Equal(e.Sector, a.Sector, "Sector")
	r.Equal(e.Industry, a.Industry, "Industry")
	r.Equal(e.TradingHours, a.TradingHours, "TradingHours")
	r.Equal(e.TickSize, a.TickSize, "TickSize")
	r.Equal(e.TickValue, a.TickValue, "TickValue")
	r.Equal(e.ExchangeFee, a.ExchangeFee, "ExchangeFee")
}
