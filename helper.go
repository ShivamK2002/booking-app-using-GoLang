package main

import "strings"

func validateUserInput(firstName string, lastName string, userEmail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(userEmail, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

//if we need to divide go file into multiple files we just need to name package and all the input parameters, However at time of running the code we need to execute whole folder rather than one single main.go file
