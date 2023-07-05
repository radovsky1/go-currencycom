package go_currencycom

import (
	stdjson "encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WsHandler func(message []byte)

type ErrHandler func(err error)

type payload map[string]interface{}

type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

type WsRequest struct {
	Destination   string  `json:"destination"`
	CorrelationID int     `json:"correlationId"`
	Payload       payload `json:"payload"`
}

func newWsRequest(destination string, correlationID int, p payload) *WsRequest {
	correlationID++
	return &WsRequest{
		Destination:   destination,
		CorrelationID: correlationID,
		Payload:       p,
	}
}

var wsServe = func(config *WsConfig, requests chan WsRequest, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	c, _, err := Dialer.Dial(config.Endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	c.SetReadLimit(655350)
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)
		if WebsocketKeepAlive {
			keepAlive(c, WebsocketTimeout)
		}
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			err := c.Close()
			if err != nil {
				return
			}
		}()

		// Wait for the requests channel to send messages in a websocket.
		go func() {
			for {
				select {
				case request := <-requests:
					msg, err := stdjson.Marshal(request)
					if err != nil {
						errHandler(err)
						return
					}
					err = c.WriteMessage(websocket.TextMessage, msg)
					if err != nil {
						errHandler(err)
						return
					}
				case <-doneC:
					return
				}
			}
		}()

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				err := c.Close()
				if err != nil {
					return
				}
				return
			}
		}
	}()
}
