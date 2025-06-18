package repositories

import (
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type PostRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewPostRepository(logger *log.Logger, db *gorm.DB) PostRepository {
	return PostRepository{
		db:     db,
		logger: logger,
	}
}

func (r PostRepository) Create(ctx context.Context, post models.Post) error {
	return r.db.WithContext(ctx).Create(&post).Error
}

func (r PostRepository) BatchSave(ctx context.Context, posts []*models.Post) error {
	return r.db.WithContext(ctx).CreateInBatches(&posts, 100).Error
}
