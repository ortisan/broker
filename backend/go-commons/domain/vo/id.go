package vo

import (
	"ortisan-broker/go-user-service/domain/entity"

	uuid "github.com/satori/go.uuid"
)

type id struct {
	value string
}

type Id interface {
	Value() string
}

func (i *id) Value() string {
	return i.value
}

func NewIdFromValue(value string) (Id, error) {
	if value == "" {
		return nil, entity.NewBadArgumentError("id value is invalid")
	}
	return &id{
		value: value,
	}, nil
}

func NewId() Id {
	uuidStr := uuid.NewV4().String()
	id, _ := NewIdFromValue(uuidStr)
	return id
}
