package assets

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"google.golang.org/grpc"
	"time"
)

type Service struct {
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterAssetServiceServer(server, s)
}

func (s *Service) Set(_ context.Context, asset *chipmunkApi.Asset) (*api.Void, error) {
	tmp := new(repository.Asset)
	tmp.Name = asset.Name
	tmp.Symbol = asset.Symbol
	tmp.ID, _ = uuid.Parse(asset.ID)
	if asset.IssueDate == 0 {
		return nil, fmt.Errorf("issue date must be declared")
	} else {
		tmp.IssueDate = time.Unix(asset.IssueDate, 0)
	}
	_, err := repository.Assets.Set(tmp)
	return &api.Void{}, err
}

func (s *Service) Get(_ context.Context, req *chipmunkApi.GetAssetRequest) (*chipmunkApi.Asset, error) {
	asset, err := repository.Assets.ReturnBySymbol(req.Name)
	if err != nil {
		return nil, err
	}
	return &chipmunkApi.Asset{
		ID:        asset.ID.String(),
		Name:      asset.Name,
		Symbol:    asset.Symbol,
		IssueDate: asset.IssueDate.Unix(),
	}, nil
}

func (s *Service) List(_ context.Context, req *chipmunkApi.GetAssetListRequest) (*chipmunkApi.Assets, error) {
	if list, err := repository.Assets.List(int(req.Page)); err != nil {
		return nil, err
	} else {
		assets := new(chipmunkApi.Assets)
		assetList := make([]*chipmunkApi.Asset, len(list))
		for i, asset := range list {
			assetList[i] = &chipmunkApi.Asset{
				ID:        asset.ID.String(),
				Name:      asset.Name,
				Symbol:    asset.Symbol,
				IssueDate: asset.IssueDate.Unix(),
			}
		}
		assets.Elements = assetList
		return assets, nil
	}
}
