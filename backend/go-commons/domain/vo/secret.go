package vo

import (
	errApp "ortisan-broker/go-commons/error"

	"golang.org/x/crypto/bcrypt"
)

type Secret interface {
	Value() string
}

type secret struct {
	value string
}

func (p *secret) Value() string {
	return p.value
}

func NewSecretFromValue(value string) (Name, error) {
	if value == "" {
		return nil, errApp.NewBadArgumentError("secret is invalid")
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(value), 14)
	return &secret{
		value: string(bytes),
	}, nil
}

func NewSecretFromValueCrypted(value string) (Name, error) {
	return &secret{
		value: value,
	}, nil
}
