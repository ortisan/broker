package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	t.Run("when password is valid, should create vo", func(t *testing.T) {
		password, err := NewPasswordFromValue("password")
		assert.Nil(t, err)
		assert.NotEqual(t, "password", password.Value())
	})

	t.Run("when password is empty, must return error", func(t *testing.T) {
		id, err := NewPasswordFromValue("")
		assert.NotNil(t, err)
		assert.Nil(t, id)
	})
}
