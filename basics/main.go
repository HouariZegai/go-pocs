package main

import (
	"fmt"
	"strings"
)

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

	// Data types and pointer
	var age int
	age = 10
	fmt.Printf("Age type is: %T\n", age)

	var sum int64 = 999999999999999999
	fmt.Println("Sum: ", sum)
	var usum uint8 = 0
	usum = usum - 2
	fmt.Println("Unsigned sum: ", usum)

	var values = []int{1, 4, 5}
	fmt.Println("Values: ", values)

	var names [10]string
	names[0] = "Mohamed"
	fmt.Scan(&names[1])
	fmt.Println("Names: ", names)
	fmt.Println("Size of the array: ", len(names))

	// Slices
	var notes []int
	notes = append(notes, 1)
	fmt.Println("Notes: ", notes)
	fmt.Println("Size of the slice: ", len(notes))
	notes = append(notes, 33)
	fmt.Println("Size of the slice: ", len(notes))

	// Other way to declare a slice
	notes2 := []int{}
	fmt.Println("Notes2: ", notes2)

	// Read user input
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Your name is: ", name)
	fmt.Println("Memory address of name: ", &name)

	index := 0
	for {
		if index == 3 {
			break
		}
		fmt.Println("Loop: ", index)
		index++
	}

	var books = []string{"Clean code", "Effective Java"}
	// Foreach
	for index, book := range books {
		fmt.Printf("Book[%v]: %v\n", index, book)
		fmt.Println("First word of the book: ", strings.Fields(book)[0]) // split the string by space
	}

	fmt.Println(append(books, "Clean Architecture"))
}
