package service

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	brokerageApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/repository"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/entity"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	marketService chipmunkApi.MarketServiceClient
	walletService chipmunkApi.WalletsServiceClient
	signalService eagleApi.SignalServiceClient
	logger        *log.Logger
	db            repository.BrokerageRepository
}

var (
	grpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, db repository.BrokerageRepository) *Service {
	if grpcService == nil {
		grpcService = new(Service)
		chipmunkConnection := grpcext.NewConnection(configs.ChipmunkGrpcAddress)
		eagleConnection := grpcext.NewConnection(configs.EagleGrpcAddress)

		grpcService.marketService = chipmunkApi.NewMarketServiceClient(chipmunkConnection)
		grpcService.walletService = chipmunkApi.NewWalletsServiceClient(chipmunkConnection)
		grpcService.signalService = eagleApi.NewSignalServiceClient(eagleConnection)
		grpcService.logger = logger
		grpcService.db = db
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
	if _, ok := api.Platform_value[req.Platform.String()]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong_platform")
	}
	if _, err := uuid.Parse(req.ResolutionID); err != nil {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_resolution")
	}
	brokerage := new(entity.Brokerage)
	mapper.Struct(req, brokerage)

	brokerage.Status = api.Status_Disable

	if err := s.db.Create(brokerage); err != nil {
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

	enables, err := s.db.ReturnEnables()
	if err != nil {
		return nil, err
	}
	for _, enable := range enables {
		if _, err = s.marketService.StopWorker(ctx, &chipmunkApi.WorkerStopReq{Platform: enable.ID.String()}); err != nil {
			return nil, err
		}
	}

	brokerage, err := s.db.ReturnByID(brokerageID)
	if err != nil {
		return nil, err
	}

	brokerage.Status = api.Status_Enable
	//if err := repository.Brokerages.ChangeStatus(core.ID); err != nil {
	//	return nil, err
	//}

	if req.CollectMarketsData {
		_, err := s.marketService.StartWorker(ctx, &chipmunkApi.WorkerStartReq{
			BrokerageID:  req.ID,
			ResolutionID: brokerage.ResolutionID.String(),
			StrategyID:   brokerage.StrategyID.String()})
		if err != nil {
			//if statusErr := repository.Brokerages.ChangeStatus(core.ID); statusErr != nil {
			//	log.WithError(statusErr).Errorf("failed ot change status of core %v to %v", core.ID, core.Status)
			//}
			return nil, err
		}
	}
	if req.StartTrading {
		if _, err = s.walletService.StartWorker(ctx, &chipmunkApi.StartWorkerRequest{
			BrokerageID: req.ID,
		}); err != nil {
			log.WithError(err).WithField("core", brokerage.ID).Error("failed to start wallet workers")
			//if statusErr := repository.Brokerages.ChangeStatus(core.ID); statusErr != nil {
			//	log.WithError(statusErr).Errorf("failed ot change status of core %v to %v", core.ID, core.Status)
			//}
			return nil, err
		}
		if _, err = s.signalService.Start(ctx, &eagleApi.SignalStartReq{
			BrokerageID: brokerage.ID.String(),
			StrategyID:  brokerage.StrategyID.String(),
			WithTrading: false,
		}); err != nil {
			if _, marketErr := s.marketService.StopWorker(ctx, &chipmunkApi.WorkerStopReq{Platform: brokerage.Platform}); marketErr != nil {
				log.WithError(marketErr).Errorf("failed to stop market workers for core %v", brokerageID)
			}
			if _, walletErr := s.walletService.StopWorker(ctx, new(api.Void)); walletErr != nil {
				log.WithError(walletErr).Errorf("failed to stop wallet workers for core %v", brokerageID)
			}
			//if statusErr := repository.Brokerages.ChangeStatus(core.ID); statusErr != nil {
			//	log.WithError(statusErr).Errorf("failed ot change status of core %v to %v", core.ID, core.Status)
			//}
			return nil, err
		}
	}
	if err := s.db.ChangeStatus(brokerage.ID); err != nil {
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

	brokerage, err := s.db.ReturnByID(brokerageID)
	if err != nil {
		return nil, err
	}

	if _, err = s.marketService.StopWorker(ctx, &chipmunkApi.WorkerStopReq{
		Platform: brokerage.Platform,
	}); err != nil {
		return nil, err
	}
	if _, err = s.walletService.StopWorker(ctx, &api.Void{}); err != nil {
		log.WithError(err).WithField("core", brokerage.ID).Error("failed to stop wallet workers")
	}
	if _, err = s.signalService.Stop(ctx, &api.Void{}); err != nil {
		log.WithError(err).WithField("core", brokerage.ID).Error("failed to stop signal workers")
	}

	brokerage.Status = api.Status_Enable
	if err := s.db.ChangeStatus(brokerage.ID); err != nil {
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
	brokerage, err := s.db.ReturnByID(brokerageID)
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	mapper.Struct(brokerage, response)
	return response, err
}

func (s *Service) Enable(_ context.Context, _ *api.Void) (*brokerageApi.Brokerage, error) {
	brokerage, err := s.db.ReturnEnable()
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
	if err := s.db.Delete(brokerageID); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) List(_ context.Context, _ *api.Void) (*brokerageApi.Brokerages, error) {
	brokerages, err := s.db.List()
	if err != nil {
		return nil, err
	}

	response := new(brokerageApi.Brokerages)
	mapper.Slice(brokerages, &response.Elements)
	return response, err
}
