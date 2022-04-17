package events

import (
	"errors"
	"time"
)

var (
	// ErrEventNameBlank validation error.
	ErrEventNameBlank      = errors.New("event name can't be blank")
	ErrEventPriceBlank     = errors.New("event price can't be zero or less")
	ErrEventStartDateBlank = errors.New("start date can't be blank")
	ErrEventEndDateBlank   = errors.New("end date can't be blank")
)

// Event represent a record stored in todos table.
type Event struct {
	ID               uint      `json:"id"`
	EventName        string    `json:"event_name"`
	EventDescription string    `json:"event_description"`
	EventLocation    string    `json:"event_location"`
	BasePrice        float64   `json:"base_price"`
	EventStartDate   time.Time `json:"event_start_date"`
	EventEndDate     time.Time `json:"event_end_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (t Event) Validate() error {
	var err error
	switch {
	case len(t.EventName) == 0:
		err = ErrEventNameBlank
	case t.BasePrice == 0:
		err = ErrEventNameBlank
	case t.EventStartDate.IsZero():
		err = ErrEventNameBlank
	case t.EventEndDate.IsZero():
		err = ErrEventNameBlank

	}

	return err
}
