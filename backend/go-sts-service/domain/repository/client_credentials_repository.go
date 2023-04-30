package repository

import (
	"context"
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-sts-service/domain/entity"
)

type ClientCredentialsRepository interface {
	CreateClientCredentials(ctx context.Context, clientCredentials entity.ClientCredentials) (entity.ClientCredentials, error)
	FindByClientId(ctx context.Context, clientId vo.Id) (entity.ClientCredentials, error)
	FindByClientName(ctx context.Context, clientName vo.Name) (entity.ClientCredentials, error)
}
