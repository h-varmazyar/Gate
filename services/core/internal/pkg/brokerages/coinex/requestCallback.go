package coinex

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
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
		r, err := NewResponse(configs)
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
	callbackDeliveries := c.queue.Consume(c.configs.CoinexCallbackQueue)
	go func() {
		for delivery := range callbackDeliveries {
			c.handleDelivery(delivery)
		}
	}()
}

func (c *Callback) handleDelivery(delivery amqp.Delivery) {
	response := new(networkAPI.Response)
	if err := proto.Unmarshal(delivery.Body, response); err != nil {
		_ = delivery.Nack(false, false)
		log.WithError(err).Error("failed to unmarshal coinex callback delivery")
		return
	}

	_ = delivery.Ack(false)

	metadata := new(brokerages.Metadata)
	if err := json.Unmarshal([]byte(response.Metadata), metadata); err != nil {
		_ = delivery.Nack(false, false)
		log.WithError(err).Error("failed to unmarshal coinex callback delivery")
		return
	}

	switch metadata.Method {
	case MethodOHLC:
		c.r.OHLC(response)
	}
}
