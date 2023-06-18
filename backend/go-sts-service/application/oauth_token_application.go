package application

import (
	"context"
	"errors"
	errapp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/domain/usecase"
)

type CreateOauthTokenApplication interface {
	CreateOauthToken(ctx context.Context, token *dto.OauthTokenRequest) (*dto.OauthTokenResponse, error)
}

type createOauthTokenApplication struct {
	adapter OauthTokenAdapter
	usecase usecase.CreateOauthTokenUseCase
}

func (c *createOauthTokenApplication) CreateOauthToken(ctx context.Context, token *dto.OauthTokenRequest) (*dto.OauthTokenResponse, error) {
	if ctx == nil {
		return nil, errapp.NewBadArgumentError("context is required")
	}
	if token == nil {
		return nil, errapp.NewBadArgumentError("token request is required")
	}

	oauthToken, err := c.adapter.AdaptFromDtoToDomain(ctx, token)
	if err != nil {
		return nil, err
	}
	oauthToken.RenewToken()

	return c.adapter.AdaptFromDomainToDto(ctx, oauthToken)
}

func NewCreateOauthTokenApplication(adapter OauthTokenAdapter, usecase usecase.CreateOauthTokenUseCase) (CreateOauthTokenApplication, error) {
	if adapter == nil {
		return nil, errors.New("adapter is required")
	}
	if usecase == nil {
		return nil, errors.New("create oauth token use case is required")
	}

	return &createOauthTokenApplication{
		adapter: adapter,
		usecase: usecase,
	}, nil
}
