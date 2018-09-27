package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/4piso/okevents/src/eventservice/rest"
	"github.com/4piso/okevents/src/lib/configuration"
	"github.com/4piso/okevents/src/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `../../src/lib/configuration/config.json`, "flag to set the path to the confiugarion file")
	flag.Parse()

	// extract the configuration file
	config, err := configuration.ExtractConfiguration(*confPath)
	if err != nil {
		log.Fatal("There is an error on the configuration part ", err)
	}

	fmt.Println("connecting to the database ...")

	// connection to the database
	dbhandler, err := dblayer.NewPersistenceLayer(config.DataBaseType, config.DBConnection)
	if err != nil {
		log.Fatal("DB connection ERROR: ", err)
	}

	log.Fatal(rest.ServeAPI(config.RestFulEndPoint, dbhandler))
}
