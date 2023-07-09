package rateLimiter

import (
	"context"
	"github.com/google/uuid"
	"net/url"
	"time"
)

type IP struct {
	ID       uuid.UUID
	Schema   string
	Address  string
	Username string
	Password string
	Port     uint16
	ProxyURL *url.URL
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
		case systemRequest := <-limiter.RequestChannel:
			for _, remoteRequest := range systemRequest.Requests {
				<-ticker.C
				go systemRequest.handleRequest(ip.ProxyURL, limiter.TimeLimit, remoteRequest)
			}

			//case <-ticker.C:
			//go func() {
			//	response := new(networkAPI.Response)
			//	systemRequest := <-limiter.RequestChannel
			//	if systemRequest.IssueTime != 0 && systemRequest.Timeout != 0 {
			//		if systemRequest.IssueTime+systemRequest.Timeout < time.Now().Unix() {
			//			//todo: must push into rabbit
			//			return
			//		}
			//	}
			//	networkURL, err := requests.New(systemRequest, ip.ProxyURL)
			//	if err != nil {
			//		log.WithError(err).Errorf("failed to create async network request")
			//		response.Code = http.StatusInternalServerError
			//		response.Body = err.Error()
			//	} else {
			//		response, err = networkURL.Do()
			//		if err != nil {
			//			log.WithError(err).Errorf("failed to do network request")
			//			response.Code = http.StatusInternalServerError
			//			response.Body = err.Error()
			//		}
			//	}
			//	response.ReferenceID = systemRequest.ReferenceID
			//	ip.sendResponse(systemRequest, response)
			//}()
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
