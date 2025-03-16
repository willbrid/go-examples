package user

import (
	"fmt"
)

type User struct {
	Name string
}

type ErrUserNotFound struct {
	User string
}

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("user %q not found", e.User)
}

var userDB map[string]*User

func FindUser(name string) (*User, error) {
	user, ok := userDB[name]
	if !ok {
		return nil, ErrUserNotFound{User: name}
	}

	return user, nil
}
