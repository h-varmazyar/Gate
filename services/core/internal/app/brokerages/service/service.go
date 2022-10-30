package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	brokerageApi "github.com/h-varmazyar/Gate/services/core/api"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/entity"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	marketService chipmunkApi.MarketServiceClient
	walletService chipmunkApi.WalletsServiceClient
	signalService eagleApi.SignalServiceClient
	logger        *log.Logger
}

var (
	grpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs) *Service {
	if grpcService == nil {
		grpcService = new(Service)
		chipmunkConnection := grpcext.NewConnection(configs.ChipmunkGrpcAddress)
		eagleConnection := grpcext.NewConnection(configs.EagleGrpcAddress)

		grpcService.marketService = chipmunkApi.NewMarketServiceClient(chipmunkConnection)
		grpcService.walletService = chipmunkApi.NewWalletsServiceClient(chipmunkConnection)
		grpcService.signalService = eagleApi.NewSignalServiceClient(eagleConnection)
		grpcService.logger = logger
	}
	return grpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	brokerageApi.RegisterBrokerageServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, req *brokerageApi.BrokerageCreateReq) (*brokerageApi.Brokerage, error) {
	if _, ok := api.AuthType_value[req.Auth.Type.String()]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong_auth_type")
	}
	if _, ok := brokerageApi.Platform_value[req.Platform.String()]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong_platform")
	}
	if _, err := uuid.Parse(req.ResolutionID); err != nil {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_resolution")
	}
	brokerage := new(entity.Brokerage)
	mapper.Struct(req, brokerage)

	brokerage.Status = api.Status_Disable

	if err := repository.Brokerages.Create(brokerage); err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	mapper.Struct(brokerage, response)
	return response, nil
}

func (s *Service) Start(ctx context.Context, req *brokerageApi.BrokerageStartReq) (*brokerageApi.Brokerage, error) {
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

	brokerage.Status = api.Status_Enable
	//if err := repository.Brokerages.ChangeStatus(brokerage.ID); err != nil {
	//	return nil, err
	//}

	if req.CollectMarketsData {
		_, err := s.marketService.StartWorker(ctx, &chipmunkApi.WorkerStartReq{
			BrokerageID:  req.ID,
			ResolutionID: brokerage.ResolutionID.String(),
			StrategyID:   brokerage.StrategyID.String()})
		if err != nil {
			//if statusErr := repository.Brokerages.ChangeStatus(brokerage.ID); statusErr != nil {
			//	log.WithError(statusErr).Errorf("failed ot change status of brokerage %v to %v", brokerage.ID, brokerage.Status)
			//}
			return nil, err
		}
	}
	if req.StartTrading {
		if _, err = s.walletService.StartWorker(ctx, &chipmunkApi.StartWorkerRequest{
			BrokerageID: req.ID,
		}); err != nil {
			log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to start wallet worker")
			//if statusErr := repository.Brokerages.ChangeStatus(brokerage.ID); statusErr != nil {
			//	log.WithError(statusErr).Errorf("failed ot change status of brokerage %v to %v", brokerage.ID, brokerage.Status)
			//}
			return nil, err
		}
		if _, err = s.signalService.Start(ctx, &eagleApi.SignalStartReq{
			BrokerageID: brokerage.ID.String(),
			StrategyID:  brokerage.StrategyID.String(),
			WithTrading: false,
		}); err != nil {
			if _, marketErr := s.marketService.StopWorker(ctx, &chipmunkApi.WorkerStopReq{BrokerageID: brokerageID.String()}); marketErr != nil {
				log.WithError(marketErr).Errorf("failed to stop market worker for brokerage %v", brokerageID)
			}
			if _, walletErr := s.walletService.StopWorker(ctx, new(api.Void)); walletErr != nil {
				log.WithError(walletErr).Errorf("failed to stop wallet worker for brokerage %v", brokerageID)
			}
			//if statusErr := repository.Brokerages.ChangeStatus(brokerage.ID); statusErr != nil {
			//	log.WithError(statusErr).Errorf("failed ot change status of brokerage %v to %v", brokerage.ID, brokerage.Status)
			//}
			return nil, err
		}
	}
	if err := repository.Brokerages.ChangeStatus(brokerage.ID); err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	mapper.Struct(brokerage, response)
	return response, nil
}

func (s *Service) Stop(ctx context.Context, req *brokerageApi.BrokerageStopReq) (*brokerageApi.Brokerage, error) {
	brokerageID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}

	brokerage, err := repository.Brokerages.ReturnByID(brokerageID)
	if err != nil {
		return nil, err
	}

	if _, err = s.marketService.StopWorker(ctx, &chipmunkApi.WorkerStopReq{
		BrokerageID: req.ID,
	}); err != nil {
		return nil, err
	}
	if _, err = s.walletService.StopWorker(ctx, &api.Void{}); err != nil {
		log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to stop wallet worker")
	}
	if _, err = s.signalService.Stop(ctx, &api.Void{}); err != nil {
		log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to stop signal worker")
	}

	brokerage.Status = api.Status_Enable
	if err := repository.Brokerages.ChangeStatus(brokerage.ID); err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	mapper.Struct(brokerage, response)
	return response, nil
}

