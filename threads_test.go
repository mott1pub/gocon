package main

import (
	"fmt"
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

func TestTheDeleteChannelEvent(t *testing.T) {
	Setup(5)
	StartThreads()

	devents := events.NewEvents(GetRandomId(), "Test Events", "Thetype")

	id := GetRandomId()
	event := events.NewEvent(id, "Test Event to Delete")

	devents.AddEvent(event)

}

func TestHashUUIDto32bit(t *testing.T) {

	myhash := GetRandomId()

	fmt.Printf("Returned a 32 Bit hashed version of the UUID:%d/n", myhash)

}
