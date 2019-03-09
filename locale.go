package monetary

import (
	"strings"
)

type LocaleName string

type LocaleNames []LocaleName

func GetSupportedLocales() LocaleNames {
	return localeNames
}

func GetIsSupported(name string) bool {
	_, ok := metaData[strings.ToLower(name)]
	return ok
}
