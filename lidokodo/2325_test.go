package main

import (
	"testing"
)

func TestCheckDecodeMessage(t *testing.T) {
	want := "this is a secret"
	got := decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv")

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
