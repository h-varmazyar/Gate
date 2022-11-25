package repository

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"time"
)

type resolutionPostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewResolutionPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (ResolutionRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &resolutionPostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (r *resolutionPostgresRepository) Set(resolution *entity.Resolution) error {
	found := new(entity.Resolution)
	tx := r.db.Model(new(entity.Resolution))
	if resolution.ID.String() == "" {
		tx.Where("value = ?", resolution.Value).
			Where("Platform = ?", resolution.Platform).
			Where("duration = ?", resolution.Duration).
			Where("label = ?", resolution.Label)
	} else {
		tx.Where("id = ?", resolution.ID)
	}
	if err := tx.First(found).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.db.Model(new(entity.Resolution)).Create(resolution).Error
		}
		return err
	}
	return r.db.Model(new(entity.Resolution)).Where("id = ?", found.ID).Save(resolution).Error
}

func (r *resolutionPostgresRepository) Return(id uuid.UUID) (*entity.Resolution, error) {
	resolution := new(entity.Resolution)
	return resolution, r.db.Model(new(entity.Resolution)).
		Where("id = ?", id).
		First(resolution).Error
}

func (r *resolutionPostgresRepository) ReturnByDuration(duration time.Duration, platform api.Platform) (*entity.Resolution, error) {
	resolution := new(entity.Resolution)
	return resolution, r.db.Model(new(entity.Resolution)).
		Where("duration = ", duration).
		Where("Platform = ?", platform).
		First(resolution).Error
}

func (r *resolutionPostgresRepository) List(platform api.Platform) ([]*entity.Resolution, error) {
	resolutions := make([]*entity.Resolution, 0)
	err := r.db.Model(new(entity.Resolution)).Where("Platform = ?", platform).Find(&resolutions).Error
	return resolutions, err
}
