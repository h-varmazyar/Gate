package repository

import (
	"github.com/mrNobody95/Gate/pkg/gormext"
	"gorm.io/gorm"
)

type BrokerageName string

type Brokerage struct {
	gormext.UniversalModel
	Name              BrokerageName `gorm:"size:50"`
	Token             string        `gorm:"size:150"`
	Username          string        `gorm:"size:50"`
	Password          string        `gorm:"size:100"`
	DefaultConfigPath string
}

type BrokerageRepository struct {
	db *gorm.DB
}

func (repository *BrokerageRepository) Create(brokerage *Brokerage) error {
	return repository.db.Model(new(Brokerage)).Save(brokerage).Error
}

func (repository *BrokerageRepository) ReturnByName(name string) (*Brokerage, error) {
	brokerage := new(Brokerage)
	return brokerage, repository.db.Where("name LIKE ?", name).First(brokerage).Error
}
