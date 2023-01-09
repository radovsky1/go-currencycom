package go_currencycom

import (
	"context"
	"net/http"
)

type DepthService struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *DepthService) Symbol(symbol string) *DepthService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *DepthService) Limit(limit int) *DepthService {
	s.limit = &limit
	return s
}

// Do send request
func (s *DepthService) Do(ctx context.Context, opts ...RequestOption) (res *DepthResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "api/v2/depth",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	j, err := newJSON(data)
	if err != nil {
		return nil, err
	}

	res = new(DepthResponse)
	res.LastUpdateID = j.Get("lastUpdateId").MustInt64()
	// bids
	bidsLen := len(j.Get("bids").MustArray())
	res.Bids = make([]Bid, bidsLen)
	for i := 0; i < bidsLen; i++ {
		bidItem := j.Get("bids").GetIndex(i)
		res.Bids[i].Price = bidItem.GetIndex(0).MustFloat64()
		res.Bids[i].Quantity = bidItem.GetIndex(1).MustFloat64()
	}
	// asks
	asksLen := len(j.Get("asks").MustArray())
	res.Asks = make([]Ask, asksLen)
	for i := 0; i < asksLen; i++ {
		askItem := j.Get("asks").GetIndex(i)
		res.Asks[i].Price = askItem.GetIndex(0).MustFloat64()
		res.Asks[i].Quantity = askItem.GetIndex(1).MustFloat64()
	}

	return res, nil
}

// DepthResponse define depth info with bids and asks
type DepthResponse struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

// Ask is a type alias for PriceLevel.
type Ask = PriceLevel

// Bid is a type alias for PriceLevel.
type Bid = PriceLevel
