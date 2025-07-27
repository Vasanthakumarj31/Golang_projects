package main

import (
	"fmt"
	"strings"
	"time"
	
)

func validateUserInput(firstName string, secondName string, email string, Tickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(secondName) >= 2
	isValidEmail := strings.Contains(email, "@gmail.com")
	isValidTickets := Tickets > 0 && Tickets <= uint(remainingTickets)
	return isValidName, isValidEmail, isValidTickets
}
func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")
}

/*
	func getFirstName(booking []map[string]string) []string {
		var firstNames []string
		for _, booking := range booking {
			firstname := booking["firstname"]
			firstNames = append(firstNames,firstname)
		}
		return firstNames
	}
*/
func getUserInput() (string, string, string, uint) {

	var firstName string
	var secondName string
	var Tickets uint
	var email string

	fmt.Printf("enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Printf("enter your last name : ")
	fmt.Scan(&secondName)

	fmt.Printf("enter the number of tickets : ")
	fmt.Scan(&Tickets)

	fmt.Printf("enter your email :")
	fmt.Scan(&email)

	return firstName, secondName, email, Tickets
}
func bookTicket(Tickets uint, firstName string, secondName string, email string) []string {
	userData := UserData{
		firstName:  firstName,
		secondName: secondName,
		email:      email,
		Tickets:    Tickets,
	}
	remainingTickets = remainingTickets - Tickets
	booking = append(booking, userData)
	fmt.Printf("List of booking %v\n", booking)
	fmt.Printf("Thank you %v %v for purchasing %v tickets. A verification email will be sent to %v\n", firstName, secondName, Tickets, email)
	fmt.Printf("%v tickets are remaining out of %v tickets\n", remainingTickets, conferenceTickets)

	var firstNames []string
	for _, Booking := range booking {

		firstNames = append(firstNames, Booking.firstName)
	}
	return firstNames
}

func sendUserTicket(firstName string, secondName string, Tickets uint, email string) {
	time.Sleep(10 * time.Second)
	var Ticket = fmt.Sprintf("%v tickets for %v %v", Tickets, firstName, secondName)

	fmt.Println("########################################################")
	fmt.Printf("sending tickets : %v \nemail address : %v\n", Ticket, email)
	fmt.Println("########################################################")
	wg.Done()
}
