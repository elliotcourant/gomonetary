package monetary

import (
	"github.com/elliotcourant/gomonetary/internal"
	"strconv"
	"strings"
)

type localeInfo struct {
	InternationalCurrencySymbol   string
	CurrencySymbol                string
	DecimalPoint                  string
	ThousandsSeparator            string
	Grouping                      []int8
	PositiveSign                  string
	NegativeSign                  string
	InternationalFractionalDigits uint8
	FractionalDigits              uint8
}

type localFieldIndex int

const (
	intCurrSymbol localFieldIndex = iota
	currencySymbol
	monDecimalPoint
	monThousandsSep
	monGrouping
	positiveSign
	negativeSign
	intFracDigits
	fracDigits
)

/*
	LC_MONETARY
	#
	int_curr_symbol     "<U><S><D>"
	currency_symbol     "<dollar-sign>"
	mon_decimal_point   "<period>"
	mon_thousands_sep   "<comma>"
	mon_grouping        <3>
	positive_sign       "<plus-sign>"
	negative_sign       "<hyphen>"
	int_frac_digits      <2>
	frac_digits          <2>
	p_cs_precedes        <1>
	p_sep_by_space       <2>
	n_cs_precedes        <1>
	n_sep_by_space       <2>
	p_sign_posn          <3>
	n_sign_posn          <3>
	debit_sign           "<D><B>"
	credit_sign          "<C><R>"
	left_parenthesis     "<left-parenthesis>"
	right_parenthesis    "<right-parenthesis>"
	#
	END LC_MONETARY
*/
func readMonetaryBytes(data []byte) (result localeInfo, err error) {
	items := internal.SplitMonetaryBytes(data)
	for i, m := range items {
		item := string(m)
		switch localFieldIndex(i) {
		case intCurrSymbol:
			result.InternationalCurrencySymbol = strings.TrimSpace(item)
		case currencySymbol:
			result.CurrencySymbol = item
		case monDecimalPoint:
			result.DecimalPoint = item
		case monThousandsSep:
			result.ThousandsSeparator = item
		case monGrouping:
			groups := strings.Split(item, ";")
			result.Grouping = make([]int8, len(groups))
			for x, g := range groups {
				group, err := strconv.ParseInt(g, 10, 8)
				if err != nil {
					return result, err
				}
				result.Grouping[x] = int8(group)
			}
		case positiveSign:
			result.PositiveSign = item
		case negativeSign:
			result.NegativeSign = item
		case intFracDigits, fracDigits:
			digits, err := strconv.ParseUint(item, 10, 8)
			if err != nil {
				return result, err
			}
			if localFieldIndex(i) == intFracDigits {
				result.InternationalFractionalDigits = uint8(digits)
			} else {
				result.FractionalDigits = uint8(digits)
			}
		}
	}
	return result, nil
}
