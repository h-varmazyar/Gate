package workers

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type PrimaryData struct {
	db      repository.CandleRepository
	buffer  *buffer.CandleBuffer
	configs *Configs
	queue   *amqpext.Queue
	Started bool
}

func NewPrimaryDataWorker(_ context.Context, db repository.CandleRepository, configs *Configs, buffer *buffer.CandleBuffer) (*PrimaryData, error) {
	ohlcQueue, err := amqpext.Client.QueueDeclare(configs.PrimaryDataQueue)
	if err != nil {
		return nil, err
	}
	return &PrimaryData{
		db:      db,
		configs: configs,
		queue:   ohlcQueue,
		buffer:  buffer,
	}, nil
}

func (w *PrimaryData) Start() {
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
	w.Started = true
}

func (w *PrimaryData) handle(delivery amqp.Delivery) {
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
	}
	if err := w.db.BulkInsert(localCandles); err != nil {
		log.WithError(err).Errorf("failed to save candles")
	}
	for _, candle := range localCandles {
		w.buffer.Push(candle)
	}
	_ = delivery.Ack(false)
}
