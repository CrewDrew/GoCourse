package statistic

import "testing"

type testpair struct {
	values          []float64
	calculatedValue float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var testsSumm = []testpair{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1, 1}, 7},
	{[]float64{-1, 1, 2}, 2},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.calculatedValue {
			t.Error(
				"For function Average ", pair.values,
				"expected", pair.calculatedValue,
				"got", v,
			)
		}
	}
}

func TestSummArraySet(t *testing.T) {
	for _, pair := range testsSumm {
		v := SummArray(pair.values)
		if v != pair.calculatedValue {
			t.Error(
				"For function SummArray ", pair.values,
				"expected", pair.calculatedValue,
				"got", v,
			)
		}
	}
}
