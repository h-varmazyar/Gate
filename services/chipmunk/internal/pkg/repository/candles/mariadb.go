package candles

import (
	"gorm.io/gorm"
)

type CandleMariadbRepository struct {
	DB *gorm.DB
}

func (r *CandleMariadbRepository) Save(candle *Candle) error {
	item := new(Candle)
	err := r.DB.Model(new(Candle)).
		Where("time = ?", candle.Time).
		Where("market_id = ?", candle.MarketID).
		Where("resolution_id = ?", candle.ResolutionID).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.DB.Model(new(Candle)).Create(candle).Error
		}
		return err
	}
	candle.ID = item.ID
	return r.DB.Save(candle).Error
}

func (r *CandleMariadbRepository) ReturnLast(marketID, resolutionID uint32) (*Candle, error) {
	item := new(Candle)
	return item, r.DB.Model(new(Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").First(item).Error
}

func (r *CandleMariadbRepository) ReturnList(marketID, resolutionID uint32, offset int) ([]*Candle, error) {
	items := make([]*Candle, 0)
	return items, r.DB.Model(new(Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").Offset(offset).Limit(1000).Find(items).Error
}
