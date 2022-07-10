package repository

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/gopack/validator"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Client struct {
	gormext.UniversalModel
	TelegramAccountID int64 `gorm:"unique;index"`
	Role              Role  `gorm:"type:varchar(50);default:user"`
}

type clientRepository struct {
	db *gorm.DB
}

func (r *clientRepository) Create(c *Client) error {
	if err := validator.Struct(c); err != nil {
		return err
	}
	return r.db.Create(c).Error
}
