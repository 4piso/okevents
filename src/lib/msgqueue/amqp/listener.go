package amqp

import (
	"fmt"

	"github.com/4piso/okevents/src/lib/msgqueue"
	"github.com/streadway/amqp"
)

const eventNameHeader = "x-event-name"

type amqpEventListener struct {
	connection *amqp.Connection
	exchange   string
	queue      string
	mapper     msgqueue.EventMapper
}

// NewAMQPEventListener contructor functions
func NewAMQPEventListener(conn *amqp.Connection, exchange string, queue string) (msgqueue.EventListener, error) {
	// amqp listener struct
	listener := amqpEventListener{
		connection: conn,
		exchange:   exchange,
		queue:      queue,
		mapper:     msgqueue.NewEventMapper(),
	}

	err := listener.setup()
	if err != nil {
		return nil, err
	}

	return &listener, nil
}

// setup
func (a *amqpEventListener) setup() error {
	// instance the channel call
	channel, err := a.connection.Channel()
	if err != nil {
		return nil
	}

	defer channel.Close()

	if err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil); err != nil {
		return err
	}

	_, err = channel.QueueDeclare(a.queue, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("could not declare queue %s: %s ", a.queue, err)
	}

	fmt.Printf("Setup finish with ExchangeDeclar: %s and QueueDeclare: %s \n ", a.exchange, a.queue)
	return nil
}

// Listen fn to listen all the events on the cicle
func (a *amqpEventListener) Listen(eventsNames ...string) (<-chan msgqueue.Event, <-chan error, error) {
	// open the channel connection
	channel, err := a.connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	// creating the binding between queue and exchange for each listener event type
	for _, event := range eventsNames {
		if err := channel.QueueBind(a.queue, event, a.exchange, false, nil); err != nil {
			return nil, nil, fmt.Errorf("could not bind event %s to the queue %s: %s", event, a.queue, err)
		}
	}

	msgs, err := channel.Consume(a.queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("could not consume queue: %s ", err)
	}

	fmt.Printf("consume setup finished with QUEUE: %s \n ", a.queue)

	events := make(chan msgqueue.Event)
	errors := make(chan error)

	// starting the gorutine
	go func() {
		// loop throw all the events messsage using the msgs variables
		for msg := range msgs {

			rawEventName, ok := msg.Headers[eventNameHeader]
			if !ok {
				errors <- fmt.Errorf("message did not contain %s header", eventNameHeader)
				msg.Nack(false, false)
				continue
			}

			eventName, ok := rawEventName.(string)
			if !ok {
				errors <- fmt.Errorf("header %s did not contain string ", eventNameHeader)
				msg.Nack(false, false)
				continue
			}
			event, err := a.mapper.MapEvent(eventName, msg.Body)
			if err != nil {
				errors <- fmt.Errorf("could not unmarshal %s: %s ", eventName, err)
				msg.Nack(false, false)
				continue
			}

			events <- event
			msg.Ack(false)

		}
	}()

	// return everything for the pages
	return events, errors, nil

}

func (a *amqpEventListener) Mapper() msgqueue.EventMapper {
	return a.mapper
}
