package main

import (
	"fmt"

	"nenya.me/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
