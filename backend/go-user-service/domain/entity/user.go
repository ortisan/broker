package entity

import (
	"ortisan-broker/go-commons/domain/vo"
	errApp "ortisan-broker/go-commons/error"
)

type User interface {
	Id() vo.Id
	Name() vo.Name
	Email() vo.Email
	Username() vo.Name
	Password() vo.Password
	FederationId() vo.Id
	ProfileAvatarUrl() vo.Url
}

type user struct {
	id               vo.Id
	name             vo.Name
	email            vo.Email
	username         vo.Name
	password         vo.Password
	federationId     vo.Id
	profileAvatarUrl vo.Url
}

func (u *user) Id() vo.Id {
	return u.id
}
func (u *user) Name() vo.Name {
	return u.name
}
func (u *user) Email() vo.Email {
	return u.email
}
func (u *user) Username() vo.Name {
	return u.name
}
func (u *user) Password() vo.Password {
	return u.password
}
func (u *user) FederationId() vo.Id {
	return u.federationId
}
func (u *user) ProfileAvatarUrl() vo.Url {
	return u.profileAvatarUrl
}

func NewUser(id vo.Id,
	name vo.Name,
	email vo.Email,
	username vo.Name,
	password vo.Password,
	federationId vo.Id,
	profileAvatarUrl vo.Url) (User, error) {
	if id == nil {
		return nil, errApp.NewBadArgumentError("id is required")
	}
	if name == nil {
		return nil, errApp.NewBadArgumentError("name is required")
	}
	if email == nil {
		return nil, errApp.NewBadArgumentError("email is required")
	}
	if username == nil {
		return nil, errApp.NewBadArgumentError("username is required")
	}
	if password == nil {
		return nil, errApp.NewBadArgumentError("password is required")
	}
	if federationId == nil {
		return nil, errApp.NewBadArgumentError("federation id is required")
	}
	return &user{
		id:               id,
		name:             name,
		email:            email,
		username:         username,
		password:         password,
		federationId:     federationId,
		profileAvatarUrl: profileAvatarUrl,
	}, nil
}
