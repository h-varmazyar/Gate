package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/internal/repository"
	"github.com/h-varmazyar/Gate/services/indicators/internal/workers"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/calculator"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/storage"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	log               *log.Logger
	configs           Configs
	repository        repository.Repository
	calculatorWorker  *workers.IndicatorCalculator
	marketService     chipmunkAPI.MarketServiceClient
	resolutionService chipmunkAPI.ResolutionServiceClient
}

type Dependencies struct {
	CalculatorWorker *workers.IndicatorCalculator
}

var (
	grpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs Configs, dependencies *Dependencies) *Service {
	if grpcService == nil {
		chipmunkConn := grpcext.NewConnection(configs.ChipmunkAddress)
		grpcService = new(Service)
		grpcService.log = logger
		grpcService.configs = configs
		grpcService.calculatorWorker = dependencies.CalculatorWorker
		grpcService.marketService = chipmunkAPI.NewMarketServiceClient(chipmunkConn)
		grpcService.resolutionService = chipmunkAPI.NewResolutionServiceClient(chipmunkConn)
	}
	return grpcService
}

func (s Service) RegisterServer(server *grpc.Server) {
	indicatorsAPI.RegisterIndicatorServiceServer(server, s)
}

func (s Service) Register(ctx context.Context, req *indicatorsAPI.IndicatorRegisterReq) (*indicatorsAPI.Indicator, error) {
	if req.Type == indicatorsAPI.Type_NOTHING {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_indicator_type")
	}

	if _, err := uuid.Parse(req.MarketId); err != nil {
		return nil, err
	}

	market, err := s.marketService.Return(ctx, &chipmunkAPI.MarketReturnReq{ID: req.MarketId})
	if err != nil {
		return nil, err
	}

	if _, err = uuid.Parse(req.ResolutionId); err != nil {
		return nil, err
	}

	resolution, err := s.resolutionService.ReturnByID(ctx, &chipmunkAPI.ResolutionReturnByIDReq{ID: req.ResolutionId})
	if err != nil {
		return nil, err
	}

	indicator := new(entities.Indicator)
	mapper.Struct(req, indicator)

	if err = s.repository.Create(ctx, indicator); err != nil {
		return nil, err
	}

	calculatorIndicator, err := calculator.NewIndicator(ctx, indicator, market, resolution)
	if err != nil {
		return nil, err
	}
	s.calculatorWorker.AddIndicator(ctx, calculatorIndicator)

	res := new(indicatorsAPI.Indicator)
	mapper.Struct(indicator, res)
	return res, nil
}

func (s Service) Values(ctx context.Context, req *indicatorsAPI.IndicatorValuesReq) (*indicatorsAPI.IndicatorValues, error) {
	values, err := storage.GetValues(ctx, uint(req.Id), req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	indicatorValues := &indicatorsAPI.IndicatorValues{Values: values}
	return indicatorValues, nil
}
