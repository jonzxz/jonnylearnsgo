package main

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	split := strings.Split(s, " ")
	result := make([]string, len(split))

	for _, v := range split {
		word := v[0 : len(v)-1]
		order := string(v[len(v)-1])
		intValue, _ := strconv.Atoi(order)
		result[intValue-1] = word
	}

	var builder strings.Builder

	for i, v := range result {
		builder.WriteString(v)
		if i != len(result)-1 {
			builder.WriteString(" ")
		}
	}

	return builder.String()

}
