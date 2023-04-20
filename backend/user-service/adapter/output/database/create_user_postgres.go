package database

import (
	"errors"
	"user-service/application/repository"
	"user-service/domain/entity"

	"gorm.io/gorm"
)

type CreateUserPostgresRepository struct {
	db *gorm.DB
}

func (cug *CreateUserPostgresRepository) Create(user entity.User) (entity.User, error) {
	return nil, nil
}

func NewCreateUserPostgresRepository(db *gorm.DB) (repository.CreateUser, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	return &CreateUserPostgresRepository{db: db}, nil
}
