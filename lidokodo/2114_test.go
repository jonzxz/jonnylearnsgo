package main

import "testing"

func TestMostWordsFound(t *testing.T) {
	subTests := []struct {
		sentences []string
		want      int
	}{
		{
			sentences: []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"},
			want:      6,
		},
		{
			sentences: []string{"please wait", "continue to fight", "continue to win"},
			want:      3,
		},
	}

	for _, subTest := range subTests {
		if test := mostWordsFound(subTest.sentences); test != subTest.want {
			t.Errorf("Wanted %v, got %v\n", subTest.want, test)
		}
	}
}
