package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl(t *testing.T) {
	urlStr := "google.com"
	t.Run("when url is valid, should create vo", func(t *testing.T) {
		url, err := NewUrlFromValue(urlStr)
		assert.Nil(t, err)
		assert.NotEqual(t, urlStr, url.Value())
	})

	t.Run("when url is empty, must return error", func(t *testing.T) {
		id, err := NewUrlFromValue("")
		assert.NotNil(t, err)
		assert.Nil(t, id)
	})

	t.Run("when url is not valid, must return error", func(t *testing.T) {
		id, err := NewUrlFromValue("url")
		assert.NotNil(t, err)
		assert.Nil(t, id)
	})
}
