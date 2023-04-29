package application

import (
	"errors"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/domain/usecase"
)

type CreateClientCredentialsApplication interface {
	CreateClientCredentials(dto.ClientCredentialsRequest) (*dto.ClientCredentials, error)
}

type createClientCredentialsApplication struct {
	adapter ClientCredentialsAdapter
	usecase usecase.CreateClientCredentials
}

func NewCreateClientCredentialsApplication(adapter ClientCredentialsAdapter, usecase usecase.CreateClientCredentials) (CreateClientCredentialsApplication, error) {
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

func (ccca *createClientCredentialsApplication) CreateClientCredentials(input dto.ClientCredentialsRequest) (*dto.ClientCredentials, error) {
	clientCredentials, err := ccca.adapter.AdaptFromDtoToDomain(input)
	if err != nil {
		return nil, err
	}
	clientCredentialsCreated, err := ccca.usecase.CreateClientCredentials(clientCredentials)
	if err != nil {
		return nil, err
	}
	return ccca.adapter.AdaptFromDomainToDto(clientCredentialsCreated)
}

type CreateOauthTokenApplication interface {
	CreateOauthToken(dto.TokenRequest) (*dto.TokenResponse, error)
}

type createOauthTokenApplication struct {
	usecase usecase.CreateOauthTokenUseCase
}
