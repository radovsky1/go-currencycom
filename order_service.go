package go_currencycom

import (
	"context"
	"net/http"
)

type CreateOrderService struct {
	c                  *Client
	accountID          *int64
	expireTimestamp    *int64
	guaranteedStopLoss *bool
	leverage           *int32
	newOrderRespType   *NewOrderRespType
	price              *float64
	profitDistance     *float64
	quantity           float64
	side               SideType
	stopDistance       *float64
	stopLoss           *float64
	symbol             string
	takeProfit         *float64
	trailingStopLoss   *bool
	orderType          OrderType
}

// Symbol set symbol
func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

// AccountID set account id
func (s *CreateOrderService) AccountID(accountID int64) *CreateOrderService {
	s.accountID = &accountID
	return s
}

// ExpireTimestamp set expire timestamp
func (s *CreateOrderService) ExpireTimestamp(expireTimestamp int64) *CreateOrderService {
	s.expireTimestamp = &expireTimestamp
	return s
}

// GuaranteedStopLoss set guaranteed stop loss
func (s *CreateOrderService) GuaranteedStopLoss(guaranteedStopLoss bool) *CreateOrderService {
	s.guaranteedStopLoss = &guaranteedStopLoss
	return s
}

// Leverage set leverage
func (s *CreateOrderService) Leverage(leverage int32) *CreateOrderService {
	s.leverage = &leverage
	return s
}

