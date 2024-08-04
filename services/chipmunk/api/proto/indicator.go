package proto

//
//import (
//	"database/sql/driver"
//	"encoding/json"
//)
//
//func (IndicatorType) InRange(v interface{}) bool {
//	i, ok := IndicatorType_value[v.(IndicatorType).String()]
//	return ok && i > 0
//}
//func (d *IndicatorType) Scan(value interface{}) error {
//	*d = IndicatorType(IndicatorType_value[value.(string)])
//	return nil
//}
//func (d IndicatorType) Value() (driver.Value, error) {
//	return d.String(), nil
//}
//func (d *IndicatorType) UnmarshalJSON(b []byte) error {
//	var str string
//	if err := json.Unmarshal(b, &str); err != nil {
//		return err
//	}
//	*d = IndicatorType(IndicatorType_value[str])
//	return nil
//}
//func (d IndicatorType) MarshalJSON() ([]byte, error) {
//	return json.Marshal(d.String())
//}
//
//// ********************************************************
//func (Source) InRange(v interface{}) bool {
//	i, ok := Source_value[v.(Source).String()]
//	return ok && i > 0
//}
//func (d *Source) Scan(value interface{}) error {
//	*d = Source(Source_value[value.(string)])
//	return nil
//}
//func (d Source) Value() (driver.Value, error) {
//	return d.String(), nil
//}
//func (d *Source) UnmarshalJSON(b []byte) error {
//	var str string
//	if err := json.Unmarshal(b, &str); err != nil {
//		return err
//	}
//	*d = Source(Source_value[str])
//	return nil
//}
//func (d Source) MarshalJSON() ([]byte, error) {
//	return json.Marshal(d.String())
//}
//
////********************************************************
////func ToIndicatorValue(data interface{}) *IndicatorValue {
////	switch data.(type) {
////	case *repository.RSIValue:
////		return ToRSIValue(data.(*repository.RSIValue))
////	case *repository.StochasticValue:
////		return ToStochasticValue(data.(*repository.StochasticValue))
////	case *repository.BollingerBandsValue:
////		return ToBollingerBandsValue(data.(*repository.BollingerBandsValue))
////	case *repository.MovingAverageValue:
////		return ToMovingAverageValue(data.(*repository.MovingAverageValue))
////	}
////	return nil
////}
////
////func ToRSIValue(data *repository.RSIValue) *IndicatorValue {
////	return &IndicatorValue{
////		Type: Indicator_RSI,
////		Value: &IndicatorValue_RSI{
////			RSI: &RSI{
////				RSI: data.RSI,
////			},
////		},
////	}
////}
//
//func GetRSIValue(data *IndicatorValue) *RSI {
//	switch t := data.Value.(type) {
//	case *IndicatorValue_RSI:
//		return t.RSI
//	}
//	return nil
//}
//
////func ToStochasticValue(data *repository.StochasticValue) *IndicatorValue {
////	return &IndicatorValue{
////		Type: Indicator_Stochastic,
////		Value: &IndicatorValue_Stochastic{
////			Stochastic: &Stochastic{
////				IndexK: data.IndexK,
////				IndexD: data.IndexD,
////			},
////		},
////	}
////}
//
//func GetStochasticValue(data *IndicatorValue) *Stochastic {
//	switch t := data.Value.(type) {
//	case *IndicatorValue_Stochastic:
//		return t.Stochastic
//	}
//	return nil
//}
//
////func ToBollingerBandsValue(data *repository.BollingerBandsValue) *IndicatorValue {
////	return &IndicatorValue{
////		Type: Indicator_BollingerBands,
////		Value: &IndicatorValue_BollingerBands{
////			BollingerBands: &BollingerBands{
////				UpperBand: data.UpperBand,
////				LowerBand: data.LowerBand,
////				MA:        data.MA,
////			},
////		},
////	}
////}
//
//func GetBollingerBandsValue(data *IndicatorValue) *BollingerBands {
//	switch t := data.Value.(type) {
//	case *IndicatorValue_BollingerBands:
//		return t.BollingerBands
//	}
//	return nil
//}
//
////func ToMovingAverageValue(data *repository.MovingAverageValue) *IndicatorValue {
////	return &IndicatorValue{
////		Type: Indicator_MovingAverage,
////		Value: &IndicatorValue_MovingAverage{
////			MovingAverage: &MovingAverage{
////				Simple:      data.Simple,
////				Exponential: data.Exponential,
////			},
////		},
////	}
////}
//
//func GetMovingAverageValue(data *IndicatorValue) *MovingAverage {
//	switch t := data.Value.(type) {
//	case *IndicatorValue_MovingAverage:
//		return t.MovingAverage
//	}
//	return nil
//}
