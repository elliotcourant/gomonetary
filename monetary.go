package monetary

import (
	"fmt"
	"strings"
	"sync"
)

var (
	defaultLocale = "C"
	localeSync    = new(sync.Mutex)
)

func GetDefaultLocale() string {
	localeSync.Lock()
	defer localeSync.Unlock()
	return defaultLocale
}

func SetDefaultLocale(locale string) error {
	_, ok := metaData[strings.ToLower(locale)]
	if !ok {
		return fmt.Errorf("locale [%s] is not valid or is not supported", locale)
	}
	localeSync.Lock()
	defer localeSync.Unlock()
	defaultLocale = locale
	return nil
}

func ParseDefault(value string) (float64, error) {
	return Parse(value, GetDefaultLocale())
}

func FormatDefault(value float64) (string, error) {
	return Format(value, GetDefaultLocale())
}

func Parse(value, locale string) (float64, error) {
	meta, ok := metaData[strings.ToLower(locale)]
	if !ok {
		return 0, fmt.Errorf("locale [%s] is not valid or is not supported", locale)
	}
	return parseMoney(meta, value)
}

func Format(value float64, locale string) (string, error) {
	meta, ok := metaData[strings.ToLower(locale)]
	if !ok {
		return "", fmt.Errorf("locale [%s] is not valid or is not supported", locale)
	}
	return formatMoney(meta, value)
}
