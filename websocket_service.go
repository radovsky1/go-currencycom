package go_currencycom

import (
	"errors"
	"time"
)

const (
	baseWsURL     = "wss://api-adapter.backend.currency.com/connect"
	baseWsDemoURL = "wss://demo-api-adapter.backend.currency.com/connect"
)

var (
	WebsocketTimeout   = 30 * time.Second
	WebsocketKeepAlive = true
	CorrelationID      = -1
)

func getWsEndpoint() string {
	if UseDemo {
		return baseWsDemoURL
	}
	return baseWsURL
}

type WsMarketDataEvent struct {
	SymbolName string  `json:"symbolName"`
	Bid        float64 `json:"bid"`
	Ofr        float64 `json:"ofr"`
	BidQty     float64 `json:"bidQty"`
	OfrQty     float64 `json:"ofrQty"`
	Timestamp  int64   `json:"timestamp"`
}

type WsMarketDataHandler func(event *WsMarketDataEvent)

func WsMarketDataServe(symbols []string, handler WsMarketDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	return wsMarketDataServe(getWsEndpoint(), symbols, handler, errHandler)
}

func wsMarketDataServe(endpoint string, symbols []string, handler WsMarketDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	config := newWsConfig(endpoint)
	requests := make(chan WsRequest)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		status, err := j.Get("status").String()
		if err != nil {
			errHandler(err)
			return
		}
		if status != "OK" {
			errHandler(errors.New(status))
			return
		}
		j = j.Get("payload")
		event := new(WsMarketDataEvent)
		event.SymbolName = j.Get("symbolName").MustString()
		event.Timestamp = j.Get("timestamp").MustInt64()
		event.Bid = j.Get("bid").MustFloat64()
		event.Ofr = j.Get("ofr").MustFloat64()
		event.BidQty = j.Get("bidQty").MustFloat64()
		event.OfrQty = j.Get("ofrQty").MustFloat64()
		handler(event)
	}
	doneC, stopC, err = wsServe(config, requests, wsHandler, errHandler)
	requests <- *newWsRequest("marketData.subscribe", CorrelationID, payload{"symbols": symbols})
	return doneC, stopC, err
}

type WsOHLCMarketDataEvent struct {
	Symbol    string  `json:"symbol"`
	Interval  string  `json:"interval"`
	Type      string  `json:"type"`
	Open      float64 `json:"o"`
	High      float64 `json:"h"`
	Low       float64 `json:"l"`
	Close     float64 `json:"c"`
	Timestamp int64   `json:"t"`
}

type WsOHLCMarketDataHandler func(event *WsOHLCMarketDataEvent)

func WsOHLCMarketDataServe(symbols []string, intervals []string, handler WsOHLCMarketDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	return wsOHLCMarketDataServe(getWsEndpoint(), symbols, intervals, handler, errHandler)
}

func wsOHLCMarketDataServe(endpoint string, symbols []string, intervals []string, handler WsOHLCMarketDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	config := newWsConfig(endpoint)
	requests := make(chan WsRequest)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		status, err := j.Get("status").String()
		if err != nil {
			errHandler(err)
			return
		}
		if status != "OK" {
			errHandler(errors.New(status))
			return
		}
		j = j.Get("payload")
		event := new(WsOHLCMarketDataEvent)
		event.Symbol = j.Get("symbol").MustString()
		event.Type = j.Get("type").MustString()
		event.Interval = j.Get("interval").MustString()
		event.Timestamp = j.Get("t").MustInt64()
		event.Open = j.Get("o").MustFloat64()
		event.High = j.Get("h").MustFloat64()
		event.Low = j.Get("l").MustFloat64()
		event.Close = j.Get("c").MustFloat64()
		handler(event)
	}
	doneC, stopC, err = wsServe(config, requests, wsHandler, errHandler)
	requests <- *newWsRequest("OHLCMarketData.subscribe", CorrelationID, payload{"symbols": symbols, "intervals": intervals})
	return doneC, stopC, err
}

type WsTradesEvent struct {
	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
	ID            int64   `json:"id"`
	Timestamp     int64   `json:"ts"`
	Symbol        string  `json:"symbol"`
	OrderID       string  `json:"orderId"`
	ClientOrderID string  `json:"clientOrderId"`
	Buyer         bool    `json:"buyer"`
}

type WsTradesHandler func(event *WsTradesEvent)

func WsTradesServe(symbols []string, handler WsTradesHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	return wsTradesServe(getWsEndpoint(), symbols, handler, errHandler)
}

func wsTradesServe(endpoint string, symbols []string, handler WsTradesHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	config := newWsConfig(endpoint)
	requests := make(chan WsRequest)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		status, err := j.Get("status").String()
		if err != nil {
			errHandler(err)
			return
		}
		if status != "OK" {
			errHandler(errors.New(status))
			return
		}
		j = j.Get("payload")
		event := new(WsTradesEvent)
		event.Symbol = j.Get("symbol").MustString()
		event.ID = j.Get("id").MustInt64()
		event.Timestamp = j.Get("ts").MustInt64()
		event.Price = j.Get("price").MustFloat64()
		event.Size = j.Get("size").MustFloat64()
		event.OrderID = j.Get("orderId").MustString()
		event.ClientOrderID = j.Get("clientOrderId").MustString()
		event.Buyer = j.Get("buyer").MustBool()
		handler(event)
	}
	doneC, stopC, err = wsServe(config, requests, wsHandler, errHandler)
	requests <- *newWsRequest("trades.subscribe", CorrelationID, payload{"symbols": symbols})
	return doneC, stopC, err
}
