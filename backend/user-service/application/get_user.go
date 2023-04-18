package application

import (
	"monolith/adapter/dto"
	errApp "monolith/application/error"
	"monolith/domain/usecase"
	"monolith/domain/vo"
)

type GetUserApplication interface {
	GetUser(userId string) (dto.User, error)
}

type getUserApplication struct {
	getUserUseCase usecase.GetUser
}

func (gua *getUserApplication) GetUser(userId string) (dto.User, error) {
	id, errId := vo.NewIdFromValue(userId)
	if errId != nil {
		return dto.User{}, errId
	}
	user, errUseCase := gua.getUserUseCase.GetUserById(id)
	if errUseCase != nil {
		return dto.User{}, errUseCase
	}
	userDto, errAdaptDto := AdaptUserDomainToUserDto(user)
	if errAdaptDto != nil {
		return dto.User{}, errAdaptDto
	}
	return userDto, nil
}

func NewGetUserApplication(getUserUseCase usecase.GetUser) (GetUserApplication, error) {
	if getUserUseCase == nil {
		return nil, errApp.NewBadArgumentError("getUserUseCase is required")
	}
	return &getUserApplication{
		getUserUseCase: getUserUseCase,
	}, nil
}
