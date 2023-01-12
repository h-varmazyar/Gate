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
	marketService   chipmunkApi.MarketServiceClient
	walletService   chipmunkApi.WalletsServiceClient
	strategyService eagleApi.StrategyServiceClient
	logger          *log.Logger
	db              repository.BrokerageRepository
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
		grpcService.strategyService = eagleApi.NewStrategyServiceClient(eagleConnection)
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
	brokerage := new(entity.Brokerage)
	mapper.Struct(req, brokerage)

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

	brokerage, err := s.db.ReturnByID(brokerageID)
	if err != nil {
		return nil, err
	}

	if brokerage.Status == api.Status_Enable {
		err := errors.New(ctx, codes.FailedPrecondition).AddDetailF("brokerage %v started before", req.ID)
		s.logger.WithError(err)
		return nil, err
	}
	if req.WithTrading {
		tradingMarkets, err := s.marketService.List(ctx, &chipmunkApi.MarketListReq{
			Platform: brokerage.Platform,
		})
		if err != nil {
			log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to get markets")
			return nil, err
		}

		if _, err = s.walletService.StartWorker(ctx, &chipmunkApi.StartWorkerRequest{
			BrokerageID: req.ID,
		}); err != nil {
			log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to start wallet workers")
			return nil, err
		}

		if _, err = s.strategyService.StartSignalChecker(ctx, &eagleApi.StrategySignalCheckStartReq{
			Platform:    brokerage.Platform,
			BrokerageID: brokerage.ID.String(),
			StrategyID:  brokerage.StrategyID.String(),
			WithTrading: req.WithTrading,
			Markets:     tradingMarkets,
		}); err != nil {
			if _, walletErr := s.walletService.StopWorker(ctx, new(api.Void)); walletErr != nil {
				log.WithError(walletErr).Errorf("failed to stop wallet workers for core %v", brokerageID)
			}
			return nil, err
		}
	}
	brokerage.Status = api.Status_Enable
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

	if _, err = s.walletService.StopWorker(ctx, &api.Void{}); err != nil {
		log.WithError(err).WithField("core", brokerage.ID).Error("failed to stop wallet workers")
		return nil, err
	}
	if _, err = s.strategyService.StopSignalChecker(ctx, &eagleApi.StrategySignalCheckStopReq{
		BrokerageID: req.ID,
	}); err != nil {
		log.WithError(err).WithField("core", brokerage.ID).Error("failed to stop signal workers")
		return nil, err
	}

	brokerage.Status = api.Status_Disable
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
