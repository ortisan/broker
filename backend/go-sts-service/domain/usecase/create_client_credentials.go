package usecase

import "ortisan-broker/go-sts-service/domain/entity"

type CreateClientCredentials interface {
	CreateClientCredentials(user entity.ClientCredentials) (entity.ClientCredentials, error)
}

type createClientCredentials struct {
}

type CreateOauthTokenUseCase interface {
	CreateOauthToken(user entity.ClientCredentials) (entity.OauthToken, error)
}
