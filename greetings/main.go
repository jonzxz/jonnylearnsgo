package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// func funcName paramName paramType returnType
func Hello(name string) string {
	message := fmt.Sprintf("Hello %v!", name)
	return message
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {
		message, err := HelloRandom(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// function returns a string and an nil-able error
func HelloWithError(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf("Hello %v!", name)
	// return both, nil for error
	return message, nil
}

func HelloRandom(name string) (string, error) {
	formats := []string{
		"Hi %v",
		"Hello %v",
		"Yo %v",
	}

	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf(formats[rand.Intn(len(formats))], name)
	return message, nil
}
