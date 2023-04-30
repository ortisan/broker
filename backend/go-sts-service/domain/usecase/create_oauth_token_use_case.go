package usecase

import (
	"context"
	"errors"
	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-sts-service/domain/entity"
)

type CreateOauthTokenUseCase interface {
	CreateOauthToken(ctx context.Context, clientCredentials entity.ClientCredentials) (entity.OauthToken, error)
}

type createOauthTokenUseCase struct {
	getClientCredentialsUseCase GetClientCredentialsUseCase
}

func (c *createOauthTokenUseCase) CreateOauthToken(ctx context.Context, clientCredentials entity.ClientCredentials) (entity.OauthToken, error) {
	if ctx == nil {
		return nil, errApp.NewBadArgumentError("context is required")
	}
	if clientCredentials == nil {
		return nil, errApp.NewBadArgumentError("client credentials entity is required")
	}

	clientCredentialsPersisted, err := c.getClientCredentialsUseCase.GetClientCredentialsByClientId(ctx, clientCredentials)
	if err != nil {
		return nil, err
	}

	oauthToken, err := entity.NewOauthToken(clientCredentialsPersisted)
	if err != nil {
		return nil, errApp.NewBadArgumentErrorWithCause("error to generate oauth token", err)
	}
	return oauthToken, nil
}

func NewCreateOauthTokenUseCase(getClientCredentialsUseCase GetClientCredentialsUseCase) (CreateOauthTokenUseCase, error) {
	if getClientCredentialsUseCase == nil {
		return nil, errors.New("get client credentials use case is required")
	}

	return &createOauthTokenUseCase{
		getClientCredentialsUseCase: getClientCredentialsUseCase,
	}, nil
}
