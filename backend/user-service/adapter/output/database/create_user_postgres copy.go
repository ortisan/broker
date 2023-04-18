package database

import (
	"user-service/application/repository"
	"user-service/domain/entity"
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
