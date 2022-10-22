package wallets

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"time"
)

type worker struct {
	functionsService brokerageApi.FunctionsServiceClient
	marketService    chipmunkApi.MarketServiceClient
	cancellation     context.CancelFunc
}

var (
	Worker *worker
)

func InitializeWorker() {
	Worker = new(worker)
	brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
	chipmunkConn := grpcext.NewConnection(fmt.Sprintf(":%v", configs.Variables.GrpcPort))
	Worker.functionsService = brokerageApi.NewFunctionsServiceClient(brokerageConn)
	Worker.marketService = chipmunkApi.NewMarketServiceClient(chipmunkConn)
}

func (w *worker) Start(brokerage *brokerageApi.Brokerage) error {
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

func (w *worker) Stop() {
	if w.cancellation != nil {
		w.cancellation()
		w.cancellation = nil
		buffer.Wallets.Flush()
	}
}

func (w *worker) run(ctx context.Context, brokerage *brokerageApi.Brokerage) {
	time.Sleep(time.Second)
	ticker := time.NewTicker(configs.Variables.WalletWorkerHeartbeat)

LOOP:
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			break LOOP
		case <-ticker.C:
			wallets, err := w.functionsService.WalletsBalance(ctx, new(api.Void))
			if err != nil {
				log.WithError(err).Error("failed to update wallets")
			} else {
				buffer.Wallets.AddOrUpdateList(wallets.Elements)
				w.calculateReferenceValue(ctx, brokerage, wallets.Elements)
			}
		}
	}
}

func (w *worker) calculateReferenceValue(ctx context.Context, brokerage *brokerageApi.Brokerage, wallets []*chipmunkApi.Wallet) {
	references := make(map[string]*chipmunkApi.Reference)
	for _, wallet := range wallets {
		list, err := w.marketService.ReturnBySource(ctx, &chipmunkApi.MarketListBySourceRequest{
			BrokerageID: brokerage.ID,
			Source:      wallet.AssetName,
		})
		if err != nil {
			log.WithError(err).Error("failed to get markets list")
			continue
		}
		for _, market := range list.Elements {
			statistics, err := w.functionsService.MarketStatistics(ctx, &brokerageApi.MarketStatisticsReq{
				MarketName: market.Name,
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
	buffer.Wallets.UpdateReferences(references)
}
