package db

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

func newPostgres(ctx context.Context, dsn string) (*gorm.DB, error) {
	db, err := gormext.Open(gormext.PostgreSQL, dsn)
	if err != nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("can not load repository configs")
	}

	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = gormext.EnableExtensions(tx,
			gormext.UUIDExtension,
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.New(ctx, codes.Internal).AddDetails("failed to migrate database")
	}

	return db, nil
}
