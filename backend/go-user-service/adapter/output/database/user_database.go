package database

import (
	"errors"
	"fmt"
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-commons/infrastructure/log"

	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-user-service/domain/entity"
	"ortisan-broker/go-user-service/domain/repository"

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

type getUserPostgresRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func (gup *getUserPostgresRepository) GetById(id vo.Id) (entity.User, error) {
	gup.logger.Infof("Getting user by id: %v", id.Value())
	var user entity.User
	gup.db.First(user, id.Value())
	if user == nil {
		return nil, errApp.NewNotFoundError(fmt.Sprintf("User not found for id %s", id.Value()))
	}
	return user, nil
}

func NewGetUserPostgresRepository(db *gorm.DB) (repository.GetUser, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	return &getUserPostgresRepository{db: db}, nil
}
