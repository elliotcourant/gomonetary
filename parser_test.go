package monetary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMoney(t *testing.T) {
	russianLocale := metaData["ru_RU"]
	result := parseMoney(russianLocale, `5,50 руб.`)
	assert.Equal(t, 5.5, result)
}
