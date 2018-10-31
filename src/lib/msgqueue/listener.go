package msgqueue

// EventListener flow events
type EventListener interface {
	Listen(events ...string) (<-chan Event, <-chan error, error)
	Mapper() EventMapper
}
