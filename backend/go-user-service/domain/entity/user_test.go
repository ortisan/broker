package entity

import (
	"ortisan-broker/go-commons/domain/vo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	id := vo.NewId()
	name, _ := vo.NewName("name")
	email, _ := vo.NewEmail("test@test.com")
	username, _ := vo.NewName("username")
	secret, _ := vo.NewSecretFromValue("secret")
	federationId := vo.NewId()
	profilePhotoUrl, _ := vo.NewUrlFromValue("http://my-profile.com/123")

	t.Run("when valid input params should create valid user", func(t *testing.T) {
		user, err := NewUser(id, name, email, username, secret, federationId, profilePhotoUrl)
		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("when nil id should return error", func(t *testing.T) {
		user, err := NewUser(nil, name, email, username, secret, federationId, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("when nil name should return error", func(t *testing.T) {
		user, err := NewUser(id, nil, email, username, secret, federationId, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("when nil email should return error", func(t *testing.T) {
		user, err := NewUser(id, name, nil, username, secret, federationId, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("when nil username should return error", func(t *testing.T) {
		user, err := NewUser(id, name, email, nil, secret, federationId, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("when nil secret should return error", func(t *testing.T) {
		user, err := NewUser(id, name, email, username, nil, federationId, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("when nil federationId should return error", func(t *testing.T) {
		user, err := NewUser(id, name, email, username, secret, nil, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})
}
