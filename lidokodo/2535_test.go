package main

import (
	"testing"
)

func TestDifferenceOfSumStruct(t *testing.T) {
	subTests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 15, 6, 3},
			want: 9,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: 0,
		},
	}

	for _, subTest := range subTests {
		if got := differenceOfSum(subTest.nums); got != subTest.want {
			t.Errorf("wanted %v, got %v\n", subTest.want, got)
		}
	}
}
