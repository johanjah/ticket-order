package payment_type

import (
	"context"

	"github.com/go-rel/rel"
)

type delete struct {
	repository rel.Repository
}

func (d delete) Delete(ctx context.Context, paymentType *PaymentType) {
	d.repository.MustDelete(ctx, paymentType)
}
