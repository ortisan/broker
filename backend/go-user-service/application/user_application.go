package application

import (
	"ortisan-broker/go-commons/domain/vo"
	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-user-service/adapter/dto"
	"ortisan-broker/go-user-service/domain/usecase"
)

type CreateUserApplication interface {
	CreateUser(dto.User) (*dto.User, error)
}

type createUserApplication struct {
	usecase usecase.CreateUser
}

func (cua *createUserApplication) CreateUser(user dto.User) (*dto.User, error) {
	user.ID = vo.NewId().Value()
	userEntity, errAdaptDomain := AdaptUserDtoToUserDomain(user)
	if errAdaptDomain != nil {
		return nil, errAdaptDomain
	}
	createdUser, errUseCase := cua.usecase.CreateUser(userEntity)
	if errUseCase != nil {
		return nil, errUseCase
	}
	createdUserDto, errAdaptDto := AdaptUserDomainToUserDto(createdUser)
	if errAdaptDto != nil {
		return nil, errAdaptDto
	}
	return &createdUserDto, nil
}

func NewCreateUserApplication(usecase usecase.CreateUser) (CreateUserApplication, error) {
	if usecase == nil {
		return nil, errApp.NewBadArgumentError("createUserUseCase is required")
	}
	return &createUserApplication{
		usecase: usecase,
	}, nil
}

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
