package database

import (
	"user-service/application/repository"
	"user-service/domain/entity"
	"user-service/domain/vo"
)

type getUserPostgresRepository struct {
	users []entity.User
}

func (gup *getUserPostgresRepository) GetById(vo.Id) (entity.User, error) {
	return nil, nil
}

func NewGetUserPostgresRepository() repository.GetUser {
	return &getUserPostgresRepository{}
}
