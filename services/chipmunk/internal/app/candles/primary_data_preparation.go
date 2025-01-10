package candles

import (
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
	marketService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/service"
	resolutionsService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strings"
	"time"
)

func preparePlatformPairs(ctx context.Context, marketService *marketService.Service, resolutionService *resolutionsService.Service) []*workers.PlatformPairs {
	pp := make([]*workers.PlatformPairs, 0)
	platforms := []api.Platform{api.Platform_Coinex}
	for _, platform := range platforms {
		markets, err := marketService.List(ctx, &chipmunkApi.MarketListReq{Platform: platform})
		if err != nil {
			log.WithError(err).Errorf("failed to load markets of platform %v", platform)
			continue
		}
		log.Infof("market len: %v", len(markets.Elements))

		resolutions, err := resolutionService.List(ctx, &chipmunkApi.ResolutionListReq{Platform: platform})
		if err != nil {
			log.WithError(err).Errorf("failed to load resolutions of platform %v", platform)
			continue
		}

		pairs := make([]*workers.Pair, 0)
		for _, market := range markets.Elements {
			for _, resolution := range resolutions.Elements {
				pairs = append(pairs, &workers.Pair{
					Market:     market,
					Resolution: resolution,
				})
			}
		}

		pp = append(pp, &workers.PlatformPairs{
			Platform: platform,
			Pairs:    pairs,
		})
	}

	return pp
}

//func preparePrimaryDataRequests(nextRequestTrigger chan bool, db repository.CandleRepository, platformsPairs []*workers.PlatformPairs, functionsService coreApi.FunctionsServiceClient) {
//	log.Infof("preparing primary candles")
//	interval := time.Second * 30
//	ticker := time.NewTicker(interval)
//	lastRun := time.Now()
//	for _, platformPair := range platformsPairs {
//		for _, pair := range platformPair.Pairs {
//			select {
//			case <-ticker.C:
//				fmt.Println("new tick")
//				if lastRun.Add(interval).Before(time.Now()) {
//					fmt.Println("sent true")
//					nextRequestTrigger <- true
//				}
//			case sig := <-nextRequestTrigger:
//				if sig {
//					lastRun = time.Now()
//					log.Infof("market: %v", pair.Market.Name)
//					item, err := prepareLocalCandlesItem(db, pair)
//					if err != nil {
//						log.WithError(err).Errorf("failed to prepare local candle item")
//						continue
//					}
//					if item == nil {
//						continue
//					}
//					asyncResp, err := functionsService.AsyncOHLC(context.Background(), &coreApi.AsyncOHLCReq{
//						Items:    []*coreApi.OHLCItem{item},
//						Platform: platformPair.Platform,
//					})
//					if err != nil {
//						log.WithError(err).Errorf("failed to create primary async OHLC request for Platform %v", platformPair.Platform)
//						continue
//					}
//					log.Infof("create new bulk request with id %v for %v. estimated execution time: %v", asyncResp.LastRequestID, platformPair.Platform, time.Duration(asyncResp.PredictedIntervalTime))
//				}
//			}
//		}
//	}
//}

func preparePrimaryDataRequests(_ chan bool, db repository.CandleRepository, platformsPairs []*workers.PlatformPairs, functionsService coreApi.FunctionsServiceClient) {
	log.Infof("preparing primary candles")
	for _, platformPair := range platformsPairs {
		for _, pair := range platformPair.Pairs {
			log.Infof("market: %v", pair.Market.Name)
			item, err := prepareLocalCandlesItem(db, pair)
			if err != nil {
				log.WithError(err).Errorf("failed to prepare local candle item")
				continue
			}
			if item == nil {
				continue
			}
			asyncResp, err := functionsService.OHLC(context.Background(), &coreApi.OHLCReq{
				Item:     item,
				Platform: platformPair.Platform,
			})
			if err != nil {
				log.WithError(err).Errorf("failed to create primary async OHLC request for Platform %v", platformPair.Platform)
				continue
			}

			candles := make([]*entity.Candle, 0)
			for _, candle := range asyncResp.Elements {
				tmp := new(entity.Candle)
				mapper.Struct(candle, tmp)
				tmp.MarketID = uuid.MustParse(pair.Market.ID)
				tmp.ResolutionID = uuid.MustParse(pair.Resolution.ID)
			}
			if err := db.BulkInsert(candles); err != nil {
				log.WithError(err).Errorf("failed to insert candles")
			}
			log.Infof("candle insertion done(%v): %v - %v", len(candles), pair.Market.Name, pair.Resolution.Label)
		}
	}
}

func prepareLocalCandlesItem(db repository.CandleRepository, pair *workers.Pair) (*coreApi.OHLCItem, error) {
	var from time.Time
	last, err := db.ReturnLast(uuid.MustParse(pair.Market.ID), uuid.MustParse(pair.Resolution.ID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			last = nil
		} else {
			log.WithError(err).Errorf("failed to return last")
			return nil, err
		}
	}
	if last == nil {
		from = time.Unix(pair.Market.IssueDate, 0)
	} else {
		from = last.Time
	}
	to := time.Now()

	if int64(to.Sub(from)) < pair.Resolution.Duration {
		return nil, nil
	}

	item := &coreApi.OHLCItem{
		Resolution: pair.Resolution,
		Market:     pair.Market,
		From:       from.Unix(),
		To:         time.Now().Unix(),
		IssueTime:  time.Now().Unix(),
	}
	return item, nil
}
