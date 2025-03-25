package chipmunk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/h-varmazyar/Gate/services/indicators/internal/domain"
	"github.com/valyala/fasthttp"
	"golang.org/x/net/context"
)

func (c Chipmunk) CandleList(_ context.Context, marketID, resolutionID uint) ([]domain.Candle, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(fmt.Sprintf("%v/v1/candles/markets/%v/resolutions/%v", c.cfg.BaseURL, marketID, resolutionID))
	req.Header.SetMethod(fasthttp.MethodGet)
	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)
	err := c.client.Do(req, httpResp)
	if err != nil {
		return nil, err
	}

	resp := new(Response[[]domain.Candle])
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
