package main

import "testing"

func TestNumJewelsInStonesOld(t *testing.T) {
	subTests := []struct {
		jewels string
		stones string
		result int
	}{
		{
			jewels: "aA",
			stones: "aAAbbbb",
			result: 3,
		},
		{
			jewels: "z",
			stones: "ZZ",
			result: 0,
		},
	}

	for _, subTest := range subTests {
		if test := numJewelsInStonesOld(subTest.jewels, subTest.stones); test != subTest.result {
			t.Errorf("numJewelsInStonesOld: wanted %v, got %v", subTest.result, test)
		}
		if test := numJewelsInStones(subTest.jewels, subTest.stones); test != subTest.result {
			t.Errorf("numJewelsInStones: wanted %v, got %v", subTest.result, test)
		}
	}
}
