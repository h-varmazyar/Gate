package coinex

import (
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"net/http"
	"time"
)

type Coinex struct {
	cfg    configs.CoinexAdapter
	client *http.Client
	//client *fasthttp.Client
}

func NewCoinex(cfg configs.CoinexAdapter) *Coinex {
	//readTimeout, _ := time.ParseDuration("10s")
	//writeTimeout, _ := time.ParseDuration("10s")
	//maxIdleConnDuration, _ := time.ParseDuration("1h")
	//client := &fasthttp.Client{
	//	ReadTimeout:                   readTimeout,
	//	WriteTimeout:                  writeTimeout,
	//	MaxIdleConnDuration:           maxIdleConnDuration,
	//	NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
	//	DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
	//	DisablePathNormalizing:        true,
	//	// increase DNS cache time to an hour instead of default minute
	//	Dial: (&fasthttp.TCPDialer{
	//		Concurrency:      4096,
	//		DNSCacheDuration: time.Hour,
	//	}).Dial,
	//}

	return &Coinex{
		client: &http.Client{
			Timeout:   time.Second * 10,
			Transport: &http.Transport{},
		},
		cfg: cfg,
	}
}

type baseResponse[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}
