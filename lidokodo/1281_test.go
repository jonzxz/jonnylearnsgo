package main

import "testing"

func TestSubtractProductAndSum(t *testing.T) {
	subTests := []struct {
		n    int
		want int
	}{
		{
			n:    234,
			want: 15,
		},
		{
			n:    4421,
			want: 21,
		},
	}

	for _, subTest := range subTests {
		if test := subtractProductAndSum(subTest.n); test != subTest.want {
			t.Errorf("Wanted %v, got %v", subTest.want, test)
		}
	}
}
