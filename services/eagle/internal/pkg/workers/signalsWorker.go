package workers

import (
	"context"
	"errors"
	"fmt"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/buffers"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/indicators"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies"
	"time"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 04.12.21
* Github: https://github.com/h-varmazyar
* Email: hossein.varmazyar@yahoo.com
**/

type signalsWorker struct {
	heartbeatInterval time.Duration
}

type SignalsSettings struct {
	Context    context.Context
	Market     *api.Market
	Resolution *api.Resolution
	Config     *indicators.Configuration
	Strategy   strategies.Strategy
}

var (
	SignalsWorker              *signalsWorker
	signalsWorkerCancellations map[string]context.CancelFunc
)

func init() {
	signalsWorkerCancellations = make(map[string]context.CancelFunc)
	IndicatorWorker = new(indicatorWorker)
	IndicatorWorker.heartbeatInterval = configs.Variables.SignalsWorkerHeartbeat
}

func (worker *signalsWorker) AddMarket(settings *SignalsSettings) {
	go worker.run(settings)
}

func (worker *signalsWorker) CancelWorker(marketID, resolutionID string) error {
	fn, ok := signalsWorkerCancellations[fmt.Sprintf("%s > %s", marketID, resolutionID)]
	if !ok {
		return errors.New("worker stopped before")
	}
	fn()
	delete(signalsWorkerCancellations, fmt.Sprintf("%s > %s", marketID, resolutionID))
	buffers.Candles.RemoveList(marketID, resolutionID)
	return nil
}

func (worker *signalsWorker) run(settings *SignalsSettings) {
	ticker := time.NewTicker(worker.heartbeatInterval)

LOOP:
	for {
		select {
		case <-settings.Context.Done():
			break LOOP
		case <-ticker.C:
			//todo: check wallet "walletResponse.Wallet.ActiveBalance > 0"
			worker.checkForSignals(settings)
		}
	}
}

func (worker *signalsWorker) checkForSignals(settings *SignalsSettings) bool {
	//walletResponse := thread.Requests.WalletInfo(brokerages.WalletInfoParams{
	//	WalletName: thread.Market.Destination.Symbol,
	//})
	//if walletResponse.Error != nil {
	//	return
	//}
	//if walletResponse.Wallet.ActiveBalance > 0 {
	candles := buffers.Candles.GetLastNCandle(2, settings.Market.ID, settings.Resolution.ID)
	rsi := settings.Strategy.RsiSignal(candles[0], candles[1])
	bb := settings.Strategy.BollingerBandSignal(float64(settings.Market.MakerFeeRate), float64(settings.Market.TakerFeeRate), candles[0], candles[1])
	stochastic := settings.Strategy.StochasticSignal(candles[0], candles[1])
	return rsi+bb+stochastic == 3
	//}
}
