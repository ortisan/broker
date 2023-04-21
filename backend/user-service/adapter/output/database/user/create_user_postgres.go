package database

import (
	"errors"
	"user-service/application/repository"
	"user-service/domain/entity"
	"user-service/infrastructure/log"

	"gorm.io/gorm"
)

type CreateUserPostgresRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func (cug *CreateUserPostgresRepository) Create(user entity.User) (entity.User, error) {
	cug.logger.Infof("Creating user %v", user)
	result := cug.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func NewCreateUserPostgresRepository(db *gorm.DB) (repository.CreateUser, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	return &CreateUserPostgresRepository{db: db}, nil
}
