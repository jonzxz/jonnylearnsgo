package main

import (
	"fmt"
)

func main() {
	ifElseSimple(true)
	ifElseShortDeclareReturnSmall(3)
	forLoopSum(5)
	forArray([4]int{1, 2, 3, 4})
	forMap(map[string]int{"keyOne": 1, "keyTwo": 2})
	switchInt(2)
}

func ifElseSimple(a bool) bool {
	if a {
		return true
	} else {
		return false
	}
}

func ifElseShortDeclareReturnSmall(b int) int {
	if a := 2 * 4; a <= b {
		return a
	} else {
		return b
	}
}

func forLoopSum(a int) int {
	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}
	return sum
}

// works for slice too
func forArray(a [4]int) int {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	return sum
}

func forMap(m map[string]int) {
	for k, v := range m {
		fmt.Printf("%s -- %d\n", k, v)
	}
}

func switchInt(a int) string {
	var val string

	switch a {
	case 0, 1, 2:
		val = "zeroonetwo"
	case 3, 4, 5:
		val = "threefourfive"
	}
	return val
}
