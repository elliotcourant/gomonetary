// +build windows

package support

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocaleSupported(t *testing.T) {
	supported := LocaleSupported()
	assert.False(t, supported)
}
