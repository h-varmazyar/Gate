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
	client := &http.Client{
		Timeout: 60 * time.Second,
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
