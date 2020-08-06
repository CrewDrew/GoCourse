package main

import "testing"

type testpair struct {
	values          []float64
	calculatedValue [2]float64
}

var tests = []testpair{
	{[]float64{1, -5, 9}, [2]float64{}},
	{[]float64{1, -4, 4}, [2]float64{2, 2}},
	{[]float64{1, 3, -4}, [2]float64{1, -4}},
}

func TestSqrtValueSet(t *testing.T) {
	for _, pair := range tests {
		v := SqrtValue(pair.values[0], pair.values[1], pair.values[2])
		if v != pair.calculatedValue {
			t.Error(
				"For function SqrtValue ", pair.values,
				"expected", pair.calculatedValue,
				"got", v,
			)
		}
	}
}
