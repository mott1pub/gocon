package main

import (
	"fmt"
	"strings"

	"softott.net/gocon/events"

	"github.com/panjf2000/ants"
)

var threadPool ThreadPool

// ThreadPool
type ThreadPool struct {
	poolSize   int
	addThreads *ants.Pool
	addChannel chan *events.Event

	deleteThreads *ants.Pool
	deleteChannel chan *events.Event
	//ventsThreads *ants.Pool
	persistanceThreads *ants.Pool
	persistanceChannel chan *events.Event
}

func setPoolSize(size int) {
	threadPool.poolSize = size
}
func setAddThreads(threads *ants.Pool) {
	threadPool.addThreads = threads
}
func setAddChannels(channels chan *events.Event) {
	threadPool.addChannel = channels
}
func setDeleteThreads(threads *ants.Pool) {
	threadPool.deleteThreads = threads
}
func setDeleteChannels(channels chan *events.Event) {
	threadPool.deleteChannel = channels
}

// func EventsThreads(threads []Thread) {
// threadPool.EventsThreads = threads
// }
func setPersistanceThreads(threads *ants.Pool) {
	threadPool.persistanceThreads = threads
}
func setPersistanceChannel(channels chan *events.Event) {
	threadPool.persistanceChannel = channels
}

func getPoolSize() int {
	return threadPool.poolSize
}

func getAddPool() *ants.Pool {
	return threadPool.addThreads
}

func getAddChannel() chan *events.Event {
	fmt.Printf("display Thread %p\n", &threadPool)
	return threadPool.addChannel
}

func getDeletePool() *ants.Pool {
	return threadPool.deleteThreads
}
func getPersistancePool() *ants.Pool {
	return threadPool.persistanceThreads
}

func getPoolStistics() string {

	mystring :=
		fmt.Sprintf(
			"Size of Add Pool: %d\nDelete Pool: %d\nPersistance Pool: %d\n",
			threadPool.addThreads.Cap(), threadPool.deleteThreads.Cap(), threadPool.persistanceThreads.Cap(),
		)

	mystring2 := fmt.Sprintf(
		"Add Pool Running: %d\nDelete Pool Running: %d\nPersistance Pool Running: %d\n",
		threadPool.addThreads.Running(), threadPool.deleteThreads.Running(), threadPool.persistanceThreads.Running(),
	)
	return mystring + mystring2
}

// Setup initializes the thread pool for event processing

func Setup(poolsize int) {
	fmt.Printf("Setting up threads...%p/n", &threadPool)
	setPoolSize(poolsize)
}

// StartThreads starts the threads for event processing
func StartThreads() {
	fmt.Println("*** Threads Starting!***")
	StartStopthreadMaintence(true)

}

// 1. Check that at leaset one thread is running for adding events
// 2. Check that at least one thread is running for delete events
// ??3. Check that at least one thread is running for Event Persistance
// 4. Check that at least one thread is running for Event Processing
func StartStopthreadMaintence(startupShutdown bool) {
	if startupShutdown {
		// Create the add Thread Pool
		setAddThreads(createThreadPool())
		// Add the AddElements function to the add Thread Pool
		setAddChannels(createChannelPool(threadPool.poolSize))
		submitAddEventstoWorkerPool()

		// Create the delete Thread Pool
		setDeleteThreads(createThreadPool())
		setDeleteChannels(createChannelPool(threadPool.poolSize))
		// create the Events Thread Pool
		// threadPool.EventsThreads := createThreadPool()
		// Create the Persistance Thread Pool
		setPersistanceThreads(createThreadPool())
		setPersistanceChannel(createChannelPool(threadPool.poolSize))
		fmt.Printf("Threads Started Successfully! %d\n", threadPool.poolSize)
	} else {
		// Stop all threads in the thread pools

		closeChannelPool(threadPool.addChannel)
		closeChannelPool(threadPool.deleteChannel)
		closeChannelPool(threadPool.persistanceChannel)
		releasePool(threadPool.addThreads)
		releasePool(threadPool.deleteThreads)
		releasePool(threadPool.persistanceThreads)
		// releasePool(threadPool.EventsThreads)
		fmt.Println("Channels/Threads Stopped Successfully!")
	}
}

