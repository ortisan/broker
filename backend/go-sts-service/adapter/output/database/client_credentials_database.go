package database

import (
	"errors"
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-commons/infrastructure/log"
	"ortisan-broker/go-sts-service/domain/entity"
	"ortisan-broker/go-sts-service/domain/repository"

	"gorm.io/gorm"
)

type clientCredentialsRepository struct {
	db      *gorm.DB
	logger  log.Logger
	adapter ClientCredentialsAdapter
}

func (ccr clientCredentialsRepository) CreateClientCredentials(cr entity.ClientCredentials) (entity.ClientCredentials, error) {
	err := ccr.db.Create(cr).Error
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (ccr clientCredentialsRepository) FindByClientId(clientId vo.Id) (entity.ClientCredentials, error) {
	var clientCredentials ClientCredentials
	err := ccr.db.Where("client_id = ?", clientId).First(&clientCredentials).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &clientCredentials, err
}

func (ccr clientCredentialsRepository) FindByClientName(clientName vo.Name) (entity.ClientCredentials, error) {
	var clientCredentials ClientCredentials
	err := ccr.db.Where("client_name = ?", clientName).First(&clientCredentials).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &clientCredentials, err
}

func NewClientCredentialsRepository(adapter ClientCredentialsAdapter) (repository.ClientCredentialsRepository, error) {
	if adapter == nil {
		return nil, errors.New("client credentials adapter is required")
	}

	return &clientCredentialsRepository{
		adapter: adapter,
	}, nil
}
