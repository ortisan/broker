package usecase

import (
	"context"
	"errors"
	"ortisan-broker/go-sts-service/domain/entity"
	"ortisan-broker/go-sts-service/domain/repository"
)

type CreateClientCredentialsUseCase interface {
	CreateClientCredentials(context.Context, entity.ClientCredentials) (entity.ClientCredentials, error)
}

type createClientCredentialsUseCase struct {
	clientCredentialsRepository repository.ClientCredentialsRepository
}

func (cccu createClientCredentialsUseCase) CreateClientCredentials(ctx context.Context, clientCredentials entity.ClientCredentials) (*entity.ClientCredentials, error) {
	cccu.clientCredentialsRepository(ctx, clientCredentials)
}

func NewCreateClientCredentialsUseCase(clientCredentialsRepository repository.ClientCredentialsRepository) (CreateClientCredentialsUseCase, error) {
	if clientCredentialsRepository == nil {
		return nil, errors.New("client credentials repository is required")
	}
	return &createClientCredentialsUseCase{
		clientCredentialsRepository: clientCredentialsRepository,
	}, nil
}

type CreateOauthTokenUseCase interface {
	CreateOauthToken(user entity.ClientCredentials) (entity.OauthToken, error)
}
