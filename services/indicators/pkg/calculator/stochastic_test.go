package calculator

import (
	"fmt"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
	"golang.org/x/net/context"
	"testing"
)

func TestStochastic_Calculate(t *testing.T) {
	calculate := func(t testing.TB, conf *Stochastic, candles []*chipmunkAPI.Candle, want []*indicatorsAPI.StochasticValue) {
		t.Helper()
		values, err := conf.Calculate(context.Background(), candles)
		if err != nil {
			t.Errorf("failed to calculate rsi: %v", err)
			return
		}

		if len(candles) != len(want) {
			t.Errorf("value and want length mismatch (values(%v) != want(%v))", len(candles), len(want))
		}

		if len(values.Values) != len(want) {
			t.Errorf("value and candles length mismatch (values(%v) != candles(%v))", len(values.Values), len(want))
		}

		for i, value := range values.Values {
			if value == nil && want[i] == nil {
				continue
			}
			if value == nil {
				t.Errorf("nil value: %v", i)
				continue
			}

			if want[i] == nil && value.GetStochastic() == nil {
				continue
			}
			if fmt.Sprintf("%.2f", value.GetStochastic().IndexK) != fmt.Sprintf("%.2f", want[i].IndexK) {
				t.Errorf("K index mismatch: %v - %v", value.GetStochastic().IndexK, want[i].IndexK)
			}
			if fmt.Sprintf("%.2f", value.GetStochastic().IndexD) != fmt.Sprintf("%.2f", want[i].IndexD) {
				t.Errorf("D index mismatch: %v - %v", value.GetStochastic().IndexD, want[i].IndexD)
			}
		}
	}

	t.Run("calculating", func(t *testing.T) {
		candles := []*chipmunkAPI.Candle{
			{Open: 0.086959, High: 0.190000, Low: 0.070000, Close: 0.088595, Time: 1704902000}, //1
			{Open: 0.089000, High: 0.127221, Low: 0.080540, Close: 0.113548, Time: 1704903000}, //2
			{Open: 0.112413, High: 0.112413, Low: 0.090910, Close: 0.106478, Time: 1704904000}, //3
			{Open: 0.106478, High: 0.127221, Low: 0.095371, Close: 0.102785, Time: 1704905000}, //4
			{Open: 0.103303, High: 0.125000, Low: 0.092000, Close: 0.101691, Time: 1704906000}, //5
			{Open: 0.101656, High: 0.107555, Low: 0.084339, Close: 0.088481, Time: 1704907000}, //6
			{Open: 0.088481, High: 0.111100, Low: 0.086542, Close: 0.099272, Time: 1704908000}, //7
			{Open: 0.099242, High: 0.119999, Low: 0.088975, Close: 0.100330, Time: 1704909000}, //8
			{Open: 0.100330, High: 0.104828, Low: 0.093010, Close: 0.101711, Time: 1704910000}, //9
			{Open: 0.101711, High: 0.137196, Low: 0.097000, Close: 0.127273, Time: 1704911000}, //10
			{Open: 0.127881, High: 0.142000, Low: 0.110316, Close: 0.117460, Time: 1704912000}, //11
			{Open: 0.117388, High: 0.123608, Low: 0.112000, Close: 0.113066, Time: 1704913000}, //12
			{Open: 0.113066, High: 0.168000, Low: 0.113065, Close: 0.166468, Time: 1704914000}, //13
			{Open: 0.167296, High: 0.176291, Low: 0.130000, Close: 0.170083, Time: 1704915000}, //14
			{Open: 0.170083, High: 0.195000, Low: 0.142331, Close: 0.142484, Time: 1704916000}, //15
			{Open: 0.142484, High: 0.175174, Low: 0.128000, Close: 0.146316, Time: 1704917000}, //16
			{Open: 0.145589, High: 0.184000, Low: 0.145000, Close: 0.145755, Time: 1704918000}, //17
			//{Close: 0.146820, Time: 1704919000}, //18
			//{Close: 0.179320, Time: 1704920000}, //19
			//{Close: 0.167225, Time: 1704921000}, //20
			//{Close: 0.144266, Time: 1704922000}, //21
			//{Close: 0.123643, Time: 1704923000}, //22
			//{Close: 0.109257, Time: 1704924000}, //23
		}

		results := []*indicatorsAPI.StochasticValue{
			nil, //1
			nil, //2
			nil, //3
			nil, //4
			nil, //5
			nil, //6
			nil, //7
			nil, //8
			nil, //9
			nil, //10
			nil, //11
			nil, //12
			nil, //13
			{
				IndexK: 83.4,
				IndexD: 0,
			}, //14
			{
				IndexK: 54.12,
				IndexD: 0,
			}, //15
			{
				IndexK: 56.01,
				IndexD: 64.51,
			}, //16
			{
				IndexK: 55.50,
				IndexD: 55.21,
			}, //17
		}

		conf := &entities.StochasticConfigs{
			Period:  14,
			SmoothK: 1,
			SmoothD: 3,
		}

		st, _ := NewStochastic(10, conf, nil, nil)
		calculate(t, st, candles, results)
	})
}

