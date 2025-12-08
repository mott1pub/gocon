
// Ants - A go thread pool management Routes
// https://erfansahaf.medium.com/managing-goroutines-with-gouroutine-pooling-in-go-9b3596e23225

1. Create Events
	Needs method for Locking Map before Add and Delete.
2. Create multiple threads
	Thise threads adds Event to the Events Structure
	.. Takes an Event and addes to Events.myevents..
	.. Locks the Events.myevents Map
	.. Updates and 
	.. Unlocks the Map
3. Create Thread For Events
	* Writes Event to DB.
	* Removes Event from Events..