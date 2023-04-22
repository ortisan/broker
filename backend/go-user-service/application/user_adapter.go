package application

import (
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-user-service/adapter/dto"
	"ortisan-broker/go-user-service/domain/entity"
)

func AdaptUserDtoToUserDomain(user dto.User) (entity.User, error) {
	id := user.ID
	name := user.Name
	username := user.Username
	email := user.Email
	password := user.Password
	federationId := user.FederationId
	profilePhotoUrl := user.ProfileAvatarUrl

	idVo, err := vo.NewIdFromValue(id)
	if err != nil {
		return nil, err
	}
	nameVo, err := vo.NewName(name)
	if err != nil {
		return nil, err
	}
	emailVo, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}
	usernameVo, err := vo.NewName(username)
	if err != nil {
		return nil, err
	}
	passwordVo, err := vo.NewPasswordFromValue(password)
	if err != nil {
		return nil, err
	}
	federationIdVo, err := vo.NewIdFromValue(federationId)
	if err != nil {
		return nil, err
	}
	var profilePhotoUrlVo vo.Url
	if profilePhotoUrl != "" {
		profilePhotoUrlVo, err = vo.NewUrlFromValue(profilePhotoUrl)
		if err != nil {
			return nil, err
		}
	}
	userEntity, err := entity.NewUser(idVo, nameVo, emailVo, usernameVo, passwordVo, federationIdVo, profilePhotoUrlVo)
	return userEntity, err
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
