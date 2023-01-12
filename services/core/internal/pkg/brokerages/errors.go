package brokerages

import "errors"

var (
	ErrNilResolution  = errors.New("resolution must be declared")
	ErrNilMarket      = errors.New("market must be declared")
	ErrWrongStartTime = errors.New("start time must be grater than 0")
	ErrWrongEndTime   = errors.New("end time must be grater than 0")
)
