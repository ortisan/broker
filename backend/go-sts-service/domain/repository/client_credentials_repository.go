package repository

import (
	"context"
	"ortisan-broker/go-sts-service/domain/entity"
)

type ClientCredentialsRepository interface {
	CreateClientCredentials(ctx context.Context, clientCredentials entity.ClientCredentials) (entity.ClientCredentials, error)
	FindByClientId(ctx context.Context, clientId string) (entity.ClientCredentials, error)
	FindByClientName(ctx context.Context, clientName string) (entity.ClientCredentials, error)
}
