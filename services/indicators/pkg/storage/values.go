package storage

import (
	"github.com/h-varmazyar/Gate/pkg/errors"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"sync"
)

type ValueStorage struct {
	IndicatorId uint
	Type        indicatorsAPI.Type
	Values      []*indicatorsAPI.IndicatorValue
}

var (
	values           map[uint]*ValueStorage //key is indicator id
	valueStorageLock *sync.RWMutex
)

func init() {
	values = make(map[uint]*ValueStorage)
	valueStorageLock = new(sync.RWMutex)
}

func AddValue(ctx context.Context, indicatorId uint, value *indicatorsAPI.IndicatorValue) error {
	if v, ok := values[indicatorId]; !ok {
		var indicatorType indicatorsAPI.Type
		switch _ := value.GetValue().(type) {
		case *indicatorsAPI.IndicatorValue_BollingerBands:
			indicatorType = indicatorsAPI.Type_BOLLINGER_BANDS
		case *indicatorsAPI.IndicatorValue_SMA:
			indicatorType = indicatorsAPI.Type_SMA
		case *indicatorsAPI.IndicatorValue_EMA:
			indicatorType = indicatorsAPI.Type_EMA
		case *indicatorsAPI.IndicatorValue_Stochastic:
			indicatorType = indicatorsAPI.Type_STOCHASTIC
		case *indicatorsAPI.IndicatorValue_RSI:
			indicatorType = indicatorsAPI.Type_RSI
		default:
			return errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_indicator_type")
		}
		v = &ValueStorage{
			IndicatorId: indicatorId,
			Type:        indicatorType,
			Values:      []*indicatorsAPI.IndicatorValue{value},
		}
		valueStorageLock.Lock()
		values[indicatorId] = v
		valueStorageLock.Unlock()
	} else {
		valueStorageLock.Lock()
		last := len(v.Values) - 1
		if v.Values[last].Time == value.Time {
			values[indicatorId].Values[last] = value
		} else if v.Values[last].Time < value.Time {
			values[indicatorId].Values = append(values[indicatorId].Values, value)
		}
		valueStorageLock.Unlock()
	}

	return nil
}

func GetValues(ctx context.Context, indicatorId uint, page, pageSize uint32) ([]*indicatorsAPI.IndicatorValue, error) {
	if values, ok := values[indicatorId]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.NotFound, "indicator_storage_not_found")
	} else {
		from := len(values.Values) - int(pageSize*page)
		to := len(values.Values) - int(pageSize*(page-1))

		if from < 0 {
			from = 0
		}

		if to < 0 {
			to = 0
		}
		return values.Values[from:to], nil
	}
}
