package wallets

import (
	"context"
	marketsService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/workers"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

type AppDependencies struct {
	MarketService *marketsService.Service
}

func NewApp(ctx context.Context, logger *log.Logger, configs *Configs, dependencies *AppDependencies) (*App, error) {
	walletBuffer := buffer.NewWalletInstance(configs.BufferConfigs)

	walletCheckWorker := workers.InitializeWorker(ctx, configs.WorkerConfigs, dependencies.MarketService)

	serviceDependencies := &service.Dependencies{
		Buffer: walletBuffer,
		Worker: walletCheckWorker,
	}

	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, serviceDependencies),
	}, nil
}
