package repository

import (
	"user-service/domain/entity"
	"user-service/domain/vo"
)

type GetUser interface {
	GetById(vo.Id) (entity.User, error)
}
