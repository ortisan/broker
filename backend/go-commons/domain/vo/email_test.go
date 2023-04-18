package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	const email = "teste@gmail.com"
	t.Run("when email is valid, should return email vo", func(t *testing.T) {
		e, err := NewEmail(email)
		assert.Nil(t, err)
		assert.Equal(t, email, e.Value())
	})

	t.Run("when email is empty, should return error", func(t *testing.T) {
		e, err := NewEmail("")
		assert.NotNil(t, err)
		assert.Nil(t, e)
	})

	t.Run("when email is invalid, should return error", func(t *testing.T) {
		e, err := NewEmail("email")
		assert.NotNil(t, err)
		assert.Nil(t, e)
	})

}
