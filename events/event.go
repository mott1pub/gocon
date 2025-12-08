package events

import "time"

type Event struct {
	id        int
	optr      *Events // This is the Owner Events Struct Pointer
	message   string
	timestamp time.Time
}

func NewEvent(aid int, amessage string) *Event {

	aevent := &Event{
		id:        aid,
		message:   amessage,
		timestamp: time.Now(),
	}

	return aevent
}

func (e *Event) GetId() int {
	return e.id
}

func (e *Event) SetOptr(aoid *Events) {
	e.optr = aoid
}

// This returns the owner Events Struct Pointer
func (e *Event) GetOptr() *Events {
	return e.optr
}

func (e *Event) SetMessage(amessage string) {

	e.message = amessage
}

func (e *Event) GetMessage() string {
	return e.message
}

func (e *Event) SetTimestamp() {
	e.timestamp = time.Now()
}
