package monetary

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseMoneySimple(t *testing.T) {
	values := []struct {
		Locale   string
		Input    string
		Expected float64
	}{
		{
			Locale:   "C",
			Input:    "5.51",
			Expected: 5.51,
		},
		{
			Locale:   "POSIX",
			Input:    "12.98",
			Expected: 12.98,
		},
		{
			Locale:   "en_US",
			Input:    "$489.54",
			Expected: 489.54,
		},
		{
			Locale:   "en_US",
			Input:    "$10,589.56",
			Expected: 10589.56,
		},
		{
			Locale:   "en_US",
			Input:    "$(10,589.56)",
			Expected: -10589.56,
		},
		{
			Locale:   "ru_RU",
			Input:    "43,32 руб.",
			Expected: 43.32,
		},
	}
	for _, val := range values {
		t.Run(fmt.Sprintf("Locale: [%s]", val.Locale), func(t *testing.T) {
			result, err := parseMoney(metaData[strings.ToLower(val.Locale)], val.Input)
			assert.NoError(t, err, "no error should be returned")
			assert.Equal(t, val.Expected, result, "should match expected")
		})
	}
}

func TestParseMoneyOverflow(t *testing.T) {
	values := []string{
		`2491275678378686887185.51`,
		`-2491278912378686887185.51`,
	}
	for _, val := range values {
		t.Run(fmt.Sprintf("Overflow: [%s]", val), func(t *testing.T) {
			result, err := parseMoney(metaData["c"], val)
			assert.Equal(t, float64(0), result, "error result should be 0")
			assert.EqualError(t, err, fmt.Sprintf(`value "%s" is out of range for type money`, val))
		})
	}
}
