package main

import (
	"testing"
)

func TestBuildArray(t *testing.T) {
	want := []int{0, 1, 2, 4, 5, 3}
	got := buildArray([]int{0, 2, 1, 5, 3, 4})

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("wanted %v, got %v", want, got)
			break
		}
	}
}
