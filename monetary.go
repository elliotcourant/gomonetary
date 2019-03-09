package monetary

import (
	"fmt"
	"strings"
)

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
