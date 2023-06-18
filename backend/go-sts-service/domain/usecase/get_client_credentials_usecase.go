package usecase

import (
	"context"
	"errors"
	errapp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-sts-service/domain/entity"
	"ortisan-broker/go-sts-service/domain/repository"
)

type GetClientCredentialsUseCase interface {
	GetClientCredentialsByClientId(ctx context.Context, clientCredentials entity.ClientCredentials) (entity.ClientCredentials, error)
}

type getClientCredentialsUseCase struct {
	clientCredentialsRepository repository.ClientCredentialsRepository
}

func (g *getClientCredentialsUseCase) GetClientCredentialsByClientId(ctx context.Context, clientCredentials entity.ClientCredentials) (entity.ClientCredentials, error) {
	if ctx == nil {
		return nil, errapp.NewBadArgumentError("context is required")
	}
	if clientCredentials == nil {
		return nil, errapp.NewBadArgumentError("client credentials entity is required")
	}

	clientCredentialsPersisted, err := g.clientCredentialsRepository.FindByClientId(ctx, clientCredentials.ClientId())
	if err != nil {
		return nil, errapp.NewBadArgumentErrorWithCause("error to getting client credentials", err)
	}
	if clientCredentialsPersisted == nil {
		return nil, errapp.NewNotFoundError("client credentials not founded")
	}
	return clientCredentialsPersisted, nil
}

func NewGetClientCredentialsUseCase(clientCredentialsRepository repository.ClientCredentialsRepository) (GetClientCredentialsUseCase, error) {
	if clientCredentialsRepository == nil {
		return nil, errors.New("client credentials repository is required")
	}

	return &getClientCredentialsUseCase{
		clientCredentialsRepository: clientCredentialsRepository,
	}, nil
}
