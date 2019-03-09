package monetary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSupportedLocales(t *testing.T) {
	locales := GetSupportedLocales()
	assert.NotEmpty(t, locales, "there should always be locales")
}

func TestGetIsSupported(t *testing.T) {
	t.Run("Supported", func(t *testing.T) {
		assert.True(t, GetIsSupported("en_US"))
		assert.True(t, GetIsSupported("en_us")) // Case insensitivity
	})
	t.Run("Unsupported", func(t *testing.T) {
		assert.False(t, GetIsSupported("bob the builder"))
	})
}
