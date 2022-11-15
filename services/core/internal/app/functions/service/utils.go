package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

func loadRequest(configs *Configs, brokerage *coreApi.Brokerage) brokerages.Requests {
	switch brokerage.Platform {
	case coreApi.Platform_Coinex:
		auth := &api.Auth{
			Type:      api.AuthType_StaticToken,
			AccessID:  brokerage.Auth.AccessID,
			SecretKey: brokerage.Auth.SecretKey,
		}
		coinexInstance := coinex.NewRequest(configs.Coinex, auth)
		return coinexInstance
	case coreApi.Platform_Nobitex:
		return nil
	}
	return nil
}

func (s *Service) loadBrokerage(ctx context.Context, id string) (*coreApi.Brokerage, error) {
	brokerageID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	brokerage, err := s.brokerageService.Return(ctx, &coreApi.BrokerageReturnReq{
		ID: brokerageID.String(),
	})
	if err != nil {
		return nil, err
	}
	return brokerage, nil
}

func (s *Service) createOHLCParams(req *coreApi.OHLCReq) *brokerages.OHLCParams {
	return &brokerages.OHLCParams{
		Resolution: req.Resolution,
		Market:     req.Market,
		From:       time.Unix(req.From, 0),
		To:         time.Unix(req.To, 0),
	}
}

func (s *Service) doAsyncRequest(request *networkAPI.Request) {
	go func(request *networkAPI.Request) {
		resp, err := s.requestService.Do(context.Background(), request)
		if err != nil {
			log.WithError(err).Errorf("failed to do request: %v", request)
			return
		}
		log.Infof("resp is : %v", resp)
	}(request)
}
