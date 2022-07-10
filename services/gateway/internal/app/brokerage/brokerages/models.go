package brokerages

import (
	"github.com/h-varmazyar/Gate/api"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
)

type Brokerage struct {
	ID           string                  `json:"id"`
	Title        string                  `json:"title"`
	Description  string                  `json:"description"`
	Platform     brokerageApi.Platform   `json:"platform"`
	Markets      *chipmunkApi.Markets    `json:"markets"`
	ResolutionID string                  `json:"resolution_id"`
	Resolution   *chipmunkApi.Resolution `json:"resolution"`
	StrategyID   string                  `json:"strategy_id"`
	Status       api.Status              `json:"status"`
}

type Start struct {
	ID      string `json:"id"`
	Trading bool   `json:"trading"`
	OHLC    bool   `json:"ohlc"`
}
