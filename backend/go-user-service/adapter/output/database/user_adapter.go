package database

import (
	"database/sql"
	"ortisan-broker/go-user-service/domain/entity"
)

func AdaptUserToModel(user entity.User) (*User, error) {
	var profileAvatarUrl sql.NullString
	if user.ProfileAvatarUrl() != nil {
		profileAvatarUrl.Valid = true
		profileAvatarUrl.String = user.ProfileAvatarUrl().Value()
	}
	return &User{
		ID:               user.Id().Value(),
		Email:            user.Email().Value(),
		FederationId:     user.FederationId().Value(),
		ProfileAvatarUrl: profileAvatarUrl,
	}, nil
}
