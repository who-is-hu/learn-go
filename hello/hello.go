package main

import (
	"example.com/greetings"
	"fmt"
	"golang.org/x/example/hello/reverse"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"david", "alex", "paul"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	fmt.Println(reverse.String("Hello"), reverse.Int(1234))
}
