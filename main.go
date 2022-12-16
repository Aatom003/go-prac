package main

import (
	"fmt"
	"strings"
)

var remainingTickets int = 50

const conferenceNumber uint = 50

func main() {
	var conferenceName = "Go Conference"
	for {

		fmt.Printf("Welcome to %v booking app!!!\n", conferenceName)
		fmt.Printf("We have total of %v tickets and %v tickets remaining \n", conferenceNumber, remainingTickets)
		fmt.Println("Get your tickets to attend the event")

		var bookingsNames []string
		firstNames := make([]string, 0)

		var firstName string
		var lastName string
		var email string
		var userTickets int
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your Email Address")
		fmt.Scan(&email)
		fmt.Println("How many ticket you want?")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry only %v tickets left!! \n\n\n", remainingTickets)
			continue
		} else {
			fmt.Printf("User %v booked %v tickets!!\n", firstName, userTickets)
			fmt.Printf("Thank you %v  %v for booking %v tickets. You'll soon receive a message on %v very soon\n\n", firstName, lastName, userTickets, email)
			bookingsNames = append(bookingsNames, firstName+""+lastName)
			for _, booking := range bookingsNames {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all first names of the booking%v \n\n", firstNames)
			remainingTickets = remainingTickets - userTickets
			noTicketsLeft := remainingTickets == 0
			if noTicketsLeft {
				fmt.Println("Sorry we are full booked. Come back next year.")
				break
			}
		}

	}
}