func (s *Service) Return(_ context.Context, req *brokerageApi.BrokerageReturnReq) (*brokerageApi.Brokerage, error) {
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

func (s *Service) Enable(_ context.Context, req *api.Void) (*brokerageApi.Brokerage, error) {
	brokerage, err := repository.Brokerages.ReturnEnable()
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	mapper.Struct(brokerage, response)
	return response, err
}

func (s *Service) Delete(ctx context.Context, req *brokerageApi.BrokerageDeleteReq) (*api.Void, error) {
	brokerageID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	if _, err = s.Stop(ctx, &brokerageApi.BrokerageStopReq{ID: req.ID}); err != nil {
		return nil, err
	}
	if err := repository.Brokerages.Delete(brokerageID); err != nil {
		return nil, err
	}
	return new(api.Void), nil
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

//func (s *Service) ChangeStatus(ctx context.Context, req *brokerageApi.StatusChangeRequest) (*brokerageApi.BrokerageStatus, error) {
//	//enables, err := repository.Brokerages.ReturnEnables()
//	//if err != nil {
//	//	return nil, err
//	//}
//	//for _, enable := range enables {
//	//	for _, market := range enable.Markets {
//	//		if _, err := s.candleService.DeleteMarket(ctx, &chipmunkApi.DeleteMarketRequest{
//	//			MarketID: market.ID,
//	//		}); err != nil {
//	//			log.WithError(err).WithField("markets", market.ID).WithField("brokerage", enable.ID).Error("failed to stop markets")
//	//		}
//	//		if _, err = s.walletService.CancelWorker(ctx, new(api.Void)); err != nil {
//	//			return nil, err
//	//		}
//	//	}
//	//}
//	//
//	//brokerage, err := repository.Brokerages.ReturnByID(req.ID)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//switch brokerage.Status {
//	//case api.Status_Enable.String():
//	//	brokerage.Status = api.Status_Disable.String()
//	//case api.Status_Disable.String():
//	//	brokerage.Status = api.Status_Enable.String()
//	//}
//	//if brokerage.Status == api.Status_Enable.String() {
//	//	if req.OHLC {
//	//		resolution := new(brokerageApi.Resolution)
//	//		mapper.Struct(brokerage.Resolution, resolution)
//	//
//	//		for _, market := range brokerage.Markets {
//	//			m := new(brokerageApi.Market)
//	//			mapper.Struct(market, m)
//	//			if _, err := s.ohlcService.AddMarket(ctx, &chipmunkApi.AddMarketRequest{
//	//				BrokerageID: uint32(brokerage.ID),
//	//				Market:      m,
//	//			}); err != nil {
//	//				log.WithError(err).WithField("markets", market.ID).WithField("brokerage", brokerage.ID).Error("failed to add markets")
//	//				return nil, err
//	//			}
//	//		}
//	//	}
//	//	if req.Trading {
//	//		if _, err = s.walletService.StartWorker(ctx, &chipmunkApi.StartWorkerRequest{
//	//			BrokerageID: uint32(brokerage.ID),
//	//		}); err != nil {
//	//			log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to start wallet worker")
//	//		}
//	//		//todo: add trading worker
//	//	}
//	//}
//	//if err := repository.Brokerages.ChangeStatus(brokerage); err != nil {
//	//	return nil, err
//	//}
//	//return &brokerageApi.BrokerageStatus{Status: api.Status(api.Status_value[brokerage.Status])}, nil
//
//	brokerageID, err := uuid.Parse(req.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	enables, err := repository.Brokerages.ReturnEnables()
//	if err != nil {
//		return nil, err
//	}
//	for _, enable := range enables {
//		if _, err = s.marketService.StopWorker(ctx, &chipmunkApi.WorkerStopReq{BrokerageID: enable.ID.String()}); err != nil {
//			return nil, err
//		}
//	}
//
//	brokerage, err := repository.Brokerages.ReturnByID(brokerageID)
//	if err != nil {
//		return nil, err
//	}
//
//	switch brokerage.Status {
//	case api.Status_Enable:
//		brokerage.Status = api.Status_Disable
//	case api.Status_Disable:
//		brokerage.Status = api.Status_Enable
//	}
//	if brokerage.Status == api.Status_Enable {
//		if req.OHLC {
//			_, err := s.marketService.StartWorker(ctx, &chipmunkApi.WorkerStartReq{
//				BrokerageID:  req.ID,
//				ResolutionID: brokerage.ResolutionID.String(),
//				StrategyID:   brokerage.StrategyID.String()})
//			if err != nil {
//				return nil, err
//			}
//		}
//		if req.Trading {
//			if _, err = s.walletService.StartWorker(ctx, &chipmunkApi.StartWorkerRequest{
//				BrokerageID: req.ID,
//			}); err != nil {
//				log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to start wallet worker")
//			}
//			//todo: add trading worker
//		}
//	}
//	if err := repository.Brokerages.ChangeStatus(brokerage); err != nil {
//		return nil, err
//	}
//	return &brokerageApi.BrokerageStatus{Status: brokerage.Status}, nil
//}
