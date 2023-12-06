package main

import (
	"testing"
)

func TestCheckIfPangram(t *testing.T) {
	want := true
	got := checkIfPangram("thequickbrownfoxjumpsoverthelazydog")

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
