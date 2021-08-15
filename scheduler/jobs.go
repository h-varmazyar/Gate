package scheduler

import (
	"errors"
	"fmt"
	"github.com/gocraft/work"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
	"time"
)

type Job struct {
	Name        JobName
	RunAt       time.Time
	RunInSecond time.Duration
	Period      time.Duration
	Args        map[string]interface{}
}

type JobName string

const (
	SingleOHLC = "single_ohlc"
	RangeOHLC  = "range_ohlc"
)

func (j *Job) EnqueueContinuously() error {
	j.Args["continuously"] = true
	return j.Enqueue()
}

func (j *Job) EnqueuePeriodically() error {
	//todo: redesign with limits of brokerage...
	var err error
	fmt.Println("args:", j.Args)
	if j.Period > 0 {
		second := int64(j.Period / time.Second)
		minute := int64(j.Period / time.Minute)
		hour := int64(j.Period / time.Hour)
		period := "@every "
		if hour > 0 {
			period = fmt.Sprintf("%s%dh", period, hour)
		}
		if minute > 0 {
			period = fmt.Sprintf("%s%dm", period, minute)
		}
		if second > 0 {
			period = fmt.Sprintf("%s%ds", period, second)
		}
		pool.PeriodicallyEnqueue(period, string(j.Name))
	} else {

	}
	return err
}

func (j *Job) Enqueue() error {
	//todo: redesign with limits of brokerage...
	var err error
	fmt.Println("args:", j.Args)
	if j.RunInSecond > 0 {
		t := int64(j.RunInSecond / time.Second)
		_, err = enqueue.EnqueueIn(string(j.Name), t, j.Args)
	} else {
		_, err = enqueue.Enqueue(string(j.Name), j.Args)
	}
	return err
}

func (c *Context) singleOhlc(job *work.Job) error {
	brokerage, ok := (job.Args["brokerage"]).(brokerages.Brokerage)
	if !ok {
		return errors.New("brokerage argument in not valid")
	}
	resolution, ok := (job.Args["resolution"]).(models.Resolution)
	if !ok {
		return errors.New("resolution not set")
	}
	callback, ok := (job.Args["callback"]).(chan api.OHLCResponse)
	if !ok {
		return errors.New("callback channel not available")
	}
	symbol, ok := (job.Args["symbol"]).(brokerages.Symbol)
	if !ok {
		return errors.New("symbol not set")
	}
	t := time.Now().Unix()
	ohlcResponse, err := brokerage.OHLC(symbol, &resolution, t, t)
	if err != nil {
		return err
	}
	if ohlcResponse.Status != "ok" {
		return errors.New(ohlcResponse.Status)
	}
	callback <- *ohlcResponse
	return nil
}

func (c *Context) rangeOhlc(job *work.Job) error {
	return nil
}
