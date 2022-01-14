package workers

import (
	"context"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/chipmunk/configs"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository/candles"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 02.12.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type ohlcWorker struct {
	heartbeatInterval time.Duration
	candleService     brokerageApi.CandleServiceClient
}

type Settings struct {
	Context    context.Context
	Market     *brokerageApi.Market
	Resolution *brokerageApi.Resolution
}

var (
	OHLCWorker          *ohlcWorker
	workerCancellations map[string]context.CancelFunc
)

func init() {
	workerCancellations = make(map[string]context.CancelFunc)
	OHLCWorker = new(ohlcWorker)
	OHLCWorker.heartbeatInterval = configs.Variables.OHLCWorkerHeartbeat
	candleConnection := grpcext.NewConnection(fmt.Sprintf(":%v", configs.Variables.GrpcAddresses.Brokerage))
	OHLCWorker.candleService = brokerageApi.NewCandleServiceClient(candleConnection)
}

func (worker *ohlcWorker) AddMarket(settings *Settings) {
	ctx := context.Background()
	if ctx == nil {
		panic("nil ctx1")
	}
	var fn context.CancelFunc
	ctx, fn = context.WithCancel(ctx)
	if ctx == nil {
		panic("nil ctx2")
	}
	settings.Context = ctx
	workerCancellations[fmt.Sprintf("%s > %v", settings.Market.ID, settings.Resolution.ID)] = fn
	buffer.Candles.AddList(settings.Market.ID, settings.Resolution.ID)
	go worker.run(settings)
}

func (worker *ohlcWorker) CancelWorker(marketID string, resolutionID uint32) error {
	fn, ok := workerCancellations[fmt.Sprintf("%s > %v", marketID, resolutionID)]
	if !ok {
		return errors.New("worker stopped before")
	}
	fn()
	delete(workerCancellations, fmt.Sprintf("%s > %v", marketID, resolutionID))
	buffer.Candles.RemoveList(marketID, resolutionID)
	return nil
}

func (worker *ohlcWorker) run(settings *Settings) {
	if err := worker.loadPrimaryData(settings); err != nil {
		_ = worker.CancelWorker(settings.Market.ID, settings.Resolution.ID)
		log.WithError(err).Error("load primary failed")
		return
	}
	ticker := time.NewTicker(worker.heartbeatInterval)
	last, err := repository.Candles.ReturnLast(settings.Market.ID, settings.Resolution.ID)
	if err != nil {
		_ = worker.CancelWorker(settings.Market.ID, settings.Resolution.ID)
		log.WithError(err).Error("load last failed")
		return
	}
	lastTime := last.Time
LOOP:
	for {
		select {
		case <-settings.Context.Done():
			break LOOP
		case <-ticker.C:
			to := time.Now()
			if to.Sub(lastTime) <= time.Second {
				continue
			}
			if err := worker.getCandle(settings, lastTime.Unix(), to.Unix()); err != nil {
				time.Sleep(time.Minute)
				log.WithError(err).Error("get candle failed")
			} else {
				lastTime = to
			}
		}
	}
}

func (worker *ohlcWorker) loadPrimaryData(ws *Settings) error {
	last, err := repository.Candles.ReturnLast(ws.Market.ID, ws.Resolution.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			last = new(candles.Candle)
			last.Time = time.Unix(ws.Market.StartTime, 0)
		} else {
			log.WithError(err).Error("load last candle failed")
			return err
		}
	}
	from := last.Time
	end := false
	for !end {
		to := from.Add(time.Duration(1000*ws.Resolution.Duration) * time.Second)
		if to.After(time.Now()) {
			to = time.Now()
			end = true
		}
		if err := worker.getCandle(ws, from.Unix(), to.Unix()); err != nil {
			fmt.Println("from:", from, "to:", to, "duration:", ws.Resolution.Duration)
			log.WithError(err).Error("save candle failed")
			return err
		}
		from = to
	}
	return nil
}

func (worker *ohlcWorker) getCandle(ws *Settings, from, to int64) error {
	c, err := worker.candleService.OHLC(ws.Context, &brokerageApi.OhlcRequest{
		Resolution: ws.Resolution,
		Market:     ws.Market,
		From:       from,
		To:         to,
	})
	if err != nil {
		return err
	}
	for _, candle := range c.Candles {
		tmp := new(candles.Candle)
		mapper.Struct(candle, tmp)
		tmp.MarketID = ws.Market.ID
		tmp.ResolutionID = ws.Resolution.ID
		err := repository.Candles.Save(tmp)
		if err != nil {
			log.WithError(err).Error("save candle failed")
		}
		buffer.Candles.Enqueue(tmp)
	}
	return nil
}
