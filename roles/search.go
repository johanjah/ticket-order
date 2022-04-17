package roles

import (
	"context"

	"github.com/go-rel/rel"
)

// Filter for search.
type Filter struct {
	Keyword string
}

type search struct {
	repository rel.Repository
}

func (s search) Search(ctx context.Context, role *[]Role, filter Filter) error {
	//var (
	//	query = rel.Select().SortAsc("order")
	//)
	//
	//if filter.Keyword != "" {
	//	query = query.Where(rel.Like("rolename", "%"+filter.Keyword+"%"))
	//}

	s.repository.MustFindAll(ctx, role)
	return nil
}
