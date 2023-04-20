package application

import (
	"user-service/adapter/dto"
	"user-service/domain/entity"
	"user-service/domain/vo"
)

func AdaptUserDtoToUserDomain(user dto.User) (entity.User, error) {
	id := user.ID
	name := user.Name
	email := user.Email
	password := user.Password
	federationId := user.FederationId
	profilePhotoUrl := user.ProfilePhotoUrl

	idVo, errId := vo.NewIdFromValue(id)
	if errId != nil {
		return nil, errId
	}
	nameVo, errName := vo.NewName(name)
	if errName != nil {
		return nil, errName
	}
	emailVo, errEmail := vo.NewEmail(email)
	if errEmail != nil {
		return nil, errEmail
	}
	passwordVo, errPassword := vo.NewPasswordFromValue(password)
	if errPassword != nil {
		return nil, errPassword
	}
	federationIdVo, errFederationId := vo.NewIdFromValue(federationId)
	if errFederationId != nil {
		return nil, errFederationId
	}
	var profilePhotoUrlVo vo.Url
	var errProfilePhotoUrl error
	if profilePhotoUrl != "" {
		profilePhotoUrlVo, errProfilePhotoUrl = vo.NewUrlFromValue(profilePhotoUrl)
		if errProfilePhotoUrl != nil {
			return nil, errProfilePhotoUrl
		}
	}
	userEntity, err := entity.NewUser(idVo, nameVo, emailVo, passwordVo, federationIdVo, profilePhotoUrlVo)
	return userEntity, err
}

func AdaptUserDomainToUserDto(user entity.User) (dto.User, error) {
	var profileUrlStr string
	if user.ProfilePhotoUrl() != nil {
		profileUrlStr = user.ProfilePhotoUrl().Value()
	}

	return dto.NewUser(
		user.Id().Value(),
		user.FederationId().Value(),
		user.Name().Value(),
		user.Email().Value(),
		profileUrlStr), nil
}
