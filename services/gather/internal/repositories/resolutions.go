package repositories

import (
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ResolutionRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewResolutionRepository(logger *log.Logger, db *gorm.DB) ResolutionRepository {
	return ResolutionRepository{
		db:     db,
		logger: logger,
	}
}

func (r ResolutionRepository) ReturnByID(ctx context.Context, id uint) (models.Resolution, error) {
	resolution := models.Resolution{}
	return resolution, r.db.WithContext(ctx).First(&resolution, "id = ?", id).Error
}

func (r ResolutionRepository) All(ctx context.Context) ([]models.Resolution, error) {
	resolutions := make([]models.Resolution, 0)
	return resolutions, r.db.WithContext(ctx).Find(&resolutions).Error
}
