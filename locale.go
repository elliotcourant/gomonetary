package monetary

import (
	"strings"
)

func GetSupportedLocales() []string {
	return localeNames
}

func GetIsSupported(name string) bool {
	_, ok := metaData[strings.ToLower(name)]
	return ok
}
