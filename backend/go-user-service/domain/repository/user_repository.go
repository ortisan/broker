package repository

import (
	"context"
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-user-service/domain/entity"
)

type GetUser interface {
	GetById(ctx context.Context, id vo.Id) (entity.User, error)
}
type CreateUser interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
}
