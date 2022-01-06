package repository

import (
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/gormext"
	"gorm.io/gorm"
)

//type BrokerageName string

type Brokerage struct {
	gormext.UniversalModel
	Title        string `gorm:"not null"`
	Description  string `gorm:"type:varchar(1000)"`
	Name         string `gorm:"type:varchar(25);not null"`
	AuthType     string `gorm:"type:varchar(25);not null"`
	Username     string `gorm:"type:varchar(100)"`
	Password     string `gorm:"type:varchar(100)"`
	AccessID     string `gorm:"type:varchar(100)"`
	SecretKey    string `gorm:"type:varchar(100)"`
	Status       string `gorm:"type:varchar(25);not null"`
	ResolutionID uint
	Resolution   Resolution `gorm:"foreignKey:ResolutionID"`
	Markets      []*Market  `gorm:"many2many:brokerage_markets"`
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
	return brokerage, repository.db.Joins("Resolution").Preload("Markets").
		Where("brokerages.id = ?", id).First(brokerage).Error
}

func (repository *BrokerageRepository) List() ([]*Brokerage, error) {
	brokerages := make([]*Brokerage, 0)
	return brokerages, repository.db.Joins("Resolution").Find(&brokerages).Error
}

func (repository *BrokerageRepository) ChangeStatus(brokerage *Brokerage) error {
	return repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("brokerages").Where("status LIKE ?", api.Status_Enable.String()).
			Update("status", api.Status_Disable.String()).Error; err != nil {
			return err
		}
		return tx.Table("brokerages").Where("id = ?", brokerage.ID).Update("status", brokerage.Status).Error
	})
}
