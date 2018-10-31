package msgqueue

// Event using the event_created struct and method
type Event interface {
	EventName() string
}
