package contracts

import "time"

// EventCreatedEvent struct definition
type EventCreatedEvent struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LocationID string    `json:"location_id"`
	Start      time.Time `json:"start_time"`
	End        time.Time `json:"end_time"`
}

// EventName implemente the EventCreatedEvent struct
func (c *EventCreatedEvent) EventName() string {
	return "event.created"
}
