# TicketBookingWithGo

package main

import (
	"fmt"
)

func main() {
	// create a map to store available tickets
	tickets := make(map[string]int)
	tickets["Concert"] = 10
	tickets["Movie"] = 20
	tickets["Play"] = 15

	// display available tickets
	fmt.Println("Available tickets:")
	for event, count := range tickets {
		fmt.Printf("%s: %d\n", event, count)
	}

	// prompt user to select an event and number of tickets
	var selectedEvent string
	var numTickets int
	fmt.Println("Select an event to book tickets:")
	fmt.Scanln(&selectedEvent)
	fmt.Println("Enter number of tickets to book:")
	fmt.Scanln(&numTickets)

	// check if selected event and number of tickets are valid
	if count, ok := tickets[selectedEvent]; ok {
		if count >= numTickets {
			// book tickets and update available tickets count
			tickets[selectedEvent] = count - numTickets
			fmt.Printf("Successfully booked %d tickets for %s.\n", numTickets, selectedEvent)
			fmt.Printf("Available tickets for %s: %d\n", selectedEvent, tickets[selectedEvent])
		} else {
			fmt.Printf("Sorry, only %d tickets are available for %s.\n", count, selectedEvent)
		}
	} else {
		fmt.Println("Invalid event selected.")
	}
}
