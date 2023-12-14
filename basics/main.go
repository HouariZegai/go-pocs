package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 3; i++ {
		fmt.Println("For loop", i)
	}

	// While loop
	i := 0
	for i < 2 {
		fmt.Println("While loop", i)
		i++
	}
}
