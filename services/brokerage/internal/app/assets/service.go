package assets

//import (
//	"context"
//	"fmt"
//	"github.com/h-varmazyar/Gate/api"
//	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
//	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
//	"google.golang.org/grpc"
//	"time"
//)
//
//type Service struct {
//}
//
//var (
//	GrpcService *Service
//)
//
//func NewService() *Service {
//	if GrpcService == nil {
//		GrpcService = new(Service)
//	}
//	return GrpcService
//}
//
//func (s *Service) RegisterServer(server *grpc.Server) {
//	brokerageApi.RegisterAssetServiceServer(server, s)
//}
//
//func (s *Service) Set(_ context.Context, asset *brokerageApi.Asset) (*api.Void, error) {
//	tmp := new(repository.Asset)
//	tmp.Name = asset.Name
//	tmp.Symbol = asset.Symbol
//	tmp.ID = uint(asset.ID)
//	if asset.IssueDate == 0 {
//		return nil, fmt.Errorf("issue date must be declared")
//	} else {
//		tmp.IssueDate = time.Unix(asset.IssueDate, 0)
//	}
//	_, err := repository.Assets.Set(tmp)
//	return &api.Void{}, err
//}
//
//func (s *Service) GetByName(_ context.Context, req *brokerageApi.GetAssetRequest) (*brokerageApi.Asset, error) {
//	asset, err := repository.Assets.ReturnBySymbol(req.Name)
//	if err != nil {
//		return nil, err
//	}
//	return &brokerageApi.Asset{
//		ID:        uint32(asset.ID),
//		Name:      asset.Name,
//		Symbol:    asset.Symbol,
//		IssueDate: asset.IssueDate.Unix(),
//	}, nil
//}
//
//func (s *Service) List(_ context.Context, req *brokerageApi.GetAssetListRequest) (*brokerageApi.Assets, error) {
//	if list, err := repository.Assets.List(int(req.Page)); err != nil {
//		return nil, err
//	} else {
//		assets := new(brokerageApi.Assets)
//		assetList := make([]*brokerageApi.Asset, len(list))
//		for i, asset := range list {
//			assetList[i] = &brokerageApi.Asset{
//				ID:        uint32(asset.ID),
//				Name:      asset.Name,
//				Symbol:    asset.Symbol,
//				IssueDate: asset.IssueDate.Unix(),
//			}
//		}
//		assets.Assets = assetList
//		return assets, nil
//	}
//}
