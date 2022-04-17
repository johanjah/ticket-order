package shoppingcharts

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, shoppingChart *ShoppingChart, changes rel.Changeset) error {
	if err := shoppingChart.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	shoppingChart.UpdatedAt = time.Now()

	u.repository.MustUpdate(ctx, shoppingChart, changes)
	return nil
}
