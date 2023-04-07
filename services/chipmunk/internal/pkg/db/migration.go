package db

import "github.com/h-varmazyar/Gate/pkg/gormext"

type Migration struct {
	gormext.UniversalModel
	Table       string
	Tag         string
	Description string
}

const MigrationTable = "migrations"
