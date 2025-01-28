package chipmunk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"golang.org/x/net/context"
	"time"
)

type Candle struct {
	ID     uint      `json:"id"`
	Time   time.Time `json:"time"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume float64   `json:"volume"`
}

func (c Chipmunk) CandleList(_ context.Context, marketID, resolutionID uint) ([]Candle, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(fmt.Sprintf(c.cfg.CandleListAddress, marketID, resolutionID))
	req.Header.SetMethod(fasthttp.MethodGet)
	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)
	err := c.client.Do(req, httpResp)
	if err != nil {
		return nil, err
	}

	resp := new(Response[[]Candle])
	err = json.Unmarshal(httpResp.Body(), resp)
	if err != nil {
		return nil, err
	}

	if resp.IsSuccess {
		return resp.Data, nil
	} else {
		return nil, errors.New(resp.Error)
	}
}
