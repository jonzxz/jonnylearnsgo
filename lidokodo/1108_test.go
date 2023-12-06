package main

import "testing"

func TestDefangIPaddr(t *testing.T) {
	want := "1[.]1[.]1[.]1"
	got := defangIPaddr("1.1.1.1")

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
