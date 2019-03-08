package monetary

import (
	"fmt"
	"strings"
)

func Parse(moneyString, locale string) (float64, error) {
	meta, ok := metaData[strings.ToLower(locale)]
	if !ok {
		return 0, fmt.Errorf("locale [%s] is not valid or is not supported", locale)
	}
	return parseMoney(meta, moneyString)
}

func ParseBytes(moneyBytes, locale string) (float64, error) {
	return Parse(string(moneyBytes), locale)
}
