package gormext

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	UUIDExtension = "uuid-ossp"
)

func EnableExtensions(db *gorm.DB, names ...string) error {
	for _, name := range names {
		if err := db.Exec(fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS \"%v\";", name)).Error; err != nil {
			return err
		}
	}
	return nil
}
