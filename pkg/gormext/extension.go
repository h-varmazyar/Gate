package gormext

import "gorm.io/gorm"

const (
	UUIDExtension = "uuid-ossp"
)

func EnableExtensions(db *gorm.DB, names ...string) error {
	for _, name := range names {
		if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "%v"`, name).Error; err != nil {
			return err
		}
	}
	return nil
}
