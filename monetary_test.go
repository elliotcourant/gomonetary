package monetary

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDefaultLocale(t *testing.T) {
	locale := GetDefaultLocale()
	assert.Equal(t, "C", locale)
}

func TestSetDefaultLocale(t *testing.T) {
	original := GetDefaultLocale()
	newDefault := "POSIX"
	err := SetDefaultLocale(newDefault)
	assert.NoError(t, err)
	locale := GetDefaultLocale()
	assert.Equal(t, newDefault, locale)
	err = SetDefaultLocale(original)
	assert.NoError(t, err)
}

func TestFormatDefault(t *testing.T) {
	value := 12.43
	locale := GetDefaultLocale()
	formattedNormal, err := Format(value, locale)
	assert.NoError(t, err)
	formattedDefault, err := FormatDefault(value)
	assert.NoError(t, err)
	assert.Equal(t, formattedNormal, formattedDefault)
}

func TestFormatDefault1(t *testing.T) {
	value := 12.43
	formattedNormal, err := Format(value, "ja_JP")
	assert.NoError(t, err)
	parsed, err := Parse(formattedNormal, "ja_JP")
	assert.Equal(t, value, parsed)
}

func TestParseDefault(t *testing.T) {
	value := 12.43
	formattedDefault, err := FormatDefault(value)
	assert.NoError(t, err)
	parsed, err := ParseDefault(formattedDefault)
	assert.NoError(t, err)
	assert.Equal(t, value, parsed)
}

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
		t.Run(fmt.Sprintf("Locale:%s", val.Locale), func(t *testing.T) {
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

func TestFormat(t *testing.T) {
	values := []struct {
		Locale        string
		Input         float64
		Expected      string
		ExpectedError error
	}{
		{
			Locale:        "bogus",
			Input:         1421321,
			Expected:      "",
			ExpectedError: fmt.Errorf("locale [%s] is not valid or is not supported", "bogus"),
		},
		{
			Locale:        "C",
			Input:         679432670410421062173920121,
			Expected:      "",
			ExpectedError: fmt.Errorf(`strconv.ParseInt: parsing "%s": value out of range`, "67943267041042103694209843200"),
		},
	}
	for _, val := range values {
		t.Run(fmt.Sprintf("Locale:%s", val.Locale), func(t *testing.T) {
			result, err := Format(val.Input, val.Locale)
			if val.ExpectedError != nil {
				assert.EqualError(t, err, val.ExpectedError.Error())
			} else {
				assert.NoError(t, err, "no error should be returned")
			}
			assert.Equal(t, val.Expected, result, "should match expected")
		})
	}
}
