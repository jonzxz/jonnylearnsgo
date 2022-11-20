package main

import (
	"testing"
)

func TestDoCopySlice(t *testing.T) {
	want := []int{4, 1, 2, 0, 69}
	got := doCopySlice([]int{0, 1, 2})

	for i, _ := range got {
		if want[i] != got[i] {
			t.Errorf("wanted %v, got %v", want, got)
			break
		}
	}
}

func TestDoReplaceValueInArrayOrSlice(t *testing.T) {
	sampleArray, sampleSlice := [3]int{0, 1, 2}, []int{0, 1, 2}
	doReplaceValueInArrayOrSlice(sampleArray, sampleSlice)

	wantArray, wantSlice := [3]int{0, 1, 2}, []int{4, 1, 2}

	if wantArray[0] != sampleArray[0] {
		t.Errorf("Update by copy for array accidentally applied, should not be"+
			"updated. Value should have been %v, but it is %v", sampleArray, wantArray)
	}

	if wantSlice[0] != sampleSlice[0] {
		t.Errorf("Update by reference for slice failed, should have been updated in place"+
			"Value should have been %v, but it is %v", sampleSlice, wantSlice)
	}
}

func TestDoUpdateMapByReference(t *testing.T) {
	mapA := map[string]int{
		"KeyOne": 1,
	}
	mapB := make(map[string]int)
	mapB = mapA

	doUpdateMapByReference(mapA, mapB)

	if mapA["KeyOne"] != 2 {
		t.Errorf("Update by reference for map failed. mapB should have updated mapA's KeyOne to increment"+
			"Value should have been %v but it is %[1]v", mapA["KeyOne"])
	}

}
