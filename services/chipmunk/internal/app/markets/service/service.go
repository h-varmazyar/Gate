package service

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	assets "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/assets/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/workers"
	resolutions "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Service struct {
	networkService     networkAPI.RequestServiceClient
	strategyService    eagleApi.StrategyServiceClient
	functionsService   coreApi.FunctionsServiceClient
	brokerageService   coreApi.BrokerageServiceClient
	resolutionsService *resolutions.Service
	assetsService      *assets.Service
	//indicatorsService  *indicators.Service
	logger *log.Logger
	db     repository.MarketRepository
}

var (
	GrpcService *Service
)

type Dependencies struct {
	AssetsService *assets.Service
	//IndicatorsService  *indicators.Service
	ResolutionsService *resolutions.Service
	//PrimaryDataWorker             *workers.PrimaryDataWorker
	StatisticsWorker *workers.StatisticsWorker
}

func NewService(_ context.Context, logger *log.Logger, configs *Configs, db repository.MarketRepository, dependencies *Dependencies) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		networkConn := grpcext.NewConnection(configs.NetworkAddress)
		eagleConn := grpcext.NewConnection(configs.EagleAddress)
		coreConn := grpcext.NewConnection(configs.CoreAddress)
		GrpcService.networkService = networkAPI.NewRequestServiceClient(networkConn)
		GrpcService.strategyService = eagleApi.NewStrategyServiceClient(eagleConn)
		GrpcService.functionsService = coreApi.NewFunctionsServiceClient(coreConn)
		GrpcService.assetsService = dependencies.AssetsService
		//GrpcService.indicatorsService = dependencies.IndicatorsService
		GrpcService.resolutionsService = dependencies.ResolutionsService
		GrpcService.db = db
		GrpcService.logger = logger
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterMarketServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, req *chipmunkApi.MarketCreateReq) (*chipmunkApi.Market, error) {
	market := new(entity.Market)
	mapper.Struct(req, market)
	var (
		err         error
		destination *chipmunkApi.Asset
		source      *chipmunkApi.Asset
	)
	destination, err = s.assetsService.ReturnBySymbol(ctx, &chipmunkApi.AssetReturnBySymbolReq{Symbol: req.DestinationSymbol})
	if err != nil {

		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			destination, err = s.assetsService.Create(ctx, &chipmunkApi.AssetCreateReq{
				Name:   req.DestinationSymbol,
				Symbol: req.DestinationSymbol,
			})
		} else {
			return nil, errors.Cast(ctx, err).AddDetailF("invalid destination %v", req.DestinationSymbol)
		}
	}

	source, err = s.assetsService.ReturnBySymbol(ctx, &chipmunkApi.AssetReturnBySymbolReq{Symbol: req.SourceSymbol})
	if err != nil {
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			source, err = s.assetsService.Create(ctx, &chipmunkApi.AssetCreateReq{
				Name:   req.SourceSymbol,
				Symbol: req.SourceSymbol,
			})
		} else {
			return nil, errors.Cast(ctx, err).AddDetailF("invalid source %v", req.DestinationSymbol)
		}
	}
	market.DestinationID = uuid.MustParse(destination.ID)
	market.SourceID = uuid.MustParse(source.ID)
	market.Status = req.Status
	if err := s.db.Create(market); err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Market)
	mapper.Struct(market, response)
	return response, nil
}

func (s *Service) Return(_ context.Context, req *chipmunkApi.MarketReturnReq) (*chipmunkApi.Market, error) {
	marketID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	market, err := s.db.ReturnByID(marketID)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Market)
	mapper.Struct(market, response)
	return response, nil
}

func (s *Service) List(_ context.Context, req *chipmunkApi.MarketListReq) (*chipmunkApi.Markets, error) {
	response, err := s.db.List(req.Platform)
	if err != nil {
		return nil, err
	}
	markets := new(chipmunkApi.Markets)
	mapper.Slice(response, &markets.Elements)
	return markets, nil
}

