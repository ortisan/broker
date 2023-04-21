package repository

import (
	"ortisan-broker/go-user-service/domain/entity"
	"ortisan-broker/go-user-service/domain/vo"
)

type GetUser interface {
	GetById(vo.Id) (entity.User, error)
}

type CreateUser interface {
	Create(user entity.User) (entity.User, error)
}
