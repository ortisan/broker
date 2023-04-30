package application

import (
	"context"
	"errors"
	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/domain/usecase"
)

type CreateClientCredentialsApplication interface {
	CreateClientCredentials(ctx context.Context, clientCredentials *dto.ClientCredentialsRequest) (*dto.ClientCredentials, error)
}

type createClientCredentialsApplication struct {
	adapter ClientCredentialsAdapter
	usecase usecase.CreateClientCredentialsUseCase
}

func NewCreateClientCredentialsApplication(adapter ClientCredentialsAdapter, usecase usecase.CreateClientCredentialsUseCase) (CreateClientCredentialsApplication, error) {
	if adapter == nil {
		return nil, errors.New("adapter is required")
	}
	if usecase == nil {
		return nil, errors.New("usecase is required")
	}

	return &createClientCredentialsApplication{
		adapter: adapter,
		usecase: usecase,
	}, nil
}

func (ccca *createClientCredentialsApplication) CreateClientCredentials(ctx context.Context, input *dto.ClientCredentialsRequest) (*dto.ClientCredentials, error) {
	if ctx == nil {
		return nil, errApp.NewBadArgumentError("context is required")
	}
	if input == nil {
		return nil, errApp.NewBadArgumentError("client credentials request is required")
	}

	clientCredentials, err := ccca.adapter.AdaptFromDtoToDomain(ctx, input)
	if err != nil {
		return nil, err
	}
	clientCredentialsCreated, err := ccca.usecase.CreateClientCredentials(ctx, clientCredentials)
	if err != nil {
		return nil, err
	}
	return ccca.adapter.AdaptFromDomainToDto(ctx, clientCredentialsCreated)
}
