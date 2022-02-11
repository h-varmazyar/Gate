package indicators

import "github.com/google/uuid"

type Indicator interface {
	Calculate()
	Update() interface{}
	GetID() string
}

type basicConfig struct {
	MarketName string
	id         uuid.UUID
}

type Source string

const (
	SourceCustom Source = "custom"
	SourceOHLC4  Source = "ohlc4"
	SourceClose  Source = "close"
	SourceOpen   Source = "open"
	SourceHigh   Source = "high"
	SourceHLC3   Source = "hlc3"
	SourceLow    Source = "low"
	SourceHL2    Source = "hl2"
)
