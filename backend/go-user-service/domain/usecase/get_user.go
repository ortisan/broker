package usecase

import (
	"context"
	"errors"
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-user-service/domain/entity"
	"ortisan-broker/go-user-service/domain/repository"

	log "github.com/sirupsen/logrus"
)

type GetUser interface {
	GetUserById(ctx context.Context, id vo.Id) (entity.User, error)
}

type getUser struct {
	getUserRepository repository.GetUser
}

func (cu getUser) GetUserById(ctx context.Context, id vo.Id) (entity.User, error) {
	return cu.getUserRepository.GetById(ctx, id)
}

func NewGetUserUseCase(getUserRepository repository.GetUser) (GetUser, error) {
	log.WithFields(log.Fields{
		"layer": "usecase",
	}).Info("Getting user...")

	if getUserRepository == nil {
		return nil, errors.New("getUserRepository is required")
	}
	return &getUser{
		getUserRepository: getUserRepository,
	}, nil
}
