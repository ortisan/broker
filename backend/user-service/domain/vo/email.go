package vo

import (
	errApp "monolith/application/error"
	"net/mail"
)

type email struct {
	value string
}

type Email interface {
	Value() string
}

func (e *email) Value() string {
	return e.value
}

func validate(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func NewEmailFromValue(value string) (Email, error) {
	if !validate(value) {
		return nil, errApp.NewBadArgumentError("id value is invalid")
	}
	return &email{
		value: value,
	}, nil
}
