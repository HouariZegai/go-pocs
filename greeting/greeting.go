package main

import (
	"fmt"

	"houarizegai.com/hello"
)

func main() {
	// Call method from another module
	message := hello.Greeting("Houari")
	fmt.Println(message)
}
