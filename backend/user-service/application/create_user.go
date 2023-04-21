package application

import (
	"user-service/adapter/dto"
	errApp "user-service/application/error"
	"user-service/domain/usecase"
	"user-service/domain/vo"
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
