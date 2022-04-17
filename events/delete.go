package events

import (
	"context"

	"github.com/go-rel/rel"
)

type delete struct {
	repository rel.Repository
}

func (d delete) Delete(ctx context.Context, event *Event) {
	d.repository.MustDelete(ctx, event)
}
