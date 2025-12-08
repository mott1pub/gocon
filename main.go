package main

import (
	"fmt"

	"softott.net/gocon/events"
)

func main() {
	var myevents = events.NewEvents(1, "MySystem", "Error")

	myevents.SetSystem("theSystem")

	fmt.Println("Events Struct:", myevents.GetSystem())

	var myevent = events.NewEvent(1, "This is an event message")

	myevents.AddEvent(myevent)

	fmt.Println("Event added to Events struct")
}
