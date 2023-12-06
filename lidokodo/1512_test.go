package main

import "testing"

func TestNumIdenticalPairs(t *testing.T) {

	subTests := []struct {
		item   []int
		result int
	}{
		{
			item:   []int{1, 2, 3, 1, 1, 3},
			result: 4,
		}, {
			item:   []int{1, 1, 1, 1},
			result: 6,
		}, {
			item:   []int{1, 2, 3},
			result: 0,
		},
	}

	for _, subTest := range subTests {
		if test := numIdenticalPairs(subTest.item); test != subTest.result {
			t.Errorf("wanted %v, got %v\n", subTest.result, test)
		}
	}
}
