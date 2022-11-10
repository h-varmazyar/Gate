package rateLimiter

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"net/url"
	"sync"
	"time"
)

type Manager struct {
	Limiters        map[uuid.UUID]*Limiter
	IPs             map[uuid.UUID]*IP
	cancelFunctions map[uuid.UUID]context.CancelFunc
	lock            *sync.Mutex
}

func NewManager(ctx context.Context, Limiters []*networkAPI.RateLimiter, IPs []*networkAPI.IP) (*Manager, error) {
	manager := new(Manager)
	manager.Limiters = make(map[uuid.UUID]*Limiter)
	manager.IPs = make(map[uuid.UUID]*IP)
	manager.cancelFunctions = make(map[uuid.UUID]context.CancelFunc)

	for _, ip := range IPs {
		proxyAddress := ""
		if ip.Username != "" && ip.Password != "" {
			proxyAddress = fmt.Sprintf("%v:%v@%v:%v", ip.Username, ip.Password, ip.Address, ip.Port)
		} else {
			proxyAddress = fmt.Sprintf("%v:%v", ip.Address, ip.Port)
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
			ID:       ipID,
			Address:  ip.Address,
			Username: ip.Username,
			Password: ip.Password,
			Port:     uint16(ip.Port),
			proxyURL: proxyURL,
			ctx:      ipCtx,
		}
		manager.cancelFunctions[ipID] = cancelFunc
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
			RequestChannel:    make(chan *networkAPI.Request),
		}
		manager.assignLimiterToIPs(tmpLimiter)
		manager.Limiters[limiterID] = tmpLimiter
	}
	return manager, nil
}

func (m *Manager) AddNewRequest(ctx context.Context, request *networkAPI.Request) error {
	limiterID, err := uuid.Parse(request.RateLimiterID)
	if err != nil {
		return err
	}
	_, ok := m.Limiters[limiterID]
	if !ok {
		return errors.New(ctx, codes.NotFound).AddDetails("rate limiter with id %v not found", limiterID.String())
	}
	m.Limiters[limiterID].RequestChannel <- request
	return nil
}

func (m *Manager) assignLimiterToIPs(limiter *Limiter) {
	for _, ip := range m.IPs {
		go func(ip *IP) {
			switch limiter.Type {
			case networkAPI.RateLimiter_Spread:
				ip.spreadAlgorithm(limiter)
			case networkAPI.RateLimiter_Immediate:
				log.Warnf("limiter type not implementerd")
			}
		}(ip)
	}
}
