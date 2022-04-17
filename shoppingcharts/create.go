package shoppingcharts

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type create struct {
	repository rel.Repository
}

func (c create) Create(ctx context.Context, shoppingChart *ShoppingChart) error {
	if err := shoppingChart.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	shoppingChart.CreatedAt = time.Now()

	c.repository.MustInsert(ctx, shoppingChart)
	return nil
}
