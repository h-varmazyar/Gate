package requests

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type Bucket struct {
	ID            uuid.UUID
	CallbackQueue string
	ReferenceID   string
	Requests      []*BucketRequest
	responses     []*BucketResponse
}

type BucketRequest struct {
	Request  *networkAPI.Request
	BucketID uuid.UUID
}

type BucketResponse struct {
	Response *networkAPI.Response
	BucketID uuid.UUID
}

func (r *Bucket) addResponse(ctx context.Context, response *BucketResponse) (isEnded bool, err error) {
	if response.BucketID != r.ID {
		err = errors.New(ctx, codes.InvalidArgument).AddDetails("invalid bucket id")
		return
	}

	r.responses = append(r.responses, response)
	if len(r.responses) == len(r.Requests) {
		queue := new(amqpext.Queue)
		queue, err = amqpext.Client.QueueDeclare(r.CallbackQueue)
		if err != nil {
			log.WithError(err).Errorf("failed to declare amqp queue")
			return
		}
		responses := &networkAPI.AsyncResponses{
			ReferenceID: r.ReferenceID,
		}
		for _, bucketResponse := range r.responses {
			bucketResponse.Response.ReferenceID = r.ReferenceID
			responses.Responses = append(responses.Responses, bucketResponse.Response)
		}
		bytes := make([]byte, 0)
		bytes, err = proto.Marshal(responses)
		if err != nil {
			log.WithError(err).Errorf("failed to marshal response")
			return
		}

		log.Infof("publishing bucket response: %v - %v", len(responses.Responses), len(bytes))
		if err = queue.Publish(bytes, grpcext.ProtobufContentType); err != nil {
			log.WithError(err).Errorf("failed to publish response")
		} else {
			isEnded = true
		}
	}
	return
}

//func (r *Bucket) handleRequest(proxyURL *url.URL, timeLimit time.Duration, remoteRequest *networkAPI.Request) {
//	response := &networkAPI.Response{
//		Method:      remoteRequest.Method,
//		ReferenceID: r.ReferenceID,
//	}
//	if remoteRequest.IssueTime == 0 {
//		remoteRequest.IssueTime = time.Now().Unix()
//	}
//	if remoteRequest.Timeout == 0 {
//		remoteRequest.Timeout = int64(timeLimit)
//	}
//
//	if remoteRequest.IssueTime+remoteRequest.Timeout < time.Now().Unix() {
//		log.Errorf("request time out exceed")
//		response.Code = http.StatusRequestTimeout
//		response.Body = "request timeout exceed"
//	} else {
//		networkRequest, err := New(remoteRequest, proxyURL)
//		if err != nil {
//			log.WithError(err).Errorf("failed to create async network request")
//			response.Code = http.StatusInternalServerError
//			response.Body = err.Error()
//		} else {
//			response, err = networkRequest.Do()
//			if err != nil {
//				log.WithError(err).Errorf("failed to do network request")
//				response.Code = http.StatusInternalServerError
//				response.Body = err.Error()
//			}
//		}
//	}
//	r.addResponse(response)
//}
