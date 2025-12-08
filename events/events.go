package events

import "sync"

type Events struct {
	id       int
	system   string
	atype    string
	myevents map[int]Event
	mylock   sync.RWMutex // Mutex for concurrent access to myevents map
}

func NewEvents(id int, systemname string, thetype string) *Events {
	return &Events{
		id:       id,
		system:   systemname,
		atype:    thetype,
		myevents: make(map[int]Event),
	}
}

func (e *Events) SetSystem(mysystem string) {

	e.system = mysystem

}

func (e *Events) GetSystem() string {

	return e.system
}

func (e *Events) GetEventsCnt() int {

	return len(e.myevents)
}

// locks the Events struct to allow for safe concurrent access
func (e *Events) Lock() {
	e.mylock.Lock()
}
func (e *Events) Unlock() {
	e.mylock.Unlock()
}

// AddEvent adds an event to the Events struct
//  1. If an event with the same ID already exists, it will still update the previous event with new data,
//  2. The event is added to the mnts map using its ID as the key.
func (e *Events) AddEvent(ae *Event) {

	if ae.GetOptr() == nil {
		ae.SetOptr(e)
	}

	e.myevents[ae.id] = *ae
}

func (e *Events) GetEvent(id int) Event {
	return e.myevents[id]
}

func (e *Events) RemoveEvent(eventID int) bool {

	if _, exists := e.myevents[eventID]; exists {
		delete(e.myevents, eventID)
		return true
	}

	return false
}
