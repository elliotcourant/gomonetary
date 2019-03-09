package main

import (
	"github.com/elliotcourant/gomonetary/lcgenerator/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMonetaryMetadata(t *testing.T) {
	if !support.LocaleSupported() {
		t.Skip("monetary metadata cannot be generated on the current platform")
	}
	metaData := GetMonetaryMetadata()
	assert.NotEmpty(t, metaData)
}

func TestGenerateGo(t *testing.T) {
	if !support.LocaleSupported() {
		t.Skip("monetary metadata cannot be generated on the current platform")
	}
	metaData := GetMonetaryMetadata()
	assert.NotEmpty(t, metaData)
	result := GenerateGo(metaData)
	assert.NotEmpty(t, result)
}
