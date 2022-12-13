package main

import "testing"

func TestInterpret(t *testing.T) {
	subTests := []struct {
		in   string
		want string
	}{
		{
			in:   "G()(al)",
			want: "Goal",
		},
		{
			in:   "G()()()()(al)",
			want: "Gooooal",
		},
	}

	for _, subTest := range subTests {
		if test := interpret(subTest.in); test != subTest.want {
			t.Errorf("Wanted %v, got %v", subTest.want, subTest.in)
			return
		}
	}
}
