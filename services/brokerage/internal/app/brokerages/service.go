package brokerages

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/brokerage/configs"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	//candleService chipmunkApi.CandleServiceClient
	marketService chipmunkApi.MarketServiceClient
	walletService chipmunkApi.WalletsServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		chipmunkConnection := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
		//GrpcService.candleService = chipmunkApi.NewCandleServiceClient(chipmunkConnection)
		GrpcService.marketService = chipmunkApi.NewMarketServiceClient(chipmunkConnection)
		GrpcService.walletService = chipmunkApi.NewWalletsServiceClient(chipmunkConnection)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	brokerageApi.RegisterBrokerageServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, req *brokerageApi.CreateBrokerageReq) (*brokerageApi.Brokerage, error) {
	if _, ok := api.AuthType_value[req.Auth.Type.String()]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong_auth_type")
	}
	if _, ok := brokerageApi.Platform_value[req.Platform.String()]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong_platform")
	}
	if _, err := uuid.Parse(req.ResolutionID); err != nil {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_resolution")
	}
	brokerage := new(repository.Brokerage)
	mapper.Struct(req, brokerage)

	brokerage.Status = api.Status_Disable

	if err := repository.Brokerages.Create(brokerage); err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	mapper.Struct(brokerage, response)
	return response, nil
}

func (s *Service) List(_ context.Context, _ *api.Void) (*brokerageApi.Brokerages, error) {
	brokerages, err := repository.Brokerages.List()
	if err != nil {
		return nil, err
	}

	response := new(brokerageApi.Brokerages)
	mapper.Slice(brokerages, &response.Elements)
	return response, err
}

func (s *Service) Return(_ context.Context, req *brokerageApi.ReturnBrokerageReq) (*brokerageApi.Brokerage, error) {
	brokerageID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	brokerage, err := repository.Brokerages.ReturnByID(brokerageID)
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	mapper.Struct(brokerage, response)
	return response, err
}

func (s *Service) Delete(_ context.Context, req *brokerageApi.DeleteBrokerageReq) (*api.Void, error) {
	brokerageID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	if err := repository.Brokerages.Delete(brokerageID); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) ChangeStatus(ctx context.Context, req *brokerageApi.StatusChangeRequest) (*brokerageApi.BrokerageStatus, error) {
	//enables, err := repository.Brokerages.ReturnEnables()
	//if err != nil {
	//	return nil, err
	//}
	//for _, enable := range enables {
	//	for _, market := range enable.Markets {
	//		if _, err := s.candleService.DeleteMarket(ctx, &chipmunkApi.DeleteMarketRequest{
	//			MarketID: market.ID,
	//		}); err != nil {
	//			log.WithError(err).WithField("markets", market.ID).WithField("brokerage", enable.ID).Error("failed to stop markets")
	//		}
	//		if _, err = s.walletService.CancelWorker(ctx, new(api.Void)); err != nil {
	//			return nil, err
	//		}
	//	}
	//}
	//
	//brokerage, err := repository.Brokerages.ReturnByID(req.ID)
	//if err != nil {
	//	return nil, err
	//}
	//switch brokerage.Status {
	//case api.Status_Enable.String():
	//	brokerage.Status = api.Status_Disable.String()
	//case api.Status_Disable.String():
	//	brokerage.Status = api.Status_Enable.String()
	//}
	//if brokerage.Status == api.Status_Enable.String() {
	//	if req.OHLC {
	//		resolution := new(brokerageApi.Resolution)
	//		mapper.Struct(brokerage.Resolution, resolution)
	//
	//		for _, market := range brokerage.Markets {
	//			m := new(brokerageApi.Market)
	//			mapper.Struct(market, m)
	//			if _, err := s.ohlcService.AddMarket(ctx, &chipmunkApi.AddMarketRequest{
	//				BrokerageID: uint32(brokerage.ID),
	//				Market:      m,
	//			}); err != nil {
	//				log.WithError(err).WithField("markets", market.ID).WithField("brokerage", brokerage.ID).Error("failed to add markets")
	//				return nil, err
	//			}
	//		}
	//	}
	//	if req.Trading {
	//		if _, err = s.walletService.StartWorker(ctx, &chipmunkApi.StartWorkerRequest{
	//			BrokerageID: uint32(brokerage.ID),
	//		}); err != nil {
	//			log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to start wallet worker")
	//		}
	//		//todo: add trading worker
	//	}
	//}
	//if err := repository.Brokerages.ChangeStatus(brokerage); err != nil {
	//	return nil, err
	//}
	//return &brokerageApi.BrokerageStatus{Status: api.Status(api.Status_value[brokerage.Status])}, nil

	brokerageID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}

	enables, err := repository.Brokerages.ReturnEnables()
	if err != nil {
		return nil, err
	}
	for _, enable := range enables {
		if _, err = s.marketService.StopWorker(ctx, &chipmunkApi.WorkerStopReq{BrokerageID: enable.ID.String()}); err != nil {
			return nil, err
		}
	}

	brokerage, err := repository.Brokerages.ReturnByID(brokerageID)
	if err != nil {
		return nil, err
	}

	switch brokerage.Status {
	case api.Status_Enable:
		brokerage.Status = api.Status_Disable
	case api.Status_Disable:
		brokerage.Status = api.Status_Enable
	}
	if brokerage.Status == api.Status_Enable {
		if req.OHLC {
			_, err := s.marketService.StartWorker(ctx, &chipmunkApi.WorkerStartReq{
				BrokerageID:  req.ID,
				ResolutionID: brokerage.ResolutionID.String(),
				StrategyID:   brokerage.StrategyID.String()})
			if err != nil {
				return nil, err
			}
		}
		if req.Trading {
			if _, err = s.walletService.StartWorker(ctx, &chipmunkApi.StartWorkerRequest{
				BrokerageID: req.ID,
			}); err != nil {
				log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to start wallet worker")
			}
			//todo: add trading worker
		}
	}
	if err := repository.Brokerages.ChangeStatus(brokerage); err != nil {
		return nil, err
	}
	return &brokerageApi.BrokerageStatus{Status: brokerage.Status}, nil
}
