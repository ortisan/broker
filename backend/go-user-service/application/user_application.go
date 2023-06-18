package application

import (
	"context"
	"errors"
	"ortisan-broker/go-commons/domain/vo"
	errapp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-user-service/adapter/dto"
	"ortisan-broker/go-user-service/domain/usecase"
)

type CreateUserApplication interface {
	CreateUser(ctx context.Context, user dto.User) (*dto.User, error)
}

type createUserApplication struct {
	usecase usecase.CreateUser
}

func (cua *createUserApplication) CreateUser(ctx context.Context, user dto.User) (*dto.User, error) {
	user.ID = vo.NewId().Value()
	userEntity, err := AdaptUserDtoToUserDomain(ctx, user)
	if err != nil {
		return nil, err
	}
	createdUser, err := cua.usecase.CreateUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}
	createdUserDto, err := AdaptUserDomainToUserDto(ctx, createdUser)
	if err != nil {
		return nil, err
	}
	return &createdUserDto, nil
}

func NewCreateUserApplication(usecase usecase.CreateUser) (CreateUserApplication, error) {
	if usecase == nil {
		return nil, errors.New("create user usecase is required")
	}
	return &createUserApplication{
		usecase: usecase,
	}, nil
}

type GetUserApplication interface {
	GetUser(ctx context.Context, userId string) (*dto.User, error)
}

type getUserApplication struct {
	getUserUseCase usecase.GetUser
}

func (g *getUserApplication) GetUser(ctx context.Context, userId string) (*dto.User, error) {
	id, err := vo.NewIdFromValue(userId)
	if err != nil {
		return nil, err
	}
	user, err := g.getUserUseCase.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	userDto, err := AdaptUserDomainToUserDto(ctx, user)
	if err != nil {
		return nil, err
	}
	return &userDto, nil
}

func NewGetUserApplication(getUserUseCase usecase.GetUser) (GetUserApplication, error) {
	if getUserUseCase == nil {
		return nil, errapp.NewBadArgumentError("get user use case is required")
	}
	return &getUserApplication{
		getUserUseCase: getUserUseCase,
	}, nil
}
