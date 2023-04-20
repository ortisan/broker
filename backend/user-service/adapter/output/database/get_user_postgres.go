package database

import (
	"errors"
	"user-service/application/repository"
	"user-service/domain/entity"
	"user-service/domain/vo"

	"gorm.io/gorm"
)

type getUserPostgresRepository struct {
	db *gorm.DB
}

func (gup *getUserPostgresRepository) GetById(vo.Id) (entity.User, error) {
	return nil, nil
}

func NewGetUserPostgresRepository(db *gorm.DB) (repository.GetUser, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}

	return &getUserPostgresRepository{db: db}, nil
}
