package indicators

import (
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
)

type Indicator interface {
	Calculate([]*repository.Candle, interface{}) error
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
