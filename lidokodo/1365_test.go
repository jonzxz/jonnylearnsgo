package main

import "testing"

func TestSmallerNumbersThanCurrent(t *testing.T) {

	subTests := []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{8, 1, 2, 2, 3},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			in:   []int{6, 5, 4, 8},
			want: []int{2, 1, 0, 3},
		},
		{
			in:   []int{7, 7, 7, 7},
			want: []int{0, 0, 0, 0},
		},
	}

	for _, subTest := range subTests {
		test := smallerNumbersThanCurrent(subTest.in)
		for i, _ := range test {
			if test[i] != subTest.want[i] {
				t.Errorf("Wanted %v, got %v", subTest.want, test)
				return
			}
		}
	}
}
