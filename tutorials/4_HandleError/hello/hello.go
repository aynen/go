package main

import (
	"fmt"
	"log"

	"nenya.me/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	message, error := greetings.Hello("John")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
}
