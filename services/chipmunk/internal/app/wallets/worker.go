package wallets

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"time"
)

type worker struct {
	WalletService brokerageApi.WalletServiceClient
	MarketService brokerageApi.MarketServiceClient
	cancellation  context.CancelFunc
}

var (
	Worker *worker
)

func init() {
	Worker = new(worker)
	brokerageApiConnection := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
	Worker.WalletService = brokerageApi.NewWalletServiceClient(brokerageApiConnection)
	Worker.MarketService = brokerageApi.NewMarketServiceClient(brokerageApiConnection)
}

func (w *worker) Start(brokerage *brokerageApi.Brokerage) error {
	if w.cancellation == nil {
		ctx, fn := context.WithCancel(context.Background())
		w.cancellation = fn

		if brokerage == nil {
			return errors.New(context.Background(), codes.NotFound)
		}

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
	ticker := time.NewTicker(configs.Variables.WalletWorkerHeartbeat)

LOOP:
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			break LOOP
		case <-ticker.C:
			wallets, err := w.WalletService.UpdateWallets(ctx, &brokerageApi.UpdateWalletRequest{BrokerageID: brokerage.ID})
			if err != nil {
				log.WithError(err).Error("failed to update wallets")
			} else {
				buffer.Wallets.AddOrUpdateList(wallets.Wallets)
				w.calculateReferenceValue(ctx, brokerage, wallets.Wallets)
			}
		}
	}
}

func (w *worker) calculateReferenceValue(ctx context.Context, brokerage *brokerageApi.Brokerage, wallets []*brokerageApi.Wallet) {
	references := make(map[string]*buffer.Reference)
	for _, wallet := range wallets {
		list, err := w.MarketService.ReturnBySource(ctx, &brokerageApi.MarketListBySourceRequest{
			BrokerageName: brokerage.Name.String(),
			Source:        wallet.AssetName,
		})
		if err != nil {
			log.WithError(err).Error("failed to get market list")
			continue
		}
		for _, market := range list.Markets {
			statistics, err := w.MarketService.MarketStatistics(ctx, &brokerageApi.MarketStatisticsRequest{
				BrokerageID: brokerage.ID,
				MarketName:  market.Name,
			})
			if err != nil {
				log.WithError(err).Error("failed to fetch market statistics")
				continue
			}
			reference, ok := references[market.Destination.Name]
			if reference == nil || !ok {
				reference = new(buffer.Reference)
			}
			reference.Blocked += statistics.Close * wallet.BlockedBalance
			reference.Active += statistics.Close * wallet.ActiveBalance
			reference.Total += statistics.Close * wallet.TotalBalance
			references[market.Destination.Name] = reference
		}
	}
	buffer.Wallets.UpdateReferences(references)
}
