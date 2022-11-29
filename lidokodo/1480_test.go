package main

import "testing"

func TestRunningSum(t *testing.T) {
	want := []int{1, 3, 6, 10}
	got := runningSum([]int{1, 2, 3, 4})

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("wanted %v, got %v", want, got)
			break
		}
	}
}