// NewOrderRespType set new order response type
func (s *CreateOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// Price set price
func (s *CreateOrderService) Price(price float64) *CreateOrderService {
	s.price = &price
	return s
}

// ProfitDistance set profit distance
func (s *CreateOrderService) ProfitDistance(profitDistance float64) *CreateOrderService {
	s.profitDistance = &profitDistance
	return s
}

// Quantity set quantity
func (s *CreateOrderService) Quantity(quantity float64) *CreateOrderService {
	s.quantity = quantity
	return s
}

// Side set side
func (s *CreateOrderService) Side(side SideType) *CreateOrderService {
	s.side = side
	return s
}

// StopDistance set stop distance
func (s *CreateOrderService) StopDistance(stopDistance float64) *CreateOrderService {
	s.stopDistance = &stopDistance
	return s
}

// StopLoss set stop loss
func (s *CreateOrderService) StopLoss(stopLoss float64) *CreateOrderService {
	s.stopLoss = &stopLoss
	return s
}

// TakeProfit set take profit
func (s *CreateOrderService) TakeProfit(takeProfit float64) *CreateOrderService {
	s.takeProfit = &takeProfit
	return s
}

// TrailingStopLoss set trailing stop loss
func (s *CreateOrderService) TrailingStopLoss(trailingStopLoss bool) *CreateOrderService {
	s.trailingStopLoss = &trailingStopLoss
	return s
}

// Type set order type
func (s *CreateOrderService) Type(orderType OrderType) *CreateOrderService {
	s.orderType = orderType
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"quantity": s.quantity,
		"type":     s.orderType,
	}
	if s.accountID != nil {
		m["accountId"] = *s.accountID
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.expireTimestamp != nil {
		m["expireTimestamp"] = *s.expireTimestamp
	}
	if s.guaranteedStopLoss != nil {
		m["guaranteedStopLoss"] = *s.guaranteedStopLoss
	}
	if s.leverage != nil {
		m["leverage"] = *s.leverage
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.profitDistance != nil {
		m["profitDistance"] = *s.profitDistance
	}
	if s.stopDistance != nil {
		m["stopDistance"] = *s.stopDistance
	}
	if s.stopLoss != nil {
		m["stopLoss"] = *s.stopLoss
	}
	if s.takeProfit != nil {
		m["takeProfit"] = *s.takeProfit
	}
	if s.trailingStopLoss != nil {
		m["trailingStopLoss"] = *s.trailingStopLoss
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

type CreateOrderResponse struct {
	ExecutedQty        string          `json:"executedQty"`
	ExpireTimestamp    int64           `json:"expireTimestamp"`
	GuaranteedStopLoss bool            `json:"guaranteedStopLoss"`
	Margin             float64         `json:"margin"`
	OrderID            string          `json:"orderId"`
	OrigQty            string          `json:"origQty"`
	Price              string          `json:"price"`
	ProfitDistance     float64         `json:"profitDistance"`
	RejectMessage      string          `json:"rejectMessage"`
	Side               SideType        `json:"side"`
	Status             OrderStatusType `json:"status"`
	StopDistance       float64         `json:"stopDistance"`
	StopLoss           float64         `json:"stopLoss"`
	Symbol             string          `json:"symbol"`
	TakeProfit         float64         `json:"takeProfit"`
	TimeInForce        TimeInForceType `json:"timeInForce"`
	TrailingStopLoss   bool            `json:"trailingStopLoss"`
	TransactTime       int64           `json:"transactTime"`
	Type               OrderType       `json:"type"`
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/api/v2/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CancelOrderService struct {
	c       *Client
	symbol  string
	orderID *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set order id
func (s *CancelOrderService) OrderID(orderID string) *CancelOrderService {
	s.orderID = &orderID
	return s
}

type CancelOrderResponse struct {
	ExecutedQty string          `json:"executedQty"`
	OrderID     string          `json:"orderId"`
	OrigQty     string          `json:"origQty"`
	Price       string          `json:"price"`
	Side        SideType        `json:"side"`
	Status      OrderStatusType `json:"status"`
	Symbol      string          `json:"symbol"`
	TimeInForce TimeInForceType `json:"timeInForce"`
	Type        OrderType       `json:"type"`
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/api/v2/order",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":  s.symbol,
		"orderId": *s.orderID,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type EditExchangeOrderService struct {
	c               *Client
	expireTimestamp *ExpireTimestampType
	orderID         string
	price           *string
}

// ExpireTimestamp set expire timestamp
func (s *EditExchangeOrderService) ExpireTimestamp(expireTimestamp ExpireTimestampType) *EditExchangeOrderService {
	s.expireTimestamp = &expireTimestamp
	return s
}

// OrderID set order id
func (s *EditExchangeOrderService) OrderID(orderID string) *EditExchangeOrderService {
	s.orderID = orderID
	return s
}

// Price set price
func (s *EditExchangeOrderService) Price(price string) *EditExchangeOrderService {
	s.price = &price
	return s
}

type EditExchangeOrderResponse struct {
	OrderID string `json:"orderId"`
}

// Do send request
func (s *EditExchangeOrderService) Do(ctx context.Context, opts ...RequestOption) (res *EditExchangeOrderResponse, err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/api/v2/order",
		secType:  secTypeSigned,
	}
	m := params{
		"orderId": s.orderID,
	}
	if s.expireTimestamp != nil {
		m["expireTimestamp"] = *s.expireTimestamp
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EditExchangeOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ListOpenOrdersService struct {
	c *Client
}

type QueryOrderResponse struct {
	AccountID          string          `json:"accountId"`
	ExecutedQty        string          `json:"executedQty"`
	ExpireTimestamp    int64           `json:"expireTimestamp"`
	GuaranteedStopLoss bool            `json:"guaranteedStopLoss"`
	IcebergQty         string          `json:"icebergQty"`
	Leverage           bool            `json:"leverage"`
	Margin             float64         `json:"margin"`
	OrderID            string          `json:"orderId"`
	OrigQty            string          `json:"origQty"`
	Price              string          `json:"price"`
	Side               SideType        `json:"side"`
	Status             OrderStatusType `json:"status"`
	StopLoss           float64         `json:"stopLoss"`
	Symbol             string          `json:"symbol"`
	TakeProfit         float64         `json:"takeProfit"`
	Time               int64           `json:"time"`
	TimeInForce        TimeInForceType `json:"timeInForce"`
	TrailingStopLoss   bool            `json:"trailingStopLoss"`
	Type               OrderType       `json:"type"`
	UpdateTime         int64           `json:"updateTime"`
	Working            bool            `json:"working"`
}

// Do send request
func (s *ListOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*QueryOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v2/openOrders",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*QueryOrderResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type FetchOrderService struct {
	c       *Client
	symbol  string
	orderID string
}

// Symbol set symbol
func (s *FetchOrderService) Symbol(symbol string) *FetchOrderService {
	s.symbol = symbol
	return s
}

// OrderID set order id
func (s *FetchOrderService) OrderID(orderID string) *FetchOrderService {
	s.orderID = orderID
	return s
}

type FetchOrderResponse struct {
	AccountID          string          `json:"accountId"`
	ExecPrice          float64         `json:"execPrice"`
	ExecQuantity       float64         `json:"execQuantity"`
	ExpireTime         int64           `json:"expireTime"`
	GuaranteedStopLoss bool            `json:"guaranteedStopLoss"`
	Margin             float64         `json:"margin"`
	OrderID            string          `json:"orderId"`
	Price              float64         `json:"price"`
	Quantity           float64         `json:"quantity"`
	RejectReason       string          `json:"rejectReason"`
	Side               SideType        `json:"side"`
	Status             OrderStatusType `json:"status"`
	StopLoss           float64         `json:"stopLoss"`
	TakeProfit         float64         `json:"takeProfit"`
	TimeInForceType    TimeInForceType `json:"timeInForceType"`
	Timestamp          int64           `json:"timestamp"`
	TrailingStopLoss   bool            `json:"trailingStopLoss"`
	Type               OrderType       `json:"type"`
}

// Do send request
func (s *FetchOrderService) Do(ctx context.Context, opts ...RequestOption) (res *FetchOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v2/fetchOrder",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":  s.symbol,
		"orderId": s.orderID,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FetchOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type UpdateTradingOrderService struct {
	c                  *Client
	orderID            string
	guaranteedStopLoss *bool
	newPrice           *float64
	profitDistance     *float64
	stopDistance       *float64
	stopLoss           *float64
	takeProfit         *float64
	trailingStopLoss   *bool
}

// OrderID set order id
func (s *UpdateTradingOrderService) OrderID(orderID string) *UpdateTradingOrderService {
	s.orderID = orderID
	return s
}

// GuaranteedStopLoss set guaranteed stop loss
func (s *UpdateTradingOrderService) GuaranteedStopLoss(guaranteedStopLoss bool) *UpdateTradingOrderService {
	s.guaranteedStopLoss = &guaranteedStopLoss
	return s
}

// NewPrice set new price
func (s *UpdateTradingOrderService) NewPrice(newPrice float64) *UpdateTradingOrderService {
	s.newPrice = &newPrice
	return s
}

// ProfitDistance set profit distance
func (s *UpdateTradingOrderService) ProfitDistance(profitDistance float64) *UpdateTradingOrderService {
	s.profitDistance = &profitDistance
	return s
}

// StopDistance set stop distance
func (s *UpdateTradingOrderService) StopDistance(stopDistance float64) *UpdateTradingOrderService {
	s.stopDistance = &stopDistance
	return s
}

// StopLoss set stop loss
func (s *UpdateTradingOrderService) StopLoss(stopLoss float64) *UpdateTradingOrderService {
	s.stopLoss = &stopLoss
	return s
}

// TakeProfit set take profit
func (s *UpdateTradingOrderService) TakeProfit(takeProfit float64) *UpdateTradingOrderService {
	s.takeProfit = &takeProfit
	return s
}

// TrailingStopLoss set trailing stop loss
func (s *UpdateTradingOrderService) TrailingStopLoss(trailingStopLoss bool) *UpdateTradingOrderService {
	s.trailingStopLoss = &trailingStopLoss
	return s
}

type UpdateTradingOrderResponse struct {
	RequestID int64  `json:"requestId"`
	State     string `json:"state"`
}

// Do send request
func (s *UpdateTradingOrderService) Do(ctx context.Context, opts ...RequestOption) (res *UpdateTradingOrderResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v2/updateTradingOrder",
		secType:  secTypeSigned,
	}
	m := params{
		"orderId": s.orderID,
	}
	if s.guaranteedStopLoss != nil {
		m["guaranteedStopLoss"] = *s.guaranteedStopLoss
	}
	if s.newPrice != nil {
		m["newPrice"] = *s.newPrice
	}
	if s.profitDistance != nil {
		m["profitDistance"] = *s.profitDistance
	}
	if s.stopDistance != nil {
		m["stopDistance"] = *s.stopDistance
	}
	if s.stopLoss != nil {
		m["stopLoss"] = *s.stopLoss
	}
	if s.takeProfit != nil {
		m["takeProfit"] = *s.takeProfit
	}
	if s.trailingStopLoss != nil {
		m["trailingStopLoss"] = *s.trailingStopLoss
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UpdateTradingOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
