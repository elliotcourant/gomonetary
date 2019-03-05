package monetary

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestReadMonetaryBytesSimple(t *testing.T) {
	data, err := ioutil.ReadFile("test/locales/en_US/LC_MONETARY")
	assert.NoError(t, err, "should be able to read test LC_MONETARY successfully")
	// fmt.Println(hex.Dump(data))
	meta, err := readMonetaryBytes(data)
	assert.NoError(t, err, "reading monetary bytes should have succeeded")
	assert.Equal(t, "USD", meta.InternationalCurrencySymbol, "international currency symbol should be USD")
	assert.Equal(t, "$", meta.CurrencySymbol, "currency symbol should be $")
	assert.Equal(t, ".", meta.DecimalPoint)
	assert.Equal(t, ",", meta.ThousandsSeparator)
	assert.Equal(t, []int8{3, 3}, meta.Grouping)
}
