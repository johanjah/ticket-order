package shoppingcharts

import (
	"errors"
	"time"
)

var (
	ErrShoppingChartTicketCountBlank = errors.New("ticket can't be zero or less")
)

// ShoppingChart represent a record stored in todos table.
type ShoppingChart struct {
	ID          uint      `json:"id"`
	UserID      int       `json:"user_id"`
	EventID     int       `json:"event_id"`
	TicketCount int       `json:"ticket_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t ShoppingChart) Validate() error {
	var err error
	switch {
	case t.TicketCount <= 0:
		err = ErrShoppingChartTicketCountBlank
	}

	return err
}
