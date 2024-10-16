package main

import "fmt"

// func printname() {
// 	fmt.Println("Hello, World", 7+10)
// }

func main() {
	var confrenceName = "Go Conference"
	const confrenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to our %v booking application\n", confrenceName)
	fmt.Println("We have total of", confrenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}
