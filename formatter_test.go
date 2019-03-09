package monetary

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFormatterSimple(t *testing.T) {
	result, err := formatMoney(metaData["en_us"], 12.56)
	assert.NoError(t, err)
	assert.Equal(t, "$12.56", result)
}

func TestFormatterSimple1(t *testing.T) {
	result, err := formatMoney(metaData["en_us"], 132142112.56)
	assert.NoError(t, err)
	assert.Equal(t, "$132,142,112.56", result)
}

func TestFormatterSimple2(t *testing.T) {
	result, _ := formatMoney(metaData["es_es"], 12.56)
	fmt.Printf("%s\n", result)
}

func TestFormatter(t *testing.T) {
	for _, name := range localeNames {
		t.Run(fmt.Sprintf("Locale:%s", name), func(t *testing.T) {
			var result string
			locale := metaData[strings.ToLower(string(name))]
			input := 0.0
			for i := 0; i < 2; i++ {
				name := "Positive"
				if input = 123459687.12; i == 1 {
					input = -input
					name = "Negative"
				}
				t.Run(name, func(t *testing.T) {
					t.Run("Format", func(t *testing.T) {
						res, err := formatMoney(locale, input)
						assert.NoError(t, err)
						assert.NotEmpty(t, res)
						result = res
					})
					t.Run("Parse", func(t *testing.T) {
						parsed, err := parseMoney(locale, result)
						assert.NoError(t, err, "[%s] failed", locale.Locale)
						assert.Equal(t, input, parsed, "[%s] does not match", locale.Locale)
					})
				})
			}
		})
	}
}
