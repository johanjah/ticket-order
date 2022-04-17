package roles

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, role *Role, changes rel.Changeset) error {
	if err := role.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	role.UpdatedAt = time.Now()

	u.repository.MustUpdate(ctx, role, changes)
	return nil
}
