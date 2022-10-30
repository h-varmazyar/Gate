package markets

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	networkService   networkAPI.RequestServiceClient
	strategyService  eagleApi.StrategyServiceClient
	functionsService brokerageApi.FunctionsServiceClient
	brokerageService brokerageApi.BrokerageServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		networkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Network)
		eagleConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Eagle)
		brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		GrpcService.networkService = networkAPI.NewRequestServiceClient(networkConn)
		GrpcService.strategyService = eagleApi.NewStrategyServiceClient(eagleConn)
		GrpcService.functionsService = brokerageApi.NewFunctionsServiceClient(brokerageConn)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterMarketServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, req *chipmunkApi.CreateMarketReq) (*chipmunkApi.Market, error) {
	market := new(repository.Market)
	mapper.Struct(req, market)
	var err error
	if req.Destination == nil {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetailF("destination not found")
	}
	market.DestinationID, err = uuid.Parse(req.Destination.ID)
	if err != nil {
		return nil, errors.Cast(ctx, err).AddDetailF("invalid destination id %v", req.Destination.ID)
	}
	if req.Source == nil {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetailF("source not found")
	}
	market.SourceID, err = uuid.Parse(req.Source.ID)
	if err != nil {
		return nil, errors.Cast(ctx, err).AddDetailF("invalid source id %v", req.Source.ID)
	}
	market.Status = req.Status
	if err := repository.Markets.SaveOrUpdate(market); err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Market)
	mapper.Struct(market, response)
	return response, nil
}

func (s *Service) Return(_ context.Context, req *chipmunkApi.ReturnMarketRequest) (*chipmunkApi.Market, error) {
	marketID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	market, err := repository.Markets.ReturnByID(marketID)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Market)
	mapper.Struct(market, response)
	return response, nil
}

func (s *Service) List(_ context.Context, req *chipmunkApi.MarketListRequest) (*chipmunkApi.Markets, error) {
	brID, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	response, err := repository.Markets.List(brID)
	if err != nil {
		return nil, err
	}
	markets := new(chipmunkApi.Markets)
	mapper.Slice(response, &markets.Elements)
	return markets, nil
}

func (s *Service) ReturnBySource(_ context.Context, req *chipmunkApi.MarketListBySourceRequest) (*chipmunkApi.Markets, error) {
	brID, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Markets)
	if markets, err := repository.Markets.ListBySource(brID, req.Source); err != nil {
		return nil, err
	} else {
		list := make([]*chipmunkApi.Market, 0)
		mapper.Slice(markets, &list)
		response.Elements = list
		return response, nil
	}
}

func (s *Service) StartWorker(ctx context.Context, req *chipmunkApi.WorkerStartReq) (*api.Void, error) {
	var (
		err                error
		brokerageID        uuid.UUID
		resolutionID       uuid.UUID
		markets            []*repository.Market
		resolution         *repository.Resolution
		strategyIndicators *eagleApi.StrategyIndicators
		loadedIndicators   map[uuid.UUID]indicators.Indicator
	)

	if brokerageID, err = uuid.Parse(req.BrokerageID); err != nil {
		return nil, err
	}

	if markets, err = repository.Markets.List(brokerageID); err != nil {
		return nil, err
	}

	if resolutionID, err = uuid.Parse(req.ResolutionID); err != nil {
		return nil, err
	}
	if resolution, err = repository.Resolutions.Return(resolutionID); err != nil {
		return nil, err
	}

	log.Infof("loaded resolution: %v", resolution)

	if strategyIndicators, err = s.strategyService.Indicators(ctx, &eagleApi.StrategyIndicatorReq{StrategyID: req.StrategyID}); err != nil {
		return nil, err
	}

	log.Infof("loaded strategies count: %v", len(strategyIndicators.Elements))

	if loadedIndicators, err = loadIndicators(ctx, strategyIndicators); err != nil {
		return nil, err
	}
	for _, market := range markets {
		settings := &WorkerSettings{
			Market:      market,
			Resolution:  resolution,
			Indicators:  loadedIndicators,
			BrokerageID: brokerageID,
		}
		worker.AddMarket(settings)
		log.Infof("new market added: %v", market.Name)
	}
	return new(api.Void), nil
}

