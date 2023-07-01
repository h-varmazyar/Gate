package service

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	"google.golang.org/grpc/codes"
	"time"
)

func loadRequest(configs *Configs, brokerage *coreApi.Brokerage) brokerages.Requests {
	switch brokerage.Platform {
	case api.Platform_Coinex:
		var auth *api.Auth
		if brokerage.Auth != nil {
			auth = &api.Auth{
				Type:      api.AuthType_StaticToken,
				AccessID:  brokerage.Auth.AccessID,
				SecretKey: brokerage.Auth.SecretKey,
			}
		}
		coinexInstance := coinex.NewRequest(configs.Coinex, auth)
		return coinexInstance
	case api.Platform_Nobitex:
		return nil
	}
	return nil
}

func loadResponse(configs *Configs, brokerage *coreApi.Brokerage) (brokerages.Responses, error) {
	switch brokerage.Platform {
	case api.Platform_Coinex:
		return coinex.NewResponse(configs.Coinex, false)
		//case api.Platform_Nobitex:
		//return nil, nobitex.NewResponse(configs.Nobitex, false)
	}
	return nil, errors.New(context.Background(), codes.Unimplemented)
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

func (s *Service) doNetworkRequest(request *networkAPI.Request) (*networkAPI.Response, error) {
	resp, err := s.requestService.Do(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Service) createOHLCParams(req *coreApi.OHLCReq) *brokerages.OHLCParams {
	return &brokerages.OHLCParams{
		Resolution: req.Resolution,
		Market:     req.Market,
		From:       time.Unix(req.From, 0),
		To:         time.Unix(req.To, 0),
	}
}

func (s *Service) createMarketInfoParams(req *coreApi.MarketInfoReq) *brokerages.MarketInfoParams {
	return &brokerages.MarketInfoParams{
		Market: req.Market,
	}
}
