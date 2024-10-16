package main

import "fmt"

// func printname() {
// 	fmt.Println("Hello, World", 7+10)
// }

func main() {
	var confrenceName = "Go Conference"
	const confrenceTickets = 50
	var remainingTickets uint = 50
	// define array
	// var bookings [50]string

	// slice
	var bookings []string
	//print type of variables
	// fmt.Printf("confrenceTickets is %T ,remainingTickets is %T, confrenceName is %T\n", confrenceTickets, remainingTickets, confrenceName)
	fmt.Printf("Welcome to our %v booking application\n", confrenceName)
	fmt.Println("We have total of", confrenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	//data types
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// take user input
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email")
	fmt.Scan(&email)
	fmt.Println("How many Ticket You want to book")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName

	//put element in slice
	bookings = append(bookings, firstName+" "+lastName)
	// array opernation
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)
}
