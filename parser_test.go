package monetary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMoney(t *testing.T) {
	t.Run("en_US", func(t *testing.T) {
		result := parseMoney(metaData["en_US"], `$5.51`)
		assert.Equal(t, 5.51, result)
	})
	t.Run("ru_RU", func(t *testing.T) {
		result := parseMoney(metaData["ru_RU"], `5,54 руб.`)
		assert.Equal(t, 5.54, result)
	})
}
