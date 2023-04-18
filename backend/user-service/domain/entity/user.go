package entity

import (
	errApp "monolith/application/error"
	"monolith/domain/vo"
)

type User interface {
	Id() vo.Id
	Name() vo.Name
	Email() vo.Email
	Password() vo.Password
	FederationId() vo.Id
	ProfilePhotoUrl() vo.Url
}

type user struct {
	id              vo.Id
	name            vo.Name
	email           vo.Email
	password        vo.Password
	federationId    vo.Id
	profilePhotoUrl vo.Url
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
func (u *user) Password() vo.Password {
	return u.password
}
func (u *user) FederationId() vo.Id {
	return u.federationId
}
func (u *user) ProfilePhotoUrl() vo.Url {
	return u.profilePhotoUrl
}

func NewUser(id vo.Id,
	name vo.Name,
	email vo.Email,
	password vo.Password,
	federationId vo.Id,
	profilePhotoUrl vo.Url) (User, error) {
	if name == nil {
		return nil, errApp.NewBadArgumentError("name is required")
	}
	if email == nil {
		return nil, errApp.NewBadArgumentError("email is required")
	}
	if password == nil {
		return nil, errApp.NewBadArgumentError("password is required")
	}
	if federationId == nil {
		return nil, errApp.NewBadArgumentError("federation id is required")
	}
	return &user{
		id:              id,
		name:            name,
		email:           email,
		password:        password,
		federationId:    federationId,
		profilePhotoUrl: profilePhotoUrl,
	}, nil
}
