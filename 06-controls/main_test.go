package main

import (
	"testing"
)

func TestIfElseSimplePositive(t *testing.T) {
	want := true
	got := ifElseSimple(true)

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestIfElseSimpleNegative(t *testing.T) {
	want := false
	got := ifElseSimple(false)

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestIfElseShortDeclareReturnSmallPositive(t *testing.T) {
	want := 8
	got := ifElseShortDeclareReturnSmall(16)

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestIfElseShortDeclareReturnSmallNegative(t *testing.T) {
	want := 3
	got := ifElseShortDeclareReturnSmall(3)

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestForLoopSum(t *testing.T) {
	want := 15
	got := forLoopSum(5)

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestForArray(t *testing.T) {
	want := 10
	got := forArray([4]int{1, 2, 3, 4})
	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestSwitchInt(t *testing.T) {
	want := "threefourfive"
	got := switchInt(4)
	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}
