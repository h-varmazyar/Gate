package wallets

//import (
//	"context"
//	"fmt"
//	"github.com/h-varmazyar/Gate/api"
//	"github.com/h-varmazyar/Gate/pkg/grpcext"
//	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
//	"github.com/h-varmazyar/Gate/services/brokerage/configs"
//	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/brokerages/coinex"
//	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
//	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
//	"google.golang.org/grpc"
//)
//
//type Service struct {
//	brokerageService brokerageApi.BrokerageServiceClient
//	networkService   networkAPI.RequestServiceClient
//}
//
//var (
//	GrpcService *Service
//)
//
//func NewService(configs *configs.Configs) *Service {
//	if GrpcService == nil {
//		GrpcService = new(Service)
//		brokerageConnection := grpcext.NewConnection(fmt.Sprintf(":%d", configs.GrpcPort))
//		networkConnection := grpcext.NewConnection(fmt.Sprintf(":%d", configs.NetworkGrpcPort))
//		GrpcService.brokerageService = brokerageApi.NewBrokerageServiceClient(brokerageConnection)
//		GrpcService.networkService = networkAPI.NewRequestServiceClient(networkConnection)
//	}
//	return GrpcService
//}
//
//func (s *Service) RegisterServer(server *grpc.Server) {
//	brokerageApi.RegisterWalletServiceServer(server, s)
//}
//
//func (s *Service) UpdateWallets(ctx context.Context, req *brokerageApi.UpdateWalletRequest) (*brokerageApi.Wallets, error) {
//	br := new(coinex.Service)
//	brokerage, err := repository.Brokerages.ReturnByID(req.BrokerageID)
//	if err != nil {
//		return nil, err
//	}
//	br.Auth = &api.Auth{
//		Type:      api.AuthType(api.AuthType_value[brokerage.AuthType]),
//		Username:  brokerage.Username,
//		Password:  brokerage.Password,
//		AccessID:  brokerage.AccessID,
//		SecretKey: brokerage.SecretKey,
//	}
//	wallets, err := br.WalletList(ctx, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
//		resp, err := s.networkService.Do(ctx, request)
//		return resp, err
//	})
//	if err != nil {
//		return nil, err
//	}
//	return wallets, nil
//}
