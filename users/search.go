package users

import (
	"context"

	"github.com/go-rel/rel"
)

// Filter for search.
type Filter struct {
	FirstName string
	LastName  string
	Email     string
}

type search struct {
	repository rel.Repository
}

func (s search) Search(ctx context.Context, user *[]User, filter Filter) error {
	var (
		query = rel.Select().SortAsc("username")
	)

	if filter.FirstName != "" {
		query = query.Where(rel.Like("first_name", "%"+filter.FirstName+"%"))
	}

	if filter.LastName != "" {
		query = query.Where(rel.Like("last_name", "%"+filter.LastName+"%"))
	}

	if filter.Email != "" {
		query = query.Where(rel.Like("email", "%"+filter.Email+"%"))
	}

	s.repository.MustFindAll(ctx, user, query)
	return nil
}
