package main

import "fmt"

// func printname() {
// 	fmt.Println("Hello, World", 7+10)
// }

func main() {
	var confrenceName = "Go Conference"
	const confrenceTickets = 50
	var remainingTickets = 50

	//print type of variables
	// fmt.Printf("confrenceTickets is %T ,remainingTickets is %T, confrenceName is %T\n", confrenceTickets, remainingTickets, confrenceName)
	fmt.Printf("Welcome to our %v booking application\n", confrenceName)
	fmt.Println("We have total of", confrenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	//data types
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email")
	fmt.Scan(&email)
	fmt.Println("How many Ticket You want to book")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v has booked %v tickets.\n", firstName, userTickets)
}
