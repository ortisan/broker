package application

import (
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/domain/entity"
)

func AdaptClientCredentialsDtoToClientCredentialsDomain(clientCredentials dto.ClientCredentialsRequest) (entity.ClientCredentials, error) {
	clientName := clientCredentials.ClientName
	clientNameVo, err := vo.NewName(clientName)
	if err != nil {
		return nil, err
	}
	clientIdVo := vo.NewId()
	clientSecret, err := vo.NewSecret()
		return nil, err
	}
	clientCredentialsEntity, err := entity.NewClientCredentials(clientName, clientId, clientSecret)
	if err != nil {
		return nil, err
	}
	return clientCredentialsEntity, err
}

func AdaptUserDomainToUserDto(user entity.User) (dto.User, error) {
	var profileUrlStr string
	if user.ProfileAvatarUrl() != nil {
		profileUrlStr = user.ProfileAvatarUrl().Value()
	}
	return dto.User{
		ID:               user.Id().Value(),
		Name:             user.Name().Value(),
		Email:            user.Email().Value(),
		Username:         user.Username().Value(),
		Password:         user.Password().Value(),
		FederationId:     user.FederationId().Value(),
		ProfileAvatarUrl: profileUrlStr,
	}, nil
}
