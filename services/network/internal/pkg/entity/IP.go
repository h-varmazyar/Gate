package entity

import "github.com/h-varmazyar/Gate/pkg/gormext"

type IP struct {
	gormext.IncrementalModel
	Address  string
	Port     uint16
	Username string
	Password string
}
