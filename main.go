package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	var eventName = "Go Conference"

	// eventName := "Go Conference"           //Syntatic Sugar -- this and previos line both are same
	const totalTickets = 50

	var remainingTickets = 50

	// var bookings[50]string    //for Array
	// var bookings = []string{} //for slice
	var bookings = make([]map[string]string, 0) //slice of map

	greetUsers(eventName, totalTickets, remainingTickets)

	for {

		firstName, lastName, emailID, numberOfTickets := getUserInput()

		isValidName, isValidMail, isValidTickets := validateUserInput(firstName, lastName, emailID, numberOfTickets, remainingTickets)

		if isValidName && isValidMail && isValidTickets {

			remainingTickets = remainingTickets - numberOfTickets

			fmt.Printf("Congratulations %v %v, Your %v Seats have been reserved for %v and a confirmation mail with e-tickets have been sent to %v \n", firstName, lastName, numberOfTickets, eventName, emailID)

			fmt.Printf("We have Now %v reamining tickets \n", remainingTickets)

			wg.Add(1)
			go sendTickets(firstName, lastName, emailID, numberOfTickets)

			// creating a map of UserData
			var UserData = make(map[string]string)
			UserData["FirstName"] = firstName
			UserData["LastName"] = lastName
			UserData["EmailID"] = emailID
			UserData["NumberofTickets"] = strconv.FormatInt(int64(numberOfTickets), 10)

			bookings = append(bookings, UserData)

			// fmt.Println("List of bookings is:", bookings)   // check the map

			printFirstNames(bookings)

			if remainingTickets == 0 {
				fmt.Println("Sorry!!! Tickets are all sold out for this event.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("FirstName or LastName is too Short!!Please Enter a Valid Name")
			}
			if !isValidMail {
				fmt.Println("Enter a vaid Email id!!")
			}
			if !isValidTickets {
				fmt.Printf("We have only %v tickets left, So you can't book %v tickets.\n", remainingTickets, numberOfTickets)
				fmt.Println("Please try again with a Valid number of tickets!!")
			}

		}
	}
	wg.Wait()
}

// encapsulating our logic in functions

func greetUsers(evtNam string, totTicks int, remTicks int) {
	fmt.Println("Welcome to our Nexus Booking App")
	fmt.Println("You can book your tickets for", evtNam, "from Here!!!!")
	fmt.Printf("There are total of %v seats and %v are available for booking.\n", totTicks, remTicks)
}

func printFirstNames(bookings []map[string]string) {

	firstNames := []string{} // Declaring an empty slice

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["FirstName"])
	}
	fmt.Printf("The first names of our bookings are: %v\n", firstNames)
}

//returning multiple values in a function

func validateUserInput(firstName string, lastName string, emailID string, numberOfTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidMail := strings.Contains(emailID, "@")
	isValidTickets := numberOfTickets > 0 && numberOfTickets <= remainingTickets

	return isValidName, isValidMail, isValidTickets
}

func getUserInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var emailID string
	var numberOfTickets int

	fmt.Println("Enter Your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email ID:")
	fmt.Scan(&emailID)

	fmt.Println("Number od tickets you want to book:")
	fmt.Scan(&numberOfTickets)

	return firstName, lastName, emailID, numberOfTickets
}

//Simulating the tickets sending process to understand Goroutines-Concurrency

func sendTickets(firstName string, lastName string, emailId string, numberofTickets int) {
	time.Sleep(50 * time.Second)
	var Ticket = fmt.Sprintf("%v tickets for %v %v", numberofTickets, firstName, lastName)
	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", Ticket, emailId)
	fmt.Println("-----------------------------------------------------------")
	wg.Done()
}
