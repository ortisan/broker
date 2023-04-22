package database

import (
	"errors"

	"gorm.io/gorm"
)

type clientCredentialsRepository struct {
	db *gorm.DB
}

func (ccr clientCredentialsRepository) CreateClientCredentials(cr *ClientCredentials) (*ClientCredentials, error) {
	err := ccr.db.Create(cr).Error
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (ccr clientCredentialsRepository) FindByClientId(clientId string) (*ClientCredentials, error) {
	var cr ClientCredentials
	err := ccr.db.Where("client_id = ?", clientId).First(&cr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cr, err
}

func (ccr clientCredentialsRepository) FindByClientName(clientName string) (*ClientCredentials, error) {
	var cr ClientCredentials
	err := ccr.db.Where("client_name = ?", clientName).First(&cr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cr, err
}
