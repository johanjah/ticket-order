package shoppingcharts

import (
	"context"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "shoppingChart")))
)

//go:generate mockery --name=Service --case=underscore --output shoppingchartstest --outpkg shoppingchartstest

// Service instance for shoppingchart's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	Search(ctx context.Context, shoppingCharts *[]ShoppingChart, filter Filter) error
	Create(ctx context.Context, shoppingChart *ShoppingChart) error
	Update(ctx context.Context, shoppingChart *ShoppingChart, changes rel.Changeset) error
	Delete(ctx context.Context, shoppingChart *ShoppingChart)
	Clear(ctx context.Context)
}

// beside embedding the struct, you can also declare the function directly on this struct.
// the advantage of embedding the struct is it allows spreading the implementation across multiple files.
type service struct {
	search
	create
	update
	delete
	clear
}

var _ Service = (*service)(nil)

// New ShoppingCharts service.
func New(repository rel.Repository) Service {
	return service{
		search: search{repository: repository},
		create: create{repository: repository},
		update: update{repository: repository},
		delete: delete{repository: repository},
		clear:  clear{repository: repository},
	}
}
