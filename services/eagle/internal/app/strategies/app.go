package strategies

import (
	"context"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/repository"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/service"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/workers"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies/automatedStrategy"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
	logger  *log.Logger
	configs *Configs
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs *Configs) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		logger.WithError(err).Error("failed to create new repository")
		return nil, err
	}

	app := &App{
		logger:  logger,
		configs: configs,
	}
	var signalCheckerWorker *workers.SignalCheckWorker
	if signalCheckerWorker, err = app.initiateSignalCheckerWorker(ctx, repositoryInstance, logger); err != nil {
		logger.WithError(err).Error("failed to initiate signal checker worker")
		return nil, err
	}
	dependencies := &service.Dependencies{
		SignalCheckWorker: signalCheckerWorker,
	}

	app.Service = service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance, dependencies)
	return app, nil
}

func (app *App) initiateSignalCheckerWorker(ctx context.Context, db repository.StrategyRepository, logger *log.Logger) (*workers.SignalCheckWorker, error) {
	logger.Infof("initializing signal check worker")
	worker := workers.SignalCheckWorkerInstance(logger)

	marketService := chipmunkApi.NewMarketServiceClient(grpcext.NewConnection(app.configs.ChipmunkAddress))

	strategies, err := db.ReturnActives(ctx)
	if err != nil {
		app.logger.WithError(err).Error("failed to fetch active strategies")
		return nil, err
	}

	for _, strategy := range strategies {
		switch strategy.Type {
		case eagleApi.StrategyType_Automated:
			automated, err := automatedStrategy.NewAutomatedStrategy(strategy, app.configs.AutomatedWorker, logger)
			if err != nil {
				app.logger.WithError(err).Errorf("failed to create new instance of automated strategy")
				return nil, err
			}

			marketListReq := &chipmunkApi.MarketListReq{Platform: api.Platform_Coinex}
			markets, err := marketService.List(ctx, marketListReq)
			if err != nil {
				app.logger.WithError(err).Errorf("failed to get markets")
				return nil, err
			}

			worker.Start(automated, markets.Elements, strategy.BrokerageID)
		default:
		}
	}
	return worker, nil
}
