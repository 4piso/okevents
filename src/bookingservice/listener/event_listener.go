package listener

import (
	"log"

	"github.com/4piso/okevents/src/contracts"
	"github.com/4piso/okevents/src/lib/msgqueue"
	"github.com/4piso/okevents/src/lib/persistence"
)

// EventProcessor struct definition
type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      persistence.DatabaseHandler
}

// ProcessEvents fn () => error
func (p *EventProcessor) ProcessEvents() error {

	log.Println("Listening to events...")

	recieved, errors, err := p.EventListener.Listen("event.created")
	if err != nil {
		return err
	}

	// loop to all the events message to get the events
	for {
		select {
		case evt := <-recieved:
			p.handleEvent(evt)
		case err = <-errors:
			log.Printf("recieved error, while processing msg: %s ", err)
		}
	}
}

// handlerEvent (evt) =>
func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created: %s ", e.ID, e)
		//p.Database.AddEvent(persistence.Event{ID: bson.ObjectId(e.ID)})
	default:
		log.Printf("Unknow event: %t ", e)
	}
}
