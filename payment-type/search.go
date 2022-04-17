package payment_type

import (
	"context"

	"github.com/go-rel/rel"
)

// Filter for search.
type Filter struct {
	PaymentName        string
	PaymentDescription string
}

type search struct {
	repository rel.Repository
}

func (s search) Search(ctx context.Context, paymentType *[]PaymentType, filter Filter) error {
	var (
		query = rel.Select()
	)

	if filter.PaymentName != "" {
		query = query.Where(rel.Like("payment_name", "%"+filter.PaymentName+"%"))
	}

	if filter.PaymentDescription != "" {
		query = query.Where(rel.Like("payment_description", "%"+filter.PaymentDescription+"%"))
	}

	s.repository.MustFindAll(ctx, paymentType, query)
	return nil
}