// StopThreads stops all the threads for event processing
//  1. this includes:
//     a. Stopping threads for adding events
//     b. Stopping threads for deleting events
func StopThreads() {
	println("*** Threads Starting!***")
	StartStopthreadMaintence(false)
	println("***  Threads Ending!***")
}

// createThreadPool creates a thread pool for event processing
func createThreadPool() *ants.Pool {
	p, _ := ants.NewPool(getPoolSize())
	//defer p.Release() //
	return p
}
func createChannelPool(size int) chan *events.Event {
	return make(chan *events.Event, size)
}
func closeChannelPool(channel chan *events.Event) {
	close(channel)
	channel = nil
}

func releasePool(p *ants.Pool) {
	p.Release()
	p = nil
}

// 1. submitAddEventstoWorkerPool submits AddEvents to the worker pool
func submitAddEventstoWorkerPool() {
	for i := 0; i < getPoolSize(); i++ {
		// Submit a task closure that invokes AddEvent with the channel
		threadPool.addThreads.Submit(func() { AddEvent(threadPool.addChannel) })
	}
}

//  1. AddEvent handles Add a new event to the System and includes:
//     a. call preAddEvents
//     b. add event to storage
//     c. call postAddEvents
func AddEvent(channel chan *events.Event) {

	addEvent := <-channel

	status := preEvent(addEvent)

	// if preEvent is not successful, log and return
	//  1. log pre event failure can happen with the following:
	//   a. event validation failure
	//   b. event already exists
	//   c. system is not ready to accept new events (Already lockeced by another process)
	if !status {
		fmt.Println("item for :", addEvent.GetMessage())
		return
	}

	// Add event to storage
	aeevents := addEvent.GetOptr()

	// lock the Events struct for safe concurrent access
	aeevents.Lock()
	defer aeevents.Unlock() // Ensure unlock after adding event
	aeevents.AddEvent(addEvent)

	postEvents(addEvent)

}

//  1. DeleteEvent handles deleting events from the system and includes:
//     a. call preDeleteEvents
//     b. delete event from storage
//     c. call postDeleteEvents
func DeleteEvent(channel chan *events.Event) {
	event := <-channel
	preEvent(event)

	// Delete event from storage)
	postEvents(event)
}

// preEvent handles actions before deleting an event
//  1. pre Delete Events include:
//     a. Check if event exists
//     b. Validate event can be deleted
//     c. Log pre delete event actions
//  2. Lock the Event to be deleted and includes:
//     a. Acquire lock on event
//     b. Ensure no other operations are being performed on the event
//     *Note:
func preEvent(event *events.Event) bool {
	// Placeholder for post event processing logic

	// Check if event is valid
	if event.GetId() <= 0 {
		mysysmsg := fmt.Sprintf("Invalid Event ID:%d", event.GetId())
		event.SetMessage(mysysmsg + event.GetMessage())
		return false
	}
	// is blank..
	if strings.TrimSpace(event.GetMessage()) == "" {
		mysysmsg := fmt.Sprintf("Invalid Event ID:%d", event.GetId())
		event.SetMessage(mysysmsg + event.GetMessage())
		return false
	}

	// Is this Event already in the system
	if event.GetOptr() == nil {
		mysysmsg := fmt.Sprintf("Event ID:%p No Parent Events", event.GetOptr())
		event.SetMessage(mysysmsg + event.GetMessage())
		return false
	}

	// Additional pre-event processing can be added here

	return true
}

// 1.
func postEvents(event *events.Event) {
	// Placeholder for post event logic
}
