package go_currencycom

import (
	"context"
	"fmt"
	"net/http"
)

type KlinesService struct {
	c         *Client
	startTime *int64
	endTime   *int64
	interval  string
	limit     *int
	symbol    string
	priceType *string
	ktype     *string
}

func (s *KlinesService) StartTime(startTime int64) *KlinesService {
	s.startTime = &startTime
	return s
}

func (s *KlinesService) EndTime(endTime int64) *KlinesService {
	s.endTime = &endTime
	return s
}

func (s *KlinesService) Interval(interval string) *KlinesService {
	s.interval = interval
	return s
}

func (s *KlinesService) Limit(limit int) *KlinesService {
	s.limit = &limit
	return s
}

func (s *KlinesService) Symbol(symbol string) *KlinesService {
	s.symbol = symbol
	return s
}

func (s *KlinesService) PriceType(priceType string) *KlinesService {
	s.priceType = &priceType
	return s
}

func (s *KlinesService) Type(ktype string) *KlinesService {
	s.ktype = &ktype
	return s
}

func (s *KlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "api/v2/klines",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.priceType != nil {
		r.setParam("priceType", *s.priceType)
	}
	if s.ktype != nil {
		r.setParam("type", *s.ktype)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Kline{}, err
	}
	j, err := newJSON(data)
	if err != nil {
		return []*Kline{}, err
	}
	num := len(j.MustArray())
	res = make([]*Kline, num)
	for i := 0; i < num; i++ {
		item := j.GetIndex(i)
		if len(item.MustArray()) < 6 {
			err = fmt.Errorf("invalid kline data: %s", item)
			return []*Kline{}, err
		}
		res[i] = &Kline{
			OpenTime: item.GetIndex(0).MustInt64(),
			Open:     item.GetIndex(1).MustString(),
			High:     item.GetIndex(2).MustString(),
			Low:      item.GetIndex(3).MustString(),
			Close:    item.GetIndex(4).MustString(),
			Volume:   item.GetIndex(5).MustString(),
		}
	}
	return res, nil
}

// Kline define kline info
type Kline struct {
	OpenTime int64  `json:"openTime"`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
}
