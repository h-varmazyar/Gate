package calculator

import (
	"fmt"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"golang.org/x/net/context"
	"testing"
)

func TestBollingerBands_Calculate(t *testing.T) {
	calculate := func(t testing.TB, conf *BollingerBands, candles []*chipmunkAPI.Candle, want []*BollingerBandsValue) {
		t.Helper()
		values := make([]*BollingerBandsValue, len(candles))
		err := conf.Calculate(context.Background(), candles, values)
		if err != nil {
			t.Errorf("failed to calculate rsi: %v", err)
			return
		}

		if len(candles) != len(want) {
			t.Errorf("value and want length mismatch (values(%v) != want(%v))", len(candles), len(want))
		}

		if len(values) != len(candles) {
			t.Errorf("value and candles length mismatch (values(%v) != candles(%v))", len(values), len(candles))
		}

		for i, value := range values {
			if value == nil && want[i] == nil {
				continue
			}
			if value == nil {
				t.Errorf("nil value: %v", i)
				continue
			}
			if want[i] == nil {
				t.Errorf("nil want: %v", i)
				continue
			}
			if fmt.Sprintf("%.6f", value.MA) != fmt.Sprintf("%.6f", want[i].MA) {
				t.Errorf("ma mismatch: %v - %v", value.MA, want[i].MA)
			}
			if fmt.Sprintf("%.6f", value.UpperBand) != fmt.Sprintf("%.6f", want[i].UpperBand) {
				t.Errorf("upper band mismatch: %v - %v", value.UpperBand, want[i].UpperBand)
			}
			if fmt.Sprintf("%.6f", value.LowerBand) != fmt.Sprintf("%.6f", want[i].LowerBand) {
				t.Errorf("lower band mismatch: %v - %v", value.LowerBand, want[i].LowerBand)
			}
		}
	}

	t.Run("calculating", func(t *testing.T) {
		candles := []*chipmunkAPI.Candle{
			{Close: 0.088595, Time: 1704902000}, //1
			{Close: 0.113548, Time: 1704903000}, //2
			{Close: 0.106478, Time: 1704904000}, //3
			{Close: 0.102785, Time: 1704905000}, //4
			{Close: 0.101691, Time: 1704906000}, //5
			{Close: 0.088481, Time: 1704907000}, //6
			{Close: 0.099272, Time: 1704908000}, //7
			{Close: 0.100330, Time: 1704909000}, //8
			{Close: 0.101711, Time: 1704910000}, //9
			{Close: 0.127273, Time: 1704911000}, //10
			{Close: 0.117460, Time: 1704912000}, //11
			{Close: 0.113066, Time: 1704913000}, //12
			//{Close: 0.166468, Time: 1704914000}, //13
			//{Close: 0.170083, Time: 1704915000}, //14
			//{Close: 0.142484, Time: 1704916000}, //15
			//{Close: 0.146316, Time: 1704917000}, //16
			//{Close: 0.145755, Time: 1704918000}, //17
			//{Close: 0.146820, Time: 1704919000}, //18
			//{Close: 0.179320, Time: 1704920000}, //19
			//{Close: 0.167225, Time: 1704921000}, //20
			//{Close: 0.144266, Time: 1704922000}, //21
			//{Close: 0.123643, Time: 1704923000}, //22
			//{Close: 0.109257, Time: 1704924000}, //23
		}

		results := []*BollingerBandsValue{
			nil, //1
			nil, //2
			nil, //3
			nil, //4
			nil, //5
			nil, //6
			nil, //7
			nil, //8
			nil, //9
			&BollingerBandsValue{
				UpperBand: 0.124513,
				LowerBand: 0.081520,
				MA:        0.103016,
			}, //11
			&BollingerBandsValue{
				UpperBand: 0.126616,
				LowerBand: 0.085190,
				MA:        0.105903,
			}, //12
			&BollingerBandsValue{
				UpperBand: 0.126499,
				LowerBand: 0.085211,
				MA:        0.105855,
			}, //13
		}

		conf, _ := NewBollingerBands(10, 2, indicatorsAPI.Source_CLOSE)
		calculate(t, conf, candles, results)
	})
}
