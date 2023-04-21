package database

import (
	"errors"
	"fmt"
	errApp "user-service/application/error"
	"user-service/application/repository"
	"user-service/domain/entity"
	"user-service/domain/vo"
	"user-service/infrastructure/log"

	"gorm.io/gorm"
)

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
