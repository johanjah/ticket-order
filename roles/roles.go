package roles

import (
	"errors"
	"time"
)

var (
	// ErrRoleNameBlank validation error.
	ErrRoleNameBlank = errors.New("roles name can't be blank")
)

// Role represent a record stored in todos table.
type Role struct {
	ID        uint      `json:"id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Role) Validate() error {
	var err error
	switch {
	case len(t.RoleName) == 0:
		err = ErrRoleNameBlank
	}

	return err
}
