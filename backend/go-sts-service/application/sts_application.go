package application

import (
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/domain/usecase"
)

type CreateClientCredentialsApplication interface {
	CreateClientCredentials(dto.ClientCredentialsRequest) (*dto.ClientCredentials, error)
}

type createClientCredentialsApplication struct {
	usecase usecase.CreateClientCredentials
}

func (ccca *createClientCredentialsApplication) CreateClientCredentials(input dto.ClientCredentialsRequest) (*dto.ClientCredentials, error) {

	ccca.usecase.CreateClientCredentials()
}

type CreateOauthTokenApplication interface {
	CreateOauthToken(dto.TokenRequest) (*dto.TokenResponse, error)
}

type createOauthTokenApplication struct {
	usecase usecase.CreateOauthTokenUseCase
}
