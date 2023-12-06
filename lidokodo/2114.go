package main

import (
	"strings"
)

func mostWordsFound(sentences []string) int {
	highestCount := 0

	for _, sentence := range sentences {
		length := len(strings.Split(sentence, " "))
		if length > highestCount {
			highestCount = length
		}
	}
	return highestCount
}
