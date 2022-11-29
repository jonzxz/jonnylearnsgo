package main

import (
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	want := []int{1, 2, 1, 1, 2, 1}
	got := getConcatenation([]int{1, 2, 1})

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("Want %v, got %v", want, got)
			break
		}

	}
}
