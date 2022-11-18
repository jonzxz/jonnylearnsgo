package main

import (
	"fmt"
	"os"
)

// execute with go run main.go for stdin
// execute with go run main.go < nums.txt or cat nums.txt | go run main.go to pipe in
func main() {
	// defaults to 0 or 0.0
	var sum float64
	var n int

	for {
		var val float64
		// &val will get the pointer to the float val, read line
		_, err := fmt.Fscanln(os.Stdin, &val)

		if err != nil {
			fmt.Println(err)
			break
		}

		// can do n++ as well, but there's only post-increment in golang
		n += 1
		sum += val
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	// since floats and int don't go well together in golang
	// explicit typecasting is required
	fmt.Printf("average is: %.2f\n", countAverage(sum, n))
}

func countAverage(sum float64, n int) float64 {
	return sum / float64(n)
}

func declareStuff() {
	a := 2
	b := 3.1

	fmt.Printf("a: %T, a: %v\n", a, a)
	fmt.Printf("b: %T, b: %v\n", b, b)

	// %[1] says to reuse the first parameter, so in this case instead
	// of ("%T %v", a , a) it can just be ("%T %V", a) because %V uses arg number 1
	fmt.Printf("a: %T  %[1]v\n", a)

	// can't do a = b in Golang, needs to typecast
	// typecasting float64 b into int so that it can be assigned to a
	fmt.Printf("Type of a before typecast: %T, value: %[1]v\n", a)
	a = int(b)
	fmt.Printf("Type of a after typecast: %T, value: %[1]v\n\n", a)

	// the fractional part .1 is gone because of the typecasting, as it should be
	fmt.Printf("Type of b before typecast: %T, value: %[1]v\n", b)
	b = float64(a)
	fmt.Printf("Type of b after typecast: %T, value: %[1]v\n\n", b)

	// some other basic types
	// bool, can't convert this to integer simply in go, eg. 0 = false, etc
	// so typecasting can't happen here, have to do a conditional / expression
	var booleanExample bool = false // or true
	fmt.Printf("Type of booleanExample %T\n", booleanExample)

	// error
	// can only be nil, or non-nil.
	// If it is non nil then it'll contain an error message
	var errorExample error = nil
	fmt.Printf("Value of errorExample %v\n", errorExample)

	/* pointers
	// an address of something
	// pointer may be nil or non nil
	 no pointer manipulation except through package `unsafe`
	*/
	// placeholder just to close this in vsc
	fmt.Printf("\n")

}
