package database

import (
	"monolith/application/repository"
	"monolith/domain/entity"
	"monolith/domain/vo"
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
