package vo

import (
	errapp "ortisan-broker/go-commons/error"
	"time"
)

type date struct {
	value time.Time
}

type Date interface {
	Value() time.Time
	ValueAsString() string
}

func (d *date) Value() time.Time {
	return d.value
}

func (d *date) ValueAsString() string {
	return d.value.Format(time.DateOnly)
}

func NewDate() Date {
	return &date{
		value: time.Now(),
	}
}

func NewDateFromValueAsString(valueStr string) (DateTime, error) {
	if valueStr == "" {
		return nil, errapp.NewBadArgumentError("date is invalid")
	}
	d, err := time.Parse(time.DateOnly, valueStr)
	if err != nil {
		errapp.NewBadArgumentErrorWithCause("date has invalid format", err)
	}
	return &date{value: d}, nil
}
