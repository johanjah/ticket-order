package shoppingcharts

import (
	"context"

	"github.com/go-rel/rel"
)

type delete struct {
	repository rel.Repository
}

func (d delete) Delete(ctx context.Context, shoppingChart *ShoppingChart) {
	d.repository.MustDelete(ctx, shoppingChart)
}
