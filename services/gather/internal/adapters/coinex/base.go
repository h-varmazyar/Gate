package coinex

import (
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"golang.org/x/net/context"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"time"
)

type Coinex struct {
	cfg    configs.CoinexAdapter
	client *http.Client
	//client *fasthttp.Client
}

func NewCoinex(cfg configs.CoinexAdapter) (*Coinex, error) {
	dialer, err := proxy.SOCKS5("tcp", cfg.SocksProxyAddress, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	httpTransport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	client := &http.Client{
		Transport: httpTransport,
		Timeout:   10 * time.Second,
	}

	return &Coinex{
		client: client,
		cfg:    cfg,
	}, nil
}

//func NewCoinex(cfg configs.CoinexAdapter) *Coinex {
//	readTimeout, _ := time.ParseDuration("10s")
//	writeTimeout, _ := time.ParseDuration("10s")
//	maxIdleConnDuration, _ := time.ParseDuration("1h")
//	dialer := fasthttpproxy.FasthttpSocksDialerDualStack(cfg.SocksProxyAddress)
//	client := &fasthttp.Client{
//		ReadTimeout:         readTimeout,
//		WriteTimeout:        writeTimeout,
//		MaxIdleConnDuration: maxIdleConnDuration,
//		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
//		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
//		DisablePathNormalizing: true,
//		// increase DNS cache time to an hour instead of default minute
//		Dial: dialer,
//	}
//
//	return &Coinex{
//		client: client,
//		cfg:    cfg,
//	}
//}

type baseResponse[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}
