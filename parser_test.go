package monetary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMoneySimple(t *testing.T) {
	t.Run("C", func(t *testing.T) {
		result := parseMoney(metaData["C"], `5.51`)
		assert.Equal(t, 5.51, result)
	})
	t.Run("POSIX", func(t *testing.T) {
		result := parseMoney(metaData["POSIX"], `5.51`)
		assert.Equal(t, 5.51, result)
	})
	t.Run("en_US", func(t *testing.T) {
		result := parseMoney(metaData["en_US"], `$5.51`)
		assert.Equal(t, 5.51, result)
	})
	t.Run("ru_RU", func(t *testing.T) {
		result := parseMoney(metaData["ru_RU"], `5,54 руб.`)
		assert.Equal(t, 5.54, result)
	})
}
