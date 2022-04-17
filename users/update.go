package users

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, user *User, changes rel.Changeset) error {
	if err := user.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	user.UpdatedAt = time.Now()

	u.repository.MustUpdate(ctx, user, changes)
	return nil
}
