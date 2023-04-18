package infrastructure

import (
	"fmt"
	"log"
	"user-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password %s", config.C.Database.Host, config.C.Database.Port, config.C.Database.DatabaseName, config.C.Database.User, config.C.Database.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
