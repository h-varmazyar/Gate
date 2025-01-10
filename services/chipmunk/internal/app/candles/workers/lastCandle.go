package workers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

type LastCandles struct {
	db               repository.CandleRepository
	configs          *Configs
	functionsService coreApi.FunctionsServiceClient
	ctx              context.Context
	cancelFunc       context.CancelFunc
	logger           *log.Logger
	Started          bool
}

func NewLastCandles(_ context.Context, db repository.CandleRepository, configs *Configs, logger *log.Logger) *LastCandles {
	coreConn := grpcext.NewConnection(configs.CoreAddress)
	return &LastCandles{
		db:               db,
		logger:           logger,
		configs:          configs,
		functionsService: coreApi.NewFunctionsServiceClient(coreConn),
	}
}

func (w *LastCandles) Start(platformsPairs []*PlatformPairs) {
	if !w.Started {
		w.logger.Infof("starting last candle worker(%v)", len(platformsPairs))
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())

		go w.run(platformsPairs)
		w.Started = true
	}
}

func (w *LastCandles) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *LastCandles) run(platformsPairs []*PlatformPairs) {
	before := len(platformsPairs[0].Pairs)
	for _, platformPair := range platformsPairs {
		w.fillEmptyBuffer(platformPair)
	}
	fmt.Println("pairs before:", before, "pairs after:", len(platformsPairs[0].Pairs))

	fmt.Println("starting last candle worker loop")
	ticker := time.NewTicker(w.configs.LastCandlesInterval)
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Infof("last candle stopped")
			ticker.Stop()
			return
		case <-ticker.C:
			fmt.Println("tickkkkk")
			for _, platformPair := range platformsPairs {
				w.checkForLastCandle(platformPair)
			}
		}
	}
}

func (w *LastCandles) fillEmptyBuffer(platformPair *PlatformPairs) {
	newPairs := make([]*Pair, 0)
	for _, pair := range platformPair.Pairs {
		fmt.Println("getting for:", pair.Market.ID, pair.Resolution.ID)
		last := buffer.CandleBuffer.Last(pair.Market.ID, pair.Resolution.ID)
		if last != nil {
			continue
		}

		var err error
		last, err = w.db.ReturnLast(uuid.MustParse(pair.Market.ID), uuid.MustParse(pair.Resolution.ID))
		if err != nil {
			continue
		}
		buffer.CandleBuffer.Push(last)

		newPairs = append(newPairs, pair)

		//end := false
		//limit := 1000000
		//candles := make([]*entity.Candle, 0)
		//for offset := 0; !end; offset += limit {
		//	list, err := w.db.List(uuid.MustParse(pair.Market.ID), uuid.MustParse(pair.Resolution.ID), limit, offset)
		//	if err != nil {
		//		w.logger.WithError(err).Errorf("failed to fetch candle list")
		//		return err
		//	}
		//	if len(list) < limit {
		//		end = true
		//	}
		//	candles = append(candles, list...)
		//}

	}
	platformPair.Pairs = newPairs
}

func (w *LastCandles) checkForLastCandle(platformPair *PlatformPairs) {
	//items := make([]*coreApi.OHLCItem, 0)
	for _, pair := range platformPair.Pairs {
		last := buffer.CandleBuffer.Last(pair.Market.ID, pair.Resolution.ID)
		if last == nil {
			fmt.Println("other fucking empty")
			w.fillEmptyBuffer(platformPair)
			return
		}
		//items = append(items, &coreApi.OHLCItem{
		//	Resolution: pair.Resolution,
		//	Market:     pair.Market,
		//	From:       last.Time.Unix(),
		//	To:         time.Now().Unix(),
		//	Timeout:    int64(w.configs.LastCandlesInterval),
		//	IssueTime:  time.Now().Unix(),
		//})

		resp, err := w.functionsService.OHLC(context.Background(), &coreApi.OHLCReq{
			Item: &coreApi.OHLCItem{
				Resolution: pair.Resolution,
				Market:     pair.Market,
				From:       last.Time.Unix(),
				To:         time.Now().Unix(),
				Timeout:    int64(w.configs.LastCandlesInterval),
				IssueTime:  time.Now().Unix(),
			},
			Platform: platformPair.Platform,
		})
		if err != nil {
			w.logger.Errorf("last candle error: %v", err)
			continue
		}

		fmt.Println("pushing")

		if resp.Elements == nil {
			fmt.Println("nil resp")
			continue
		}

		for _, element := range resp.Elements {
			candle := &entity.Candle{
				Time:         time.Unix(element.Time, 0),
				Open:         element.Open,
				High:         element.High,
				Low:          element.Low,
				Close:        element.Close,
				Volume:       element.Volume,
				Amount:       element.Amount,
				MarketID:     uuid.MustParse(pair.Market.ID),
				ResolutionID: uuid.MustParse(pair.Resolution.ID),
			}

			buffer.CandleBuffer.Push(candle)
		}

	}

	//_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.AsyncOHLCReq{
	//	Items:    items,
	//	Platform: platformPair.Platform,
	//})
	//if err != nil {
	//	w.logger.WithError(err).Errorf("failed to create last candle request for %v", platformPair.Platform)
	//}
}
