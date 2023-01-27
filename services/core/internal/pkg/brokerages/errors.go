package brokerages

import "errors"

var (
	ErrNilResolution      = errors.New("resolution must be declared")
	ErrNilMarket          = errors.New("market must be declared")
	ErrUnImplemented      = errors.New("method not implemented")
	ErrNotAvailable       = errors.New("method not available for this platform")
	ErrWrongStartTime     = errors.New("start time must be grater than 0")
	ErrWrongEndTime       = errors.New("end time must be grater than 0")
	ErrInvalidCandleCount = errors.New("candle count is invalid")
)
