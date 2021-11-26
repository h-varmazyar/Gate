package wallets

import (
	"context"
	"fmt"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/configs"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/brokerages"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
	"google.golang.org/grpc"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Service struct {
	brokerageService brokerageApi.BrokerageServiceClient
	networkService   networkAPI.RequestServiceClient
}

var (
	GrpcService *Service
)

func NewService(configs *configs.Configs) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		brokerageConnection := grpcext.NewConnection(fmt.Sprintf(":%d", configs.GrpcPort))
		networkConnection := grpcext.NewConnection(fmt.Sprintf(":%d", configs.NetworkGrpcPort))
		GrpcService.brokerageService = brokerageApi.NewBrokerageServiceClient(brokerageConnection)
		GrpcService.networkService = networkAPI.NewRequestServiceClient(networkConnection)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	brokerageApi.RegisterWalletServiceServer(server, s)
}

func (s *Service) List(ctx context.Context, req *brokerageApi.WalletListRequest) (*brokerageApi.Wallets, error) {
	br, err := s.brokerageService.GetInternal(ctx, &brokerageApi.BrokerageIDReq{ID: req.BrokerageID})
	if err != nil {
		return nil, err
	}
	brokerage, err := brokerages.Fetch(ctx, br)
	if err != nil {
		return nil, err
	}

	wallets, err := brokerage.WalletList(func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.networkService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	response := make([]*brokerageApi.Wallet, 0)
	mapper.Slice(wallets, response)
	return &brokerageApi.Wallets{Wallets: response}, nil
}
