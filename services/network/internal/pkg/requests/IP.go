package requests

import (
	"context"
	"github.com/google/uuid"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

type IP struct {
	ID           uuid.UUID
	Schema       string
	Address      string
	Username     string
	Password     string
	Port         uint16
	ProxyURL     *url.URL
	ctx          context.Context
	responseChan chan *BucketResponse
}

func (ip *IP) spreadAlgorithm(limiter *Limiter) {
	interval := limiter.TimeLimit / time.Duration(limiter.RequestCountLimit)
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ip.ctx.Done():
			ticker.Stop()
			return
			//case systemRequest := <-limiter.RequestChannel:
			//	for _, remoteRequest := range systemRequest.Requests {
			//		<-ticker.C
			//		go systemRequest.handleRequest(ip.ProxyURL, limiter.TimeLimit, remoteRequest)
			//	}

		case <-ticker.C:
			go func() {
				response := new(networkAPI.Response)
				bucketRequest := <-limiter.RequestChannel
				if bucketRequest.Request.IssueTime != 0 && bucketRequest.Request.Timeout != 0 {
					if bucketRequest.Request.IssueTime+bucketRequest.Request.Timeout < time.Now().Unix() {
						//todo: must push into rabbit
						return
					}
				}
				networkRequest, err := NewNetworkRequest(bucketRequest.Request, ip.ProxyURL)
				if err != nil {
					log.WithError(err).Errorf("failed to create async network request")
					response.Code = http.StatusInternalServerError
					response.Body = err.Error()
				} else {
					response, err = networkRequest.Do()
					if err != nil {
						log.WithError(err).Errorf("failed to do network request")
						response.Code = http.StatusInternalServerError
						response.Body = err.Error()
					}
				}
				bucketResponse := &BucketResponse{
					Response: response,
					BucketID: bucketRequest.BucketID,
				}

				ip.responseChan <- bucketResponse
			}()
		}
	}
}

//func (ip *IP) sendResponse(request *networkAPI.Request, response *networkAPI.Response) {
//	queue, err := amqpext.Client.QueueDeclare(request.CallbackQueue)
//	if err != nil {
//		log.WithError(err).Errorf("failed to declare amqp queue")
//		return
//	}
//	bytes, err := proto.Marshal(response)
//	if err != nil {
//		log.WithError(err).Errorf("failed to marshal response")
//		return
//	}
//
//	if err = queue.Publish(bytes, grpcext.ProtobufContentType); err != nil {
//		log.WithError(err).Errorf("failed to publish response")
//	}
//}
