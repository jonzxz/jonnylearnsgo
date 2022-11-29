package main

import "testing"

func TestShuffle(t *testing.T) {
	want := []int{2, 3, 5, 4, 1, 7}
	got := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("wanted %v, got %v", want, got)
			break
		}
	}
}
