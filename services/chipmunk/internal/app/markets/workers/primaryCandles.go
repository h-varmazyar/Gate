package workers

//
//import (
//	"context"
//	"github.com/h-varmazyar/Gate/pkg/errors"
//	"github.com/h-varmazyar/Gate/pkg/grpcext"
//	"github.com/h-varmazyar/Gate/pkg/mapper"
//	"github.com/h-varmazyar/Gate/services/Dolphin/configs"
//	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
//	candles "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
//	resolutions "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions/service"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
//	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
//	log "github.com/sirupsen/logrus"
//	"google.golang.org/grpc/codes"
//	"gorm.io/gorm"
//	"time"
//)
//
//type PrimaryCandlesWorker struct {
//	candlesService    *candles.Service
//	resolutionService *resolutions.Service
//	functionsService  coreApi.FunctionsServiceClient
//}
//
//func NewPrimaryCandlesWorker(ctx context.Context, configs *Configs, candlesService *candles.Service) (*PrimaryCandlesWorker, error) {
//	if candlesService == nil {
//		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetails("candles service must be initialized")
//	}
//
//	coreConn := grpcext.NewConnection(configs.CoreAddress)
//	return &PrimaryCandlesWorker{
//		candlesService:   candlesService,
//		functionsService: coreApi.NewFunctionsServiceClient(coreConn),
//	}, nil
//}
//
//func (w *PrimaryCandlesWorker) AddMarket() {
//
//}
//
//func (w *PrimaryCandlesWorker) hasCandles(ws *WorkerSettings) (*entity.Candle, bool) {
//	candle, err := w.candlesService.ReturnLastNCandles(ws.ctx, &chipmunkApi.BufferedCandlesRequest{
//		ResolutionID: ws.Resolution.ID.String(),
//		MarketID:     ws.Market.ID.String(),
//		Count:        1,
//	})
//	if err != nil && errors.Code(ws.ctx, err) == codes.NotFound {
//		return nil, false
//	}
//	if candle == nil {
//		return nil, false
//	}
//	response := new(entity.Candle)
//	mapper.Struct(candle, response)
//	return response, true
//}
//
//func (w *PrimaryCandlesWorker) findMarketIssueDate(ws *WorkerSettings) error {
//	if ws.Market.StartTime.Unix() != int64(0) {
//	}
//
//	issueResolution, err := w.resolutionService.GetByDuration(ws.ctx, &chipmunkApi.GetResolutionByDurationRequest{
//		BrokerageName: coreApi.Platform_Coinex.String(),
//		Duration:      int64(time.Hour),
//	})
//	if err != nil {
//		return err
//	}
//
//	to := time.Now()
//	from := to.Add(time.Duration(issueResolution.Duration) * 1000 * -1)
//	ticker := time.NewTicker(time.Millisecond * 100)
//	for {
//		select {
//		case <-ws.ctx.Done():
//			return
//		case <-ticker.C:
//			w.functionsService.OHLC(ws.ctx, &coreApi.OHLCReq{
//				Resolution:  issueResolution,
//				Market:      ws.Market,
//				From:        from.Unix(),
//				To:          to.Unix(),
//				BrokerageID: "",
//			})
//		}
//	}
//}
//
//func (w *PrimaryCandlesWorker) loadPrimaryData(ws *WorkerSettings) error {
//	totalCandles := make([]*entity.Candle, 0)
//	end := false
//	limit := 10000
//	var from time.Time
//
//	for i := 0; ; i += limit {
//		list, err := w.candlesService.ReturnLastNCandles(ws.Market.ID, ws.Resolution.ID, limit, i)
//		if err != nil && err != gorm.ErrRecordNotFound {
//			log.WithError(err).Error("load primary candles failed")
//			return err
//		}
//		totalCandles = append(totalCandles, list...)
//		if len(list) < limit {
//			break
//		}
//	}
//
//	if len(totalCandles) == 0 {
//		from = ws.Market.StartTime
//	} else {
//		from = totalCandles[len(totalCandles)-1].Time
//	}
//
//	for !end {
//		to := from.Add(1000 * ws.Resolution.Duration * time.Second)
//		if to.After(time.Now()) {
//			to = time.Now()
//			end = true
//		}
//		if candles, err := worker.downloadCandlesInfo(ws, from.Unix(), to.Unix()); err != nil {
//			log.WithError(err).Error("get candles failed")
//			return err
//		} else {
//			from = to
//			totalCandles = append(totalCandles, candles...)
//		}
//	}
//	for _, candle := range totalCandles {
//		candle.IndicatorValues = entity.NewIndicatorValues()
//	}
//	for _, indicator := range ws.Indicators {
//		err := indicator.Calculate(totalCandles)
//		if err != nil {
//			return err
//		}
//	}
//	i := len(totalCandles) - configs.Variables.CandleBufferLength
//	if i < 0 {
//		i = 0
//	}
//	for ; i < len(totalCandles); i++ {
//		buffer.Markets.Push(ws.Market.ID, totalCandles[i])
//	}
//	return nil
//}
