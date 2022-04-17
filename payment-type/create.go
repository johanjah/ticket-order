package payment_type

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type create struct {
	repository rel.Repository
}

func (c create) Create(ctx context.Context, paymentType *PaymentType) error {
	if err := paymentType.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	paymentType.CreatedAt = time.Now()

	c.repository.MustInsert(ctx, paymentType)
	return nil
}
