package repository

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"gorm.io/gorm"
	"time"
)

type Resolution struct {
	gormext.UniversalModel
	BrokerageName string
	Duration      time.Duration
	Label         string
	Value         string
}

type ResolutionRepository struct {
	db *gorm.DB
}

func (r *ResolutionRepository) Set(resolution *Resolution) error {
	found := new(Resolution)
	tx := r.db.Model(new(Resolution))
	if resolution.ID.String() == "" {
		tx.Where("value LIKE ?", resolution.Value).
			Where("brokerage_name LIKE ?", resolution.BrokerageName).
			Where("duration = ?", resolution.Duration).
			Where("label LIKE ?", resolution.Label)
	} else {
		tx.Where("id = ?", resolution.ID)
	}
	if err := tx.First(found).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.db.Model(new(Resolution)).Create(resolution).Error
		}
		return err
	}
	return r.db.Model(new(Resolution)).Where("id = ?", found.ID).Save(resolution).Error
}

func (r *ResolutionRepository) Return(id uuid.UUID) (*Resolution, error) {
	resolution := new(Resolution)
	return resolution, r.db.Model(new(Resolution)).
		Where("id = ?", id).
		First(resolution).Error
}

func (r *ResolutionRepository) GetByDuration(duration time.Duration, brokerageName string) (*Resolution, error) {
	resolution := new(Resolution)
	return resolution, r.db.Model(new(Resolution)).
		Where("duration = ", duration).
		Where("brokerage_name = ?", brokerageName).
		First(resolution).Error
}

func (r *ResolutionRepository) List(brokerageName string) ([]*Resolution, error) {
	resolutions := make([]*Resolution, 0)
	err := r.db.Model(new(Resolution)).Where("brokerage_Name LIKE ?", brokerageName).Find(&resolutions).Error
	return resolutions, err
}
