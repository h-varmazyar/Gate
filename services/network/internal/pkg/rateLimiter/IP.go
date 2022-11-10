package rateLimiter

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/requests"
	log "github.com/sirupsen/logrus"
	"net/url"
	"time"
)

type IP struct {
	ID       uuid.UUID
	Address  string
	Username string
	Password string
	Port     uint16
	proxyURL *url.URL
	ctx      context.Context
}

func (ip *IP) spreadAlgorithm(limiter *Limiter) {
	interval := limiter.TimeLimit / time.Duration(limiter.RequestCountLimit)
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ip.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			systemRequest := <-limiter.RequestChannel
			networkURL, err := requests.New(systemRequest, ip.proxyURL)
			if err != nil {
				log.WithError(err).Errorf("failed to create async network request")
				continue
			}
			response, err := networkURL.Do()
			if err != nil {
				log.WithError(err).Errorf("failed to do network request")
				continue
			}
			ip.sendResponse(systemRequest, response)
		}
	}
}

func (ip *IP) sendResponse(request *networkAPI.Request, response *networkAPI.Response) {
	queue, err := amqpext.Client.QueueDeclare(request.CallbackQueue)
	if err != nil {
		log.WithError(err).Errorf("failed to declare amqp queue")
		return
	}
	bytes, err := proto.Marshal(response)
	if err != nil {
		log.WithError(err).Errorf("failed to marshal response")
		return
	}

	if err = queue.Publish(bytes, grpcext.ProtobufContentType); err != nil {
		log.WithError(err).Errorf("failed to publish response")
	}
}
