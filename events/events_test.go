package events

import (
	"testing"
)

func TestAddCreateEvent(t *testing.T) {
	var myevents = *NewEvents(1, "systemA", "typeX")

	if myevents.GetSystem() != "systemA" {
		t.Errorf("Expected system name to be 'systemA', got '%s'", myevents.GetSystem())

	}

}

func TestAddEventToEvents(t *testing.T) {

	hashedInit := 1

	var myevents = NewEvents(hashedInit, "systemA", "typeX")

	var myevent = NewEvent(hashedInit, "This is a test event")

	myevents.AddEvent(myevent)

	// verify new event as	added..
	if myevents.GetEventsCnt() != 1 {
		t.Errorf("Expected 1 event in Events, got %d", myevents.GetEventsCnt())
	}

	if myevents.myevents[1].id != 1 {
		t.Errorf("Expected event descriptio id to be 'This is a test event', got '%d'", myevents.myevents[1].id)
	}

}

func TestGetEvent(t *testing.T) {

	hashedInit := 1

	var myevents = NewEvents(hashedInit, "systemA", "typeX")

	var myevent = NewEvent(hashedInit, "This is a test event")

	myevents.AddEvent(myevent)

	myreturnevent := myevents.GetEvent(hashedInit)

	if myreturnevent.id != 1 {
		t.Errorf("Returned Wrong Event Expected %d got %d/n!", myreturnevent.id, myreturnevent.id)
	}

}

// Duplicate Kyes overwrite the existing Key/Values
func TestAddDuplicateEvent(t *testing.T) {

	var myevents = NewEvents(1, "systemA", "typeX")

	var myevent1 = NewEvent(1, "This is a test event")

	myevents.AddEvent(myevent1)

	var myevent2 = NewEvent(1, "This is a duplicate test event")

	myevents.AddEvent(myevent2)

	// verify new event as	added..
	if myevents.GetEventsCnt() != 1 {
		t.Errorf("Expected 2 event in Events after adding duplicate, got %d\n", myevents.GetEventsCnt())
		t.Error("Show Map:", myevents.myevents)
	}

}

func TestCountEvents(t *testing.T) {

	var myevents = NewEvents(1, "systemA", "typeX")

	var myevent = NewEvent(1, "This is a test event")

	myevents.AddEvent(myevent)

	// verify new event as	added..
	if myevents.GetEventsCnt() != 1 {
		t.Errorf("Expected 1 event in Events, got %d", myevents.GetEventsCnt())
	}

}

func TestRemoveEvent(t *testing.T) {

	var myevents = NewEvents(1, "systemA", "typeX")

	var myevent = NewEvent(1, "This is a test event")

	myevents.AddEvent(myevent)

	isRemoved := myevents.RemoveEvent(1)

	if isRemoved == false {
		t.Errorf("Expected event to be removed successfully")
	}

	if myevents.GetEventsCnt() != 0 {
		t.Errorf("Expected 0 events in Events after removal, got %d", myevents.GetEventsCnt())
	}
}

func TestRemoveNonExistentEvent(t *testing.T) {

	var myevents = NewEvents(1, "systemA", "typeX")

	var myevent = NewEvent(1, "This is a test event")

	myevents.AddEvent(myevent)

	removed := myevents.RemoveEvent(2) // trying to remove non-existent event

	if removed {
		t.Errorf("Expected removal of non-existent event to fail")
	}

	if myevents.GetEventsCnt() != 1 {
		t.Errorf("Expected 1 event in Events after failed removal, got %d", myevents.GetEventsCnt())
	}
}
