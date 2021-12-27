package main

import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	remainingTickets := conferenceTickets
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and still %v are available\n",conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var fName string
	var lName string
	var email string
	var userTickets uint
	fmt.Print("Enter your first name:")
	fmt.Scan(&fName)
	fmt.Print("Enter your last name:")
	fmt.Scan(&lName)
	fmt.Print("Enter your email address:")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets:")
	fmt.Scan(&userTickets)
	remainingTickets= remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n",fName, lName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
