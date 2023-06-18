package database

import (
	"context"
	"ortisan-broker/go-commons/domain/vo"
	errapp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-sts-service/domain/entity"
)

type ClientCredentialsAdapter interface {
	AdaptFromDomainToModel(ctx context.Context, clientCredentials entity.ClientCredentials) (*ClientCredentials, error)
	AdaptFromModelToDomain(ctx context.Context, clientCredentials *ClientCredentials) (entity.ClientCredentials, error)
}

type clientCredentialsAdapter struct {
}

func (*clientCredentialsAdapter) AdaptFromDomainToModel(ctx context.Context, clientCredentials entity.ClientCredentials) (*ClientCredentials, error) {
	if ctx == nil {
		return nil, errapp.NewBadArgumentError("context is required")
	}
	if clientCredentials == nil {
		return nil, errapp.NewBadArgumentError("client credentials is required")
	}

	return &ClientCredentials{
		ClientName:   clientCredentials.ClientName().Value(),
		ClientId:     clientCredentials.ClientId().Value(),
		ClientSecret: clientCredentials.ClientSecret().Value(),
	}, nil
}

func (*clientCredentialsAdapter) AdaptFromModelToDomain(ctx context.Context, clientCredentials *ClientCredentials) (entity.ClientCredentials, error) {
	if ctx == nil {
		return nil, errapp.NewBadArgumentError("context is required")
	}
	if clientCredentials == nil {
		return nil, errapp.NewBadArgumentError("client credentials is required")
	}

	clientName, err := vo.NewName(clientCredentials.ClientName)
	if err != nil {
		return nil, err
	}
	clientId, err := vo.NewIdFromValue(clientCredentials.ClientId)
	if err != nil {
		return nil, err
	}

	clientSecret, err := vo.NewSecretFromValueCrypted(clientCredentials.ClientSecret)
	if err != nil {
		return nil, err
	}

	return entity.NewClientCredentials(clientName, clientId, clientSecret)
}

func NewClientCredentialsAdapter() (ClientCredentialsAdapter, error) {
	return &clientCredentialsAdapter{}, nil
}
