package vo

import (
	nurl "net/url"
	errapp "ortisan-broker/go-commons/error"
)

type url struct {
	value *nurl.URL
}

type Url interface {
	Value() string
}

func (u *url) Value() string {
	return u.value.String()
}

func NewUrlFromValue(value string) (Url, error) {
	if value == "" {
		return nil, errapp.NewBadArgumentError("url value is invalid")
	}
	urlN, err := nurl.ParseRequestURI(value)
	if err != nil {
		return nil, errapp.NewBadArgumentErrorWithCause("url value is invalid", err)
	}
	return &url{
		value: urlN,
	}, nil
}
