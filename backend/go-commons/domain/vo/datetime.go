package vo

import (
	errapp "ortisan-broker/go-commons/error"
	"time"
)

type datetime struct {
	value time.Time
}

type DateTime interface {
	Value() time.Time
	ValueAsString() string
}

func (d *datetime) Value() time.Time {
	return d.value
}

func (d *datetime) ValueAsString() string {
	return d.value.Format(time.RFC3339)
}

func NewDateTime() DateTime {
	return &datetime{
		value: time.Now(),
	}
}

func NewDateTimeFromValueAsString(valueStr string) (DateTime, error) {
	if valueStr == "" {
		return nil, errapp.NewBadArgumentError("date time is invalid")
	}
	dt, err := time.Parse(time.RFC3339, valueStr)
	if err != nil {
		errapp.NewBadArgumentErrorWithCause("date time has invalid format", err)
	}
	return &datetime{value: dt}, nil
}
