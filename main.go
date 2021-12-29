package main

import (
	"fmt"
	"sync"
)


const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets = conferenceTickets
var bookings = make([]UserData, 0)
type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}
var wg = sync.WaitGroup{}
func main() {
	greetUsers()
	userInput()
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and still %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func userInput() {
	for {
		// get input fromn user
		fName, lName, email, userTickets := getUserInput()
		// validate user inputs
		isValidName, isValidEmail, isValidTicket := validateUserInput(fName, lName, email, userTickets)
		if isValidName && isValidEmail && isValidTicket {
			// book a ticket
			bookTicket(userTickets, fName, lName, email)
			wg.Add(1)
			go sendTicket(userTickets, fName, lName, email)
			firstNames := getFirstName()
			fmt.Printf("These are all our bookings: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println(" ****** first name or last name you entered is too short ****** ")
			}
			if !isValidEmail {
				fmt.Println(" ****** email address you entered does not contain @ sign ****** ")
			}
			if !isValidTicket {
				fmt.Println(" ****** number of tickets you entered is invalid ****** ")
			}
			fmt.Println(" ****** Please try again ****** ")
		}
	}
}

func getUserInput() (string, string, string, uint) {
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
	return fName, lName, email, userTickets
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(userTickets uint, fName string, lName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName: fName,
		lastName: lName,
		email: email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", fName, lName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, fName string, lName string, email string){
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, fName, lName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}