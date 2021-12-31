package repository

import (
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/gormext"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
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
	BrokerageName   string
	PricingDecimal  int
	TradingDecimal  int
	TakerFeeRate    float64
	MakerFeeRate    float64
	DestinationID   uuid.UUID
	StartTime       time.Time
	MinAmount       float64
	SourceID        uuid.UUID
	IsAMM           bool
	Name            string
	Status          api.Status
	SourceName      string `gorm:"-"`
	DestinationName string `gorm:"-"`
}

type MarketRepository struct {
	db *gorm.DB
}

func (repository *MarketRepository) Info(brokerageName, marketName string) (*Market, error) {
	market := new(Market)
	return market, repository.db.Model(new(Market)).
		Where("brokerage_name LIKE ?", brokerageName).
		Where("name LIKE ?", marketName).
		First(market).Error
}

func (repository *MarketRepository) List(brokerageName string) ([]*Market, error) {
	markets := make([]*Market, 0)
	tx := repository.db.Model(new(Market))
	if brokerageName != brokerageApi.Names_All.String() {
		tx = tx.Where("brokerage_name LIKE ?", brokerageName)
	}
	return markets, tx.Find(&markets).Error
}

func (repository *MarketRepository) ReturnByID(id uuid.UUID) (*Market, error) {
	market := new(Market)
	return market, repository.db.Model(new(Market)).
		Where("id LIKE ?", id).
		Find(market).Error
}

func (repository *MarketRepository) SaveOrUpdate(market *Market) error {
	count := int64(0)
	err := repository.db.Model(new(Market)).Where("name LIKE ?", market.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return repository.db.Model(new(Market)).Create(market).Error
	}
	return repository.db.Updates(market).Error
}
