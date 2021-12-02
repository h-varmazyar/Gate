package candles

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 02.12.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type CandleMariadbRepository struct {
	DB *gorm.DB
}

func (r *CandleMariadbRepository) Save(candle *Candle) error {
	item := new(Candle)
	err := r.DB.Model(new(Candle)).
		Where("time = ?", candle.Time).
		Where("market_id LIKE ?", candle.MarketID).
		Where("resolution_id LIKE ?", candle.ResolutionID).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("not found:", candle)
			candle.ID = uuid.New()
			return r.DB.Model(new(Candle)).Create(candle).Error
		}
		return err
	}
	candle.ID = item.ID
	return r.DB.Save(candle).Error
}

func (r *CandleMariadbRepository) ReturnLast(marketID, resolutionID string) (*Candle, error) {
	item := new(Candle)
	return item, r.DB.Model(new(Candle)).
		Where("market_id LIKE ?", marketID).
		Where("resolution_id LIKE ?", resolutionID).Order("time desc").First(item).Error
}
