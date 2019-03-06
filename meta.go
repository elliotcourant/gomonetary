package monetary

type localeInfo struct {
	InternationalCurrencySymbol   string
	CurrencySymbol                []byte
	DecimalPoint                  string
	ThousandsSeparator            string
	Grouping                      []int8
	PositiveSign                  string
	NegativeSign                  string
	InternationalFractionalDigits uint8
	FractionalDigits              uint8
}
