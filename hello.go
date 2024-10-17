package main

import (
	"fmt"
	"strings"
)

func main() {
	var confrenceName = "Go Conference"
	const confrenceTickets = 50
	var remainingTickets uint = 50
	// define array
	// var bookings [50]string

	// slice
	var bookings []string

	fmt.Printf("Welcome to our %v booking application\n", confrenceName)
	fmt.Println("We have total of", confrenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	//data types
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		// take user input
		greetUsers()
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your Email")
		fmt.Scan(&email)
		fmt.Println("How many Ticket You want to book")
		fmt.Scan(&userTickets)

		//validate the user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		// check user enter more ticket than remaining ticket
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName

			//put element in slice
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

			// Logic to Extract the first name
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
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
	fmt.Println("Welcome to our confrence")
}
