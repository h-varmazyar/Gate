package calculator

import (
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"testing"
)

func TestCalculate(t *testing.T) {
	calculate := func(t testing.TB, conf SMA, candles []*chipmunkAPI.Candle, want []float64) {
		t.Helper()
		conf.Calculate(candles)

		if len(candles) != len(want) {
			t.Errorf("value and want length mismatch (values(%v) != want(%v))", len(candles), len(want))
		}

		values := conf.Values(0, len(candles))

		for i, value := range values {
			if value.Value != want[i] {
				t.Errorf("want and value mismatch(%v) - want %v but got %v", i, want[i], value.Value)
			}
		}
	}

	t.Run("calculating", func(t *testing.T) {
		candles := []*chipmunkAPI.Candle{
			{Close: 7191, Time: 1704902000},
			{Close: 6993, Time: 1704903000},
			{Close: 6623, Time: 1704904000},
			{Close: 6356, Time: 1704905000},
			{Close: 5555, Time: 1704906000},
			{Close: 5473, Time: 1704907000},
			{Close: 5425, Time: 1704908000},
			{Close: 5568, Time: 1704909000},
			{Close: 5077, Time: 1704910000},
			{Close: 5130, Time: 1704911000},
			{Close: 5509, Time: 1704912000},
			{Close: 5426, Time: 1704913000},
			{Close: 5897, Time: 1704914000},
			{Close: 6224, Time: 1704915000},
			{Close: 7169, Time: 1704916000},
			{Close: 6795, Time: 1704917000},
			{Close: 7295, Time: 1704918000},
			{Close: 7462, Time: 1704919000},
		}

		results := []float64{
			719,
			1418,
			2080,
			2716,
			3271,
			3819,
			4361,
			4918,
			5426,
			5939,
			5896,
			5849,
			5854,
			5891,
			6019,
			6096,
			6216,
			6341,
		}

		conf := SMA{
			Period: 10,
			Source: indicatorsAPI.Source_CLOSE,
		}
		calculate(t, conf, candles, results)
	})
}
