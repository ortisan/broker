package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecret(t *testing.T) {
	t.Run("when secret is valid, should create vo", func(t *testing.T) {
		secret, err := NewSecretFromValue("secret")
		assert.Nil(t, err)
		assert.NotEqual(t, "secret", secret.Value())
	})

	t.Run("when secret is empty, must return error", func(t *testing.T) {
		id, err := NewSecretFromValue("")
		assert.NotNil(t, err)
		assert.Nil(t, id)
	})
}
