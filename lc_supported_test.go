// +build darwin linux

package monetary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocaleSupported(t *testing.T) {
	supported := LocaleSupported()
	assert.True(t, supported)
}
