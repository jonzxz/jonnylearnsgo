package main

import (
	"testing"
)

func TestCountAverage(t *testing.T) {
	want := 2.0
	got := countAverage(6, 3)

	if want != got {
		t.Errorf("wanted %f, got %f", want, got)
	}
}

func TestCountAverageMultiple(t *testing.T) {
	subtests := []struct {
		sum    float64
		n      int
		result float64
	}{
		{
			sum:    6,
			n:      3,
			result: 2.0,
		},
		{
			sum:    66.32,
			n:      4,
			result: 16.58,
		},
	}

	for _, subtest := range subtests {
		if s := countAverage(subtest.sum, subtest.n); s != subtest.result {
			t.Errorf("wanted %.2f [%v]], got %.2f", subtest.result, subtest.sum, s)
		}
	}
}
