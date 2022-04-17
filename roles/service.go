package roles

import (
	"context"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "roles")))
)

//go:generate mockery --name=Service --case=underscore --output rolestest --outpkg rolestest

// Service instance for roles's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	Search(ctx context.Context, role *[]Role, filter Filter) error
	Create(ctx context.Context, role *Role) error
	Update(ctx context.Context, role *Role, changes rel.Changeset) error
	Delete(ctx context.Context, role *Role)
	Clear(ctx context.Context)
}

// beside embeding the struct, you can also declare the function directly on this struct.
// the advantage of embedding the struct is it allows spreading the implementation across multiple files.
type service struct {
	search
	create
	update
	delete
	clear
}

var _ Service = (*service)(nil)

// New Roles service.
func New(repository rel.Repository) Service {
	return service{
		search: search{repository: repository},
		create: create{repository: repository},
		update: update{repository: repository},
		delete: delete{repository: repository},
		clear:  clear{repository: repository},
	}
}
