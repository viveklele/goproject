package main

import (
    "fmt"
    "sync"
    "time"
)

var conferenceName = "Go Congerence"

const conferenceTicket = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
    firstNames      string
    lastName        string
    email           string
    numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

    greetUser()

    firstName, lastName, email, userTickets := getUserInput()
    isValidName, isValidEmail, isvalidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

    if isValidName && isValidEmail && isvalidTicketNumber {

        bookTicket(userTickets, firstName, lastName, email)

        wg.Add(1)
        go sendTicket(userTickets, firstName, lastName, email)

        firstNames := getFirstName()
        fmt.Printf("First names of bookings are: %v\n", firstNames)

        if remainingTickets == 0 {

            fmt.Println("Our conference is booked out. Come back next year.")
        }
    } else {
        if !isValidName {
            fmt.Println("First Name or last name you entered is short")
        }
        if !isValidEmail {
            fmt.Println("email address you entered dos't contain @ sign")
        }
        if !isvalidTicketNumber {
            fmt.Println("number of tickets you entered is invalid")
            fmt.Println("Enter number of tickets: ")
        }
    }
    wg.Wait()
}

func greetUser() {

    fmt.Printf("Welcome to %v bookings application\n", conferenceName)
    fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
    fmt.Println("Get your ticket here to attend")
}

func getFirstName() []string {

    firstNames := []string{}
    for _, booking := range bookings {
        firstNames = append(firstNames, booking.firstNames)
    }
    return firstNames
}

func getUserInput() (string, string, string, uint) {
    var firstName string
    var lastName string
    var email string
    var userTickets uint

    fmt.Println("enter your first name: ")
    fmt.Scan(&firstName)

    fmt.Println("enter your last name: ")
    fmt.Scan(&lastName)

    fmt.Println("enter your email address: ")
    fmt.Scan(&email)

    fmt.Println("Enter number of tickets: ")
    fmt.Scan(&userTickets)
    return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
    remainingTickets = remainingTickets - userTickets

    var userData = UserData{
        firstNames:      firstName,
        lastName:        lastName,
        email:           email,
        numberOfTickets: userTickets,
    }

    bookings = append(bookings, userData)
    fmt.Printf("List of booking is %v\n", bookings)

    fmt.Printf("Thenk you %v %v for bookings %v tickets.You will receiverd conformation email at %v\n", firstName, lastName, userTickets, email)
    fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
    time.Sleep(10 * time.Second)
    var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
    fmt.Println("##########################")
    fmt.Printf("Sending tickets:\n%vto email address %v\n", ticket, email)
    fmt.Println("##########################")
    wg.Done()
}
