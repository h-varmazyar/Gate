package repository

import (
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/pkg/gormext"
	"gorm.io/gorm"
	"time"
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
* Date: 19.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Market struct {
	gormext.UniversalModel
	PricingDecimal float64
	TradingDecimal float64
	BrokerageID    uuid.UUID
	TakerFeeRate   float64
	MakerFeeRate   float64
	DestinationID  uuid.UUID
	StartTime      time.Time
	MinAmount      float64
	SourceID       uuid.UUID
	IsAMM          bool
	Name           string
	Status         string
}

type MarketRepository struct {
	db *gorm.DB
}

func (repository *MarketRepository) Info(brokerageId, marketName string) (*Market, error) {
	market := new(Market)
	return market, repository.db.Model(new(Market)).
		Where("brokerage_id LIKE ?", brokerageId).
		Where("name LIKE ?", marketName).
		First(market).Error
}

func (repository *MarketRepository) List(brokerageId string) ([]*Market, error) {
	markets := make([]*Market, 0)
	return markets, repository.db.Model(new(Market)).
		Where("brokerage_id LIKE ?", brokerageId).
		Find(&markets).Error
}

func (repository *MarketRepository) ReturnByID(id uuid.UUID) (*Market, error) {
	market := new(Market)
	return market, repository.db.Model(new(Market)).
		Where("id LIKE ?", id).
		Find(market).Error
}

func (repository *MarketRepository) Update(market *Market) error {
	return repository.db.Save(market).Error
}
