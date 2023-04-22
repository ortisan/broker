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
	userEntity, err := AdaptUserDtoToUserDomain(user)
	if err != nil {
		return nil, err
	}
	createdUser, err := cua.usecase.CreateUser(userEntity)
	if err != nil {
		return nil, err
	}
	createdUserDto, err := AdaptUserDomainToUserDto(createdUser)
	if err != nil {
		return nil, err
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
	id, err := vo.NewIdFromValue(userId)
	if err != nil {
		return nil, err
	}
	user, err := gua.getUserUseCase.GetUserById(id)
	if err != nil {
		return nil, err
	}
	userDto, err := AdaptUserDomainToUserDto(user)
	if err != nil {
		return nil, err
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
