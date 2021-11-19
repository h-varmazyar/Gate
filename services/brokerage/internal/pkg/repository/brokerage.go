package repository

import (
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/pkg/gormext"
	"gorm.io/gorm"
)

type BrokerageName string

type Brokerage struct {
	gormext.UniversalModel
	Name     BrokerageName `gorm:"size:50"`
	AuthType string
	Token    string `gorm:"size:150"`
	Username string `gorm:"size:50"`
	Password string `gorm:"size:100"`
	Status   string
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
	return brokerage, repository.db.Where("id LIKE ?", id).First(brokerage).Error
}

func (repository *BrokerageRepository) Update(brokerage *Brokerage) error {
	return repository.db.Save(brokerage).Error
}

func (repository *BrokerageRepository) ReturnByName(name string) (*Brokerage, error) {
	brokerage := new(Brokerage)
	return brokerage, repository.db.Where("name LIKE ?", name).First(brokerage).Error
}
