package vo

import "go-commons/shared/error"

type name struct {
	value string
}

type Name interface {
	Value() string
}

func (n *name) Value() string {
	return n.value
}

func NewName(value string) (Name, error) {
	if value == "" {
		return nil, errorApp.NewBadArgumentError("name value is invalid")
	}
	return &name{
		value: value,
	}, nil
}
