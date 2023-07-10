package requests

import (
	"context"
	networkApi "github.com/h-varmazyar/Gate/services/network/api/proto"
	ipService "github.com/h-varmazyar/Gate/services/network/internal/app/IPs/service"
	rateLimiterService "github.com/h-varmazyar/Gate/services/network/internal/app/rateLimiters/service"
	"github.com/h-varmazyar/Gate/services/network/internal/app/requests/service"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/requests"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

func NewApp(ctx context.Context, logger *log.Logger, rateLimiterService *rateLimiterService.Service, ipService *ipService.Service) (*App, error) {
	var limiters *networkApi.RateLimiters
	var err error
	if limiters, err = rateLimiterService.List(ctx, &networkApi.RateLimiterListReq{Type: networkApi.RateLimiter_All}); err != nil {
		return nil, err
	}
	var ips *networkApi.IPs
	if ips, err = ipService.List(ctx, new(networkApi.IPListReq)); err != nil {
		return nil, err
	}

	var rateLimiterManager *requests.Manager
	if rateLimiterManager, err = requests.NewManager(ctx, limiters.Elements, ips.Elements); err != nil {
		return nil, err
	}
	return &App{
		Service: service.NewService(ctx, logger, rateLimiterManager),
	}, nil
}
