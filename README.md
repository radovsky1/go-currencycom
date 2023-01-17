# go-currencycom

A Go client library for the [Currency.com API](https://currency.com/ru/api/).

[![GoDoc](https://godoc.org/github.com/radovsky1/go-currencycom?status.svg)](https://godoc.org/github.com/radovsky1/go-currencycom)
[![Currency.com Swagger API](https://img.shields.io/badge/Currency.com-Swagger%20API-blue.svg)](https://apitradedoc.currency.com/swagger-ui.html#/)
[![Go Report Card](https://goreportcard.com/badge/github.com/radovsky1/go-currencycom)](https://goreportcard.com/report/github.com/radovsky1/go-currencycom)
### Installation

```shell
go get github.com/radovsky1/go-currencycom
```

### Contribution
This project is fully inspired by [go-binance](https://github.com/adshao/go-binance).
All other materials are taken from the official [Currency.com API](https://currency.com/ru/api) documentation.

### Importing
```golang
import (
    currencycom "github.com/radovsky1/go-currencycom"
)
```

### REST API

#### Setup

Init client for API services. Get APIKey/SecretKey from your binance account.

```golang
var (
    apiKey = "your api key"
    secretKey = "your secret key"
)
currencycom.UseDemo = true // use demo server
client := currencycom.NewClient(apiKey, secretKey)
```

A service instance stands for a REST API endpoint and is initialized by client.NewXXXService function.

Simply call API in chain style. Call Do() in the end to send HTTP request and get response.

#### Create Order

```golang
order, err := client.NewCreateOrderService().
        Symbol("BTC/USD_LEVERAGE").
        Side(go_currencycom.SideTypeBuy).
        Type(go_currencycom.OrderTypeLimit).
        Quantity(0.03).
        Price(15000).
        Do(context.Background())
if err != nil {
    panic(err)
}
println(order.OrderID)
```

#### Fetch Order

```golang
order, err := client.NewGetOrderService().
        Symbol("BTC/USD_LEVERAGE").
        OrderID("123456789").
        Do(context.Background())
if err != nil {
    panic(err)
}
println(order.OrderID)
```

#### Cancel Order

```golang
order, err := client.NewCancelOrderService().
        Symbol("BTC/USD_LEVERAGE").
        OrderID("123456789").
        Do(context.Background())
if err != nil {
    panic(err)
}
println(order.OrderID)
```

#### List Open Orders

```golang
openOrders, err := client.NewListOpenOrdersService().Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(openOrders)
```

#### List Trading Positions

```golang
positions, err := client.NewListTradingPositionsService().Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(positions)
```

#### Get Account Information

```golang
account, err := client.NewGetAccountService().Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(account)
```

There are more services available, please check the source code.

### Websocket API

You don't need Client in websocket API. Just call currencycom.WsXxxServe(args, handler, errHandler).

#### Market Data

```golang
wsMarketDataHandler := func(event *currencycom.WsMarketDataEvent) {
    fmt.Println(event)
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, stopC, err := currencycom.WsMarketDataServe("BTC/USD_LEVERAGE", wsMarketDataHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
// use stopC to exit
go func() {
    time.Sleep(5 * time.Second)
    stopC <- struct{}{}
}()
// remove this if you do not want to be blocked here
<-doneC
``` 

#### OHLC Market Data

```golang
wsOHLCMarketDataHandler := func(event *currencycom.WsOHLCMarketDataEvent) {
    fmt.Println(event)
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, stopC, err := currencycom.WsOHLCMarketDataServe("BTC/USD_LEVERAGE", wsOHLCMarketDataHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
// use stopC to exit
go func() {
    time.Sleep(5 * time.Second)
    stopC <- struct{}{}
}()
// remove this if you do not want to be blocked here
<-doneC
```

#### Trades

```golang
wsTradesHandler := func(event *currencycom.WsTradesEvent) {
    fmt.Println(event)
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, stopC, err := currencycom.WsTradesServe("BTC/USD_LEVERAGE", wsTradesHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
// use stopC to exit
go func() {
    time.Sleep(5 * time.Second)
    stopC <- struct{}{}
}()
// remove this if you do not want to be blocked here
<-doneC
```

### Feedback

If you have any questions/suggestions, please feel free to contact me.