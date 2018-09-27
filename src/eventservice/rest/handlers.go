package rest

import (
	"net/http"

	"github.com/4piso/okevents/src/lib/persistence"
)

type eventServicesHandler struct {
	dbhandler persistence.DatabaseHandler
}

// newEventHandler contructor
func newEventHandler(databasehandler persistence.DatabaseHandler) *eventServicesHandler {
	return &eventServicesHandler{
		dbhandler: databasehandler,
	}
}

// implementing the event services definitions
func (eh *eventServicesHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	panic("still nothing")
}
func (eh *eventServicesHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {
	panic("still nothing")
}
func (eh *eventServicesHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {
	panic("still nothing")
}
