package calculator

//import (
//	"context"
//	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
//	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
//	"testing"
//)
//
//func TestCalculate(t *testing.T) {
//	calculate := func(t testing.TB, conf *SMA, candles []*chipmunkAPI.Candle, want []*float64) {
//		t.Helper()
//		values := make([]*SMAValue, len(candles))
//		conf.Calculate(context.Background(), candles, values)
//
//		if len(candles) != len(want) {
//			t.Errorf("value and want length mismatch (values(%v) != want(%v))", len(candles), len(want))
//		}
//
//		if len(values) != len(candles) {
//			t.Errorf("value and candles length mismatch (values(%v) != candles(%v))", len(values), len(candles))
//		}
//
//		for i, value := range values {
//			if value.Value != want[i] {
//				t.Errorf("want and value mismatch(%v) - want %v but got %v", i, *want[i], *value.Value)
//			}
//		}
//	}
//
//	t.Run("calculating", func(t *testing.T) {
//		candles := []*chipmunkAPI.Candle{
//			{Close: 0.088595, Time: 1704902000}, //1
//			{Close: 0.113548, Time: 1704903000}, //2
//			{Close: 0.106478, Time: 1704904000}, //3
//			{Close: 0.102785, Time: 1704905000}, //4
//			{Close: 0.101691, Time: 1704906000}, //5
//			{Close: 0.088481, Time: 1704907000}, //6
//			{Close: 0.099272, Time: 1704908000}, //7
//			{Close: 0.100330, Time: 1704909000}, //8
//			{Close: 0.101711, Time: 1704910000}, //9
//			{Close: 0.127273, Time: 1704911000}, //10
//			{Close: 0.117460, Time: 1704912000}, //11
//			{Close: 0.113066, Time: 1704913000},
//			//{Close: 5897, Time: 1704914000},
//			//{Close: 6224, Time: 1704915000},
//			//{Close: 7169, Time: 1704916000},
//			//{Close: 6795, Time: 1704917000},
//			//{Close: 7295, Time: 1704918000},
//			//{Close: 7462, Time: 1704919000},
//		}
//
//		r10 := 0.103016
//		r11 := 0.104461
//		r12 := 0.105321
//
//		results := []*float64{
//			nil,  //1
//			nil,  //2
//			nil,  //3
//			nil,  //4
//			nil,  //5
//			nil,  //6
//			nil,  //7
//			nil,  //8
//			nil,  //9
//			&r10, //10
//			&r11, //11
//			&r12, //12
//		}
//
//		conf, _ := NewSMA(10, indicatorsAPI.Source_CLOSE)
//		calculate(t, conf, candles, results)
//	})
//}
