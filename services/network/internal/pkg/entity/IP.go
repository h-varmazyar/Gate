package entity

import "github.com/h-varmazyar/Gate/pkg/gormext"

type IP struct {
	gormext.UniversalModel
	Schema   string
	Address  string
	Port     uint16
	Username string
	Password string
}
