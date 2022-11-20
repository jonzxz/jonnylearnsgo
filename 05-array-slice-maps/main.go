package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// run from stdin with file go run main.go < text.txt
// text.txt has 17 words, 15 uniques
func main() {
	fmt.Printf("Array, Slice, Maps\n====================\n")

	// Examples functions
	// arrayExamples()
	// sliceExamples()
	// copyByValueOrReferenceExample()
	// copySliceToSliceExample()
	// mapExamples()
	// updateMapByReferenceExample()
	// mapTwoResultLookupExample()
	// makingNilUseful()

	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++ // returns a string from scan and increment the word found
	}

	fmt.Printf("Unique words: %v\n", len(words))

	// struct called kv
	type kv struct {
		key string
		val int
	}

	// a slice containing struct kvs
	var ss []kv

	for k, v := range words {
		// create a new kv struct of k, v and append into ss
		ss = append(ss, kv{k, v})
	}

	fmt.Printf("Before sorting: %v\n\n", ss)
	// literal fns, like anonymous fns in js
	// compare item in index i and j
	sort.Slice(ss, func(i int, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss {
		fmt.Printf("word %v appears %v times\n", s.key, s.val)
	}
}

// Arrays
func arrayExamples() {
	// Arrays
	var exampleArrayZeroInit [4]int
	exampleArrayDeclared := [4]int{1, 2, 3, 4}
	fmt.Printf("\tArrays\n")
	fmt.Printf("Zero initialized array: %v\n", exampleArrayZeroInit)
	fmt.Printf("Declared array: %v\n", exampleArrayDeclared)
}

// End Arrays

// Slices
func sliceExamples() {
	// Slices
	var exampleSliceZeroInit []int
	exampleSliceDeclared := []int{5, 6, 7, 8}
	fmt.Printf("\tSlices\n")
	// empty because zero-initialised on dynamic list
	fmt.Printf("Zero initialized slice: %v\n", exampleSliceZeroInit)
	fmt.Printf("Declared slice: %v\n", exampleSliceDeclared)

	exampleSliceZeroInit = append(exampleSliceZeroInit, 99)
	fmt.Printf("Appended 99 to exampleSliceZeroInit -- %v\n", exampleSliceZeroInit)

	exampleSliceDeclared = append(exampleSliceDeclared, 101)
	fmt.Printf("Appended 101 to exampleSliceDeclared -- %v\n", exampleSliceDeclared)

	// Test Slice Copies
	exampleSliceOriginal := []int{1, 2, 3}
	exampleSliceCopied := []int{}

	fmt.Printf("Before copy -- original: %v, copied: %v\n", exampleSliceOriginal, exampleSliceCopied)
	exampleSliceCopied = exampleSliceOriginal
	exampleSliceOriginal = append(exampleSliceOriginal, 4)
	fmt.Printf("After copy and append to original -- original: %v, copy: %v\n",
		exampleSliceOriginal, exampleSliceCopied)

	// comparing
	fmt.Printf("0th Element of original and copy is the same? %v\n",
		exampleSliceOriginal[0] == exampleSliceCopied[0])

	// subslice test
	subSlice := [6]int{0, 1, 2, 3, 4, 5}

	// mathematically written is [3, 5)
	// indicating index 3 inclusive, index 5 exclusive
	slicedSubSlice := subSlice[3:5]
	fmt.Printf("3:5 is -- %v\n", slicedSubSlice)
}

func copyByValueOrReferenceExample() {
	a := [3]int{0, 1, 2}
	b := []int{0, 1, 2}

	fmt.Printf("Before -- Value of a is %v\nValue of b is %v\n", a, b)
	// In this case, original array does not get modified because it's copied by value
	// but the original slice is modified because it's copied by reference
	doReplaceValueInArrayOrSlice(a, b)
	fmt.Printf("After -- Value of a is %v\nValue of b is %v\n", a, b)
}

func doReplaceValueInArrayOrSlice(a [3]int, b []int) {
	a[0] = 4
	b[0] = 4
}

func copySliceToSliceExample() {
	a := []int{0, 1, 2}
	fmt.Printf("a before: %v\n", a)
	b := doCopySlice(a)
	fmt.Printf("return value of doCopySlice: %v\n", b)
}

func doCopySlice(a []int) []int {
	// make an empty slice of size 5
	b := make([]int, 5)
	// modify by reference 0th element of a to 4, so [0, 1, 2] -> [4, 1, 2]
	a[0] = 4

	b[4] = 69
	// copies a into b, b is now [0, 0, 0, 0, 69]
	// after copying a into b, b is now [4, 1, 2, 0, 69]
	copy(b, a)
	return b
}

// End Slices

// Maps
func mapExamples() {
	// Maps
	fmt.Printf("\tMaps\n")
	var exampleMapNilInit map[string]int
	exampleMapMake := make(map[string]int)
	exampleMapDeclared := map[string]int{
		"KeyOne": 1,
		"KeyTwo": 2,
	}
	fmt.Printf("Nil initialized map: %v\n", exampleMapNilInit)
	fmt.Printf("Declared map: %v\n", exampleMapDeclared)
	fmt.Printf("Make declared: %v\n", exampleMapMake)

	// PANIC, map is nil
	// exampleMapNilInit["Key3"] = 3

	// these are ok, because exampleMapDeclared and exampleMapMake is not nil
	exampleMapDeclared["Key3"] = 3
	exampleMapMake["Key3"] = 3
	fmt.Printf("Make declared, updated: %v\n", exampleMapMake)
	// incremental of value of Key3
	exampleMapMake["Key3"]++
	fmt.Printf("Make updated with increment: %v\n", exampleMapMake)
}

func updateMapByReferenceExample() {
	a := map[string]int{
		"KeyOne": 1,
	}
	// since a is "KeyOne" : 1, doUpdateMapByReference
	// function references b to a, and a manipulation in b causes a to be updated
	// this is because b is a reference to the same memory block as a
	// but if b is declared in the function then this does not behave the same way
	fmt.Printf("a is: %v\n", a)
	b := make(map[string]int)
	doUpdateMapByReference(a, b)
	fmt.Printf("a is now after updating b: %v\n", a)
}

func doUpdateMapByReference(a map[string]int, b map[string]int) {
	b = a
	b["KeyOne"]++
}

// special behavior in map to
// return a value of key and a boolean for "if key is found"
func mapTwoResultLookupExample() {
	a := make(map[string]int)

	val, ok := a["abc"]
	fmt.Printf("empty map, result of val is %v and ok is %v\n", val, ok)

	// this conditional is skipped because ok returns 0
	// conditional only goes in IF ok is truthy
	if v, ok := a["abc"]; ok {
		fmt.Printf("Value of v is %v\n", v)
	}

	a["abc"]++
	valNow, ok := a["abc"]
	fmt.Printf("map[\"abc\"] is incremented, result of valNow is %v and ok is %v\n", valNow, ok)

	// this conditional now goes in because key is found
	if v, ok := a["abc"]; ok {
		fmt.Printf("Value of v is %v\n", v)
	}

	z := make([]int, 4)
	fmt.Printf("%v\n", len(z))
}

// End maps

// Misc
func makingNilUseful() {
	var sVar []int
	sShort := []int{}
	sMake := make([]int, 2)

	// _ is index, but can't use it if its _
	// this does not iterate because it's empty
	for _, v := range sVar {
		fmt.Printf("%v\n", v)
	}

	// this does not iterate because it's empty
	for _, v := range sShort {
		fmt.Printf("%v\n", v)
	}

	// this iterates because it is initialized with zeros
	for _, v := range sMake {
		fmt.Printf("%v\n", v)
	}

}
