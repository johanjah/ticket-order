package shoppingcharts

import (
	"context"

	"github.com/go-rel/rel"
)

// Filter for search.
type Filter struct {
}

type search struct {
	repository rel.Repository
}

func (s search) Search(ctx context.Context, shoppingChart *[]ShoppingChart, filter Filter) error {

	// search not implemented yet

	s.repository.MustFindAll(ctx, shoppingChart)
	return nil
}
