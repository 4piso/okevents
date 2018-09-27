package rest

import (
	"net/http"

	"github.com/4piso/okevents/src/lib/persistence"
	"github.com/gorilla/mux"
)

// ServeAPI is going to run our web server
func ServeAPI(endpoint string, dbhandler persistence.DatabaseHandler) error {
	// accesing the eventServices definition
	handler := newEventHandler(dbhandler)
	r := mux.NewRouter()

	eventsrouter := r.PathPrefix("/events").Subrouter()

	// define the routers of the rest api
	eventsrouter.Methods("GET").Path("{/SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(endpoint, r)
}
