package sahamyab

import (
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"net/http"
	"time"
)

type Sahamyab struct {
	cfg    configs.SahamyabAdapter
	client *http.Client
}

func NewSahamyab(cfg configs.SahamyabAdapter) *Sahamyab {
	//dialer, err := proxy.SOCKS5("tcp", cfg.SocksProxyAddress, nil, proxy.Direct)
	//if err != nil {
	//	return nil, err
	//}
	//
	//httpTransport := &http.Transport{
	//	DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
	//		return dialer.Dial(network, addr)
	//	},
	//}

	client := &http.Client{
		//Transport: httpTransport,
		Timeout: 30 * time.Second,
	}
	return &Sahamyab{
		cfg:    cfg,
		client: client,
	}
}

type Data interface {
	any
}

type baseResponse struct {
	ErrorCode  string `json:"errorCode"`
	ErrorTitle string `json:"errorTitle"`
	Success    bool   `json:"success"`
	HasMore    bool   `json:"hasMore,omitempty"`
}
