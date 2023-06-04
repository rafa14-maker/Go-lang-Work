package main

import (
	"booking-app/helper"
	"fmt"
)

var UserFirstName string
var UserLastName string
var UserEmail string
var userTickets int
var bookings []string
var UserAge int

func main() {
	var conferenceName = "Go conference"
	var conferenceTickets = 50
	var remainingTickets = 50

	helper.Fun(conferenceName, conferenceTickets, remainingTickets)

	for {

	UserFirstNameFormat:

		fmt.Print("Enter user first Name :")
		fmt.Scan(&UserFirstName)
		fmt.Print("\n")

		if !helper.ValidateUserFirstName(UserFirstName) {
			println("Input corret formet of First Name")
			goto UserFirstNameFormat
		}

	UserLastNameFormat:

		fmt.Print("Enter user Last Name :")
		fmt.Scan(&UserLastName)
		fmt.Print("\n")

		if !helper.ValidateUserLastName(UserLastName) {
			println("Input corret formet of First Name")
			goto UserLastNameFormat
		}

	UserEmailFormat:

		fmt.Print("Enter user Email :")
		fmt.Scan(&UserEmail)
		fmt.Print("\n")

		if !helper.ValidateUserEmail(UserEmail) {
			println("Input corret formet of Email")
			goto UserEmailFormat
		}

	UserTicketsFormat:

		fmt.Print("Enter user Tickets :")
		fmt.Scan(&userTickets)
		fmt.Print("\n")

		if !helper.ValidateUserTickets(userTickets, remainingTickets) {
			println("Input corret formet of Tickets amount")
			goto UserTicketsFormat
		}

		fmt.Print("Enter your age :")
		fmt.Scan(&UserAge)
		fmt.Println("Total expense ", helper.GenarateTotalMoney(userTickets, UserAge))

		bookings = append(bookings, UserFirstName+" "+UserLastName)

		//firstNames := getFristNames(bookings)
		//fmt.Printf("the first names of bookings are : %v", firstNames)

		fmt.Printf("Thank you %v %v for booking %v tickets.", UserFirstName, UserLastName, userTickets)
		fmt.Printf("You will get a confirmation at %v. \n", UserEmail)
		fmt.Printf("the bookings %v\n", bookings)

		remainingTickets = remainingTickets - userTickets
		fmt.Println("Remaining tickets ", remainingTickets)

		if remainingTickets == 0 {
			println("All tickets has been booked")
			break
		}

	}

}
