package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

type mintestpair struct {
	values []float64
	min    float64
}

type maxtestpair struct {
	values []float64
	max    float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

var minTests = []mintestpair{
	{[]float64{1, 2, 3, 4, 5}, 1},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{4, 5, 2, 7, 5, 8, 9, 2, 1}, 1},
	{[]float64{}, 0},
}

var maxTests = []maxtestpair{
	{[]float64{1, 2, 3, 4, 5}, 5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{4, 5, 2, 7, 5, 8, 9, 2, 1}, 9},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}

/**
Problems
1. The test fails. Added check for empty array
2. see above
*/
