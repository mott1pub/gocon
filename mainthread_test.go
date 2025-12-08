package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"softott.net/gocon/events"
)

// Test Create Thread Pools

func TestThreadAddEvent(t *testing.T) {
	Setup(5)

	StartThreads()

	aeevents := events.NewEvents(GetRandomId(), "Test Events", "Thetype")

	aeevents.GetSystem()

	anevent := events.NewEvent(GetRandomId(), "Test Event 1")
	anevent.SetOptr(aeevents) // Set the owner to events

	// Add Event to Events

	mychannel := getAddChannel()
	mychannel <- anevent

	// Used to all ow time for the event to be processed
	time.Sleep(5 * time.Second)

	// Now try to add the same event again via preEvent
	fmt.Println("Adding duplicate event to test handling")

	StopThreads()
}

func TestThreadAddMultiEvent(t *testing.T) {
	Setup(5)

	StartThreads()

	aeevents := events.NewEvents(GetRandomId(), "Test Events", "Thetype")

	aeevents.GetSystem()

	myid := GetRandomId()
	anevent := events.NewEvent(myid, "Test Event"+strconv.Itoa(myid))
	anevent.SetOptr(aeevents) // Set the owner to events

	// Add Event to Events

	mychannel := getAddChannel()
	mychannel <- anevent

	myid = GetRandomId()
	anevent = events.NewEvent(myid, "Test Event"+strconv.Itoa(myid))
	anevent.SetOptr(aeevents) // Set the owner to events

	// Add Event to Events

	mychannel <- anevent

	// Used to all ow time for the event to be processed
	time.Sleep(60 * time.Second)

	// verify new event as	added..
	if aeevents.GetEventsCnt() != 2 {
		t.Errorf("Expected 2 event in Events after adding 2, got %d\n", aeevents.GetEventsCnt())
	}

	// Now try to add the same event again via preEvent
	fmt.Println("Adding duplicate event to test handling")

	StopThreads()
}
