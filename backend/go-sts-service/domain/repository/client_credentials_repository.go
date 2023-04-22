package repository

import "ortisan-broker/go-sts-service/domain/entity"

type ClientCredentialsRepository interface {
	CreateClientCredentials(cr entity.ClientCredentials) (entity.ClientCredentials, error)
	FindByClientId(clientId string) (entity.ClientCredentials, error)
	FindByClientName(clientName string) (entity.ClientCredentials, error)
}
