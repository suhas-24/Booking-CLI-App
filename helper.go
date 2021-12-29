package main

import "strings"

func validateUserInput(fName string, lName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(fName) > 2 && len(lName) >= 1
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicket
}
