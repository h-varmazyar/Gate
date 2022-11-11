package entity

import "github.com/h-varmazyar/Gate/pkg/gormext"

type IP struct {
	gormext.UniversalModel
	Address  string
	Port     uint16
	Username string
	Password string
}
