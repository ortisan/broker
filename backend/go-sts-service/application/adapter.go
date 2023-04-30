package application

import (
	"context"
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/domain/entity"
)

type ClientCredentialsAdapter interface {
	AdaptFromDtoToDomain(ctx context.Context, clientCredentials dto.ClientCredentialsRequest) (entity.ClientCredentials, error)
	AdaptFromDomainToDto(ctx context.Context, clientCredentials entity.ClientCredentials) (*dto.ClientCredentials, error)
}

type clientCredentialsAdapter struct {
}

func (cca clientCredentialsAdapter) AdaptFromDtoToDomain(ctx context.Context, clientCredentials dto.ClientCredentialsRequest) (entity.ClientCredentials, error) {
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
