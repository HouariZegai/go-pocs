package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	var conferenceName = "Go conference"
	const totalTickets = 10
	var remainingTickets = 10

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("We have total of %v tickets and %v are still available\n", totalTickets, remainingTickets)
	fmt.Println(quote.Go())

}
