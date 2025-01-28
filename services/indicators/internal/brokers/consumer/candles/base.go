package candles

import (
	"encoding/json"
	"fmt"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/calculator"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/storage"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"sync"
)

type indicatorsRepo interface {
	All(ctx context.Context) ([]*entities.Indicator, error)
}

type Consumer struct {
	logger        *log.Logger
	nc            *nats.Conn
	lock          *sync.Mutex
	indicatorsMap map[string][]calculator.Indicator
}

func NewConsumer(logger *log.Logger, nc *nats.Conn, indicatorsRepo indicatorsRepo) (*Consumer, error) {
	indicators, err := indicatorsRepo.All(context.Background())
	if err != nil {
		return nil, err
	}

	indicatorMap := make(map[string][]calculator.Indicator)
	for _, indicator := range indicators {
		ind, err := calculator.NewIndicator(context.Background(), indicator, nil, nil)
		if err != nil {
			return nil, err
		}
		key := getKey(indicator.MarketId, indicator.ResolutionId)
		if _, ok := indicatorMap[key]; ok {
			indicatorMap[key] = append(indicatorMap[key], ind)
		} else {
			indicatorMap[key] = []calculator.Indicator{ind}
		}
	}

	return &Consumer{
		logger:        logger,
		nc:            nc,
		lock:          &sync.Mutex{},
		indicatorsMap: indicatorMap,
	}, nil
}

func (c *Consumer) StartListening() {
	subject := "candles.market.*.*"
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
	key := getKey(payload.MarketID, payload.ResolutionID)

	c.lock.Lock()
	indicators, ok := c.indicatorsMap[key]
	if !ok {
		c.logger.Errorf("no indicator map found with key %v", key)
		return
	}
	c.lock.Unlock()

	candles := make([]*chipmunkAPI.Candle, len(payload.Candles))
	for _, candle := range payload.Candles {
		candles = append(candles, &chipmunkAPI.Candle{
			Time:   candle.Timestamp,
			Open:   candle.Open,
			High:   candle.High,
			Low:    candle.Low,
			Close:  candle.Close,
			Volume: candle.Volume,
		})
	}

	for _, indicator := range indicators {
		value := indicator.UpdateLast(context.Background(), candles[len(candles)-1])
		if err := storage.AddValue(context.Background(), indicator.GetId(), value); err != nil {
			c.logger.WithError(err).Warnf("failed to save indicator value. id: %v", indicator.GetId())
		}
	}
	c.logger.Infof("calculating ind len: %v", len(indicators))
}

func getKey(marketID, resolutionID uint) string {
	return fmt.Sprintf("%v.%v", marketID, resolutionID)
}
