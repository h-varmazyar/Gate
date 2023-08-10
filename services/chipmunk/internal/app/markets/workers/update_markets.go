package workers

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
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"strings"
	"time"
)

type UpdateMarketWorker struct {
	functionsService coreApi.FunctionsServiceClient
	assetsService    *assets.Service
	logger           *log.Logger
	db               repository.MarketRepository
}

func NewUpdateMarketWorker(db repository.MarketRepository, logger *log.Logger, assetsService *assets.Service, coreAddress string) *UpdateMarketWorker {
	coreConn := grpcext.NewConnection(coreAddress)
	return &UpdateMarketWorker{
		functionsService: coreApi.NewFunctionsServiceClient(coreConn),
		assetsService:    assetsService,
		logger:           logger,
		db:               db,
	}
}

func (w *UpdateMarketWorker) UpdateFromPlatform(ctx context.Context, platform api.Platform) (*chipmunkApi.Markets, error) {
	w.logger.Infof("updating market info for: %v", platform.String())
	markets, err := w.functionsService.MarketList(ctx, &coreApi.MarketListReq{Platform: platform})
	if err != nil {
		w.logger.WithError(err).Errorf("failed to get market list")
		return nil, err
	}
	w.logger.Infof("market list size: %v", len(markets.Elements))

	availableMarkets := make([]uuid.UUID, 0)
	for _, market := range markets.Elements {
		w.logger.Infof("getting market info: %v", market.Name)
		localMarket, err := w.db.ReturnByName(platform, market.Name)
		if err != nil {
			if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
				mapper.Struct(market, localMarket)
				source, sourceErr := w.loadOrCreateAsset(ctx, market.Source.Name)
				if sourceErr != nil {
					w.logger.WithError(err).Errorf("failed to load or create source for market %v", market.Name)
					continue
				}
				destination, destinationErr := w.loadOrCreateAsset(ctx, market.Destination.Name)
				if destinationErr != nil {
					w.logger.WithError(err).Errorf("failed to load or create destination for market %v", market.Name)
					continue
				}
				localMarket.SourceID = source.ID
				localMarket.DestinationID = destination.ID
				localMarket.Platform = platform
				localMarket.Status = api.Status_Enable
				if platform == api.Platform_Coinex {
					marketInfo, err := w.functionsService.GetMarketInfo(ctx, &coreApi.MarketInfoReq{Market: market})
					if err != nil {
						w.logger.WithError(err).Errorf("failed to get market info for %v in Platform %v", market.Name, market.Platform.String())
						continue
					}
					localMarket.IssueDate = time.Unix(marketInfo.IssueDate, 0)
				} else {
					return nil, errors.New(ctx, codes.FailedPrecondition).AddDetails("check market issue date")
				}
				err = w.db.Create(localMarket)
				if err != nil {
					w.logger.WithError(err).Error("failed to create market")
					continue
				}
			} else {
				w.logger.WithError(err).Errorf("failed to fetch market %v", market.Name)
			}
		}
		availableMarkets = append(availableMarkets, localMarket.ID)
	}
	if err = w.deleteOldMarkets(platform, availableMarkets); err != nil {
		w.logger.WithError(err).Errorf("failed to delete old markets for %v", platform.String())
		return nil, err
	}
	return markets, nil
}

func (w *UpdateMarketWorker) loadOrCreateAsset(ctx context.Context, assetName string) (*entity.Asset, error) {
	asset, err := w.assetsService.ReturnBySymbol(ctx, &chipmunkApi.AssetReturnBySymbolReq{Symbol: assetName})
	resp := new(entity.Asset)
	if err != nil {
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			setAsset := &chipmunkApi.AssetCreateReq{
				Name:   assetName,
				Symbol: assetName,
			}
			asset, err = w.assetsService.Create(ctx, setAsset)
			if err != nil {
				log.WithError(err).WithField("asset_name", assetName).Error("failed to create ips")
				return nil, err
			}
		} else {
			log.WithError(err).WithField("asset_name", assetName).Error("failed to get ips")
			return nil, err
		}
	}
	mapper.Struct(asset, resp)
	resp.ID = uuid.MustParse(asset.ID)
	return resp, nil
}

func (w *UpdateMarketWorker) deleteOldMarkets(platform api.Platform, availableMarkets []uuid.UUID) error {
	localMarkets, err := w.db.List(platform)
	if err != nil {
		w.logger.WithError(err).Errorf("failed to return list of markets for %v", platform.String())
		return err
	}
OUTER:
	for _, local := range localMarkets {
		for _, available := range availableMarkets {
			if local.ID == available {
				continue OUTER
			}
		}
		err = w.db.Delete(local)
		if err != nil {
			w.logger.WithError(err).Errorf("failed to delete market %v", local.ID)
			continue
		}
	}
	return nil
}
