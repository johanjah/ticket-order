package events

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type create struct {
	repository rel.Repository
}

func (c create) Create(ctx context.Context, event *Event) error {
	if err := event.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	event.CreatedAt = time.Now()

	c.repository.MustInsert(ctx, event)
	return nil
}
