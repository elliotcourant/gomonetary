package monetary

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestReadMonetaryBytes(t *testing.T) {
	data, err := ioutil.ReadFile("test/locales/en_US/LC_MONETARY")
	assert.NoError(t, err, "should be able to read test LC_MONETARY successfully")
	fmt.Println(hex.Dump(data))
}
