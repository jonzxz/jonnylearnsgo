package main

import "strings"

func finalValueAfterOperations(operations []string) int {
	x := 0

	for i := range operations {
		if strings.Contains(operations[i], "--") {
			x -= 1
		} else if strings.Contains(operations[i], "++") {
			x += 1
		}
	}
	return x
}
