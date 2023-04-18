package database

import (
	"monolith/application/repository"
	"monolith/domain/entity"
)

type CreateUserPostgresRepository struct {
	users []entity.User
}

func (cug *CreateUserPostgresRepository) Create(user entity.User) (entity.User, error) {
	cug.users = append(cug.users, user)
	return user, nil
}

func NewCreateUserPostgresRepository() repository.CreateUser {
	return &CreateUserPostgresRepository{}
}
