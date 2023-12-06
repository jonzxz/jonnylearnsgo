package main

import "testing"

func TestKidsWithCandies(t *testing.T) {
	subTests := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			want:         []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			want:         []bool{true, false, false, false, false},
		},
		{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			want:         []bool{true, false, true},
		},
	}

	for _, subTest := range subTests {
		test := kidsWithCandies(subTest.candies, subTest.extraCandies)
		for i, _ := range test {
			if test[i] != subTest.want[i] {
				t.Errorf("Wanted %v, got %v\n", subTest.want, test)
				return
			}
		}
	}
}
