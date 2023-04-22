package database

import "database/sql"

type User struct {
	ID               string `gorm:"primaryKey"`
	Name             string
	Email            string
	Username         string
	Password         string
	FederationId     string
	ProfileAvatarUrl sql.NullString
}

func (User) TableName() string {
	return "application.user"
}
