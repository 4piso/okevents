package msgqueue

import (
	"encoding/json"
	"fmt"

	"github.com/4piso/okevents/src/contracts"
	"github.com/mitchellh/mapstructure"
)

// StaticEventMapper get the struct definition
type StaticEventMapper struct{}

// MapEvent (eventName, serialized) => Event, error
func (a *StaticEventMapper) MapEvent(eventName string, serialized interface{}) (Event, error) {
	// instance the events structure
	var event Event

	switch eventName {
	case "event.created":
		event = &contracts.EventCreatedEvent{}
	default:
		return nil, fmt.Errorf("unknown event type %s ", eventName)
	}

	// conver the serialized mgs body
	switch s := serialized.(type) {
	case []byte:
		if err := json.Unmarshal(s, event); err != nil {
			return nil, fmt.Errorf("could not unmarshal event %s: %s ", eventName, err)
		}
	default:
		cfg := mapstructure.DecoderConfig{
			Result:  event,
			TagName: "json",
		}
		dec, err := mapstructure.NewDecoder(&cfg)
		if err != nil {
			return nil, fmt.Errorf("could not initialize decoder for event %s: %s ", eventName, err)
		}

		err = dec.Decode(s)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal event %s: %s ", eventName, err)
		}
	}

	return event, nil
}
