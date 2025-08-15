package workers

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/h-varmazyar/Gate/services/indicators/configs"
	"github.com/h-varmazyar/Gate/services/indicators/internal/domain"
	"github.com/h-varmazyar/Gate/services/indicators/internal/entities"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/calculator"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/storage"
	log "github.com/sirupsen/logrus"
)

type candlesAdapter interface {
	List(ctx context.Context, marketID, resolutionID uint) ([]domain.Candle, error)
}

type indicatorRepository interface {
	List(ctx context.Context, marketID, resolutionID uint) ([]entities.Indicator, error)
}

type IndicatorCalculator struct {
	log            *log.Logger
	configs        configs.CalculatorConfigs
	candlesChan    chan domain.CandlesConsumerPayload
	indicatorsMap  sync.Map
	done           chan bool
	candlesAdapter candlesAdapter
	indicatorRepo  indicatorRepository
}

func NewIndicatorCalculatorWorker(_ context.Context, log *log.Logger, configs configs.CalculatorConfigs) *IndicatorCalculator {
	return &IndicatorCalculator{
		indicatorsMap: sync.Map{},
		configs:       configs,
		log:           log,
	}
}

func (w *IndicatorCalculator) Start(ctx context.Context) {
	for i := 0; i < w.configs.WorkerPoolSize; i++ {
		go w.listenIncomes()
	}
}

func (w *IndicatorCalculator) listenIncomes() {
	for {
		select {
		case candlesPayload := <-w.candlesChan:
			ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
			defer cancelFunc()
			w.consumeCandles(ctx, candlesPayload)
		case <-w.done:
			w.log.Info("Income listener exited")
			return
		}
	}
}

func (w *IndicatorCalculator) consumeCandles(ctx context.Context, payload domain.CandlesConsumerPayload) {
	key := generateKey(payload.MarketID, payload.ResolutionID)
	indicators, ok := w.indicatorsMap.Load(key)
	if !ok {
		var err error
		indicators, err = w.addIndicators(ctx, payload.MarketID, payload.ResolutionID)
		if err != nil {
			w.log.WithError(err).Errorf("failed to add indicators for %v - %v", payload.MarketID, payload.ResolutionID)
			return
		}

		if err = w.calculatePrimaryIndicatorsValue(ctx, indicators.([]calculator.Indicator), payload.MarketID, payload.ResolutionID); err != nil {
			return
		}

		w.indicatorsMap.Store(key, indicators)
	}

	w.calculateMarketIndicators(ctx, indicators.([]calculator.Indicator), payload.Candles)
}

func (w *IndicatorCalculator) addIndicators(ctx context.Context, marketID, resolutionID uint) ([]calculator.Indicator, error) {
	indicators, err := w.indicatorRepo.List(ctx, marketID, resolutionID)
	if err != nil {
		return nil, err
	}
	calcIndocators := make([]calculator.Indicator, len(indicators))
	for i, ind := range indicators {
		indicator, err := calculator.NewIndicator(context.Background(), &ind, nil, nil)
		if err != nil {
			return nil, err
		}
		calcIndocators[i] = indicator
	}

	return calcIndocators, nil
}

func (w *IndicatorCalculator) calculateMarketIndicators(ctx context.Context, indicators []calculator.Indicator, candles []domain.Candle) {
	for _, indicator := range indicators {
		for _, c := range candles {
			candle := calculator.Candle{
				Time:  c.Time,
				Open:  c.Open,
				High:  c.High,
				Low:   c.Low,
				Close: c.Close,
			}
			value := indicator.UpdateLast(ctx, candle)
			if err := storage.AddValue(ctx, indicator.GetId(), value); err != nil {
				w.log.WithError(err).Warnf("failed to save indicator value. id: %v", indicator.GetId())
			}
		}
	}
}

func (w *IndicatorCalculator) calculatePrimaryIndicatorsValue(ctx context.Context, indicators []calculator.Indicator, marketID, resolutionID uint) error {
	primaryCandles, err := w.candlesAdapter.List(ctx, marketID, resolutionID)
	if err != nil {
		w.log.WithError(err).Errorf("failed to get primary candles for %v - %v", marketID, resolutionID)
		return err
	}

	indicatorCandles := make([]calculator.Candle, len(primaryCandles))

	for i, candle := range primaryCandles {
		indicatorCandles[i] = calculator.Candle{
			Time:   candle.Time,
			Open:   candle.Open,
			High:   candle.High,
			Low:    candle.Low,
			Close:  candle.Close,
			Volume: candle.Volume,
		}
	}

	for _, indicator := range indicators {
		values, err := indicator.Calculate(ctx, indicatorCandles)
		if err != nil {
			return err
		}
		for _, value := range values.Values {
			if err = storage.AddValue(ctx, indicator.GetId(), value); err != nil {
				w.log.WithError(err).Warnf("failed to save indicator value. id: %v", indicator.GetId())
				return err
			}
		}

	}

	return nil
}

func generateKey(marketID, resolutionID uint) string {
	indicatorMapKeyPattern := "%v*%v"
	return fmt.Sprintf(indicatorMapKeyPattern, marketID, resolutionID)
}
