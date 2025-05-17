package main

import (
	"fmt"

	"nenya.me/greetings"
)

func main() {
	message, _ := greetings.Hello("John")
	fmt.Println(message)
}
