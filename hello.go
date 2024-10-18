package main

import (
	"example/hello/helper"
	"fmt"
	"strconv"
)

const confrenceTickets = 50

var firstName string
var lastName string
var email string
var userTickets uint
var confrenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {
		// take user input
		firstName, lastName, email, userTickets := getUserInput()
		//validate the user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		// check user enter more ticket than remaining ticket
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName

			// create a map for user to store the value
			var userData = make(map[string]string)
			userData["firstName"] = firstName
			userData["lastName"] = lastName
			userData["email"] = email
			userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

			//put element in slice

			bookings = append(bookings, userData)
			fmt.Printf("List of booking is %v\n", bookings)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

			// call function print first names
			var firstNames = getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our confrence is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", confrenceName)
	fmt.Println("We have total of", confrenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	//data types
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email")
	fmt.Scan(&email)
	fmt.Println("How many Ticket You want to book")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
