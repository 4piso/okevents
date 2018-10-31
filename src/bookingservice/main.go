package main

import (
	"flag"
	"fmt"

	"github.com/4piso/okevents/src/bookingservice/listener"
	"github.com/4piso/okevents/src/lib/configuration"
	msgqueue_amqp "github.com/4piso/okevents/src/lib/msgqueue/amqp"
	"github.com/4piso/okevents/src/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	// get the config path
	confPath := flag.String("config", `../../src/lib/configuration/config.json`, "flag to set the configuration file for the booking services")
	flag.Parse()

	// extract the configuration file
	config, _ := configuration.ExtractConfiguration(*confPath)

	dbhandler, err := dblayer.NewPersistenceLayer(config.DataBaseType, config.DBConnection)
	if err != nil {
		fmt.Println("erro connecting to the db: ", err)
	}

	// config the message broker with RabbitMQ
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		fmt.Println("Error connecting to the message broker: ", err)
	}

	// activate the event listener
	eventListener, err := msgqueue_amqp.NewAMQPEventListener(conn, "okevents", "booking")
	if err != nil {
		fmt.Println("cannot create connection with the event listener: ", err)
	}

	// call the process events
	processor := listener.EventProcessor{eventListener, dbhandler}
	processor.ProcessEvents()
}
