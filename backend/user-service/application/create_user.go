package application

import (
	"monolith/adapter/dto"
	errApp "monolith/application/error"
	"monolith/domain/usecase"
	"monolith/domain/vo"
)

type CreateUserApplication interface {
	CreateUser(dto.User) (dto.User, error)
}

type createUserApplication struct {
	createUserUseCase usecase.CreateUser
}

func (cua *createUserApplication) CreateUser(user dto.User) (dto.User, error) {
	user.ID = vo.NewId().Value()
	userEntity, errAdaptDomain := AdaptUserDtoToUserDomain(user)
	if errAdaptDomain != nil {
		return dto.User{}, errAdaptDomain
	}
	createdUser, errUseCase := cua.createUserUseCase.CreateUser(userEntity)
	if errUseCase != nil {
		return dto.User{}, errUseCase
	}
	createdUserDto, errAdaptDto := AdaptUserDomainToUserDto(createdUser)
	if errAdaptDto != nil {
		return dto.User{}, errAdaptDto
	}
	return createdUserDto, nil
}

func NewCreateUserApplication(createUserUseCase usecase.CreateUser) (CreateUserApplication, error) {
	if createUserUseCase == nil {
		return nil, errApp.NewBadArgumentError("createUserUseCase is required")
	}
	return &createUserApplication{
		createUserUseCase: createUserUseCase,
	}, nil
}
