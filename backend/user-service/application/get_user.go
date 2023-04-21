package application

import (
	"user-service/adapter/dto"
	errApp "user-service/application/error"
	"user-service/domain/usecase"
	"user-service/domain/vo"
)

type GetUserApplication interface {
	GetUser(userId string) (*dto.User, error)
}

type getUserApplication struct {
	getUserUseCase usecase.GetUser
}

func (gua *getUserApplication) GetUser(userId string) (*dto.User, error) {
	id, errId := vo.NewIdFromValue(userId)
	if errId != nil {
		return nil, errId
	}
	user, errUseCase := gua.getUserUseCase.GetUserById(id)
	if errUseCase != nil {
		return nil, errUseCase
	}
	userDto, errAdaptDto := AdaptUserDomainToUserDto(user)
	if errAdaptDto != nil {
		return nil, errAdaptDto
	}
	return &userDto, nil
}

func NewGetUserApplication(getUserUseCase usecase.GetUser) (GetUserApplication, error) {
	if getUserUseCase == nil {
		return nil, errApp.NewBadArgumentError("getUserUseCase is required")
	}
	return &getUserApplication{
		getUserUseCase: getUserUseCase,
	}, nil
}
