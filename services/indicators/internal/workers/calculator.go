package workers

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/calculator"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/storage"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

var (
	indicatorMapKeyPattern = "%v*%v"
)

type IndicatorCalculator struct {
	candleService chipmunkAPI.CandleServiceClient
	indicatorsMap map[string][]calculator.Indicator //key is the combination of market and resolution id
	lock          *sync.Mutex
	configs       Configs
	log           *log.Logger
}

func NewIndicatorCalculatorWorker(_ context.Context, log *log.Logger, configs Configs) *IndicatorCalculator {
	chipmunkConn := grpcext.NewConnection(configs.ChipmunkAddress)
	return &IndicatorCalculator{
		candleService: chipmunkAPI.NewCandleServiceClient(chipmunkConn),
		indicatorsMap: make(map[string][]calculator.Indicator),
		lock:          new(sync.Mutex),
		configs:       configs,
		log:           log,
	}
}

func (w IndicatorCalculator) Start(ctx context.Context, indicators []calculator.Indicator) {
	for _, indicator := range indicators {
		w.addIndicator(ctx, indicator)
	}

	for key, _ := range w.indicatorsMap {
		go w.calculateMarketIndicators(key)
	}
}

func (w IndicatorCalculator) AddIndicator(ctx context.Context, indicator calculator.Indicator) {
	w.lock.Lock()
	if key, isNewKey := w.addIndicator(ctx, indicator); isNewKey {
		go w.calculateMarketIndicators(key)
	}
	w.lock.Unlock()
}

func (w IndicatorCalculator) addIndicator(_ context.Context, indicator calculator.Indicator) (string, bool) {
	isNewKey := false
	key := fmt.Sprintf(indicatorMapKeyPattern, indicator.GetMarket().ID, indicator.GetResolution().ID)
	if _, ok := w.indicatorsMap[key]; !ok {
		isNewKey = true
		w.indicatorsMap[key] = make([]calculator.Indicator, 0)
	}
	w.indicatorsMap[key] = append(w.indicatorsMap[key], indicator)
	return key, isNewKey
}

func (w IndicatorCalculator) calculateMarketIndicators(key string) {
	ids := strings.Split(key, "*")
	candleFetchReq := &chipmunkAPI.CandleListReq{
		ResolutionID: ids[0],
		MarketID:     ids[1],
		Count:        1,
	}
	ticker := time.NewTicker(w.configs.CalculatorInterval)
	for {
		select {
		case <-ticker.C:
			start := time.Now()
			ctx, fn := context.WithTimeout(context.Background(), w.configs.CalculatorInterval)

			candles, err := w.candleService.List(ctx, candleFetchReq)
			if err != nil {
				continue
			}

			w.lock.Lock()
			indicators, ok := w.indicatorsMap[key]
			if !ok {
				continue
			}
			w.lock.Unlock()

			for _, indicator := range indicators {
				value := indicator.UpdateLast(ctx, candles.Elements[0])
				if err = storage.AddValue(ctx, indicator.GetId(), value); err != nil {
					w.log.WithError(err).Warnf("failed to save indicator value. id: %v", indicator.GetId())
				}

			}

			if diff := time.Now().Sub(start); diff < w.configs.CalculatorInterval {
				time.Sleep(w.configs.CalculatorInterval - diff)
			}
			fn()
		}
	}
}
