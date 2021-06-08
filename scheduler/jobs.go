package scheduler

import (
	"github.com/gocraft/work"
	"time"
)

type Job struct {
	Name     JobName
	ExpireAt time.Duration
	Args     map[string]interface{}
}

type JobName string

const (
	SingleOHLC = "single_ohlc"
	RangeOHLC  = "range_ohlc"
)

func (c *Context) singleOhlc(job *work.Job) error {
	return nil
}
func (c *Context) rangeOhlc(job *work.Job) error {
	return nil
}
