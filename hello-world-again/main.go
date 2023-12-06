package main

import (
	"fmt"
	"log"
	// Hack: edit so that module does not have to be published
	// go mod edit -replace github.com/jonzxz/jonnylearnsgo/greetings=../greetings
	"github.com/jonzxz/jonnylearnsgo/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	InvokeHellos()
}

func InvokeHello() {
	fmt.Println("Hello world!")
	message := greetings.Hello("Jonny")
	fmt.Println(message)
}

func InvokeHelloWithError() {
	message, err := greetings.HelloWithError("JonnyWithError")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

func InvokeHelloRandom() {
	message, err := greetings.HelloRandom("Jonny")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

func InvokeHellos() {
	names := []string {
		"Jonny1",
		"Jonny2",
		"Jonny3",
	}

	messages, err := greetings.Hellos(names)
	
	if err != nil {
		log.Fatal(err)	
	}

	fmt.Println(messages)
}
