package amqp

import (
	"encoding/json"
	"fmt"

	"github.com/4piso/okevents/src/lib/msgqueue"
	"github.com/streadway/amqp"
)

// amqpEventEmitter struct definition
type amqpEventEmitter struct {
	connection *amqp.Connection
	exchange   string
}

// NewAMQPEventEmitter contructor fn: (amqp connection) => (eventEmitter, error)
func NewAMQPEventEmitter(conn *amqp.Connection, exchange string) (msgqueue.EventEmitter, error) {
	// defining the emitter events
	emitter := amqpEventEmitter{
		connection: conn,
		exchange:   exchange,
	}

	err := emitter.setup()
	if err != nil {
		return nil, err
	}

	return &emitter, nil
}

// setup method for ground braking the setup
func (a *amqpEventEmitter) setup() error {
	// getting the connectins from the user
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	fmt.Println("EXCHANGE DECLARE PUB: ", a.exchange)

	if err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil); err != nil {
		return fmt.Errorf("ERROR OCURRED WHILE SETTING EXCHANGE DECLARE: %s ", err)
	}

	return nil
}

// Emit fn: (event) => error
func (a *amqpEventEmitter) Emit(event msgqueue.Event) error {
	// parse the json format from the event
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	fmt.Println("EVENTOS: ", event)
	// parse the events on json format
	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("could not parse the json format: %s ", err)
	}
	// build the message to emit on the define structure
	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		ContentType: "application/json",
		Body:        jsonDoc,
	}

	if err = channel.Publish(a.exchange, event.EventName(), false, false, msg); err != nil {
		return fmt.Errorf("Error doing the publication: %s ", err)
	}

	return nil

}
