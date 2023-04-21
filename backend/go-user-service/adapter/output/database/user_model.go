package database

import "database/sql"

type User struct {
	ID               string `gorm:"primaryKey"`
	Name             string
	Email            string
	FederationId     string
	ProfileAvatarUrl sql.NullString
}
