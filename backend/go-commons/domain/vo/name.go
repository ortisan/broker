package vo

import errApp "ortisan-broker/go-commons/error"

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
		return nil, errApp.NewBadArgumentError("name value is invalid")
	}
	return &name{
		value: value,
	}, nil
}
