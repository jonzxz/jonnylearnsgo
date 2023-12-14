package main

import (
	"testing"
)

func TestNumberOfEmployeesWhoMetTarget(t *testing.T) {
	hours := []int{0, 1, 2, 3, 4}

	target := 2

	want := 3
	got := numberOfEmployeesWhoMetTarget(hours, target)

	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
