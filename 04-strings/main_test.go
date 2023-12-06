package main

import (
	"testing"
)

func TestSearchAndReplace(t *testing.T) {
	want := "hello jonny"
	got := searchAndReplace("hello world", "world", "jonny")

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}

	// shown only in go test -v
	t.Logf("Before %s -- after %s\n", want, got)
}
