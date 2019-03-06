package monetary

type localeInfo struct {
	InternationalCurrencySymbol   string
	CurrencySymbol                []byte
	DecimalPoint                  []byte
	ThousandsSeparator            []byte
	Grouping                      []int8
	PositiveSign                  []byte
	NegativeSign                  []byte
	InternationalFractionalDigits uint8
	FractionalDigits              uint8
}
