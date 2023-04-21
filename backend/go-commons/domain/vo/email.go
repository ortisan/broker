package vo

import (
	"net/mail"
	"ortisan-broker/go-user-service/domain/entity"
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

func NewEmail(value string) (Email, error) {
	if !validate(value) {
		return nil, entity.NewBadArgumentError("id value is invalid")
	}
	return &email{
		value: value,
	}, nil
}
