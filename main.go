package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Println("Welcome to the", conferenceName, "application")
	fmt.Println("We have total", conferenceTickets, "but only", remainingTickets, "are remaining!")

	for {
		var firstName string
		var lastName string
		var userTickets uint
		var userEmail string
		var userPhone int
		//ask username
		fmt.Println("Enter you first name please: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter you last name please: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter number of tickets you want to book: ")
		fmt.Scan(&userTickets)
		fmt.Println("Enter your mobile number: ")
		fmt.Scan(&userPhone)
		fmt.Println("Enter your email id: ")
		fmt.Scan(&userEmail)

		//To check if user enters correct values or not(user input validation)
		var isValidName = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail = strings.Contains(userEmail, "@")
		var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			//booking logic
			remainingTickets = remainingTickets - userTickets

			//arrays
			// var bookings [50]string
			// bookings[0]= firstName+" "+lastName

			//dynamic allocation
			//slices
			var bookings []string
			bookings = append(bookings, firstName+" "+lastName)
			//retreiving information from slice is same as array
			// fmt.Printf("The whole array:%v\n",bookings )
			// fmt.Printf("The first value in array:%v\n",bookings[0] )
			// fmt.Printf("Array type:%T\n",bookings )
			// //T for type of any variable or array
			// fmt.Printf("array length:%v\n",len(bookings) )

			fmt.Printf("Thanks %v %v for booking %v tickets with us! You will receive your tickets on %v and %v.\n", firstName, lastName, userTickets, userPhone, userEmail)
			fmt.Printf("Hurry! Tell your friends that only %v tickets are remaining.\n", remainingTickets)

			//loops
			//no multiple loops in GO
			// firstNames= []string{}
			// for index, bookings:=range bookings{
			// 	var names= strings.Fileds(booking)

			// }
			fmt.Printf("These are all the bookings: %v\n", bookings)

			//conditionals
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out!")
				break
			}
		} else {
			fmt.Printf("Try again with valid information\n", remainingTickets, userTickets)
			//break
			//using continue will skip below part but start for loop iteration from the beginning

		}

	}

}
