package workers

import (
	"context"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	"github.com/mrNobody95/Gate/pkg/mapper"
	chipmunkApi "github.com/mrNobody95/Gate/services/chipmunk/api"
	"github.com/mrNobody95/Gate/services/eagle/configs"
	"github.com/mrNobody95/Gate/services/eagle/internal/pkg/buffers"
	"github.com/mrNobody95/Gate/services/eagle/internal/pkg/indicators"
	"github.com/mrNobody95/Gate/services/eagle/internal/pkg/models"
	log "github.com/sirupsen/logrus"
	"sync"
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
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type indicatorWorker struct {
	heartbeatInterval time.Duration
	ohlcService       chipmunkApi.OhlcServiceClient
}

type Settings struct {
	Context    context.Context
	Market     *api.Market
	Resolution *api.Resolution
	Config     *indicators.Configuration
}

var (
	IndicatorWorker     *indicatorWorker
	workerCancellations map[string]context.CancelFunc
)

func init() {
	workerCancellations = make(map[string]context.CancelFunc)
	IndicatorWorker = new(indicatorWorker)
	IndicatorWorker.heartbeatInterval = configs.Variables.IndicatorsWorkerHeartbeat
	candleConnection := grpcext.NewConnection(fmt.Sprintf(":%v", configs.Variables.GrpcAddresses.Chipmunk))
	IndicatorWorker.ohlcService = chipmunkApi.NewOhlcServiceClient(candleConnection)
}

func (worker *indicatorWorker) AddMarket(settings *Settings) {
	go worker.run(settings)
}

func (worker *indicatorWorker) CancelWorker(marketID, resolutionID string) error {
	fn, ok := workerCancellations[fmt.Sprintf("%s > %s", marketID, resolutionID)]
	if !ok {
		return errors.New("worker stopped before")
	}
	fn()
	delete(workerCancellations, fmt.Sprintf("%s > %s", marketID, resolutionID))
	buffers.Candles.RemoveList(marketID, resolutionID)
	return nil
}

func (worker *indicatorWorker) run(settings *Settings) {
	ticker := time.NewTicker(worker.heartbeatInterval)

	if err := worker.initiateBuffer(settings); err != nil {
		log.WithError(err).Error("initiate buffer")
		return
	}

	if err := worker.initiateIndicators(settings); err != nil {
		log.WithError(err).Error("initiate indicators")
		return
	}
LOOP:
	for {
		select {
		case <-settings.Context.Done():
			break LOOP
		case <-ticker.C:
			if err := worker.updateBuffer(settings); err != nil {
				log.WithError(err).Error("updating candle failed")
				continue
			}
			worker.calculateIndicators(settings)
		}
	}
}

func (worker *indicatorWorker) initiateBuffer(settings *Settings) error {
	candles, err := worker.ohlcService.ReturnBufferedCandles(settings.Context, &chipmunkApi.BufferedCandlesRequest{
		ResolutionID: settings.Resolution.ID,
		MarketID:     settings.Market.ID,
	})
	if err != nil {
		return err
	}
	buffers.Candles.AddList(settings.Market.ID, settings.Resolution.ID)
	for _, candle := range candles.Candles {
		tmp := new(models.Candle)
		mapper.Struct(candle, tmp)
		buffers.Candles.Enqueue(tmp)
	}
	return nil
}

func (worker *indicatorWorker) updateBuffer(settings *Settings) error {
	candle, err := worker.ohlcService.ReturnLastCandle(settings.Context, &chipmunkApi.BufferedCandlesRequest{
		ResolutionID: settings.Resolution.ID,
		MarketID:     settings.Market.ID,
	})
	if err != nil {
		return err
	}
	tmp := new(models.Candle)
	mapper.Struct(candle, tmp)
	buffers.Candles.Enqueue(tmp)
	return nil
}

func (worker *indicatorWorker) initiateIndicators(settings *Settings) error {
	if err := settings.Config.CalculateBollingerBand(buffers.Candles.List(settings.Market.ID, settings.Resolution.ID)); err != nil {
		return err
	}
	if err := settings.Config.CalculateRSI(buffers.Candles.List(settings.Market.ID, settings.Resolution.ID)); err != nil {
		return err
	}
	if err := settings.Config.CalculateStochastic(buffers.Candles.List(settings.Market.ID, settings.Resolution.ID)); err != nil {
		return err
	}
	return nil
}

func (worker *indicatorWorker) calculateIndicators(settings *Settings) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := settings.Config.UpdateBollingerBand(buffers.Candles.List(settings.Market.ID, settings.Resolution.ID)); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		settings.Config.UpdateRSI(buffers.Candles.List(settings.Market.ID, settings.Resolution.ID))
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		settings.Config.UpdateStochastic(buffers.Candles.List(settings.Market.ID, settings.Resolution.ID))
	}(&wg)

	wg.Wait()
}
