package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"time"
)

type ResolutionRepository interface {
	Set(resolution *entity.Resolution) error
	Return(id uuid.UUID) (*entity.Resolution, error)
	GetByDuration(duration time.Duration, brokerageName string) (*entity.Resolution, error)
	List(brokerageName string) ([]*entity.Resolution, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (ResolutionRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Resolution)); err != nil {
		return nil, err
	}
	return NewResolutionPostgresRepository(ctx, logger, db.PostgresDB)
}
