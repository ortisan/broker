package database

import (
	"errors"
	"fmt"
	"log"
	"ortisan-broker/go-commons/config"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: l,
		})
	if err != nil {
		return nil, err
	}
	return db, nil
}
