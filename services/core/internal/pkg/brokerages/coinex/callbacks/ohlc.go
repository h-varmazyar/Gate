package callbacks

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type OHLCCallback struct {
	queue        *amqpext.Queue
	responseChan chan *coreApi.OHLCResponse
}

func ListenOHLCCallback(responseChan chan *coreApi.OHLCResponse) error {
	if callbackQueue, err := amqpext.Client.QueueDeclare(coinex.QueueOHLC); err != nil {
		log.WithError(err).Error("failed to declare coinex callback queue")
		return err
	} else {
		c := &OHLCCallback{
			queue:        callbackQueue,
			responseChan: responseChan,
		}
		c.run()
		return nil
	}
}

func (c *OHLCCallback) run() {
	counter := 0
	callbackDeliveries := c.queue.Consume(coinex.QueueOHLC)
	go func() {
		for delivery := range callbackDeliveries {
			counter++
			if counter%10 == 0 {
				log.Infof("new 10 delivery: %v", counter/10)
			}

			c.handleDelivery(context.Background(), delivery)
		}
	}()
}

func (c *OHLCCallback) handleDelivery(ctx context.Context, delivery amqp.Delivery) {
	response := new(networkAPI.AsyncResponses)
	if err := proto.Unmarshal(delivery.Body, response); err != nil {
		_ = delivery.Nack(false, false)
		log.WithError(err).Error("failed to unmarshal coinex callback delivery")
		return
	}

	r, err := coinex.NewResponse()
	if err != nil {
		_ = delivery.Nack(false, false)
		return
	}

	log.Infof("responses: %v", len(response.Responses))
	c.responseChan <- r.AsyncOHLC(ctx, response)

	_ = delivery.Ack(false)
}
