module softott.net/gocon

go 1.24.2

require (
	github.com/google/uuid v1.6.0
	github.com/panjf2000/ants v1.3.0
)

replace softott.net/events => ./events

replace softott.net/gocon/events => ../events
