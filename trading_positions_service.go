package go_currencycom

import (
	"context"
	"net/http"
)

type CloseTradingPositionService struct {
	c          *Client
	positionID string
}

func (s *CloseTradingPositionService) PositionID(positionID string) *CloseTradingPositionService {
	s.positionID = positionID
	return s
}

type RequestDto struct {
	AccountID        string `json:"accountId"`
	CreatedTimestamp int64  `json:"createdTimestamp"`
	ID               int64  `json:"id"`
	OrderID          string `json:"orderId"`
	PositionID       string `json:"positionId"`
	RejectReason     string `json:"rejectReason"`
	RqType           string `json:"rqType"`
	State            string `json:"state"`
}

type CloseTradingPositionResponse struct {
	Request []RequestDto `json:"request"`
}

// Do send request
func (s *CloseTradingPositionService) Do(ctx context.Context, opts ...RequestOption) (res *CloseTradingPositionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "api/v2/closeTradingPosition",
		secType:  secTypeSigned,
	}
	r.setParam("positionId", s.positionID)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CloseTradingPositionResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ListTradingPositionsService struct {
	c *Client
}

type TradingPositionDto struct {
	AccountID                           string  `json:"accountId"`
	ClosePrice                          float64 `json:"closePrice"`
	CloseQuantity                       float64 `json:"closeQuantity"`
	CloseTimestamp                      int64   `json:"closeTimestamp"`
	Cost                                float64 `json:"cost"`
	CreatedTimestamp                    int64   `json:"createdTimestamp"`
	Currency                            string  `json:"currency"`
	CurrentTrailingPrice                float64 `json:"currentTrailingPrice"`
	CurrenTrailingPriceUpdatedTimestamp int64   `json:"currenTrailingPriceUpdatedTimestamp"`
	Dividend                            float64 `json:"dividend"`
	Fee                                 float64 `json:"fee"`
	GuaranteedStopLoss                  bool    `json:"guaranteedStopLoss"`
	ID                                  string  `json:"id"`
	InstrumentID                        int64   `json:"instrumentId"`
	Margin                              float64 `json:"margin"`
	OpenPrice                           float64 `json:"openPrice"`
	OpenQuantity                        float64 `json:"openQuantity"`
	OpenTimestamp                       int64   `json:"openTimestamp"`
	OrderID                             string  `json:"orderId"`
	Rpl                                 float64 `json:"rpl"`
	RplConverted                        float64 `json:"rplConverted"`
	State                               string  `json:"state"`
	StopLoss                            float64 `json:"stopLoss"`
	Swap                                float64 `json:"swap"`
	SwapConverted                       float64 `json:"swapConverted"`
	Symbol                              string  `json:"symbol"`
	TakeProfit                          float64 `json:"takeProfit"`
	TrailingQuotedPrice                 float64 `json:"trailingQuotedPrice"`
	TrailingStopLoss                    bool    `json:"trailingStopLoss"`
	Type                                string  `json:"type"`
	Upl                                 float64 `json:"upl"`
	UplConverted                        float64 `json:"uplConverted"`
}

type ListTradingPositionsResponse struct {
	Positions []TradingPositionDto `json:"positions"`
}

