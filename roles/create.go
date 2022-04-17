package roles

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type create struct {
	repository rel.Repository

	// example add to other table
	//scores     scores.Service
}

func (c create) Create(ctx context.Context, role *Role) error {
	if err := role.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	role.CreatedAt = time.Now()

	c.repository.MustInsert(ctx, role)
	return nil
}
