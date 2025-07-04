package gather

import (
	"github.com/h-varmazyar/Gate/services/indicators/configs"
	"github.com/valyala/fasthttp"
)

type Gather struct {
	cfg    configs.GatherAdapter
	client *fasthttp.Client
}

type Response[D any] struct {
	Data      D      `json:"data"`
	Error     string `json:"error"`
	IsSuccess bool   `json:"is_success"`
}

func NewChipmunk(cfg configs.GatherAdapter) Gather {
	return Gather{
		cfg: cfg,
		client: &fasthttp.Client{
			Name: "Indicators-Gather-Adapter",
		},
	}
}
