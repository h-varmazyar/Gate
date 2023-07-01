package coinex

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type Callback struct {
	configs *Configs
	queue   *amqpext.Queue
	r       *Response
}

func ListenCallbacks(configs *Configs) error {
	if callbackQueue, err := amqpext.Client.QueueDeclare(configs.CoinexCallbackQueue); err != nil {
		log.WithError(err).Error("failed to declare coinex callback queue")
		return err
	} else {
		r, err := NewResponse(configs, true)
		if err != nil {
			return err
		}
		c := &Callback{
			configs: configs,
			queue:   callbackQueue,
			r:       r,
		}
		c.run()
	}
	return nil
}

func (c *Callback) run() {
	counter := 0
	callbackDeliveries := c.queue.Consume(c.configs.CoinexCallbackQueue)
	go func() {
		for delivery := range callbackDeliveries {
			counter++
			if counter%10 == 0 {
				log.Infof("new 10 delivery: %v", counter/10)
			}
			ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute)
			defer cancelFunc()
			c.handleDelivery(ctx, delivery)
		}
	}()
}

func (c *Callback) handleDelivery(ctx context.Context, delivery amqp.Delivery) {
	response := new(networkAPI.Response)
	if err := proto.Unmarshal(delivery.Body, response); err != nil {
		_ = delivery.Nack(false, false)
		log.WithError(err).Error("failed to unmarshal coinex callback delivery")
		return
	}

	metadata := new(brokerages.Metadata)
	if err := json.Unmarshal([]byte(response.Metadata), metadata); err != nil {
		_ = delivery.Nack(false, false)
		log.WithError(err).Error("failed to unmarshal coinex callback delivery")
		return
	}

	switch metadata.Method {
	case brokerages.MethodOHLC:
		if metadata.MarketID == uuid.Nil.String() || metadata.ResolutionID == uuid.Nil.String() {
			return
		}
		c.r.AsyncOHLC(ctx, response, metadata)
	}

	_ = delivery.Ack(false)
}
