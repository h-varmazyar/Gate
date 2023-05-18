package brokerages

import (
	api "github.com/h-varmazyar/Gate/api/proto"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
)

type Brokerage struct {
	ID           string                  `json:"id"`
	Title        string                  `json:"title"`
	Description  string                  `json:"description"`
	Platform     api.Platform            `json:"Platform"`
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
