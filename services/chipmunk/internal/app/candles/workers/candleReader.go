package workers

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type CandleReader struct {
	db                 repository.CandleRepository
	logger             *log.Logger
	configs            *Configs
	queue              *amqpext.Queue
	insertChan         chan *entity.Candle
	handeled           int
	lock               *sync.Mutex
	nextRequestTrigger chan bool
}

func NewCandleReaderWorker(_ context.Context, db repository.CandleRepository, configs *Configs, logger *log.Logger, nextRequestTrigger chan bool) (*CandleReader, error) {
	ohlcQueue, err := amqpext.Client.QueueDeclare(configs.PrimaryDataQueue)
	if err != nil {
		return nil, err
	}
	reader := &CandleReader{
		db:                 db,
		logger:             logger,
		configs:            configs,
		queue:              ohlcQueue,
		insertChan:         make(chan *entity.Candle, 100000),
		lock:               new(sync.Mutex),
		nextRequestTrigger: nextRequestTrigger,
	}

	return reader, nil
}

func (w *CandleReader) Start() {
	w.logger.Infof("starting candle reader worker...")
	//w.indicators = indicators
	deliveries := w.queue.Consume(w.configs.PrimaryDataQueue)
	for i := 0; i < w.configs.ConsumerCount; i++ {
		go func() {
			for delivery := range deliveries {
				w.handle(delivery)
			}
		}()
	}
	go w.handleInsert()
}

func (w *CandleReader) handle(delivery amqp.Delivery) {
	w.lock.Lock()
	w.handeled++
	w.logger.Infof("handling new delivery: %v", w.handeled)
	w.lock.Unlock()
	resp := new(coreApi.OHLCResponse)
	if err := proto.Unmarshal(delivery.Body, resp); err != nil {
		log.WithError(err).Errorf("failed to unmarshal delivery")
		_ = delivery.Nack(false, false)
		return
	}
	for _, item := range resp.Items {
		if item.Error != "" {
			w.logger.Errorf("failed to get async candle update resp for market %v: %v", item.MarketID, item.Error)
			continue
		}
		for _, candle := range item.Candles {
			tmp := new(entity.Candle)
			mapper.Struct(candle, tmp)
			tmp.MarketID = uuid.MustParse(item.MarketID)
			tmp.ResolutionID = uuid.MustParse(item.ResolutionID)
			w.insertChan <- tmp
			//buffer.CandleBuffer.Push(tmp)
		}
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
		if time.Now().Sub(lastInsert) > time.Second/10 {
			err := w.insert(candleMap)
			if err != nil {
				log.WithError(err).Errorf("failed to insert candles")
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

	w.nextRequestTrigger <- true
	if err := w.db.BulkInsert(candles); err != nil {
		return err
	}
	return nil
}
