package msgqueue

// EventMapper interface definition
type EventMapper interface {
	MapEvent(string, interface{}) (Event, error)
}

// NewEventMapper () => EventMapper
func NewEventMapper() EventMapper {
	return &StaticEventMapper{}
}
