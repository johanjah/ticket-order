package events

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, event *Event, changes rel.Changeset) error {
	if err := event.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	event.UpdatedAt = time.Now()

	u.repository.MustUpdate(ctx, event, changes)
	return nil
}
