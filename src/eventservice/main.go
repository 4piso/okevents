package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/4piso/okevents/src/eventservice/rest"
	"github.com/4piso/okevents/src/lib/configuration"
	msgqueue_amqp "github.com/4piso/okevents/src/lib/msgqueue/amqp"
	"github.com/4piso/okevents/src/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	confPath := flag.String("conf", `../../src/lib/configuration/config.json`, "flag to set the path to the confiugarion file")
	flag.Parse()

	// extract the configuration file
	config, err := configuration.ExtractConfiguration(*confPath)
	if err != nil {
		log.Fatal("There is an error on the configuration part ", err)
	}

	// add the message broker to the main functionallity
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}

	// add the new event emitter services calling the event emitter contruction function
	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn, "okevents")
	if err != nil {
		panic(err)
	}

	fmt.Println("connecting to the database ...")

	// connection to the database
	dbhandler, err := dblayer.NewPersistenceLayer(config.DataBaseType, config.DBConnection)
	if err != nil {
		log.Fatal("DB connection ERROR: ", err)
	}

	log.Fatal(rest.ServeAPI(config.RestFulEndPoint, dbhandler, emitter))
}
