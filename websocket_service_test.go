package go_currencycom

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type websocketServiceTestSuite struct {
	baseTestSuite
	origWsServe func(*WsConfig, chan WsRequest, WsHandler, ErrHandler) (chan struct{}, chan struct{}, error)
	serveCount  int
}

func TestWebsocketService(t *testing.T) {
	suite.Run(t, new(websocketServiceTestSuite))
}

func (s *websocketServiceTestSuite) SetupTest() {
	s.origWsServe = wsServe
}

func (s *websocketServiceTestSuite) TearDownTest() {
	wsServe = s.origWsServe
	s.serveCount = 0
}

func (s *websocketServiceTestSuite) mockWsServe(data []byte, err error) {
	wsServe = func(cfg *WsConfig, requests chan WsRequest, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, innerErr error) {
		s.serveCount++
		doneC = make(chan struct{})
		stopC = make(chan struct{})
		go func() {
			<-requests
			close(requests)
		}()
		go func() {
			<-stopC
			close(doneC)
		}()
		handler(data)
		if err != nil {
			errHandler(err)
		}
		return doneC, stopC, nil
	}
}

func (s *websocketServiceTestSuite) assertWsServe(count ...int) {
	e := 1
	if len(count) > 0 {
		e = count[0]
	}
	s.r().Equal(e, s.serveCount)
}

func (s *websocketServiceTestSuite) TestWsMarketDataServe() {
	data := []byte(`{
		"status":"OK",
		"destination":"internal.quote",
		"payload":{
			"symbolName":"TXN",
			"bid":139.85,
			"bidQty":2500,
			"ofr":139.92000000000002,
			"ofrQty":2500,
			"timestamp":1597850971558
		}}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarketDataServe([]string{"TXN"}, func(event *WsMarketDataEvent) {
		e := &WsMarketDataEvent{
			SymbolName: "TXN",
			Bid:        139.85,
			Ofr:        139.92000000000002,
			BidQty:     2500,
			OfrQty:     2500,
			Timestamp:  1597850971558,
		}
		s.assertWsMarketDataEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	close(stopC)
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMarketDataEventEqual(e, a *WsMarketDataEvent) {
	r := s.r()
	r.Equal(e.SymbolName, a.SymbolName, "SymbolName")
	r.Equal(e.Bid, a.Bid, "Bid")
	r.Equal(e.Ofr, a.Ofr, "Ofr")
	r.Equal(e.BidQty, a.BidQty, "BidQty")
	r.Equal(e.OfrQty, a.OfrQty, "OfrQty")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
}

func (s *websocketServiceTestSuite) TestWsOHLCMarketDataServe() {
	data := []byte(`{
		"status":"OK",
		"destination":"ohlc.event",
		"payload":{
			"interval":"1m",
			"symbol":"BTC/USD_LEVERAGE",
			"type":"classic",
			"t":1673619780000,
			"h":18940.3,
			"l":18926.55,
			"o":18938.2,
			"c":18936.2}
		}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsOHLCMarketDataServe([]string{"BTC/USD_LEVERAGE"}, []string{"1m"}, func(event *WsOHLCMarketDataEvent) {
		e := &WsOHLCMarketDataEvent{
			Interval:  "1m",
			Symbol:    "BTC/USD_LEVERAGE",
			Type:      "classic",
			Timestamp: 1673619780000,
			High:      18940.3,
			Low:       18926.55,
			Open:      18938.2,
			Close:     18936.2,
		}
		s.assertWsOHLCMarketDataEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	close(stopC)
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsOHLCMarketDataEventEqual(e, a *WsOHLCMarketDataEvent) {
	r := s.r()
	r.Equal(e.Interval, a.Interval, "Interval")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.High, a.High, "High")
	r.Equal(e.Low, a.Low, "Low")
	r.Equal(e.Open, a.Open, "Open")
	r.Equal(e.Close, a.Close, "Close")
}

func (s *websocketServiceTestSuite) TestWsTradesServe() {
	data := []byte(`{
		"status":"OK",
		"destination":"internal.trade",
		"payload":{
			"price":11400.95,
			"size":0.058,
			"id":1616651347,
			"ts":1596625079952,
			"symbol":"BTC/USD",
			"orderId":"00a02503-0079-54c4-0000-00004020316a",
			"clientOrderId":"00a02503-0079-54c4-0000-482f00003a06",
			"buyer":true
		}}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsTradesServe([]string{"BTC/USD"}, func(event *WsTradesEvent) {
		e := &WsTradesEvent{
			Price:         11400.95,
			Size:          0.058,
			ID:            1616651347,
			Timestamp:     1596625079952,
			Symbol:        "BTC/USD",
			OrderID:       "00a02503-0079-54c4-0000-00004020316a",
			ClientOrderID: "00a02503-0079-54c4-0000-482f00003a06",
			Buyer:         true,
		}
		s.assertWsTradeEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	close(stopC)
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsTradeEventEqual(e, a *WsTradesEvent) {
	r := s.r()
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Size, a.Size, "Size")
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Buyer, a.Buyer, "Buyer")
}
