package roles

import (
	"context"

	"github.com/go-rel/rel"
)

type delete struct {
	repository rel.Repository
}

func (d delete) Delete(ctx context.Context, role *Role) {
	d.repository.MustDelete(ctx, role)
}
