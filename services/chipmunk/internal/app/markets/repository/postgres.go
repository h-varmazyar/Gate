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
)

type marketPostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewMarketPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (MarketRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &marketPostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (repository *marketPostgresRepository) List(platform api.Platform) ([]*entity.Market, error) {
	markets := make([]*entity.Market, 0)
	return markets, repository.db.Model(new(entity.Market)).Where("platform = ?", platform).Preload("Source").Preload("Destination").Find(&markets).Error
}

func (repository *marketPostgresRepository) ListBySource(platform api.Platform, source string) ([]*entity.Market, error) {
	markets := make([]*entity.Market, 0)
	return markets, repository.db.Model(new(entity.Market)).
		Joins("join ips as source on source.id = markets.source_id").
		Preload("Destination").
		Where("Source.name = ?", source).
		Where("markets.Platform = ?", platform).Find(&markets).Error
}

func (repository *marketPostgresRepository) ReturnByID(id uuid.UUID) (*entity.Market, error) {
	market := new(entity.Market)
	return market, repository.db.Model(new(entity.Market)).
		Where("id = ?", id).
		Find(market).Error
}

func (repository *marketPostgresRepository) ReturnByName(platform api.Platform, marketName string) (*entity.Market, error) {
	market := new(entity.Market)
	return market, repository.db.Model(new(entity.Market)).
		Where("platform = ?", platform).
		Where("name = ?", marketName).
		First(market).Error
}

func (repository *marketPostgresRepository) Create(market *entity.Market) error {
	m := new(entity.Market)
	err := repository.db.Model(new(entity.Market)).Where("name LIKE ?", market.Name).First(m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return repository.db.Save(market).Error
}

func (repository *marketPostgresRepository) Update(market *entity.Market) error {
	return repository.db.Where("id = ?", market.ID).Updates(market).Error
}

func (repository *marketPostgresRepository) Delete(market *entity.Market) error {
	err := repository.db.Delete(market).Error
	if err != nil {
		return err
	}
	return nil
}