func (s *Service) StopWorker(ctx context.Context, req *chipmunkApi.WorkerStopReq) (*api.Void, error) {
	brokerageID, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	markets, err := repository.Markets.List(brokerageID)
	if err != nil {
		return nil, err
	}
	for _, market := range markets {
		err := worker.DeleteMarket(market.ID)
		if err != nil {
			log.WithError(err).Errorf("failed to delete market %v", market)
		}
	}
	return new(api.Void), nil
}

func (s *Service) Update(ctx context.Context, req *chipmunkApi.MarketUpdateReq) (*chipmunkApi.Markets, error) {
	brokerageID, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	markets, err := s.functionsService.MarketList(ctx, &brokerageApi.MarketListReq{BrokerageID: brokerageID.String()})
	if err != nil {
		return nil, err
	}
	for _, market := range markets.Elements {
		source, sourceErr := loadOrCreateAsset(market.Source.Name)
		if sourceErr != nil {
			log.WithError(err).Errorf("failed to load or create source for market %v", market.Name)
			continue
		}
		destination, destinationErr := loadOrCreateAsset(market.Destination.Name)
		if destinationErr != nil {
			log.WithError(err).Errorf("failed to load or create destination for market %v", market.Name)
			continue
		}
		localMarket := new(repository.Market)
		mapper.Struct(market, localMarket)
		localMarket.SourceID = source.ID
		localMarket.DestinationID = destination.ID
		localMarket.BrokerageID = brokerageID

		err = repository.Markets.SaveOrUpdate(localMarket)
		if err != nil {
			log.WithError(err).Error("failed to update markets")
			continue
		}
	}
	return markets, nil
}

func loadOrCreateAsset(assetName string) (*repository.Asset, error) {
	asset, err := repository.Assets.ReturnBySymbol(assetName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			asset := new(repository.Asset)
			asset.Name = assetName
			asset.Symbol = assetName
			asset.IssueDate = time.Now()
			err = repository.Assets.Create(asset)
			if err != nil {
				log.WithError(err).WithField("asset_name", assetName).Error("failed to create assets")
				return nil, err
			}
			return asset, nil
		}
		log.WithError(err).WithField("asset_name", assetName).Error("failed to get assets")
		return nil, err
	}
	return asset, nil
}

func loadIndicators(_ context.Context, strategyIndicators *eagleApi.StrategyIndicators) (map[uuid.UUID]indicators.Indicator, error) {
	response := make(map[uuid.UUID]indicators.Indicator)
	for _, strategyIndicator := range strategyIndicators.Elements {
		id, err := uuid.Parse(strategyIndicator.IndicatorID)
		if err != nil {
			continue
		}
		indicator, err := repository.Indicators.Return(id)
		if err != nil {
			return nil, err
		}
		var indicatorCalculator indicators.Indicator
		switch indicator.Type {
		case chipmunkApi.Indicator_RSI:
			indicatorCalculator, err = indicators.NewRSI(indicator.ID, indicator.Configs.RSI)
		case chipmunkApi.Indicator_Stochastic:
			indicatorCalculator, err = indicators.NewStochastic(indicator.ID, indicator.Configs.Stochastic)
		case chipmunkApi.Indicator_MovingAverage:
			indicatorCalculator, err = indicators.NewMovingAverage(indicator.ID, indicator.Configs.MovingAverage)
		case chipmunkApi.Indicator_BollingerBands:
			indicatorCalculator, err = indicators.NewBollingerBands(indicator.ID, indicator.Configs.BollingerBands)
		}
		if err != nil {
			return nil, err
		}
		response[indicator.ID] = indicatorCalculator
	}
	return response, nil
}
