package application

import "ortisan-broker/go-user-service/adapter/dto"

type CreateClientCredentialsApplictions interface {
	CreateClientCredentials(dto.User) (*dto.User, error)
}
