package datastore

import (
	"errors"
	"fmt"
	"ortisan-broker/go-user-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s TimeZone=America/Sao_Paulo",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DatabaseName,
		cfg.Database.User,
		cfg.Database.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
