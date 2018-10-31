package msgqueue

// EventEmitter methods: Emit
type EventEmitter interface {
	Emit(event Event) error
}
