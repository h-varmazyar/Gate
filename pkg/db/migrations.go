package db

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
)

type Migration struct {
	gorm.Model
	File    string
	Success bool
	Error   string
}

type migrationValue struct {
	Name  string
	Value string
}

const MigrationTable = "migrations"

func createMigrateTable(ctx context.Context, db *DB) error {
	err := db.WithContext(ctx).AutoMigrate(new(Migration))
	if err != nil {
		return err
	}
	return nil
}

func migrate(ctx context.Context, db *DB) error {
	fmt.Println("migrationnnnnnn")

	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println("Current directory:", dir)

	dirPath := fmt.Sprintf("../migrations")

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	migrations := make([]migrationValue, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			content, err := os.ReadFile(fmt.Sprintf("../migrations/%v", entry.Name())) // Replace with your file name
			if err != nil {
				return err
			}

			migrations = append(migrations, migrationValue{entry.Name(), string(content)})
		}
	}

	migrate := new(Migration)
	err = db.WithContext(ctx).Table(MigrationTable).Where("success = ?", true).Order("created_at desc").First(migrate).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	newMigrationIndex := 0
	for i, m := range migrations {
		if m.Name == migrate.File {
			newMigrationIndex = i + 1
			break
		}
	}

	for i := newMigrationIndex; i < len(migrations); i++ {
		err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			m := &Migration{File: migrations[i].Name}
			err = tx.Exec(migrations[i].Value).Error
			if err != nil {
				fmt.Println("failed to execute migration")
				m.Success = false
				m.Error = err.Error()
			} else {
				m.Success = true
			}
			if err = tx.Table(MigrationTable).Create(m).Error; err != nil {
				fmt.Println("failed to create migration", err)
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}

	}
	return nil
}
