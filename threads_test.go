package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"softott.net/gocon/events"
)

// Test Create Thread Pools

func TestCreateThreadPools(t *testing.T) {
	poolsize := 5
	Setup(poolsize)

	StartThreads() // Start the threads}

	if getPoolSize() != poolsize {
		t.Errorf("Expected thread pool size %d, got %d", poolsize, getPoolSize())
	}

	if getAddPool().Cap() != poolsize {
		t.Errorf("Expected Add thread pool size %d, got %d", poolsize, getAddPool().Cap())
	}

	if getAddPool().Running() <= poolsize {
		t.Errorf("Add Runnning thread count should be %d instead is %d", poolsize, getAddPool().Running())
	}

	if getDeletePool().Cap() != poolsize {
		t.Errorf("Expected Delete thread pool size %d, got %d", poolsize, getDeletePool().Cap())
	}

	if getDeletePool().Running() <= poolsize {
		t.Errorf("Delete Runnning thread count should be %d instead is %d", poolsize, getDeletePool().Running())
	}

	if getPersistancePool().Cap() != poolsize {
		t.Errorf("Expected Persistance thread pool size %d, got %d", poolsize, getPersistancePool().Cap())
	}

	fmt.Println(getPoolStistics())

	time.Sleep(2 * time.Second)

	StopThreads()

}

// this tests the adding and reading from the channel

func TestTheAddChannelEvent(t *testing.T) {
	Setup(5)
	StartThreads()

	id := GetRandomId()
	event := events.NewEvent(id, "Test Event 1")

	getAddChannel() <- event

	ea := <-getAddChannel()

	if ea.GetId() != id {
		t.Errorf("Channel Not Empty Expecting %d got %d", id, ea.GetId())
	}

	StopThreads()
}

// Test adding an event  multiple times
func TestThreadAddDuplicateEvent(t *testing.T) {

	createevents := 5

	Setup(5)

	StartThreads()

	aeevents := events.NewEvents(GetRandomId(), "Test Events", "Thetype")

	aeevents.GetSystem()

	mychannel := getAddChannel()

	for i := range createevents {
		anevent := events.NewEvent(GetRandomId(), "Test Event"+strconv.Itoa(i))
		anevent.SetOptr(aeevents) // Set the owner to events

		// Add Event to Events
		mychannel <- anevent
		time.Sleep(1 * time.Second)
	}
	// Used to all ow time for the event to be processed
	time.Sleep(10 * time.Second)

	// verify only one event exists
	if aeevents.GetEventsCnt() != createevents {
		t.Errorf("Expected %d event(s) in Events after adding duplicate, got %d\n", createevents, aeevents.GetEventsCnt())
	}
	StopThreads()
}

func TestTheDeleteChannelEvent(t *testing.T) {
	Setup(5)
	StartThreads()

	devents := events.NewEvents(GetRandomId(), "Test Events", "Thetype")

	id := GetRandomId()
	event := events.NewEvent(id, "Test Event to Delete")

	devents.AddEvent(event)

}
func TestAdding2ToChannel(t *testing.T) {
	Setup(5)

	StartThreads()

	aeevents := events.NewEvents(GetRandomId(), "Test Events", "Thetype")

	anevent := events.NewEvent(GetRandomId(), "Test Event 1")
	anevent.SetOptr(aeevents) // Set the owner to events

	mychannel := getAddChannel()
	mychannel <- anevent

	anevent2 := events.NewEvent(GetRandomId(), "Test Event 2")
	anevent2.SetOptr(aeevents) // Set the owner to events
	mychannel <- anevent2

	time.Sleep(10 * time.Second)

	// Used to all ow time for the event to be processed

	StopThreads()
}

func TestHashUUIDto32bit(t *testing.T) {

	myhash := GetRandomId()

	fmt.Printf("Returned a 32 Bit hashed version of the UUID:%d/n", myhash)

}