func (s *Service) ListBySource(_ context.Context, req *chipmunkApi.MarketListBySourceReq) (*chipmunkApi.Markets, error) {
	response := new(chipmunkApi.Markets)
	if markets, err := s.db.ListBySource(req.Platform, req.Source); err != nil {
		return nil, err
	} else {
		list := make([]*chipmunkApi.Market, 0)
		mapper.Slice(markets, &list)
		response.Elements = list
		return response, nil
	}
}

func (s *Service) Update(ctx context.Context, req *chipmunkApi.MarketUpdateReq) (*chipmunkApi.Market, error) {
	market := new(entity.Market)
	mapper.Struct(req, market)

	marketID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	market.ID = marketID

	source, err := s.loadOrCreateAsset(ctx, req.SourceSymbol)
	if err != nil {
		log.WithError(err).Errorf("failed to load or create source for market %v", req.SourceSymbol)
		return nil, err
	}
	destination, err := s.loadOrCreateAsset(ctx, req.DestinationSymbol)
	if err != nil {
		log.WithError(err).Errorf("failed to load or create destination for market %v", req.DestinationSymbol)
		return nil, err
	}

	market.SourceID = source.ID
	market.DestinationID = destination.ID

	err = s.db.Update(market)
	if err != nil {
		log.WithError(err).Error("failed to update markets")
		return nil, err
	}

	res := new(chipmunkApi.Market)
	mapper.Struct(market, res)
	return res, nil
}

//func (s *Service) Update(ctx context.Context, req *chipmunkApi.MarketUpdateReq) (*chipmunkApi.Markets, error) {
//	markets, err := s.functionsService.MarketList(ctx, &coreApi.MarketListReq{Platform: req.Platform})
//	if err != nil {
//		return nil, err
//	}
//	for _, market := range markets.Elements {
//		source, sourceErr := s.loadOrCreateAsset(ctx, market.Source.Name)
//		if sourceErr != nil {
//			log.WithError(err).Errorf("failed to load or create source for market %v", market.Name)
//			continue
//		}
//		destination, destinationErr := s.loadOrCreateAsset(ctx, market.Destination.Name)
//		if destinationErr != nil {
//			log.WithError(err).Errorf("failed to load or create destination for market %v", market.Name)
//			continue
//		}
//		localMarket := new(entities.Market)
//		mapper.Struct(market, localMarket)
//		localMarket.SourceID = source.ID
//		localMarket.DestinationID = destination.ID
//		localMarket.Platform = req.Platform
//
//		err = s.db.SaveOrUpdate(localMarket)
//		if err != nil {
//			log.WithError(err).Error("failed to update markets")
//			continue
//		}
//	}
//	return markets, nil
//}

func (s *Service) UpdateFromPlatform(ctx context.Context, req *chipmunkApi.MarketUpdateFromPlatformReq) (*chipmunkApi.Markets, error) {
	markets, err := s.functionsService.MarketList(ctx, &coreApi.MarketListReq{Platform: req.Platform})
	if err != nil {
		return nil, err
	}
	availableMarkets := make([]uuid.UUID, 0)
	for _, market := range markets.Elements {
		localMarket, err := s.db.ReturnByName(req.Platform, market.Name)
		if err != nil {
			if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
				mapper.Struct(market, localMarket)
				source, sourceErr := s.loadOrCreateAsset(ctx, market.Source.Name)
				if sourceErr != nil {
					s.logger.WithError(err).Errorf("failed to load or create source for market %v", market.Name)
					continue
				}
				destination, destinationErr := s.loadOrCreateAsset(ctx, market.Destination.Name)
				if destinationErr != nil {
					s.logger.WithError(err).Errorf("failed to load or create destination for market %v", market.Name)
					continue
				}
				localMarket.SourceID = source.ID
				localMarket.DestinationID = destination.ID
				localMarket.Platform = req.Platform
				localMarket.Status = api.Status_Enable
				//todo: change it to async request
				if req.Platform == api.Platform_Coinex {
					marketInfo, err := s.functionsService.GetMarketInfo(ctx, &coreApi.MarketInfoReq{Market: market})
					if err != nil {
						s.logger.WithError(err).Errorf("failed to get market info for %v in Platform %v", market.Name, market.Platform.String())
						return nil, err
					}
					localMarket.IssueDate = time.Unix(marketInfo.IssueDate, 0)
				} else {
					return nil, errors.New(ctx, codes.FailedPrecondition).AddDetails("check market issue date")
				}
				err = s.db.Create(localMarket)
				if err != nil {
					s.logger.WithError(err).Error("failed to create market")
					continue
				}
			} else {
				s.logger.WithError(err).Errorf("failed to fetch market %v", market.Name)
			}
		}
		availableMarkets = append(availableMarkets, localMarket.ID)
	}
	if err = s.deleteOldMarkets(req.Platform, availableMarkets); err != nil {
		s.logger.WithError(err).Errorf("failed to delete old markets for %v", req.Platform.String())
		return nil, err
	}
	return markets, nil
}

