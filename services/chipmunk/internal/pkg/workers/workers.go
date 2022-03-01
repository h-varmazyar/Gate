package workers

import (
	"context"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/chipmunk/configs"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/workers/OHLC"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/workers/wallet"
)

var (
	WalletWorker *wallet.Worker
	OHLCWorker   *OHLC.Worker
)

func init() {
	{
		WalletWorker = new(wallet.Worker)
		WalletWorker.HeartbeatInterval = configs.Variables.WalletWorkerHeartbeat
		brokerageApiConnection := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		WalletWorker.WalletService = brokerageApi.NewWalletServiceClient(brokerageApiConnection)
		WalletWorker.MarketService = brokerageApi.NewMarketServiceClient(brokerageApiConnection)
	}
	{
		OHLCWorker = new(OHLC.Worker)
		OHLCWorker.HeartbeatInterval = configs.Variables.OHLCWorkerHeartbeat
		candleConnection := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		OHLCWorker.CandleService = brokerageApi.NewCandleServiceClient(candleConnection)
		OHLCWorker.Cancellations = make(map[string]context.CancelFunc)
	}
}
