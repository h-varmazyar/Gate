package repository

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"gorm.io/gorm"
)

type Brokerage struct {
	gormext.UniversalModel
	Title        string                `gorm:"not null"`
	Description  string                `gorm:"type:varchar(1000)"`
	Platform     brokerageApi.Platform `gorm:"type:varchar(25);not null"`
	AuthType     api.AuthType          `gorm:"type:varchar(25);not null"`
	Username     string                `gorm:"type:varchar(100)"`
	Password     string                `gorm:"type:varchar(100)"`
	AccessID     string                `gorm:"type:varchar(100)"`
	SecretKey    string                `gorm:"type:varchar(100)"`
	Status       api.Status            `gorm:"type:varchar(25);not null"`
	ResolutionID *uuid.UUID            `gorm:"not null"`
	//ResolutionID *uuid.UUID             `gorm:"type:uuid REFERENCES resolutions(id);not null;index"`
	//Resolution   *Resolution           `gorm:"->;foreignkey:ResolutionID;references:ID"`
	//Markets      []*Market  `gorm:"many2many:brokerage_markets"`
	StrategyID *uuid.UUID `gorm:"not null"`
}

type BrokerageRepository struct {
	db *gorm.DB
}

func (repository *BrokerageRepository) Create(brokerage *Brokerage) error {
	return repository.db.Model(new(Brokerage)).Create(brokerage).Error
}

func (repository *BrokerageRepository) Delete(id uuid.UUID) error {
	return repository.db.Model(new(Brokerage)).Where("id LIKE ?", id).Delete(new(Brokerage)).Error
}

func (repository *BrokerageRepository) ReturnByID(id uuid.UUID) (*Brokerage, error) {
	brokerage := new(Brokerage)
	return brokerage, repository.db.Where("brokerages.id = ?", id).First(brokerage).Error
}

func (repository *BrokerageRepository) ReturnEnable() (*Brokerage, error) {
	brokerage := new(Brokerage)
	return brokerage, repository.db.
		Where("status LIKE ?", api.Status_Enable.String()).Find(&brokerage).Error
}

func (repository *BrokerageRepository) ReturnEnables() ([]*Brokerage, error) {
	brokerages := make([]*Brokerage, 0)
	return brokerages, repository.db.
		Where("status LIKE ?", api.Status_Enable.String()).Find(&brokerages).Error
}

func (repository *BrokerageRepository) List() ([]*Brokerage, error) {
	brokerages := make([]*Brokerage, 0)
	return brokerages, repository.db.Find(&brokerages).Error
}

func (repository *BrokerageRepository) ChangeStatus(brokerageID uuid.UUID) error {
	status := ""
	if err := repository.db.Model(new(Brokerage)).Where("id = ?", brokerageID).Select("status").Scan(&status).Error; err != nil {
		return err
	}
	newStatus := api.Status_Enable
	if status == api.Status_Enable.String() {
		newStatus = api.Status_Disable
	}
	return repository.db.Model(new(Brokerage)).Where("id = ?", brokerageID).Update("status", newStatus).Error
}
