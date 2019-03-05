package internal

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSplitMonetaryBytes(t *testing.T) {
	data, err := ioutil.ReadFile("../test/locales/en_US/LC_MONETARY")
	assert.NoError(t, err, "should be able to read test LC_MONETARY successfully")
	assert.Equal(t, bytes.Count(data, []byte{10})+1, len(SplitMonetaryBytes(data)), "number of items in LC_MONETARY format should be greater than 1")
}
