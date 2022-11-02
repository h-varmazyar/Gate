package service

import (
	"github.com/h-varmazyar/Gate/api"
	coreApi "github.com/h-varmazyar/Gate/services/core/api"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
)

func loadBrokerage(brokerage *coreApi.Brokerage) brokerages.Brokerage {
	switch brokerage.Platform {
	case coreApi.Platform_Coinex:
		coinexInstance := new(coinex.Service)
		coinexInstance.Auth = &api.Auth{
			Type:      api.AuthType_StaticToken,
			AccessID:  brokerage.Auth.AccessID,
			SecretKey: brokerage.Auth.SecretKey,
		}
		return coinexInstance
	case coreApi.Platform_Nobitex:
		return nil
	}
	return nil
}
