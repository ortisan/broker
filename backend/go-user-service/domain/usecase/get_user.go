package usecase

import (
	"ortisan-broker/go-commons/domain/vo"
	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-user-service/domain/entity"
	"ortisan-broker/go-user-service/domain/repository"

	log "github.com/sirupsen/logrus"
)

type GetUser interface {
	GetUserById(id vo.Id) (entity.User, error)
}

type getUser struct {
	getUserRepository repository.GetUser
}

func (cu getUser) GetUserById(id vo.Id) (entity.User, error) {
	return cu.getUserRepository.GetById(id)
}

func NewGetUserUseCase(getUserRepository repository.GetUser) (GetUser, error) {
	log.WithFields(log.Fields{
		"layer": "usecase",
	}).Info("Getting user...")

	if getUserRepository == nil {
		return nil, errApp.NewBadArgumentError("getUserRepository is required")
	}
	return &getUser{
		getUserRepository: getUserRepository,
	}, nil
}
