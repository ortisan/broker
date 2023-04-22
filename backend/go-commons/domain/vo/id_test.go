package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestId(t *testing.T) {
	t.Run("when id is valid, should create vo", func(t *testing.T) {
		id, err := NewIdFromValue("id")
		assert.Nil(t, err)
		assert.Equal(t, "id", id.Value())
	})

	t.Run("when newId is invoke, must return vo", func(t *testing.T) {
		id := NewId()
		assert.NotNil(t, id)
	})

	t.Run("when id is empty, must return error", func(t *testing.T) {
		id, err := NewIdFromValue("")
		assert.NotNil(t, err)
		assert.Nil(t, id)
	})
}
