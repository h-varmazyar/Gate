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
	logger     *log.Logger
	configs    *Configs
	queue      *amqpext.Queue
	indicators []indicatorsPkg.Indicator
	insertChan chan *entity.Candle
	runners    map[string]*Runner
}

func NewCandleReaderWorker(_ context.Context, db repository.CandleRepository, configs *Configs, logger *log.Logger) (*CandleReader, error) {
	ohlcQueue, err := amqpext.Client.QueueDeclare(configs.PrimaryDataQueue)
	if err != nil {
		return nil, err
	}
	reader := &CandleReader{
		db:         db,
		logger:     logger,
		configs:    configs,
		queue:      ohlcQueue,
		insertChan: make(chan *entity.Candle, 1000),
	}

	return reader, nil
}

func (w *CandleReader) Start(runners map[string]*Runner, indicators []indicatorsPkg.Indicator) {
	w.logger.Infof("starting candle reader worker...")
	w.indicators = indicators
	w.runners = runners
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
	go w.handleInsert()
}

//create new response struct candles, market_id, reference_id, error
func (w *CandleReader) handle(delivery amqp.Delivery) {
	resp := new(chipmunkApi.CandlesAsyncUpdate)
	if err := proto.Unmarshal(delivery.Body, resp); err != nil {
		log.WithError(err).Errorf("failed to unmarshal delivery")
		_ = delivery.Nack(false, false)
		return
	}
	if resp.ReferenceID == w.runners[resp.MarketID].LastEventID {
		w.runners[resp.MarketID].IsPrimaryCandlesLoaded = true
	}
	if resp.Error != "" {
		w.logger.Errorf("failed to get async candle update resp for market %v: %v", resp.MarketID, resp.Error)
		delivery.Nack(false, false)
		return
	}
	localCandles := make([]*entity.Candle, 0)
	for _, candle := range resp.Candles {
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
		if time.Now().Sub(lastInsert) > time.Second {
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
