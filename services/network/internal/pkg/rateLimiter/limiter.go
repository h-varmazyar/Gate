package rateLimiter

import (
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/requests"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

type Limiter struct {
	ID                uuid.UUID
	RequestCountLimit int64
	TimeLimit         time.Duration
	Type              networkAPI.RateLimiterType
	RequestChannel    chan *Request
}

type Request struct {
	CallbackQueue string
	ReferenceID   string
	Requests      []*networkAPI.Request
	responses     []*networkAPI.Response
}

func (r *Request) addResponse(response *networkAPI.Response) {
	r.responses = append(r.responses, response)
	if len(r.responses) == len(r.Requests) {
		queue, err := amqpext.Client.QueueDeclare(r.CallbackQueue)
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
}

func (r *Request) handleRequest(proxyURL *url.URL, timeLimit time.Duration, remoteRequest *networkAPI.Request) {
	response := &networkAPI.Response{
		Method:      remoteRequest.Method,
		ReferenceID: r.ReferenceID,
	}
	if remoteRequest.IssueTime == 0 {
		remoteRequest.IssueTime = time.Now().Unix()
	}
	if remoteRequest.Timeout == 0 {
		remoteRequest.Timeout = int64(timeLimit)
	}

	if remoteRequest.IssueTime+remoteRequest.Timeout < time.Now().Unix() {
		log.Errorf("request time out exceed")
		response.Code = http.StatusRequestTimeout
		response.Body = "request timeout exceed"
	} else {
		networkRequest, err := requests.New(remoteRequest, proxyURL)
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
	}
	r.addResponse(response)
}
