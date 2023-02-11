package indicators

import (
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
)

type Indicator interface {
	Calculate([]*entity.Candle) error
	Update([]*entity.Candle)
	GetType() chipmunkApi.IndicatorType
	GetLength() int
}
