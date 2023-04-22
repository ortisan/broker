package vo

import (
	nurl "net/url"
	errApp "ortisan-broker/go-commons/error"
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
		return nil, errApp.NewBadArgumentError("url value is invalid")
	}
	urlN, err := nurl.ParseRequestURI(value)
	if err != nil {
		return nil, errApp.NewBadArgumentErrorWithCause("url value is invalid", err)
	}
	return &url{
		value: urlN,
	}, nil
}
