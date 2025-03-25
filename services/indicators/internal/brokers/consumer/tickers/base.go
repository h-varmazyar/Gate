package candles

import (
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	"github.com/h-varmazyar/Gate/services/indicators/internal/domain"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/calculator"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/storage"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"sync"
	"time"
)

type indicatorsRepo interface {
	All(ctx context.Context) ([]*entities.Indicator, error)
}

type candleService interface {
	CandleList(ctx context.Context, marketId, resolutionId uint) ([]domain.Candle, error)
}
type Consumer struct {
	logger        *log.Logger
	nc            *nats.Conn
	lock          *sync.RWMutex
	indicatorsMap map[uint]map[uint][]calculator.Indicator
	candleService candleService
}

func NewConsumer(logger *log.Logger, nc *nats.Conn, indicatorsRepo indicatorsRepo, candleService candleService) (*Consumer, error) {
	indicators, err := indicatorsRepo.All(context.Background())
	if err != nil {
		logger.WithError(err).Error("failed to load indicators in tickers")
		return nil, err
	}

	indicatorMap := make(map[uint]map[uint][]calculator.Indicator)
	for _, indicator := range indicators {
		ind, err := calculator.NewIndicator(context.Background(), indicator, nil, nil)
		if err != nil {
			logger.WithError(err).Error("failed to create indicator in tickers")
			return nil, err
		}
		if _, ok := indicatorMap[indicator.MarketId]; ok {
			indicatorMap[indicator.MarketId][indicator.ResolutionId] = append(indicatorMap[indicator.MarketId][indicator.ResolutionId], ind)
		} else {
			indicatorMap[indicator.MarketId] = make(map[uint][]calculator.Indicator)
			indicatorMap[indicator.MarketId][indicator.ResolutionId] = []calculator.Indicator{ind}
		}
	}

	for marketId, resolutionMap := range indicatorMap {
		for resolutionId, indicators := range resolutionMap {
			candles, err := candleService.CandleList(context.Background(), marketId, resolutionId)
			if err != nil {
				logger.WithError(err).Error("failed to get candles in tickers")
				return nil, err
			}
			indicatorCandles := make([]calculator.Candle, 0)
			mapper.Slice(candles, &indicatorCandles)
			for _, indicator := range indicators {
				if _, err = indicator.Calculate(context.Background(), indicatorCandles); err != nil {
					return nil, err
				}
			}
		}
	}

	return &Consumer{
		logger:        logger,
		nc:            nc,
		lock:          &sync.RWMutex{},
		indicatorsMap: indicatorMap,
	}, nil
}

func (c *Consumer) StartListening() {
	subject := "tickers.market.*"
	_, err := c.nc.Subscribe(subject, func(msg *nats.Msg) {
		var payload CandlePayload
		if err := json.Unmarshal(msg.Data, &payload); err != nil {
			c.logger.Printf("Error unmarshaling message: %v", err)
			return
		}

		go c.calculateIndicators(payload)
	})

	if err != nil {
		c.logger.Fatalf("Failed to subscribe: %v", err)
	}

	c.logger.Printf("Subscribed to %v\n", subject)
}

func (c *Consumer) calculateIndicators(payload CandlePayload) {
	candles := make([]calculator.Candle, len(payload.Candles))
	for _, candle := range payload.Candles {
		candles = append(candles, calculator.Candle{
			Time:   time.Unix(candle.Timestamp, 0),
			Open:   candle.Open,
			High:   candle.High,
			Low:    candle.Low,
			Close:  candle.Close,
			Volume: candle.Volume,
		})
	}

	c.lock.Lock()
	resolutions, ok := c.indicatorsMap[payload.MarketID]
	c.lock.Unlock()
	if !ok {
		return
	}
	for _, indicators := range resolutions {
		for _, indicator := range indicators {
			value := indicator.UpdateLast(context.Background(), candles[len(candles)-1])
			if err := storage.AddValue(context.Background(), indicator.GetId(), value); err != nil {
				c.logger.WithError(err).Warnf("failed to save indicator value. id: %v", indicator.GetId())
			}
		}
	}
}

func getKey(marketID, resolutionID uint) string {
	return fmt.Sprintf("%v.%v", marketID, resolutionID)
}
