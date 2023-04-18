package repository

import (
	"monolith/domain/entity"
	"monolith/domain/vo"
)

type GetUser interface {
	GetById(vo.Id) (entity.User, error)
}
