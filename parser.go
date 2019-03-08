package monetary

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func parseMoney(locale localeInfo, input string) (float64, error) {
	var result float64
	var value int64 = 0
	var dec int64 = 0
	var sgn = 1
	var seenDot = false
	var s = input
	var fpoint int64
	var dsymbol string
	var ssymbol string
	var csymbol string
	var psymbol string
	var nsymbol string

	/*
	 * frac_digits will be CHAR_MAX in some locales, notably C.  However, just
	 * testing for == CHAR_MAX is risky, because of compilers like gcc that
	 * "helpfully" let you alter the platform-standard definition of whether
	 * char is signed or not.  If we are so unfortunate as to get compiled
	 * with a nonstandard -fsigned-char or -funsigned-char switch, then our
	 * idea of CHAR_MAX will not agree with libc's. The safest course is not
	 * to test for CHAR_MAX at all, but to impose a range check for plausible
	 * frac_digits values.
	 */
	fpoint = int64(locale.FractionalDigits)
	if fpoint < 0 || fpoint > 10 {
		fpoint = 2 /* best guess in this case, I think */
	}

	/* we restrict dsymbol to be a single byte, but not the other symbols */
	if locale.DecimalPoint != nil &&
		len(locale.DecimalPoint) > 0 {
		dsymbol = string(locale.DecimalPoint)
	} else {
		dsymbol = "."
	}

	if locale.ThousandsSeparator != nil && len(locale.ThousandsSeparator) > 0 {
		ssymbol = string(locale.ThousandsSeparator)
	} else {
		if dsymbol != "," {
			ssymbol = ","
		} else {
			ssymbol = "."
		}
	}

	if locale.CurrencySymbol != nil && len(locale.CurrencySymbol) > 0 {
		csymbol = string(locale.CurrencySymbol)
	} else {
		csymbol = "$"
	}

	if locale.PositiveSign != nil && len(locale.PositiveSign) > 0 {
		psymbol = string(locale.PositiveSign)
	} else {
		psymbol = "" // +
	}

	if locale.NegativeSign != nil && len(locale.NegativeSign) > 0 {
		nsymbol = string(locale.NegativeSign)
	} else {
		nsymbol = "-"
	}

	/* we need to add all sorts of checking here.  For now just */
	/* strip all leading whitespace and any leading currency symbol */
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, csymbol)
	s = strings.TrimSpace(s)

	/* a leading minus or paren signifies a negative number */
	/* again, better heuristics needed */
	/* XXX - doesn't properly check for balanced parens - djmc */
	if strings.HasPrefix(s, nsymbol) {
		sgn = -1
		s = strings.TrimPrefix(s, nsymbol)
	} else if strings.HasPrefix(s, "(") {
		sgn = -1
		s = strings.TrimPrefix(s, "(")
	} else if strings.HasPrefix(s, psymbol) {
		s = strings.TrimPrefix(s, psymbol)
	}

	/* allow whitespace and currency symbol after the sign, too */
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, csymbol)
	s = strings.TrimSpace(s)

	/*
	 * We accumulate the absolute amount in "value" and then apply the sign at
	 * the end.  (The sign can appear before or after the digits, so it would
	 * be more complicated to do otherwise.)  Because of the larger range of
	 * negative signed integers, we build "value" in the negative and then
	 * flip the sign at the end, catching most-negative-number overflow if
	 * necessary.
	 */
	for ; len(s) > 0; s = s[1:] {
		c := rune(s[0])
		/*
		 * We look for digits as long as we have found less than the required
		 * number of decimal places.
		 */
		if unicode.IsDigit(c) && (!seenDot || dec < fpoint) {
			digit := c - '0'

			if multiplyOverflow(value, 10, &value) ||
				subtractOverflow(value, int64(int8(digit)), &value) {
				return 0, fmt.Errorf(`value "%s" is out of range for type %s`, input, "money")
			}

			if seenDot {
				dec++
			}
		} else if strings.HasPrefix(s, dsymbol) && !seenDot {
			seenDot = true
		} else if strings.HasPrefix(s, ssymbol) {
			s = s[len(ssymbol)-1:]
		} else {
			break
		}
	}

	// ADD ROUNDING HERE
	if len(s) > 0 && unicode.IsDigit(rune(s[0])) && rune(s[0]) > '5' {
		if subtractOverflow(value, 1, &value) {
			return 0, fmt.Errorf(`value "%s" is out of range for type %s`, input, "money")
		}
	}

	/* adjust for less than required decimal places */
	for ; dec < fpoint; dec++ {
		if multiplyOverflow(value, 10, &value) {
			return 0, fmt.Errorf(`value "%s" is out of range for type %s`, input, "money")
		}
	}

	/*
	 * should only be trailing digits followed by whitespace, right paren,
	 * trailing sign, and/or trailing currency symbol
	 */
	for ; len(s) > 0 && unicode.IsDigit(rune(s[0])); s = s[1:] {
		// All this loop does is consume the trailing digits.
	}

	for len(s) > 0 {
		if unicode.IsSpace(rune(s[0])) || rune(s[0]) == ')' {
			s = s[1:]
		} else if strings.HasPrefix(s, nsymbol) {
			sgn = -1
			s = strings.TrimPrefix(s, nsymbol)
		} else if len(psymbol) > 0 && strings.HasPrefix(s, psymbol) {
			s = strings.TrimPrefix(s, psymbol)
		} else if strings.HasPrefix(s, csymbol) {
			s = strings.TrimPrefix(s, csymbol)
		} else {
			return 0, fmt.Errorf(`invalid input syntax for type %s: "%s"`, "money", input)
		}
	}

	/*
	 * If the value is supposed to be positive, flip the sign, but check for
	 * the most negative number.
	 */
	if sgn > 0 {
		if value == math.MinInt64 {
			return 0, fmt.Errorf(`value "%s" is out of range for type %s`, input, "money")
		}
		result = float64(-value)
	} else {
		result = float64(value)
	}

	return result / 100, nil
}
