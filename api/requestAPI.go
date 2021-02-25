package api

import "github.com/mrNobody95/Gate/models"

type Requests interface {
	Login(totp int) error
	OrderBook(symbol string) (*OrderBookResponse, error)
	RecentTrades(symbol string) (*RecentTradesResponse, error)
	MarketStats(destination, source string) (*MarketStatusResponse, error)
	OHLC(symbol string, resolution *models.Resolution, from, to float64) (*OHLCResponse, error)
}

type RequestType int

type Request struct {
	Type     RequestType
	Endpoint string
	Headers  map[string]interface{}
	Params   map[string]interface{}
}

type Response struct {
	Code         int
	ErrorMessage string
	Body         []byte
	Headers      map[string]interface{}
}

const (
	GET RequestType = iota
	PUT
	POST
	PATCH
	DELETE
)

func (r *Request) Execute() *Response {
	return nil
}
