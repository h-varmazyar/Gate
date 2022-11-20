package repository

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"time"
)

type ResolutionRepository interface {
	Set(resolution *entity.Resolution) error
	Return(id uuid.UUID) (*entity.Resolution, error)
	ReturnByDuration(duration time.Duration, platform api.Platform) (*entity.Resolution, error)
	List(platform api.Platform) ([]*entity.Resolution, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (ResolutionRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Resolution)); err != nil {
		return nil, err
	}
	return NewResolutionPostgresRepository(ctx, logger, db.PostgresDB)
}
