package monetary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSupportedLocales(t *testing.T) {
	locales := GetSupportedLocales()
	assert.NotEmpty(t, locales, "there should always be locales")
}
