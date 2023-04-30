package usecase

import (
	"context"
	"errors"
	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-sts-service/domain/entity"
	"ortisan-broker/go-sts-service/domain/repository"
)

type CreateClientCredentialsUseCase interface {
	CreateClientCredentials(context.Context, entity.ClientCredentials) (entity.ClientCredentials, error)
}

type createClientCredentialsUseCase struct {
	clientCredentialsRepository repository.ClientCredentialsRepository
}

func (c createClientCredentialsUseCase) CreateClientCredentials(ctx context.Context, clientCredentials entity.ClientCredentials) (entity.ClientCredentials, error) {
	if ctx == nil {
		return nil, errApp.NewBadArgumentError("context is required")
	}
	if clientCredentials == nil {
		return nil, errApp.NewBadArgumentError("client credentials entity is required")
	}

	clientCredentialsPersisted, err := c.clientCredentialsRepository.CreateClientCredentials(ctx, clientCredentials)
	if err != nil {
		return nil, errApp.NewBaseErrorWithCause("error to create client credentials", err)
	}
	return clientCredentialsPersisted, nil
}

func NewCreateClientCredentialsUseCase(clientCredentialsRepository repository.ClientCredentialsRepository) (CreateClientCredentialsUseCase, error) {
	if clientCredentialsRepository == nil {
		return nil, errors.New("client credentials repository is required")
	}
	return &createClientCredentialsUseCase{
		clientCredentialsRepository: clientCredentialsRepository,
	}, nil
}
