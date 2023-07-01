package workers

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	indicatorsPkg "github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type CandleReader struct {
	db         repository.CandleRepository
	configs    *Configs
	queue      *amqpext.Queue
	indicators []indicatorsPkg.Indicator
	insertChan chan *entity.Candle
}

func NewCandleReaderWorker(_ context.Context, db repository.CandleRepository, configs *Configs) (*CandleReader, error) {
	ohlcQueue, err := amqpext.Client.QueueDeclare(configs.PrimaryDataQueue)
	if err != nil {
		return nil, err
	}
	reader := &CandleReader{
		db:         db,
		configs:    configs,
		queue:      ohlcQueue,
		insertChan: make(chan *entity.Candle, 1000),
	}
	go reader.handleInsert()

	return reader, nil
}

func (w *CandleReader) Start(indicators []indicatorsPkg.Indicator) {
	w.indicators = indicators
	handled := int64(0)
	deliveries := w.queue.Consume(w.configs.PrimaryDataQueue)
	for i := 0; i < w.configs.ConsumerCount; i++ {
		go func() {
			for delivery := range deliveries {
				handled++
				w.handle(delivery)
				if handled%1000 == 0 {
					log.Infof("handled: %v", handled)
				}
			}
		}()
	}
}

func (w *CandleReader) handle(delivery amqp.Delivery) {
	candles := new(chipmunkApi.Candles)
	if err := proto.Unmarshal(delivery.Body, candles); err != nil {
		log.WithError(err).Errorf("failed to unmarshal delivery")
		_ = delivery.Nack(false, false)
		return
	}
	localCandles := make([]*entity.Candle, 0)
	for _, candle := range candles.Elements {
		tmp := new(entity.Candle)
		mapper.Struct(candle, tmp)
		localCandles = append(localCandles, tmp)

		w.insertChan <- tmp
	}

	for _, indicator := range w.indicators {
		indicator.Update(localCandles)
	}

	for _, candle := range localCandles {
		buffer.CandleBuffer.Push(candle)
	}
	_ = delivery.Ack(false)
}

func (w *CandleReader) handleInsert() {
	key := ""
	candleMap := make(map[string]*entity.Candle)
	lastInsert := time.Now()
	for candle := range w.insertChan {
		key = fmt.Sprintf("%v%v%v", candle.MarketID, candle.ResolutionID, candle.Time.Unix())
		candleMap[key] = candle
		if time.Now().Sub(lastInsert) > time.Minute {
			err := w.insert(candleMap)
			if err != nil {
				log.WithError(err).Errorf("failed to insert candles")
				continue
			}
			candleMap = make(map[string]*entity.Candle)
			lastInsert = time.Now()
		}
	}
}

func (w *CandleReader) insert(candleMap map[string]*entity.Candle) error {
	candles := make([]*entity.Candle, 0)
	for _, candle := range candleMap {
		candles = append(candles, candle)
	}

	if err := w.db.BulkInsert(candles); err != nil {
		return err
	}
	return nil
}
