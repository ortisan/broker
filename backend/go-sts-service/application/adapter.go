package application

import (
	"context"
	"ortisan-broker/go-commons/domain/vo"
	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/domain/entity"
)

type ClientCredentialsAdapter interface {
	AdaptFromDtoToDomain(ctx context.Context, clientCredentials *dto.ClientCredentialsRequest) (entity.ClientCredentials, error)
	AdaptFromDomainToDto(ctx context.Context, clientCredentials entity.ClientCredentials) (*dto.ClientCredentials, error)
}

type clientCredentialsAdapter struct {
}

func (cca clientCredentialsAdapter) AdaptFromDtoToDomain(ctx context.Context, clientCredentials *dto.ClientCredentialsRequest) (entity.ClientCredentials, error) {
	if ctx == nil {
		return nil, errApp.NewBadArgumentError("context is required")
	}
	if clientCredentials == nil {
		return nil, errApp.NewBadArgumentError("client credentials is required")
	}

	clientName, err := vo.NewName(clientCredentials.ClientName)
	if err != nil {
		return nil, err
	}
	clientId := vo.NewId()
	clientSecret := vo.NewSecret()
	clientCredentialsEntity, err := entity.NewClientCredentials(clientName, clientId, clientSecret)
	if err != nil {
		return nil, err
	}
	return clientCredentialsEntity, err
}

func (cca clientCredentialsAdapter) AdaptFromDomainToDto(ctx context.Context, clientCredentials entity.ClientCredentials) (*dto.ClientCredentials, error) {
	return &dto.ClientCredentials{
		ClientName:   clientCredentials.ClientName().Value(),
		ClientId:     clientCredentials.ClientId().Value(),
		ClientSecret: clientCredentials.ClientSecret().Value(),
	}, nil
}

func NewClientCredentialsAdapter() (ClientCredentialsAdapter, error) {
	return &clientCredentialsAdapter{}, nil
}

type OauthTokenAdapter interface {
	AdaptFromDtoToDomain(ctx context.Context, tokenRequest *dto.OauthTokenRequest) (entity.OauthToken, error)
	AdaptFromDomainToDto(ctx context.Context, oauthToken entity.OauthToken) (*dto.OauthTokenResponse, error)
}

type oauthTokenAdapter struct {
}

func (o *oauthTokenAdapter) AdaptFromDtoToDomain(ctx context.Context, oauthTokenRequest *dto.OauthTokenRequest) (entity.OauthToken, error) {
	if ctx == nil {
		return nil, errApp.NewBadArgumentError("context is required")
	}
	if oauthTokenRequest == nil {
		return nil, errApp.NewBadArgumentError("oauth token request is required")
	}

	clientId := vo.NewId()
	clientSecret := vo.NewSecret()
	clientCredentialsEntity, err := entity.NewClientCredentialsWithoutName(clientId, clientSecret)
	if err != nil {
		return nil, err
	}

	return entity.NewOauthToken(clientCredentialsEntity)
}

func (o *oauthTokenAdapter) AdaptFromDomainToDto(ctx context.Context, oauthToken entity.OauthToken) (*dto.OauthTokenResponse, error) {
	if ctx == nil {
		return nil, errApp.NewBadArgumentError("context is required")
	}
	if oauthToken == nil {
		return nil, errApp.NewBadArgumentError("oauth token is required")
	}

	return &dto.OauthTokenResponse{
		ClientId:     oauthToken.ClientCredentials().ClientId().Value(),
		ClientSecret: oauthToken.ClientCredentials().ClientSecret().Value(),
		Token:        oauthToken.Value(),
	}, nil
}

func NewOauthTokenAdapter() (OauthTokenAdapter, error) {
	return &oauthTokenAdapter{}, nil
}