// Do send request
func (s *ListTradingPositionsService) Do(ctx context.Context, opts ...RequestOption) (res *ListTradingPositionsResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "api/v2/tradingPositions",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ListTradingPositionsResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type UpdateTradingPositionService struct {
	c                  *Client
	positionID         string
	guaranteedStopLoss *bool
	profitDistance     *float64
	stopDistance       *float64
	stopLoss           *float64
	takeProfit         *float64
	trailingStopLoss   *bool
}

func (s *UpdateTradingPositionService) PositionID(positionID string) *UpdateTradingPositionService {
	s.positionID = positionID
	return s
}

func (s *UpdateTradingPositionService) GuaranteedStopLoss(guaranteedStopLoss bool) *UpdateTradingPositionService {
	s.guaranteedStopLoss = &guaranteedStopLoss
	return s
}

func (s *UpdateTradingPositionService) ProfitDistance(profitDistance float64) *UpdateTradingPositionService {
	s.profitDistance = &profitDistance
	return s
}

func (s *UpdateTradingPositionService) StopDistance(stopDistance float64) *UpdateTradingPositionService {
	s.stopDistance = &stopDistance
	return s
}

func (s *UpdateTradingPositionService) StopLoss(stopLoss float64) *UpdateTradingPositionService {
	s.stopLoss = &stopLoss
	return s
}

func (s *UpdateTradingPositionService) TakeProfit(takeProfit float64) *UpdateTradingPositionService {
	s.takeProfit = &takeProfit
	return s
}

func (s *UpdateTradingPositionService) TrailingStopLoss(trailingStopLoss bool) *UpdateTradingPositionService {
	s.trailingStopLoss = &trailingStopLoss
	return s
}

type UpdateTradingPositionResponse struct {
	RequestID int64  `json:"requestId"`
	State     string `json:"state"`
}

// Do send request
func (s *UpdateTradingPositionService) Do(ctx context.Context, opts ...RequestOption) (res *UpdateTradingPositionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "api/v2/updateTradingPosition",
		secType:  secTypeSigned,
	}
	r.setParam("positionId", s.positionID)
	if s.guaranteedStopLoss != nil {
		r.setParam("guaranteedStopLoss", *s.guaranteedStopLoss)
	}
	if s.profitDistance != nil {
		r.setParam("profitDistance", *s.profitDistance)
	}
	if s.stopDistance != nil {
		r.setParam("stopDistance", *s.stopDistance)
	}
	if s.stopLoss != nil {
		r.setParam("stopLoss", *s.stopLoss)
	}
	if s.takeProfit != nil {
		r.setParam("takeProfit", *s.takeProfit)
	}
	if s.trailingStopLoss != nil {
		r.setParam("trailingStopLoss", *s.trailingStopLoss)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UpdateTradingPositionResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ListHistoricalPositionsService struct {
	c      *Client
	from   *int64
	to     *int64
	symbol *string
	limit  *int
}

func (s *ListHistoricalPositionsService) From(from int64) *ListHistoricalPositionsService {
	s.from = &from
	return s
}

func (s *ListHistoricalPositionsService) To(to int64) *ListHistoricalPositionsService {
	s.to = &to
	return s
}

func (s *ListHistoricalPositionsService) Symbol(symbol string) *ListHistoricalPositionsService {
	s.symbol = &symbol
	return s
}

func (s *ListHistoricalPositionsService) Limit(limit int) *ListHistoricalPositionsService {
	s.limit = &limit
	return s
}

type FeeDetailsDto struct {
	Commission float64 `json:"commission"`
}

type HistoricalPositionDto struct {
	AccountCurrency  string        `json:"accountCurrency"`
	AccountID        int64         `json:"accountId"`
	CreatedTimestamp int64         `json:"createdTimestamp"`
	Currency         string        `json:"currency"`
	ExecID           string        `json:"execId"`
	ExecTimestamp    int64         `json:"execTimestamp"`
	ExecutionTyp     string        `json:"executionType"`
	Fee              float64       `json:"fee"`
	FeeDetails       FeeDetailsDto `json:"feeDetails"`
	FxRate           float64       `json:"fxRate"`
	GSL              bool          `json:"gSL"`
	InstrumentID     int64         `json:"instrumentId"`
	PositionID       string        `json:"positionId"`
	Price            float64       `json:"price"`
	Quantity         float64       `json:"quantity"`
	RejectReason     string        `json:"rejectReason"`
	Rpl              float64       `json:"rpl"`
	RplConverted     float64       `json:"rplConverted"`
	Source           string        `json:"source"`
	Status           string        `json:"status"`
	StopLoss         float64       `json:"stopLoss"`
	Swap             float64       `json:"swap"`
	SwapConverted    float64       `json:"swapConverted"`
	Symbol           string        `json:"symbol"`
	TakeProfit       float64       `json:"takeProfit"`
	TrailingStopLoss bool          `json:"trailingStopLoss"`
}

type ListHistoricalPositionsResponse struct {
	History []HistoricalPositionDto `json:"history"`
}

// Do send request
func (s *ListHistoricalPositionsService) Do(ctx context.Context, opts ...RequestOption) (res *ListHistoricalPositionsResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "api/v2/tradingPositionsHistory",
		secType:  secTypeSigned,
	}
	if s.from != nil {
		r.setParam("from", *s.from)
	}
	if s.to != nil {
		r.setParam("to", *s.to)
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ListHistoricalPositionsResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
