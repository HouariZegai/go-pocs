package main

import (
	"fmt"
	"strings"
)

// Package level variable
const appName = "GoLang Basics Application"

var tickets = 10

func main() {
	fmt.Println("\nCall a function")
	greeting()

	fmt.Println("\nFor loop")
	for i := 0; i < 3; i++ {
		fmt.Println("For loop", i)
	}

	fmt.Println("\nWhile loop")
	i := 0
	for i < 2 {
		fmt.Println("While loop", i)
		i++
	}

	fmt.Println("\nData types and pointer")
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

	fmt.Println("\nSlices")
	var notes []int
	notes = append(notes, 1)
	fmt.Println("Notes: ", notes)
	fmt.Println("Size of the slice: ", len(notes))
	notes = append(notes, 33)
	fmt.Println("Size of the slice: ", len(notes))

	fmt.Println("\nOther way to declare a slice")
	notes2 := []int{}
	fmt.Println("Notes2: ", notes2)

	fmt.Println("\nRead user input")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Your name is: ", name)
	fmt.Println("Memory address of name: ", &name)

	fmt.Println("\nIf Condition")
	num := 1
	if num < 1 {
		fmt.Println("Num is less than 1")
	} else if num == 1 {
		fmt.Println("Num is equal to 1")
	} else {
		fmt.Println("Num is greater than 1")
	}

	index := 0
	for {
		if index == 3 {
			break
		}
		fmt.Println("Loop: ", index)
		index++
	}

	var books = []string{"Clean code", "Effective Java"}

	fmt.Println("\nFor Each")
	for index, book := range books {
		fmt.Printf("Book[%v]: %v\n", index, book)
		fmt.Println("First word of the book: ", strings.Fields(book)[0]) // split the string by space
	}

	fmt.Println(append(books, "Clean Architecture"))

	if len(books) > 0 {
		fmt.Println("We have some books!")
	} else {
		fmt.Println("We don't have any books!")
	}

	if len(books) == 2 {
		fmt.Println("We have two books!")
	}

	fmt.Println("\nLoop with validation")
	for {
		fmt.Print("Enter your name (Mohamed): ")
		fmt.Scan(&name)
		if name != "Mohamed" {
			continue
		}
		fmt.Println("Welcome Mohamed!")
		break
	}

	for {
		var email string
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		validEmail, alwaysTrue := validEmail(email)
		if !validEmail {
			continue
		}
		fmt.Println("Always true: ", alwaysTrue)
		fmt.Println("Thanks for providing a valid email!")
		break
	}

	fmt.Println("\nWhile loop")
	counter := 0
	for counter < 2 && true {
		fmt.Println("Counter: ", counter)
		counter++
	}

	fmt.Println("\nSwitch Statement")
	city := "Algiers"
	switch city {
	case "Oran", "Tiaret", "Tlemcen":
		fmt.Println("City is located in the West of Algeria")
	case "Algiers", "Blida":
		fmt.Println("City is located in the Center of Algeria")
	default:
		fmt.Println("City is not located in West nor Center of Algeria")
	}
}

func greeting() {
	fmt.Println("Welcome to our", appName)
}