func (s *Service) loadOrCreateAsset(ctx context.Context, assetName string) (*entity.Asset, error) {
	asset, err := s.assetsService.ReturnBySymbol(ctx, &chipmunkApi.AssetReturnBySymbolReq{Symbol: assetName})
	resp := new(entity.Asset)
	if err != nil {
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			setAsset := new(chipmunkApi.AssetCreateReq)
			setAsset.Name = assetName
			setAsset.Symbol = assetName
			_, err = s.assetsService.Create(ctx, setAsset)
			if err != nil {
				log.WithError(err).WithField("asset_name", assetName).Error("failed to create ips")
				return nil, err
			}
			mapper.Struct(setAsset, resp)
			return resp, nil
		}
		log.WithError(err).WithField("asset_name", assetName).Error("failed to get ips")
		return nil, err
	}
	resp = entity.WrapAsset(asset)
	return resp, nil
}

func (s *Service) deleteOldMarkets(platform api.Platform, availableMarkets []uuid.UUID) error {
	localMarkets, err := s.db.List(platform)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to return list of markets for %v", platform.String())
		return err
	}
	//for _, local := range localMarkets {
	//	for _, remote := range newMarkets {
	//		if local.Name == remote.Name {
	//			continue
	//		}
	//	}
	//	err = s.db.Delete(local)
	//	if err != nil {
	//		s.logger.WithError(err).Errorf("failed to delete market %v", local.ID)
	//		continue
	//	}
	//}
	//return nil
OUTER:
	for _, local := range localMarkets {
		for _, available := range availableMarkets {
			if local.ID == available {
				continue OUTER
			}
		}
		err = s.db.Delete(local)
		if err != nil {
			s.logger.WithError(err).Errorf("failed to delete market %v", local.ID)
			continue
		}
	}
	return nil
}

//
//func (s *Service) loadIndicators(ctx context.Context, strategyIndicators *eagleApi.StrategyIndicators) (map[uuid.UUID]indicatorsPkg.Indicator, error) {
//	response := make(map[uuid.UUID]indicatorsPkg.Indicator)
//	for _, strategyIndicator := range strategyIndicators.Elements {
//		indicatorResp, err := s.indicatorsService.Return(ctx, &chipmunkApi.IndicatorReturnReq{ID: strategyIndicator.IndicatorID})
//		indicator := new(entities.Indicator)
//		mapper.Struct(indicatorResp, indicator)
//		if err != nil {
//			return nil, err
//		}
//		var indicatorCalculator indicatorsPkg.Indicator
//		switch indicator.Type {
//		case chipmunkApi.Indicator_RSI:
//			indicatorCalculator, err = indicatorsPkg.NewRSI(indicator.ID, indicator.Configs.RSI)
//		case chipmunkApi.Indicator_Stochastic:
//			indicatorCalculator, err = indicatorsPkg.NewStochastic(indicator.ID, indicator.Configs.Stochastic)
//		case chipmunkApi.Indicator_MovingAverage:
//			indicatorCalculator, err = indicatorsPkg.NewMovingAverage(indicator.ID, indicator.Configs.MovingAverage)
//		case chipmunkApi.Indicator_BollingerBands:
//			indicatorCalculator, err = indicatorsPkg.NewBollingerBands(indicator.ID, indicator.Configs.BollingerBands)
//		}
//		if err != nil {
//			return nil, err
//		}
//		response[indicator.ID] = indicatorCalculator
//	}
//	return response, nil
//}
