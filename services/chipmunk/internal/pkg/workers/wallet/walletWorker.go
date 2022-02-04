package wallet

import (
	"context"
	"github.com/mrNobody95/Gate/pkg/errors"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/buffer"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"time"
)

type Worker struct {
	HeartbeatInterval time.Duration
	WalletService     brokerageApi.WalletServiceClient
	MarketService     brokerageApi.MarketServiceClient
	cancellation      context.CancelFunc
}

func (w *Worker) Start(brokerage *brokerageApi.Brokerage) error {
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

func (w *Worker) Stop() {
	if w.cancellation != nil {
		w.cancellation()
		w.cancellation = nil
		buffer.Wallets.Flush()
	}
}

func (w *Worker) run(ctx context.Context, brokerage *brokerageApi.Brokerage) {
	ticker := time.NewTicker(w.HeartbeatInterval)

LOOP:
	for {
		select {
		case <-ctx.Done():
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

func (w *Worker) calculateReferenceValue(ctx context.Context, brokerage *brokerageApi.Brokerage, wallets []*brokerageApi.Wallet) {
	references := make(map[string]*buffer.Reference)
	for _, wallet := range wallets {
		list, err := w.MarketService.ReturnBySource(ctx, &brokerageApi.MarketListBySourceRequest{
			BrokerageName: brokerage.Name.String(),
			Source:        wallet.AssetName,
		})
		if err != nil {
			continue
		}
		for _, market := range list.Markets {
			statistics, err := w.MarketService.MarketStatistics(ctx, &brokerageApi.MarketStatisticsRequest{
				BrokerageID: brokerage.ID,
				MarketName:  market.Name,
			})
			if err != nil {
				continue
			}
			reference := references[market.Destination.Name]
			if reference == nil {
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
