package main

import (
	"github.com/elliotcourant/gomonetary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMonetaryMetadata(t *testing.T) {
	if !monetary.LocaleSupported() {
		t.Skip("monetary metadata cannot be generated on the current platform")
	}
	metaData := GetMonetaryMetadata()
	assert.NotEmpty(t, metaData)
}