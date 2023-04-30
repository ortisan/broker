package database

import (
	"context"
	"errors"
	"fmt"
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-commons/infrastructure/log"
	"strings"

	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-user-service/domain/entity"
	"ortisan-broker/go-user-service/domain/repository"

	"gorm.io/gorm"
)

const errUniqueConstraint = "ERROR: duplicate key value violates unique constraint"

type CreateUserPostgresRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func (cug *CreateUserPostgresRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	cug.logger.Infof("Creating user %v", user)
	userModel, err := AdaptUserEntityToUserModel(user)
	if err != nil {
		return nil, err
	}
	result := cug.db.Create(&userModel)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), errUniqueConstraint) {
			return nil, errApp.NewConflictErrorWithCause("there is another with same data.", result.Error)
		}
		return nil, result.Error
	}
	return user, nil
}

func NewCreateUserPostgresRepository(db *gorm.DB, logger log.Logger) (repository.CreateUser, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	if logger == nil {
		return nil, errors.New("logger is required")
	}
	return &CreateUserPostgresRepository{db: db, logger: logger}, nil
}

type getUserPostgresRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func (gup *getUserPostgresRepository) GetById(ctx context.Context, id vo.Id) (entity.User, error) {
	gup.logger.Infof("Getting user by id: %v", id.Value())
	var user User
	gup.db.Where("id = ?", id.Value()).Find(&user)
	if user.ID == "" {
		return nil, errApp.NewNotFoundError(fmt.Sprintf("User not found for id %s", id.Value()))
	}
	userEntity, err := AdaptUserModelToUserDomain(&user)
	if err != nil {
		return nil, errApp.NewBaseErrorWithCause("error to parse assemble user model", err)
	}
	return userEntity, nil
}

func NewGetUserPostgresRepository(db *gorm.DB, logger log.Logger) (repository.GetUser, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	if logger == nil {
		return nil, errors.New("logger is required")
	}
	return &getUserPostgresRepository{db: db, logger: logger}, nil
}
