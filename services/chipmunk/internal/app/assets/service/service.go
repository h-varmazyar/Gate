package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/assets/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

type Service struct {
	logger *log.Logger
	db     repository.AssetRepository
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, _ *Configs, db repository.AssetRepository) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
		GrpcService.db = db
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterAssetServiceServer(server, s)
}

func (s *Service) Set(_ context.Context, asset *chipmunkApi.Asset) (*api.Void, error) {
	tmp := new(entity.Asset)
	tmp.Name = asset.Name
	tmp.Symbol = asset.Symbol
	tmp.ID, _ = uuid.Parse(asset.ID)
	if asset.IssueDate == 0 {
		return nil, fmt.Errorf("issue date must be declared")
	} else {
		tmp.IssueDate = time.Unix(asset.IssueDate, 0)
	}
	_, err := s.db.Set(tmp)
	return &api.Void{}, err
}

func (s *Service) Get(_ context.Context, req *chipmunkApi.GetAssetRequest) (*chipmunkApi.Asset, error) {
	asset, err := s.db.ReturnBySymbol(req.Name)
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
	if list, err := s.db.List(int(req.Page)); err != nil {
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
