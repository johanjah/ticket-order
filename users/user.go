package users

import (
	"errors"
	"time"
)

var (
	// ErrUserUsernameBlank validation error.
	ErrUserUsernameBlank  = errors.New("username can't be blank")
	ErrUserEmailBlank     = errors.New("email can't be blank")
	ErrUserPasswordBlank  = errors.New("password can't be blank")
	ErrUserFirstnameBlank = errors.New("first Name can't be blank")
	ErrUserLastnameBlank  = errors.New("last Name can't be blank")
)

// User represent a record stored in todos table.
type User struct {
	ID        uint      `json:"id"`
	RoleID    int       `json:"role_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t User) Validate() error {
	var err error
	switch {
	case len(t.Username) == 0:
		err = ErrUserUsernameBlank
	case len(t.Email) == 0:
		err = ErrUserEmailBlank
	case len(t.Password) == 0:
		err = ErrUserPasswordBlank
	case len(t.FirstName) == 0:
		err = ErrUserFirstnameBlank
	case len(t.LastName) == 0:
		err = ErrUserLastnameBlank
	}

	return err
}
