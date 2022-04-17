package payment_type

import (
	"context"

	"github.com/go-rel/rel"
)

type clear struct {
	repository rel.Repository
}

func (c clear) Clear(ctx context.Context) {
	c.repository.MustDeleteAny(ctx, rel.From("payment_type"))
}
