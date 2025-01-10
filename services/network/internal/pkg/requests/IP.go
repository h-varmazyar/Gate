package requests

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
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
	interval := limiter.IntervalDuration()
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ip.ctx.Done():
			ticker.Stop()
			return

		case <-ticker.C:
			ctx, cancelFunc := context.WithTimeout(context.Background(), interval)
			ip.handleBucketRequest(ctx, limiter)
			cancelFunc()
		}
	}
}

func (ip *IP) handleBucketRequest(ctx context.Context, limiter *Limiter) {
	select {
	case <-ctx.Done():
		return
	case bucketRequest := <-limiter.RequestChannel:
		go ip.doRequest(bucketRequest)
	}
}

func (ip *IP) doRequest(bucketRequest *BucketRequest) {
	response := new(networkAPI.Response)
	if bucketRequest.Request.IssueTime != 0 && bucketRequest.Request.Timeout != 0 {
		if time.Unix(bucketRequest.Request.IssueTime, 0).Add(time.Duration(bucketRequest.Request.Timeout)).Before(time.Now()) {
			response.Code = http.StatusRequestTimeout
			response.Body = errors.New(context.Background(), codes.DeadlineExceeded).Error()
		}
	}
	if response.Code == 0 {
		networkRequest, err := NewNetworkRequest(bucketRequest.Request, ip.ProxyURL)
		if err != nil {
			log.WithError(err).Errorf("failed to create async network request")
			response.Code = http.StatusInternalServerError
			response.Body = err.Error()
		} else {
			var httpResp *networkAPI.Response
			httpResp, err = networkRequest.Do()
			if err != nil {
				log.WithError(err).Errorf("failed to do network request")
				response.Code = http.StatusInternalServerError
				response.Body = err.Error()
			} else {
				mapper.Struct(httpResp, response)
			}
		}
	}

	bucketResponse := &BucketResponse{
		Response: response,
		BucketID: bucketRequest.BucketID,
	}
	ip.responseChan <- bucketResponse
}
