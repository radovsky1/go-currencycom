package go_currencycom

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/bitly/go-simplejson"
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// SideType define side type of order
type SideType string

// PositionSideType define position side type of order
type PositionSideType string

// OrderType define order type
type OrderType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// NewOrderRespType define response JSON verbosity
type NewOrderRespType string

// OrderStatusType define order status
type OrderStatusType string

// ExpireTimestampType define expire time of order
type ExpireTimestampType string

// CandlestickInterval define interval of candlestick
type CandlestickInterval string

const (
	// BaseURL is the base url of currency.com api
	BaseURL = "https://api-adapter.backend.currency.com/"
	// BaseDemoURL is the base url of currency.com demo api
	BaseDemoURL = "https://demo-api-adapter.backend.currency.com/"
)

// UseDemo is the flag to use demo api
var UseDemo = true

// Redefining the standard package
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Global enums
const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"
	OrderTypeStop   OrderType = "STOP"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"
	OrderStatusTypeRejected        OrderStatusType = "REJECTED"
	OrderStatusTypeExpired         OrderStatusType = "EXPIRED"
	OrderStatusTypePendingCancel   OrderStatusType = "PENDING_CANCEL"

	TimeInForceTypeGTC TimeInForceType = "GTC"
	TimeInForceTypeIOC TimeInForceType = "IOC"
	TimeInForceTypeFOK TimeInForceType = "FOK"

	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeRESULT NewOrderRespType = "RESULT"
	NewOrderRespTypeFULL   NewOrderRespType = "FULL"

	ExpireTimestampTypeGTC ExpireTimestampType = "GTC"
	ExpireTimestampTypeFOK ExpireTimestampType = "FOK"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"

	CandlestickInterval1m  CandlestickInterval = "1m"
	CandlestickInterval5m  CandlestickInterval = "5m"
	CandlestickInterval15m CandlestickInterval = "15m"
	CandlestickInterval30m CandlestickInterval = "30m"
	CandlestickInterval1h  CandlestickInterval = "1h"
	CandlestickInterval4h  CandlestickInterval = "4h"
	CandlestickInterval1d  CandlestickInterval = "1d"
	CandlestickInterval1w  CandlestickInterval = "1w"
)

func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func currentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func getAPIEndpoint() string {
	if UseDemo {
		return BaseDemoURL
	}
	return BaseURL
}

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    getAPIEndpoint(),
		UserAgent:  "go-currencycom",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(log.Writer(), log.Prefix(), log.Flags()),
	}
}

func (c *Client) debug(format string, args ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, args...)
	}
}

func (c *Client) parseRequest(req *request, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(req)
	}
	err = req.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, req.endpoint)
	if req.recvWindow > 0 {
		req.setParam(recvWindowKey, req.recvWindow)
	}
	if req.secType == secTypeSigned {
		req.setParam(timestampKey, currentTimestamp()-c.TimeOffset)
	}
	queryString := req.query.Encode()
	body := &bytes.Buffer{}
	bodyString := req.form.Encode()
	header := http.Header{}
	if req.header != nil {
		header = req.header.Clone()
	}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if req.secType == secTypeAPIKey || req.secType == secTypeSigned {
		header.Set("X-MBX-APIKEY", c.APIKey)
	}

	if req.secType == secTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		_, err := mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", mac.Sum(nil)))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}

	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("Request: %s %s", req.method, fullURL)

	req.fullURL = fullURL
	req.header = header
	req.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, nil
	}

	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("Request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	resp, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	c.debug("Response: %s", string(data))

	if resp.StatusCode >= http.StatusBadRequest {
		apiErr := new(APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("Failed to parse error message: %s", e)
		}
		return nil, apiErr
	}

	return data, nil
}

func (c *Client) SetAPIEndpoint(endpoint string) *Client {
	c.BaseURL = endpoint
	return c
}

func (c *Client) NewDepthService() *DepthService {
	return &DepthService{c: c}
}

func (c *Client) NewGetAccountService() *GetAccountService {
	return &GetAccountService{c: c}
}

func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{c: c}
}

func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{c: c}
}

func (c *Client) NewEditExchangeOrderService() *EditExchangeOrderService {
	return &EditExchangeOrderService{c: c}
}

func (c *Client) NewFetchOrderService() *FetchOrderService {
	return &FetchOrderService{c: c}
}

func (c *Client) NewListOpenOrdersService() *ListOpenOrdersService {
	return &ListOpenOrdersService{c: c}
}

func (c *Client) NewExchangeInfoService() *ExchangeInfoService {
	return &ExchangeInfoService{c: c}
}

func (c *Client) NewCloseTradingPositionService() *CloseTradingPositionService {
	return &CloseTradingPositionService{c: c}
}

func (c *Client) NewListTradingPositionsService() *ListTradingPositionsService {
	return &ListTradingPositionsService{c: c}
}

func (c *Client) NewListHistoricalPositionsService() *ListHistoricalPositionsService {
	return &ListHistoricalPositionsService{c: c}
}

func (c *Client) NewKlinesService() *KlinesService {
	return &KlinesService{c: c}
}
