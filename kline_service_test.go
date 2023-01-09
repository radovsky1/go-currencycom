package go_currencycom

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type klineServiceTestSuite struct {
	baseTestSuite
}

func TestKlineService(t *testing.T) {
	suite.Run(t, new(klineServiceTestSuite))
}

func (s *klineServiceTestSuite) TestKline() {
	data := []byte(`[
        [
            1499040000000,
            "0.01634790",
            "0.80000000",
            "0.01575800",
            "0.01577100",
            "148976.11427815"
        ],
        [
            1499040000001,
            "0.01634790",
            "0.80000000",
            "0.01575800",
            "0.01577101",
            "148976.11427815"
        ]
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTC/USD_LEVERAGE"
	interval := "1m"
	limit := 10
	startTime := int64(1499040000000)
	endTime := int64(1499040000001)
	priceType := "spot"
	ktype := "kline"
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol":    symbol,
			"interval":  interval,
			"limit":     limit,
			"startTime": startTime,
			"endTime":   endTime,
			"priceType": priceType,
			"type":      ktype,
		})
		s.assertRequestEqual(e, r)
	})
	klines, err := s.client.NewKlinesService().
		Symbol(symbol).
		Interval(interval).
		Limit(limit).
		StartTime(startTime).EndTime(endTime).
		PriceType(priceType).Type(ktype).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(klines, 2)
	ansKlines := []*Kline{
		{
			OpenTime: 1499040000000,
			Open:     "0.01634790",
			High:     "0.80000000",
			Low:      "0.01575800",
			Close:    "0.01577100",
			Volume:   "148976.11427815",
		},
		{
			OpenTime: 1499040000001,
			Open:     "0.01634790",
			High:     "0.80000000",
			Low:      "0.01575800",
			Close:    "0.01577101",
			Volume:   "148976.11427815",
		},
	}
	for i := range klines {
		s.assertKlineEqual(klines[i], ansKlines[i])
	}
}

func (s *klineServiceTestSuite) assertKlineEqual(e, a *Kline) {
	r := s.r()
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.Open, a.Open, "Open")
	r.Equal(e.High, a.High, "High")
	r.Equal(e.Low, a.Low, "Low")
	r.Equal(e.Close, a.Close, "Close")
	r.Equal(e.Volume, a.Volume, "Volume")
}
