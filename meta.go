package monetary

type localeInfo struct {
	InternationalCurrencySymbol string
	CurrencySymbol              string
	DecimalPoint                string
	ThousandsSeparator          string
	Grouping                    [2]uint8
	PositiveSign                string
	NegativeSign                string
	IntFracDigits               uint8
	FracDigits                  uint8
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
	int_curr_symbol  "<U><S><D>"
	currency_symbol  "<dollar-sign>"
	mon_decimal_point        "<period>"
	mon_thousands_sep        "<comma>"
	mon_grouping             <3>
	positive_sign            "<plus-sign>"
	negative_sign            "<hyphen>"
	int_frac_digits  <2>
	frac_digits              <2>
	p_cs_precedes            <1>
	p_sep_by_space   <2>
	n_cs_precedes            <1>
	n_sep_by_space   <2>
	p_sign_posn              <3>
	n_sign_posn              <3>
	debit_sign               "<D><B>"
	credit_sign              "<C><R>"
	left_parenthesis         "<left-parenthesis>"
	right_parenthesis        "<right-parenthesis>"
	#
	END LC_MONETARY
*/
func readMonetaryBytes(data []byte) (localeInfo, error) {
	return localeInfo{}, nil
}
