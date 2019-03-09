package monetary

import (
	"fmt"
)

func formatMoney(locale localeInfo, input float64) (string, error) {
	value, err := floatToInt(input)
	if err != nil {
		return "", err
	}

	var buf []byte

	var points int8
	var grouping []int8
	var decimalSymbol byte
	var thousandsSymbol []byte
	var currencySymbol []byte
	var signSymbol []byte

	var signPosition int
	var currencySymbolPrecedes bool
	var separatedBySpace int8

	points = int8(locale.FractionalDigits)
	if points < 0 || points > 10 {
		points = 2 /* best guess in this case, I think */
	}

	// grouping here
	if locale.Grouping != nil &&
		len(locale.Grouping) > 0 {
		grouping = locale.Grouping
	} else {
		grouping = []int8{3}
	}

	/* we restrict dsymbol to be a single byte, but not the other symbols */
	if locale.DecimalPoint != 0 {
		decimalSymbol = locale.DecimalPoint
	} else {
		decimalSymbol = Decimal[0]
	}

	if locale.ThousandsSeparator != nil &&
		len(locale.ThousandsSeparator) > 0 &&
		!sliceEquals([]byte{decimalSymbol}, locale.ThousandsSeparator) {
		thousandsSymbol = locale.ThousandsSeparator
	} else if sliceEquals([]byte{decimalSymbol}, Comma) {
		thousandsSymbol = Decimal
	} else {
		thousandsSymbol = Comma
	}

	if locale.CurrencySymbol != nil &&
		len(locale.CurrencySymbol) > 0 {
		currencySymbol = locale.CurrencySymbol
	} else {
		currencySymbol = Dollar
	}

	if value < 0 {
		/* make the amount positive for digit-reconstruction loop */
		value = -value

		/* set up formatting data */
		if locale.NegativeSign != nil &&
			len(locale.NegativeSign) > 0 {
			signSymbol = locale.NegativeSign
		} else {
			signSymbol = Negative
		}
		signPosition = locale.NegativeSignPosition
		currencySymbolPrecedes = locale.NegativeCurrencySymbolPrecedes
		separatedBySpace = locale.NegativeSeparatedBySpace
	} else {
		signSymbol = locale.PositiveSign
		signPosition = locale.PositiveSignPosition
		currencySymbolPrecedes = locale.PositiveCurrencySymbolPrecedes
		separatedBySpace = locale.PositiveSeparatedBySpace
	}

	digitPosition := points
	for value > 0 || digitPosition >= 0 {
		g := digitPosition % grouping[0]
		if points > 0 && digitPosition == 0 {
			buf = append([]byte{decimalSymbol}, buf...)
		} else if digitPosition < 0 && g == 0 {
			if len(grouping) > 1 {
				prevGroup := grouping[0]
				grouping = grouping[1:]
				if grouping[0] == -1 {
					grouping[0] = prevGroup
				}
			}
			buf = append(thousandsSymbol, buf...)
		}
		buf = append([]byte{uint8((value % 10) + '0')}, buf...)
		value = int64(value / 10)
		digitPosition--
	}

	result, monetarySpace, signSpace := "", "", ""
	if separatedBySpace > 0 {
		monetarySpace = " "
	}
	if separatedBySpace == 2 {
		signSpace = " "
	}
	switch signPosition {
	case 0:
		if currencySymbolPrecedes {
			result = fmt.Sprintf("(%s%s%s)",
				currencySymbol, monetarySpace, buf)
		} else {
			result = fmt.Sprintf("(%s%s%s)",
				buf, monetarySpace, currencySymbol)
		}
	case 1:
		fallthrough
	default:
		if currencySymbolPrecedes {
			result = fmt.Sprintf("%s%s%s%s%s",
				signSymbol, signSpace, currencySymbol, monetarySpace, buf)
		} else {
			result = fmt.Sprintf("%s%s%s%s%s",
				signSymbol, signSpace, buf, monetarySpace, currencySymbol)
		}
	case 2:
		if currencySymbolPrecedes {
			result = fmt.Sprintf("%s%s%s%s%s",
				currencySymbol, monetarySpace, buf, signSpace, signSymbol)
		} else {
			result = fmt.Sprintf("%s%s%s%s%s",
				buf, monetarySpace, currencySymbol, signSpace, signSymbol)
		}
	case 3:
		if currencySymbolPrecedes {
			result = fmt.Sprintf("%s%s%s%s%s",
				signSymbol, signSpace, currencySymbol, monetarySpace, buf)
		} else {
			result = fmt.Sprintf("%s%s%s%s%s",
				buf, monetarySpace, signSymbol, signSpace, currencySymbol)
		}
	case 4:
		if currencySymbolPrecedes {
			result = fmt.Sprintf("%s%s%s%s%s",
				currencySymbol, signSpace, signSymbol, monetarySpace, buf)
		} else {
			result = fmt.Sprintf("%s%s%s%s%s",
				buf, monetarySpace, currencySymbol, signSpace, signSymbol)
		}
	}

	return result, nil
}
