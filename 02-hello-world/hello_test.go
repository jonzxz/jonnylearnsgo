package hello

// Execute with go tests
import (
	"testing"
)

/*
Structure
Tests will always take a pointer of testing.T

func Test(t *Testing.t) {
	want := // expectations
	got :=  // return

	if want != got { // error checking
		t.Errorf("error")
	}
}
*/

func TestSayHelloSingleDefault(t *testing.T) {
	want := "Hello, world!"
	got := Say([]string{})

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

func TestSayHelloSingleCustom(t *testing.T) {
	want := "Hello, Jonny!"
	got := Say([]string{"Jonny"})

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

func TestSayHelloSubTest(t *testing.T) {
	// subtest is a list/slice of structs
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"A", "B", "C"},
			result: "Hello, A, B, C!",
		},
	}

	for _, subtest := range subtests {
		if s := Say(subtest.items); s != subtest.result {
			t.Errorf("wanted %s [%v]], got %s", subtest.result, subtest.items, s)
		}
	}
}
