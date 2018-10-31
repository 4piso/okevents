package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/4piso/okevents/src/contracts"
	"github.com/4piso/okevents/src/lib/msgqueue"
	"github.com/4piso/okevents/src/lib/persistence"
	"github.com/gorilla/mux"
)

type eventServicesHandler struct {
	dbhandler    persistence.DatabaseHandler
	eventEmitter msgqueue.EventEmitter
}

// newEventHandler contructor
func newEventHandler(databasehandler persistence.DatabaseHandler, eventemitter msgqueue.EventEmitter) *eventServicesHandler {
	return &eventServicesHandler{
		dbhandler:    databasehandler,
		eventEmitter: eventemitter,
	}
}

// implementing the event services definitions this will be use for one or multiple events
func (eh *eventServicesHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	// get the url variables search criteria keyword
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		fmt.Fprintf(w, `{error : No search criteria found, you can either search by id /id/45 or by /name/codply/}`)
		return
	}

	// add the search key word
	searchKey, ok := vars["search"]
	if !ok {
		fmt.Fprintf(w, `{error : No search criteria found, you can either search by id /id/45 or by /name/codply/}`)
		return
	}

	// defining local variables
	var event persistence.Event
	var err error

	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchKey)
	case "id":
		id, err := hex.DecodeString(searchKey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}

	if err != nil {
		fmt.Fprintf(w, `{ERROR: "%s" }`, err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)

}

func (eh *eventServicesHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {
	// add the search for all events in the persistence layer
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		fmt.Fprintf(w, "{error: Error ocurred while trying to loop throw all the data events}")
		return
	}

	// add the json format to the header
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		fmt.Fprintf(w, `{error: Error ocurred while trying to parse the json file %s}`, err)
		return
	}

}

func (eh *eventServicesHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {
	// local properties of the event
	event := persistence.Event{}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		fmt.Fprintf(w, `{error: Error ocurred while trying to decode the json format %s }`, err)
		return
	}

	id, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		fmt.Fprintf(w, `{error: error trying to persistence event %d %s }`, id, err)
		return
	}

	fmt.Fprintf(w, `{"id": %d }`, id)

	// create the eventCreated struct to pass the events to rabbit MQ
	msg := contracts.EventCreatedEvent{
		ID:         hex.EncodeToString(id),
		Name:       event.Name,
		LocationID: string(event.Location.ID),
		Start:      time.Unix(event.StartDate/1000, 0),
		End:        time.Unix(event.EndDate/1000, 0),
	}

	if err = eh.eventEmitter.Emit(&msg); err != nil {
		fmt.Printf("ERROR: ", err)
	}
	// add the json properties to the header
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)

}
