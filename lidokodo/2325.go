package main

import (
	"slices"
	"strings"
)

func decodeMessage(key string, message string) string {
	alphabets := []rune{}

	for i := 'A'; i <= 'Z'; i++ {
		alphabets = append(alphabets, i)
	}

	key = strings.ToLower(key)

	keys := []string{}

	for i := 0; i < len(key); i++ {
		if !slices.Contains(keys, string(key[i])) && string(key[i]) != " " {
			keys = append(keys, string(key[i]))
		}
	}

	mapping := make(map[string]string)

	for i, v := range keys {
		mapping[v] = string(alphabets[i])
	}

	result := ""
	message = strings.ToLower(message)
	for _, v := range message {
		if string(v) == " " {
			result += " "
		} else {
			result += mapping[string(v)]
		}
	}
	return strings.ToLower(result)
}
