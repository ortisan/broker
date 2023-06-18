package usecase

import (
	"context"
	errapp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-user-service/domain/entity"
	"ortisan-broker/go-user-service/domain/repository"

	log "github.com/sirupsen/logrus"
)

type CreateUser interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
}

type createUser struct {
	createUserRepository repository.CreateUser
}

func (cu createUser) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	return cu.createUserRepository.Create(ctx, user)
}

func NewCreateUserUseCase(createUserRepository repository.CreateUser) (CreateUser, error) {
	log.WithFields(log.Fields{
		"layer": "usecase",
	}).Info("Creating user...")

	if createUserRepository == nil {
		return nil, errapp.NewBadArgumentError("createUserUseCase is required")
	}
	return &createUser{
		createUserRepository: createUserRepository,
	}, nil
}
