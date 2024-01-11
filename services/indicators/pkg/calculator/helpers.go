package calculator

import chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"

func cloneCandle(from *chipmunkAPI.Candle) *chipmunkAPI.Candle {
	if from == nil {
		return nil
	}

	return &chipmunkAPI.Candle{
		UpdatedAt:    from.UpdatedAt,
		CreatedAt:    from.CreatedAt,
		Volume:       from.Volume,
		Amount:       from.Amount,
		Close:        from.Close,
		Open:         from.Open,
		Time:         from.Time,
		High:         from.High,
		Low:          from.Low,
		ID:           from.ID,
		MarketID:     from.MarketID,
		ResolutionID: from.ResolutionID,
	}
}

func cloneCandles(from []*chipmunkAPI.Candle) []*chipmunkAPI.Candle {
	if from == nil {
		return nil
	}

	to := make([]*chipmunkAPI.Candle, len(from))
	for i, candle := range from {
		to[i] = cloneCandle(candle)
	}
	return to
}
