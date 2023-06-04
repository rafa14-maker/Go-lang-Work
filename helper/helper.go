package helper

import (
	"fmt"
	"strings"
)

/*
func ValidateUserInput(UserFirstName string, UserLastName string, UserEmail string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(UserFirstName) >= 2 && len(UserLastName) >= 2
	isValidEmail := strings.Contains(UserEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
*/

func Fun(conferenceName string, conferenceTickets int, remainingTickets int) {
	println("Hello")
	fmt.Println(conferenceName)
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Println("we have total of", conferenceTickets, "available are", remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func GetFristNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func ValidateUserFirstName(Name string) bool {
	for i := 0; i < len(Name); i++ {
		if Name[i] <= '9' && Name[i] >= '0' {
			return false
		}
	}
	return true
}

func ValidateUserLastName(Name string) bool {
	for i := 0; i < len(Name); i++ {
		if Name[i] <= '9' && Name[i] >= '0' {
			return false
		}
	}
	return true
}

func ValidateUserTickets(ustickets int, rmtickets int) bool {
	if ustickets <= rmtickets {
		return true
	}
	return false
}

func GenarateTotalMoney(ticket int, age int) float64 {
	var expense float64

	if age < 18 {

		expense = float64(ticket * 500)
		expense = expense - (expense * 0.25)

	} else {
		expense = float64(ticket * 500)
	}

	return expense

}

func ValidateUserEmail(email string) bool {

	var beforeAt string
	var beforeUrl string
	var urL string

	var site_Url = []string{"com", "edu", "org"}
	var site_email = []string{"gmail", "yahoo", "rocket"}

	var cnt = 0

	for i := 0; i < len(email); i++ {
		if cnt == 0 && email[i] != '@' {
			beforeAt += string(email[i])
		}
		if cnt == 1 && email[i] != '.' {
			beforeUrl += string(email[i])
		}
		if cnt == 2 {
			urL += string(email[i])
		}
		if email[i] == '@' {
			if len(beforeAt) == 0 {
				return false
			}
			cnt = 1
		}
		if email[i] == '.' {
			if cnt == 1 {

				if len(beforeUrl) == 0 {
					return false
				}

				cnt = 2
			}
		}

	}

	//println(beforeAt)
	//println(beforeUrl)
	//println(urL)

	if beforeAt[0] <= 'a' && beforeAt[0] >= 'z' {
		return false
	}

	var tr = false

	for i := 0; i < len(site_email); i++ {
		if site_email[i] == beforeUrl {
			tr = true
		}
	}

	if !tr {
		return false
	}

	tr = false

	for i := 0; i < len(site_Url); i++ {
		if site_Url[i] == urL {
			tr = true
		}
	}

	if !tr {
		return false
	}

	return true

}
