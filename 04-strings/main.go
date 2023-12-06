package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ran with go run main.go matt ed < text.txt
// replaces all instances of matt with ed
func main() {
	// other random functions written to test and show stuff
	// bytesVsRunes()
	// stringStructure()
	// stringFunctions()

	// arg count check from stdin
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]

	// kind of like a Java reader readline
	scan := bufio.NewScanner(os.Stdin)

	// function that will scan and return true if its successful
	for scan.Scan() {
		r := searchAndReplace(scan.Text(), old, new)
		fmt.Println(r)
	}
}

func bytesVsRunes() {
	// random cyrillic unicode character
	s := "Ѐlite"

	// the length of a string is the byte-string in the program that is necessary to encode the string in utf-8
	// so visually it's 5 characters, but it's physically 6 bytes in utf-8 encoding so the length is 6..
	// so the length here actually is the bytes required to represent the string rather than the actual length
	fmt.Printf("Type is %T, value is %[1]v, length is %d\n", s, len(s))

	// rune is 32 bit int, typecasting s to []rune is basically like a char array of some sort
	// which will print out the unicode / utf-8 value
	// prints 1024(cyrllic E) 108(l) 105(i) 116(t) 101(e)
	r := []rune(s)
	fmt.Printf("Type is %T, value is %[1]v, length is %d\n", r, len(r))

	// prints 208 128 (cyrllic E) 108(l) 105(i) 116(t) 101(e)
	b := []byte(s)
	fmt.Printf("Type is %T, value is %[1]v, length is %d\n", b, len(b))
}

func stringStructure() {
	s := "hello, world"
	// this hello here actually points to the
	// **same memory address** where the original "hello" in s is stored
	hello := s[:5]
	fmt.Println(hello)

	// in this case, the memory usage of d is actually
	// is going to point to a completely different memory, it just pulls the value from s but stores it separately
	d := s[:5] + "morning" + s[7:]
	fmt.Println(d)

	// since strings are immutable we can't do the following
	// s[5] = 'a' // SYNTAX ERROR

	// string concatenation
	// the original memory of s is taken as reference
	// what happens here is that the whole of original "s" is going to be copied to a new place in memory, concatenated with !!
	// so the hello and d above points to the "original" memory of s, while s is now pointed to the "new" s after the concat
	// which also means, d is basically not affected by !! because of s[7:] because it's taking the memory of s before the concat
	s += "!!"
	fmt.Println(s)
}

func stringFunctions() {
	s := "a string"

	lengthOfString := len(s) // built-in function

	strings.Contains(s, "g") // returns true
	strings.Contains(s, "z") // returns false

	strings.HasPrefix(s, "a")  // returns true
	strings.Index(s, "string") // returns 2, "a" returns 1

	// it does in a copy, so s isn't actually changed since strings are immutable
	// the "s" here is just going to point to a **different block in memory**
	// golang is gc-ed language, so the original memory used for the smallercase s will be removed
	s = strings.ToUpper(s)

	fmt.Printf("Length of string is %d, s is %s\n", lengthOfString, s)
}

func searchAndReplace(inputString string, replacedString string, newString string) string {
	s := strings.Split(inputString, replacedString)
	return strings.Join(s, newString)
}
