package main

import "testing"

func TestSortSentence(t *testing.T) {
	want := "this is a sentence"
	got := sortSentence("a3 is2 sentence4 this1")

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}
