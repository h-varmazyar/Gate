package scheduler

import (
	"fmt"
	"github.com/gocraft/work"
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
	return nil
}
func (c *Context) rangeOhlc(job *work.Job) error {
	return nil
}
