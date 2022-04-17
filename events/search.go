package events

import (
	"context"

	"github.com/go-rel/rel"
)

// Filter for search.
type Filter struct {
	EventName        string
	EventDescription string
}

type search struct {
	repository rel.Repository
}

func (s search) Search(ctx context.Context, event *[]Event, filter Filter) error {
	var (
		query = rel.Select()
	)

	if filter.EventName != "" {
		query = query.Where(rel.Like("event_name", "%"+filter.EventName+"%"))
	}

	if filter.EventDescription != "" {
		query = query.Where(rel.Like("event_description", "%"+filter.EventDescription+"%"))
	}

	s.repository.MustFindAll(ctx, event)
	return nil
}
