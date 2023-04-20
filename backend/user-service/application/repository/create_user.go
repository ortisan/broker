package repository

import "user-service/domain/entity"

type CreateUser interface {
	Create(user entity.User) (entity.User, error)
}
