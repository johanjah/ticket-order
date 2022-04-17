package payment_type

import (
	"errors"
	"time"
)

var (
	// ErrPaymentTypeNameBlank validation error.
	ErrPaymentTypeNameBlank = errors.New("payment type name can't be blank")
)

// PaymentType represent a record stored in todos table.
type PaymentType struct {
	ID                 uint      `json:"id"`
	PaymentName        string    `json:"payment_name"`
	PaymentDescription string    `json:"payment_description"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (t PaymentType) Validate() error {
	var err error
	switch {
	case len(t.PaymentName) == 0:
		err = ErrPaymentTypeNameBlank
	}

	return err
}
