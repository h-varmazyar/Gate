package db

import (
	"fmt"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func (db *DB) postsMigrations() error {
	log.Infof("migrating posts")
	var err error
	migrations := make(map[string]struct{})
	tags := make([]string, 0)
	err = db.PostgresDB.Table(MigrationTable).Where("table_name = ?", postsTable).Select("tag").Find(&tags).Error
	if err != nil {
		return err
	}

	for _, tag := range tags {
		migrations[tag] = struct{}{}
	}

	newMigrations := make([]*Migration, 0)
	err = db.PostgresDB.Transaction(func(tx *gorm.DB) error {
		if _, ok := migrations["v1.0.0"]; !ok {
			log.Infof("migrating v1")
			err = tx.AutoMigrate(new(entity.Post))
			if err != nil {
				return err
			}
			newMigrations = append(newMigrations, &Migration{
				TableName:   postsTable,
				Tag:         "v1.0.0",
				Description: fmt.Sprintf("create %v table", postsTable),
			})
		}
		if _, ok := migrations["v1.1.0"]; !ok {
			if !tx.Migrator().HasColumn(new(entity.Post), "sentiment") {
				err = tx.Migrator().AddColumn(new(entity.Post), "sentiment")
			}
			if err != nil {
				return err
			}
			newMigrations = append(newMigrations, &Migration{
				TableName:   postsTable,
				Tag:         "v1.1.0",
				Description: fmt.Sprintf("add sentiment column to %v table", postsTable),
			})
		}

		err = tx.Model(new(Migration)).CreateInBatches(newMigrations, 100).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) ReturnEmotionlessPost(_ context.Context, maxLength int64) (*entity.Post, error) {
	post := new(entity.Post)

	return post, db.PostgresDB.Model(new(entity.Post)).
		Where("sentiment = ?", chipmunkAPI.Polarity_NOT_DETECTED).
		Where("length(content) < ?", maxLength).First(post).Error
}

func (db *DB) SavePost(_ context.Context, post *entity.Post) error {
	return db.PostgresDB.Save(post).Error
}

func (db *DB) BatchSave(_ context.Context, posts []entity.Post) error {
	return db.PostgresDB.CreateInBatches(posts, 1000).Error
}

func (db *DB) OldestPost(_ context.Context, provider chipmunkAPI.Provider) (*entity.Post, error) {
	post := new(entity.Post)
	return post, db.PostgresDB.Model(new(entity.Post)).Where("provider = ?", provider).Order("id ASC").First(post).Error
}
