package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter your name [%v]: ", i)
		var name string
		fmt.Scan(&name)

		wg.Add(1)
		go sendTicket(name, fmt.Sprintf("%v@houarizegai.com", name))
	}

	wg.Wait()
}

func sendTicket(name string, email string) {
	time.Sleep(1 * time.Second)
	var ticket = fmt.Sprintf("ticket to the user %v", name)
	fmt.Println("###########")
	fmt.Printf("We send %v to this email: %v\n", ticket, email)
	fmt.Println("###########")
	wg.Done()
}
