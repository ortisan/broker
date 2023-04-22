package database

import "time"

type ClientCredentials struct {
	ClientName   string `gorm:"primarykey"`
	ClientId     string `gorm:"not null"`
	ClientSecret string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (ClientCredentials) TableName() string {
	return "application.client_credentials"
}
