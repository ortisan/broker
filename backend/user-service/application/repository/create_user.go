package repository

import "monolith/domain/entity"

type CreateUser interface {
	Create(user entity.User) (entity.User, error)
}
