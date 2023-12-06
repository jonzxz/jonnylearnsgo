package main

import "testing"

func TestMaximumWealth(t *testing.T) {
	subTests := []struct {
		item   [][]int
		result int
	}{
		{
			item:   [][]int{[]int{1, 2, 3}, []int{3, 2, 1}},
			result: 6,
		},
		{
			item:   [][]int{[]int{2, 8, 7}, []int{7, 1, 3}, []int{1, 9, 5}},
			result: 17,
		},
	}

	for _, subTest := range subTests {
		if test := maximumWealth(subTest.item); test != subTest.result {
			t.Errorf("wanted %v, got %v\n", subTest.result, test)
		}
	}
}
