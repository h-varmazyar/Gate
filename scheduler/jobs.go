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
	OHLC = "get_ohlc"
)

func (c *Context) ohlc(job *work.Job) error {
	return nil
}
