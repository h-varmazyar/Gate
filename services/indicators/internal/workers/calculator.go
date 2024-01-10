package workers

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var (
	indicatorMapKeyPattern = "%v*%v"
)

type Indicator interface {
	Calculate(ctx context.Context, candles []*chipmunkApi.Candle)
	CandleCountNeedsForCalculation(ctx context.Context) int32
	GetMarket(ctx context.Context) *chipmunkApi.Market
	GetResolution(ctx context.Context) *chipmunkApi.Resolution
}

type IndicatorCalculator struct {
	candleService chipmunkApi.CandleServiceClient
	indicatorsMap map[string][]Indicator //key is the combination of market and resolution id
	configs       Configs
	log           *log.Logger
}

func NewIndicatorCalculatorWorker(_ context.Context, log *log.Logger, configs Configs) *IndicatorCalculator {
	chipmunkConn := grpcext.NewConnection(configs.ChipmunkAddress)
	return &IndicatorCalculator{
		candleService: chipmunkApi.NewCandleServiceClient(chipmunkConn),
		indicatorsMap: make(map[string][]Indicator),
		configs:       configs,
		log:           log,
	}
}

func (w IndicatorCalculator) Start(ctx context.Context, indicators []Indicator) {
	for _, indicator := range indicators {
		w.addIndicator(ctx, indicator)
	}

	for key, indicators := range w.indicatorsMap {
		ids := strings.Split(key, "*")
		w.calculateMarketIndicators(indicators, ids[0], ids[1])
	}
}

func (w IndicatorCalculator) addIndicator(ctx context.Context, indicator Indicator) {
	key := fmt.Sprintf(indicatorMapKeyPattern, indicator.GetMarket(ctx).ID, indicator.GetResolution(ctx).ID)
	if indicators, ok := w.indicatorsMap[key]; ok {
		indicators = append(indicators, indicator)
	} else {
		w.indicatorsMap[key] = make([]Indicator, 0)
	}
}

func (w IndicatorCalculator) calculateMarketIndicators(indicators []Indicator, marketId, resolutionId string) {
	candleFetchCount := int32(2)
	candleCountCtx := context.Background()
	for _, indicator := range indicators {
		if count := indicator.CandleCountNeedsForCalculation(candleCountCtx); count > candleFetchCount {
			candleFetchCount = count
		}
	}
	candleFetchReq := &chipmunkApi.CandleListReq{
		ResolutionID: resolutionId,
		MarketID:     marketId,
		Count:        candleFetchCount,
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

			for _, indicator := range indicators {
				indicator.Calculate(ctx, candles.Elements)
			}

			if diff := time.Now().Sub(start); diff < w.configs.CalculatorInterval {
				time.Sleep(w.configs.CalculatorInterval - diff)
			}
			fn()
		}
	}
}
