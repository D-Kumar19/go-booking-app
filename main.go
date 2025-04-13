package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

const conferenceTickets uint = 50

func main() {
	// conferenceName := "Go Conference"
	// remainingTickets := 50
	// const conferenceTickets uint = 50

	greetUser()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := helper.GetUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("Email address is not valid. Please try again.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid. Please try again.")
			}
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Println("Welcome to the Booking App!")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets for", conferenceName, "here!")

	// fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	// fmt.Printf("Get your tickets for %v here!\n", conferenceName)
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Println("List of bookings:", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at your email address (%v).\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("Booked %v tickets for %v %v for %v!", userTickets, firstName, lastName, conferenceName)
	fmt.Println("\n########################")
	fmt.Printf("Sending ticket:\n%v\nto email address%v!\n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}
