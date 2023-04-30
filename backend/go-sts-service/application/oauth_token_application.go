package application

import (
	"context"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/domain/usecase"
)

type CreateOauthTokenApplication interface {
	CreateOauthToken(ctx context.Context, token dto.TokenRequest) (*dto.TokenResponse, error)
}

type createOauthTokenApplication struct {
	usecase usecase.CreateOauthTokenUseCase
}
