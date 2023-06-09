package entity

import (
	"ortisan-broker/go-commons/domain/vo"
	errapp "ortisan-broker/go-commons/error"
)

type ClientCredentials interface {
	ClientName() vo.Name
	ClientId() vo.Id
	ClientSecret() vo.Secret
}

type clientCredentials struct {
	clientName   vo.Name
	clientId     vo.Id
	clientSecret vo.Secret
}

func (c *clientCredentials) ClientName() vo.Name {
	return c.clientName
}

func (c *clientCredentials) ClientId() vo.Id {
	return c.clientId
}

func (c *clientCredentials) ClientSecret() vo.Secret {
	return c.clientSecret
}

func NewClientCredentials(clientName vo.Name, clientId vo.Id, clientSecret vo.Secret) (ClientCredentials, error) {
	if clientName == nil {
		return nil, errapp.NewBadArgumentError("client name is required")
	}
	if clientId == nil {
		return nil, errapp.NewBadArgumentError("client id is required")
	}
	if clientSecret == nil {
		return nil, errapp.NewBadArgumentError("client secret is required")
	}

	return &clientCredentials{
		clientName:   clientName,
		clientId:     clientId,
		clientSecret: clientSecret,
	}, nil

}

func NewClientCredentialsWithoutName(clientId vo.Id, clientSecret vo.Secret) (ClientCredentials, error) {
	if clientId == nil {
		return nil, errapp.NewBadArgumentError("client id is required")
	}
	if clientSecret == nil {
		return nil, errapp.NewBadArgumentError("client secret is required")
	}

	return &clientCredentials{
		clientId:     clientId,
		clientSecret: clientSecret,
	}, nil

}
