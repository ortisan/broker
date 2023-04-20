package usecase

import (
	errApp "user-service/application/error"
	"user-service/application/repository"
	"user-service/domain/entity"

	log "github.com/sirupsen/logrus"
)

type CreateUser interface {
	CreateUser(user entity.User) (entity.User, error)
}

type createUser struct {
	createUserRepository repository.CreateUser
}

func (cu createUser) CreateUser(user entity.User) (entity.User, error) {
	return cu.createUserRepository.Create(user)
}

func NewCreateUserUseCase(createUserRepository repository.CreateUser) (CreateUser, error) {
	log.WithFields(log.Fields{
		"layer": "usecase",
	}).Info("Creating user...")

	if createUserRepository == nil {
		return nil, errApp.NewBadArgumentError("createUserUseCase is required")
	}
	return &createUser{
		createUserRepository: createUserRepository,
	}, nil
}
