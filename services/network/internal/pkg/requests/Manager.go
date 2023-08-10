package requests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"math/rand"
	"net/url"
	"sync"
	"time"
)

type Manager struct {
	Limiters         map[uuid.UUID]*Limiter
	IPs              map[uuid.UUID]*IP
	buckets          map[uuid.UUID]*Bucket
	responseChan     chan *BucketResponse
	cancelFunctions  map[uuid.UUID]context.CancelFunc
	lock             *sync.Mutex
	defaultLimiterID uuid.UUID
}

func NewManager(ctx context.Context, Limiters []*networkAPI.RateLimiter, IPs []*networkAPI.IP) (*Manager, error) {
	manager := new(Manager)
	manager.Limiters = make(map[uuid.UUID]*Limiter)
	manager.IPs = make(map[uuid.UUID]*IP)
	manager.cancelFunctions = make(map[uuid.UUID]context.CancelFunc)
	manager.lock = new(sync.Mutex)
	manager.buckets = make(map[uuid.UUID]*Bucket)
	manager.responseChan = make(chan *BucketResponse, 1000000)

	for _, ip := range IPs {
		proxyAddress := ""
		if ip.Username != "" && ip.Password != "" {
			proxyAddress = fmt.Sprintf("%v://%v:%v@%v:%v", ip.Schema, ip.Username, ip.Password, ip.Address, ip.Port)
		} else {
			proxyAddress = fmt.Sprintf("%v://%v:%v", ip.Schema, ip.Address, ip.Port)
		}
		proxyURL, err := url.Parse(proxyAddress)
		if err != nil {
			return nil, err
		}
		ipID, err := uuid.Parse(ip.ID)
		if err != nil {
			return nil, err
		}
		ipCtx, cancelFunc := context.WithCancel(ctx)
		manager.IPs[ipID] = &IP{
			ID:           ipID,
			Address:      ip.Address,
			Username:     ip.Username,
			Password:     ip.Password,
			Port:         uint16(ip.Port),
			ProxyURL:     proxyURL,
			ctx:          ipCtx,
			responseChan: manager.responseChan,
		}
		manager.cancelFunctions[ipID] = cancelFunc
	}

	if len(IPs) == 0 {
		ipCtx, cancelFunc := context.WithCancel(ctx)
		id := uuid.New()

		manager.IPs[id] = &IP{
			ID:           id,
			ctx:          ipCtx,
			responseChan: manager.responseChan,
		}

		manager.cancelFunctions[id] = cancelFunc
	}

	if len(Limiters) == 0 {
		defaultLimiter := manager.getDefaultLimiter()
		Limiters = append(Limiters, defaultLimiter)
	}

	for _, limiter := range Limiters {
		limiterID, err := uuid.Parse(limiter.ID)
		if err != nil {
			return nil, err
		}
		tmpLimiter := &Limiter{
			ID:                limiterID,
			RequestCountLimit: limiter.RequestCountLimit,
			TimeLimit:         time.Duration(limiter.TimeLimit),
			Type:              limiter.Type,
			RequestChannel:    make(chan *BucketRequest, 1000000),
		}
		manager.assignLimiterToIPs(tmpLimiter)
		manager.Limiters[limiterID] = tmpLimiter
	}

	go manager.listenToResponses()

	return manager, nil
}
func (m *Manager) getDefaultLimiter() *networkAPI.RateLimiter {
	id := uuid.New()
	defaultLimiter := &networkAPI.RateLimiter{
		ID:                id.String(),
		RequestCountLimit: 400,
		TimeLimit:         int64(time.Second),
		Type:              networkAPI.RateLimiter_Spread,
	}
	m.defaultLimiterID = id
	return defaultLimiter
}

//func (m *Manager) AddNewRequest(ctx context.Context, rateLimiterID string, request *Request) error {
//	limiter, err := m.prepareLimiterForNewRequest(ctx, rateLimiterID)
//	if err != nil {
//		return err
//	}
//	limiter.RequestChannel <- request
//	return nil
//}

func (m *Manager) AddNewBucket(ctx context.Context, callbackQueue, referenceID string, request []*networkAPI.Request) (int64, error) {
	bucket := &Bucket{
		ID:            uuid.New(),
		CallbackQueue: callbackQueue,
		ReferenceID:   referenceID,
	}

	for _, req := range request {
		bucket.Requests = append(bucket.Requests, &BucketRequest{
			Request:  req,
			BucketID: bucket.ID,
		})
	}

	m.lock.Lock()
	m.buckets[bucket.ID] = bucket
	m.lock.Unlock()

	totalIntervalDuration := int64(0)
	for _, req := range bucket.Requests {
		limiter, err := m.prepareLimiterForNewRequest(ctx, req.Request.RateLimiterID)
		if err != nil {
			return 0, err
		}
		limiter.RequestChannel <- req
		totalIntervalDuration += int64(limiter.IntervalDuration())
	}

	return totalIntervalDuration, nil
}

func (m *Manager) GetRandomIP(ctx context.Context) (*IP, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	rand.Seed(time.Now().Unix())
	rndNum := rand.Intn(len(m.IPs))

	i := 0
	for _, ip := range m.IPs {
		if i == rndNum {
			return ip, nil
		}
		i++
	}
	return nil, errors.New(ctx, codes.NotFound)
}

func (m *Manager) assignLimiterToIPs(limiter *Limiter) {
	for _, ip := range m.IPs {
		switch limiter.Type {
		case networkAPI.RateLimiter_Spread:
			go ip.spreadAlgorithm(limiter)
		case networkAPI.RateLimiter_Immediate:
			log.Warnf("limiter type not implementerd")
		}
	}
}

func (m *Manager) prepareLimiterForNewRequest(ctx context.Context, rateLimiterID string) (*Limiter, error) {
	var limiterID uuid.UUID
	if rateLimiterID != "" {
		var err error
		limiterID, err = uuid.Parse(rateLimiterID)
		if err != nil {
			log.WithError(err).Errorf("invalid rate limiter id %v", rateLimiterID)
			return nil, err
		}
	} else {
		limiterID = m.defaultLimiterID
	}
	_, ok := m.Limiters[limiterID]
	if !ok {
		return nil, errors.New(ctx, codes.NotFound).AddDetails("rate limiter with id %v not found", limiterID.String())
	}
	return m.Limiters[limiterID], nil
}

func (m *Manager) listenToResponses() {
	for response := range m.responseChan {
		m.lock.Lock()
		bucket, ok := m.buckets[response.BucketID]
		if !ok {
			log.Errorf("bucket not found")
			continue
		}
		isEnded, err := bucket.addResponse(context.Background(), response)
		if err != nil {
			log.WithError(err).Error("failed to add response")
		}
		if isEnded {
			delete(m.buckets, response.BucketID)
		}
		m.lock.Unlock()
	}
}
