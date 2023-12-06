package main

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	want := 1
	got := finalValueAfterOperations([]string{"--X", "X++", "X++"})

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
