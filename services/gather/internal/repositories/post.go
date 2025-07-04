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

func (r PostRepository) GetOldest(ctx context.Context, provider models.PostProvider) (*models.Post, error) {
	post := new(models.Post)
	return post, r.db.WithContext(ctx).
		Model(&models.Post{}).
		Where("provider = ?", provider).
		Order("posted_at ASC").
		First(post).Error
}

func (r PostRepository) List(ctx context.Context) ([]models.Post, error) {
	posts := make([]models.Post, 0)
	return posts, r.db.WithContext(ctx).Where("sentiment IS NULL").Limit(10).Find(&posts).Error
}

func (r PostRepository) UpdateSentiment(ctx context.Context, postID uint, sentiment float64) error {
	return r.db.WithContext(ctx).Model(&models.Post{}).Where("id = ?", postID).UpdateColumn("sentiment", sentiment).Error
}
