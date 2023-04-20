package entity

import (
	"testing"
	"user-service/domain/vo"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	id := vo.NewId()
	name, _ := vo.NewName("name")
	email, _ := vo.NewEmail("test@test.com")
	password, _ := vo.NewPasswordFromValue("password")
	federationId := vo.NewId()
	profilePhotoUrl, _ := vo.NewUrlFromValue("http://my-profile.com/123")

	t.Run("when valid input params should create valid user", func(t *testing.T) {
		user, err := NewUser(id, name, email, password, federationId, profilePhotoUrl)
		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("when nil id should return error", func(t *testing.T) {
		user, err := NewUser(nil, name, email, password, federationId, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("when nil name should return error", func(t *testing.T) {
		user, err := NewUser(id, nil, email, password, federationId, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("when nil email should return error", func(t *testing.T) {
		user, err := NewUser(id, name, nil, password, federationId, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("when nil password should return error", func(t *testing.T) {
		user, err := NewUser(id, name, email, nil, federationId, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("when nil federationId should return error", func(t *testing.T) {
		user, err := NewUser(id, name, email, password, nil, profilePhotoUrl)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})
}
