package database

import (
	"database/sql"
	"ortisan-broker/go-commons/domain/vo"
	"ortisan-broker/go-user-service/domain/entity"
	"time"
)

func AdaptUserEntityToUserModel(user entity.User) (*User, error) {
	var profileAvatarUrl sql.NullString
	if user.ProfileAvatarUrl() != nil {
		profileAvatarUrl.Valid = true
		profileAvatarUrl.String = user.ProfileAvatarUrl().Value()
	}
	return &User{
		ID:               user.Id().Value(),
		Name:             user.Name().Value(),
		Email:            user.Email().Value(),
		Username:         user.Username().Value(),
		Secret:           user.Secret().Value(),
		FederationId:     user.FederationId().Value(),
		ProfileAvatarUrl: profileAvatarUrl,
		CreatedAt:        time.Now(),
	}, nil
}

func AdaptUserModelToUserDomain(user *User) (entity.User, error) {
	id, err := vo.NewIdFromValue(user.ID)
	if err != nil {
		return nil, err
	}
	name, err := vo.NewName(user.Name)
	if err != nil {
		return nil, err
	}
	email, err := vo.NewEmail(user.Email)
	if err != nil {
		return nil, err
	}
	username, err := vo.NewName(user.Username)
	if err != nil {
		return nil, err
	}
	password, err := vo.NewSecretFromValueCrypted(user.Secret)
	if err != nil {
		return nil, err
	}
	federationId, err := vo.NewIdFromValue(user.FederationId)
	if err != nil {
		return nil, err
	}
	var profileAvatarUrl vo.Url
	if user.ProfileAvatarUrl.Valid {
		profileAvatarUrl, err = vo.NewUrlFromValue(user.ProfileAvatarUrl.String)
		if err != nil {
			return nil, err
		}
	}
	return entity.NewUser(id, name, email, username, password, federationId, profileAvatarUrl)
}
