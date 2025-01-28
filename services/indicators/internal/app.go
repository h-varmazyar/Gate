package internal

import (
	"context"
	"github.com/h-varmazyar/Gate/services/indicators/internal/service"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/db"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

type AppDependencies struct {
}

func NewApp(ctx context.Context, logger *log.Logger, configs Configs, _ *AppDependencies, dbInstance *db.DB) (*App, error) {
	//calculatorWorker := workers.NewIndicatorCalculatorWorker(ctx, logger, configs.WorkersConfigs)
	//
	//repository := db.NewIndicators(ctx, dbInstance)
	//
	//indicators, err := repository.List(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//
	//chipmunkConn := grpcext.NewConnection(configs.ChipmunkAddress)
	//marketService := chipmunkAPI.NewMarketServiceClient(chipmunkConn)
	//resolutionService := chipmunkAPI.NewResolutionServiceClient(chipmunkConn)
	//
	//workerIndicators := make([]calculator.Indicator, 0)
	//for _, indicator := range indicators {
	//	market, err := marketService.Return(ctx, &chipmunkAPI.MarketReturnReq{ID: indicator.MarketId.String()})
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	resolution, err := resolutionService.ReturnByID(ctx, &chipmunkAPI.ResolutionReturnByIDReq{ID: indicator.ResolutionId.String()})
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	workerIndicator, err := calculator.NewIndicator(ctx, indicator, market, resolution)
	//	if err != nil {
	//		return nil, err
	//	}
	//	workerIndicators = append(workerIndicators, workerIndicator)
	//}
	//
	//calculatorWorker.Start(ctx, workerIndicators)
	//
	//serviceDependencies := &service.Dependencies{CalculatorWorker: calculatorWorker}
	//return &App{
	//	Service: service.NewService(ctx, logger, configs.ServiceConfigs, serviceDependencies),
	//}, nil

	return nil, nil
}
