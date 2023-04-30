package database

import (
	"database/sql"
	"time"
)

type User struct {
	ID               string `gorm:"primaryKey"`
	Name             string
	Email            string
	Username         string
	Secret           string
	FederationId     string
	ProfileAvatarUrl sql.NullString
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (User) TableName() string {
	return "application.user"
}
