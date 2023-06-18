package database

import (
	"context"
	"errors"
	"ortisan-broker/go-commons/domain/vo"
	errapp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-commons/infrastructure/log"
	"ortisan-broker/go-sts-service/domain/entity"
	"ortisan-broker/go-sts-service/domain/repository"

	"gorm.io/gorm"
)

type clientCredentialsPostgresRepository struct {
	db      *gorm.DB
	logger  log.Logger
	adapter ClientCredentialsAdapter
}

func (ccr clientCredentialsPostgresRepository) CreateClientCredentials(ctx context.Context, clientCredentials entity.ClientCredentials) (entity.ClientCredentials, error) {
	if ctx == nil {
		return nil, errapp.NewBadArgumentError("context is required")
	}
	if clientCredentials == nil {
		return nil, errapp.NewBadArgumentError("client id is required")
	}

	clientCredentialsModel, err := ccr.adapter.AdaptFromDomainToModel(ctx, clientCredentials)
	if err != nil {
		return nil, err
	}

	if err := ccr.db.Create(clientCredentialsModel).Error; err != nil {
		return nil, err
	}

	return ccr.adapter.AdaptFromModelToDomain(ctx, clientCredentialsModel)
}

func (ccr clientCredentialsPostgresRepository) FindByClientId(ctx context.Context, clientId vo.Id) (entity.ClientCredentials, error) {
	if ctx == nil {
		return nil, errapp.NewBadArgumentError("context is required")
	}
	if clientId == nil {
		return nil, errapp.NewBadArgumentError("client id is required")
	}

	var clientCredentials ClientCredentials
	err := ccr.db.Where("client_id = ?", clientId).First(&clientCredentials).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errapp.NewNotFoundError("client credentials not found")
	}

	return ccr.adapter.AdaptFromModelToDomain(ctx, &clientCredentials)
}

func (ccr clientCredentialsPostgresRepository) FindByClientName(ctx context.Context, clientName vo.Name) (entity.ClientCredentials, error) {
	var clientCredentials ClientCredentials
	err := ccr.db.Where("client_name = ?", clientName).First(&clientCredentials).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return ccr.adapter.AdaptFromModelToDomain(ctx, &clientCredentials)
}

func NewClientCredentialsPostgresRepository(db *gorm.DB, logger log.Logger, adapter ClientCredentialsAdapter) (repository.ClientCredentialsRepository, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	if logger == nil {
		return nil, errors.New("logger is required")
	}
	if adapter == nil {
		return nil, errors.New("client credentials adapter is required")
	}

	return &clientCredentialsPostgresRepository{
		adapter: adapter,
	}, nil
}
