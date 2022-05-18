package repository

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"gorm.io/gorm"
	"time"
)

type Market struct {
	gormext.UniversalModel
	BrokerageID     uuid.UUID
	PricingDecimal  float64
	TradingDecimal  float64
	TakerFeeRate    float64
	MakerFeeRate    float64
	DestinationID   uuid.UUID
	Destination     *Asset `gorm:"->;foreignkey:DestinationID;references:ID"`
	StartTime       time.Time
	MinAmount       float64
	SourceID        uuid.UUID
	Source          *Asset `gorm:"->;foreignkey:SourceID;references:ID"`
	IsAMM           bool
	Name            string
	Status          api.Status
	SourceName      string `gorm:"-"`
	DestinationName string `gorm:"-"`
}

type MarketRepository struct {
	db *gorm.DB
}

func (repository *MarketRepository) Info(brokerageID uuid.UUID, marketName string) (*Market, error) {
	market := new(Market)
	return market, repository.db.Model(new(Market)).
		Where("brokerage_id = ?", brokerageID).
		Where("name LIKE ?", marketName).
		First(market).Error
}

func (repository *MarketRepository) List(brokerageID uuid.UUID) ([]*Market, error) {
	markets := make([]*Market, 0)
	return markets, repository.db.Model(new(Market)).Where("brokerage_id = ?", brokerageID).Find(&markets).Error
}

func (repository *MarketRepository) ListBySource(brokerageID uuid.UUID, source string) ([]*Market, error) {
	markets := make([]*Market, 0)
	return markets, repository.db.Model(new(Market)).Joins("Source").Preload("Destination").
		Where("Source.name LIKE ?", source).
		Where("markets.brokerage_id = ?", brokerageID).Find(&markets).Error
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
	return repository.db.Updates(market).Where("name = ?", market.Name).Error
}
