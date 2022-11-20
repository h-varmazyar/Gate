package entity

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"gorm.io/gorm"
	"time"
)

type Asset struct {
	gormext.UniversalModel
	Name   string `json:"name"`
	Symbol string `json:"symbol" gorm:"unique"`
}

func WrapAsset(from *chipmunkApi.Asset) *Asset {
	assetID, _ := uuid.Parse(from.ID)
	return &Asset{
		UniversalModel: gormext.UniversalModel{
			ID:        assetID,
			CreatedAt: time.Unix(from.CreatedAt, 0),
			UpdatedAt: time.Unix(from.UpdatedAt, 0),
			DeletedAt: gorm.DeletedAt{},
		},
		Name:   from.Name,
		Symbol: from.Symbol,
	}
}

func UnWrapAsset(from *Asset) *chipmunkApi.Asset {
	return &chipmunkApi.Asset{
		ID:        from.ID.String(),
		Name:      from.Name,
		Symbol:    from.Symbol,
		CreatedAt: from.CreatedAt.Unix(),
		UpdatedAt: from.UpdatedAt.Unix(),
	}
}
