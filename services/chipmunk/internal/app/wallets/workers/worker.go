package workers

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	marketsService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/buffer"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"time"
)

type WalletCheck struct {
	functionsService coreApi.FunctionsServiceClient
	marketService    *marketsService.Service
	cancellation     context.CancelFunc
	configs          *Configs
	buffer           *buffer.WalletBuffer
}

func InitializeWorker(_ context.Context, configs *Configs, marketService *marketsService.Service) *WalletCheck {
	Worker := new(WalletCheck)
	coreConn := grpcext.NewConnection(configs.CoreAddress)
	Worker.functionsService = coreApi.NewFunctionsServiceClient(coreConn)
	Worker.marketService = marketService
	Worker.configs = configs
	return Worker
}

func (w *WalletCheck) Start(brokerage *coreApi.Brokerage) error {
	if w.cancellation == nil {
		if brokerage == nil {
			return errors.New(context.Background(), codes.NotFound)
		}

		ctx, fn := context.WithCancel(context.Background())
		w.cancellation = fn

		go w.run(ctx, brokerage)
		return nil
	}
	return errors.New(context.Background(), codes.AlreadyExists)
}

func (w *WalletCheck) Stop() {
	if w.cancellation != nil {
		w.cancellation()
		w.cancellation = nil
		w.buffer.Flush()
	}
}

func (w *WalletCheck) run(ctx context.Context, brokerage *coreApi.Brokerage) {
	time.Sleep(time.Second)
	ticker := time.NewTicker(w.configs.WalletWorkerInterval)

LOOP:
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			break LOOP
		case <-ticker.C:
			wallets, err := w.functionsService.WalletsBalance(ctx, &coreApi.WalletsBalanceReq{BrokerageID: brokerage.ID})
			if err != nil {
				log.WithError(err).Error("failed to update wallets")
			} else {
				w.buffer.AddOrUpdateList(wallets.Elements)
				w.calculateReferenceValue(ctx, brokerage, wallets.Elements)
			}
		}
	}
}

func (w *WalletCheck) calculateReferenceValue(ctx context.Context, brokerage *coreApi.Brokerage, wallets []*chipmunkApi.Wallet) {
	references := make(map[string]*chipmunkApi.Reference)
	for _, wallet := range wallets {
		list, err := w.marketService.ListBySource(ctx, &chipmunkApi.MarketListBySourceReq{
			Platform: brokerage.Platform,
			Source:   wallet.AssetName,
		})
		if err != nil {
			log.WithError(err).Error("failed to get markets list")
			continue
		}
		for _, market := range list.Elements {
			statistics, err := w.functionsService.SingleMarketStatistics(ctx, &coreApi.MarketStatisticsReq{
				MarketName: market.Name,
				Platform:   brokerage.Platform,
			})
			if err != nil {
				log.WithError(err).Error("failed to fetch markets statistics")
				continue
			}
			reference, ok := references[market.Destination.Name]
			if reference == nil || !ok {
				reference = new(chipmunkApi.Reference)
			}
			reference.AssetName = market.Destination.Name
			reference.BlockedBalance += statistics.Close * wallet.BlockedBalance
			reference.ActiveBalance += statistics.Close * wallet.ActiveBalance
			reference.TotalBalance += statistics.Close * wallet.TotalBalance
			references[market.Destination.Name] = reference
		}
	}
	w.buffer.UpdateReferences(references)
}
