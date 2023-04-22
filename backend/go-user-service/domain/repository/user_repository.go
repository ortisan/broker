package repository

import (
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-user-service/domain/entity"
)

type GetUser interface {
	GetById(vo.Id) (entity.User, error)
}

type CreateUser interface {
	Create(user entity.User) (entity.User, error)
}
