package monetary

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	values := []struct {
		Locale        string
		Input         string
		Expected      float64
		ExpectedError error
	}{
		{
			Locale:        "C",
			Input:         "5.51",
			Expected:      5.51,
			ExpectedError: nil,
		},
		{
			Locale:        "POSIX",
			Input:         "12.98",
			Expected:      12.98,
			ExpectedError: nil,
		},
		{
			Locale:        "en_US",
			Input:         "$489.54",
			Expected:      489.54,
			ExpectedError: nil,
		},
		{
			Locale:        "en_US",
			Input:         "$10,589.56",
			Expected:      10589.56,
			ExpectedError: nil,
		},
		{
			Locale:        "en_US",
			Input:         "$(10,589.56)",
			Expected:      -10589.56,
			ExpectedError: nil,
		},
		{
			Locale:        "ru_RU",
			Input:         "43,32 руб.",
			Expected:      43.32,
			ExpectedError: nil,
		},
		{
			Locale:        "bogus",
			Input:         "43,32 руб.",
			Expected:      0,
			ExpectedError: fmt.Errorf("locale [%s] is not valid or is not supported", "bogus"),
		},
		{
			Locale:        "C",
			Input:         "afhjkasdhkjas",
			Expected:      0,
			ExpectedError: fmt.Errorf(`invalid input syntax for type money: "%s"`, "afhjkasdhkjas"),
		},
	}
	for _, val := range values {
		t.Run(fmt.Sprintf("Locale: [%s]", val.Locale), func(t *testing.T) {
			result, err := Parse(val.Input, val.Locale)
			if val.ExpectedError != nil {
				assert.EqualError(t, err, val.ExpectedError.Error())
			} else {
				assert.NoError(t, err, "no error should be returned")
			}
			assert.Equal(t, val.Expected, result, "should match expected")
		})
	}
}
