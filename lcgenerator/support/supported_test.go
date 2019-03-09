// +build darwin linux

package support

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocaleSupported(t *testing.T) {
	supported := LocaleSupported()
	assert.True(t, supported)
}
