package chipmunk

import (
	"github.com/h-varmazyar/Gate/services/indicators/configs"
	"github.com/valyala/fasthttp"
)

type Chipmunk struct {
	cfg    configs.ChipmunkAdapter
	client *fasthttp.Client
}

type Response[D any] struct {
	Data      D      `json:"data"`
	Error     string `json:"error"`
	IsSuccess bool   `json:"is_success"`
}

func NewGather(cfg configs.ChipmunkAdapter) Chipmunk {
	return Chipmunk{
		cfg: cfg,
		client: &fasthttp.Client{
			DialTimeout:                   nil,
			Dial:                          nil,
			TLSConfig:                     nil,
			RetryIf:                       nil,
			RetryIfErr:                    nil,
			ConfigureClient:               nil,
			Name:                          "",
			MaxConnsPerHost:               0,
			MaxIdleConnDuration:           0,
			MaxConnDuration:               0,
			MaxIdemponentCallAttempts:     0,
			ReadBufferSize:                0,
			WriteBufferSize:               0,
			ReadTimeout:                   0,
			WriteTimeout:                  0,
			MaxResponseBodySize:           0,
			MaxConnWaitTimeout:            0,
			ConnPoolStrategy:              0,
			NoDefaultUserAgentHeader:      false,
			DialDualStack:                 false,
			DisableHeaderNamesNormalizing: false,
			DisablePathNormalizing:        false,
			StreamResponseBody:            false,
		},
	}
}
