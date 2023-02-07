package service

import (
	"context"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/assets/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"strings"
)

type Service struct {
	logger *log.Logger
	db     repository.AssetRepository
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, db repository.AssetRepository) *Service {
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

func (s *Service) Create(_ context.Context, req *chipmunkApi.AssetCreateReq) (*chipmunkApi.Asset, error) {
	asset := new(entity.Asset)
	asset.Name = req.Name
	asset.Symbol = strings.ToUpper(req.Symbol)
	err := s.db.Create(asset)
	if err != nil {
		return nil, err
	}
	resp := entity.UnWrapAsset(asset)
	return resp, nil
}

func (s *Service) ReturnByID(_ context.Context, req *chipmunkApi.AssetReturnByIDReq) (*chipmunkApi.Asset, error) {
	assetID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}

	asset, err := s.db.ReturnByID(assetID)
	if err != nil {
		return nil, err
	}
	resp := entity.UnWrapAsset(asset)
	return resp, nil
}

func (s *Service) ReturnBySymbol(_ context.Context, req *chipmunkApi.AssetReturnBySymbolReq) (*chipmunkApi.Asset, error) {
	asset, err := s.db.ReturnBySymbol(req.Symbol)
	if err != nil {
		return nil, err
	}
	resp := entity.UnWrapAsset(asset)
	return resp, nil
}

func (s *Service) List(_ context.Context, req *chipmunkApi.GetAssetListRequest) (*chipmunkApi.Assets, error) {
	if list, err := s.db.List(int(req.Page)); err != nil {
		return nil, err
	} else {
		assets := new(chipmunkApi.Assets)
		assetList := make([]*chipmunkApi.Asset, len(list))
		for i, asset := range list {
			assetList[i] = entity.UnWrapAsset(asset)
		}
		assets.Elements = assetList
		return assets, nil
	}
}
