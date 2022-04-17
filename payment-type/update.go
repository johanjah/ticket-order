package payment_type

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, paymentType *PaymentType, changes rel.Changeset) error {
	if err := paymentType.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	paymentType.UpdatedAt = time.Now()

	u.repository.MustUpdate(ctx, paymentType, changes)
	return nil
}
