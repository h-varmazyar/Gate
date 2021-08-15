package scheduler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
	"github.com/sirupsen/logrus"
	"time"
)

var exitChannels map[uuid.UUID]chan bool
var brokeragesOHLCFrequency map[brokerages.BrokerageName]time.Duration
var brokerageRegisteredSymbol map[brokerages.BrokerageName]int

func init() {
	exitChannels = map[uuid.UUID]chan bool{}
	brokeragesOHLCFrequency = map[brokerages.BrokerageName]time.Duration{}
	brokerageRegisteredSymbol = map[brokerages.BrokerageName]int{}
}

type ScheduleConfig struct {
	Uuid            uuid.UUID
	Symbol          brokerages.Symbol
	Brokerage       brokerages.Brokerage
	StartFrom       time.Time
	LastCandle      *models.Candle
	Resolution      *models.Resolution
	CallbackChannel chan api.OHLCResponse
}

func (conf *ScheduleConfig) SubscribeNewSymbol() error {
	err := conf.validate()
	if err != nil {
		return err
	}
	exitChannels[conf.Uuid] = make(chan bool, 1)
	brokerageRegisteredSymbol[conf.Brokerage.GetName()] = brokerageRegisteredSymbol[conf.Brokerage.GetName()] + 1
	err = conf.updateOHLCFrequency()
	if err != nil {
		delete(exitChannels, conf.Uuid)
		brokerageRegisteredSymbol[conf.Brokerage.GetName()] = brokerageRegisteredSymbol[conf.Brokerage.GetName()] - 1
		return err
	}
	go func(conf *ScheduleConfig) {
		for {
			time.Sleep(brokeragesOHLCFrequency[conf.Brokerage.GetName()])
			if <-exitChannels[conf.Uuid] {
				return
			}
			t := time.Now().Unix()
			response, err := conf.Brokerage.OHLC(conf.Symbol, conf.Resolution, t, t)
			if err != nil {
				logrus.Errorf("get OHLC error for symbol %s of brokerage %s: %v",
					conf.Symbol,
					conf.Brokerage.GetName(),
					err)
				//todo handle errors
			}
			conf.CallbackChannel <- *response
		}
	}(conf)
	return nil
}

func (conf *ScheduleConfig) validate() error {
	if conf.Symbol == "" {
		return errors.New("symbol is not valid")
	}
	if conf.Brokerage == nil {
		return errors.New("brokerage must be declared")
	}
	if conf.CallbackChannel == nil {
		return errors.New("callback channel must be declared")
	}
	if conf.Resolution == nil {
		return errors.New("resolution must be declared")
	}
	if conf.Uuid.String() == "" {
		return errors.New("uuid must be declared")
	}
	return nil
}

func (conf *ScheduleConfig) updateOHLCFrequency() error {
	//todo get next value from network manager
	totalRequestPerMinute := 1
	rmb := float64(totalRequestPerMinute) / float64(brokerageRegisteredSymbol[conf.Brokerage.GetName()])
	brokeragesOHLCFrequency[conf.Brokerage.GetName()] = time.Duration(float64(time.Second) / rmb)
	return nil
}