func TestStochastic_UpdateLast(t *testing.T) {
	calculate := func(t testing.TB, conf *Stochastic, primaryCandles, updatingCandles []*chipmunkAPI.Candle, want []*indicatorsAPI.StochasticValue) {
		t.Helper()
		_, err := conf.Calculate(context.Background(), primaryCandles)
		if err != nil {
			t.Errorf("failed to calculate rsi: %v", err)
			return
		}

		for i, candle := range updatingCandles {
			value := conf.UpdateLast(context.Background(), candle)

			if fmt.Sprintf("%.2f", value.GetStochastic().IndexK) != fmt.Sprintf("%.2f", want[i].IndexK) {
				t.Errorf("K index mismatch: %v - %v", value.GetStochastic().IndexK, want[i].IndexK)
			}
			if fmt.Sprintf("%.2f", value.GetStochastic().IndexD) != fmt.Sprintf("%.2f", want[i].IndexD) {
				t.Errorf("D index mismatch: %v - %v", value.GetStochastic().IndexD, want[i].IndexD)
			}
		}
	}

	t.Run("calculating", func(t *testing.T) {
		primaryCandles := []*chipmunkAPI.Candle{
			{Open: 0.086959, High: 0.190000, Low: 0.070000, Close: 0.088595, Time: 1704902000}, //1
			{Open: 0.089000, High: 0.127221, Low: 0.080540, Close: 0.113548, Time: 1704903000}, //2
			{Open: 0.112413, High: 0.112413, Low: 0.090910, Close: 0.106478, Time: 1704904000}, //3
			{Open: 0.106478, High: 0.127221, Low: 0.095371, Close: 0.102785, Time: 1704905000}, //4
			{Open: 0.103303, High: 0.125000, Low: 0.092000, Close: 0.101691, Time: 1704906000}, //5
			{Open: 0.101656, High: 0.107555, Low: 0.084339, Close: 0.088481, Time: 1704907000}, //6
			{Open: 0.088481, High: 0.111100, Low: 0.086542, Close: 0.099272, Time: 1704908000}, //7
			{Open: 0.099242, High: 0.119999, Low: 0.088975, Close: 0.100330, Time: 1704909000}, //8
			{Open: 0.100330, High: 0.104828, Low: 0.093010, Close: 0.101711, Time: 1704910000}, //9
			{Open: 0.101711, High: 0.137196, Low: 0.097000, Close: 0.127273, Time: 1704911000}, //10
			{Open: 0.127881, High: 0.142000, Low: 0.110316, Close: 0.117460, Time: 1704912000}, //11
			{Open: 0.117388, High: 0.123608, Low: 0.112000, Close: 0.113066, Time: 1704913000}, //12
			{Open: 0.113066, High: 0.168000, Low: 0.113065, Close: 0.166468, Time: 1704914000}, //13
			{Open: 0.167296, High: 0.176291, Low: 0.130000, Close: 0.170083, Time: 1704915000}, //14
			{Open: 0.170083, High: 0.195000, Low: 0.142331, Close: 0.142484, Time: 1704916000}, //15
			{Open: 0.142484, High: 0.175174, Low: 0.128000, Close: 0.146316, Time: 1704917000}, //16
			{Open: 0.145589, High: 0.184000, Low: 0.145000, Close: 0.145755, Time: 1704918000}, //17
		}
		updatingCandles := []*chipmunkAPI.Candle{
			{Open: 0.145754, High: 0.155305, Low: 0.135182, Close: 0.146820, Time: 1704919000}, //18
			{Open: 0.146821, High: 0.193215, Low: 0.139240, Close: 0.179320, Time: 1704920000}, //19
			{Open: 0.179319, High: 0.184800, Low: 0.158000, Close: 0.167225, Time: 1704921000}, //20
			{Open: 0.167227, High: 0.170467, Low: 0.114934, Close: 0.144266, Time: 1704922000}, //21
			{Open: 0.144264, High: 0.147446, Low: 0.110000, Close: 0.123643, Time: 1704923000}, //22
			{Open: 0.123641, High: 0.127406, Low: 0.101030, Close: 0.109257, Time: 1704924000}, //23
		}

		results := []*indicatorsAPI.StochasticValue{
			{
				IndexK: 56.46,
				IndexD: 55.99,
			}, //18
			{
				IndexK: 85.83,
				IndexD: 65.93,
			}, //19
			{
				IndexK: 74.39,
				IndexD: 72.23,
			}, //20
			{
				IndexK: 52.15,
				IndexD: 70.79,
			}, //21
			{
				IndexK: 30.04,
				IndexD: 52.19,
			}, //22
			{
				IndexK: 12.51,
				IndexD: 31.56,
			}, //23
		}

		conf := &entities.StochasticConfigs{
			Period:  14,
			SmoothK: 1,
			SmoothD: 3,
		}

		st, _ := NewStochastic(10, conf, nil, nil)
		calculate(t, st, primaryCandles, updatingCandles, results)
	})
}
