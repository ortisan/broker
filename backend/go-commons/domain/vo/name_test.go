package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	t.Run("when name is valid, should create vo", func(t *testing.T) {
		name, err := NewName("name")
		assert.Nil(t, err)
		assert.Equal(t, "name", name.Value())
	})

	t.Run("when name is empty, must return error", func(t *testing.T) {
		id, err := NewName("")
		assert.NotNil(t, err)
		assert.Nil(t, id)
	})
}
