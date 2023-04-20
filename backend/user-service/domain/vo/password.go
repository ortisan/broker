package vo

import (
	errApp "user-service/application/error"

	"golang.org/x/crypto/bcrypt"
)

type Password interface {
	Value() string
}

type password struct {
	value string
}

func (p *password) Value() string {
	return p.value
}

func NewPasswordFromValue(value string) (Name, error) {
	if value == "" {
		return nil, errApp.NewBadArgumentError("password is invalid")
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(value), 14)
	return &password{
		value: string(bytes),
	}, nil
}
