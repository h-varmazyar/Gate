package candles

import (
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"gorm.io/gorm"
)

type CandleMariadbRepository struct {
	DB *gorm.DB
}

func (r *CandleMariadbRepository) Save(candle *repository.Candle) error {
	item := new(repository.Candle)
	err := r.DB.Model(new(repository.Candle)).
		Where("time = ?", candle.Time).
		Where("market_id = ?", candle.MarketID).
		Where("resolution_id = ?", candle.ResolutionID).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.DB.Model(new(repository.Candle)).Create(candle).Error
		}
		return err
	}
	candle.ID = item.ID
	return r.DB.Save(candle).Error
}

func (r *CandleMariadbRepository) ReturnLast(marketID, resolutionID uint32) (*repository.Candle, error) {
	item := new(repository.Candle)
	return item, r.DB.Model(new(repository.Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").First(item).Error
}

func (r *CandleMariadbRepository) ReturnList(marketID, resolutionID uint32, offset int) ([]*repository.Candle, error) {
	items := make([]*repository.Candle, 0)
	return items, r.DB.Model(new(repository.Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").Offset(offset).Limit(1000).Find(items).Error
}
