package repository

import (
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
* Date: 25.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Resolution struct {
	gormext.IncrementalModel
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
	if err := r.db.Model(new(Resolution)).Where("id LIKE ?", resolution.ID).First(found).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.db.Model(new(Resolution)).Create(resolution).Error
		}
		return err
	}
	return r.db.Model(new(Resolution)).Where("id LIKE ?", resolution.ID).Save(resolution).Error
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
	return resolutions, r.db.Model(new(Resolution)).
		Where("brokerage_Name LIKE ?", brokerageName).
		Find(&resolutions).Error
}
