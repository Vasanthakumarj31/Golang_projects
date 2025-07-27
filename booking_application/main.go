package main

import (
	"fmt"
	"sync"
)

// package variables
const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var booking []UserData

// it is used to save list of string values
type UserData struct {
	firstName  string
	secondName string
	email      string
	Tickets    uint
}

var wg = sync.WaitGroup{}

func main() {
	//: is known as syntatic sugar
	booking = make([]UserData, 0)
	greetUsers()

	firstName, secondName, email, Tickets := getUserInput()
	isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, secondName, email, Tickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {

		firstNames := bookTicket(Tickets, firstName, secondName, email)
		wg.Add(1)
		go sendUserTicket(firstName, secondName, Tickets, email) //the thread is happened by the goroutines
		fmt.Printf("the first name of bookings are %v\n", firstNames)
		if remainingTickets == 0 {
			fmt.Printf("our conference is breaked up,come back next year\n")

		}
	} else {
		if !isValidName {
			fmt.Println("your first or second name is too short")
		}
		if !isValidEmail {
			fmt.Println("your email Id must contain @ sign")
		}

		if !isValidTickets {
			fmt.Println("Number of ticket is invalid")
		}

	}

	wg.Wait()
}
