package monetary

type localeInfo struct {
	Locale                         string
	InternationalCurrencySymbol    string
	CurrencySymbol                 []byte
	DecimalPoint                   byte
	ThousandsSeparator             []byte
	Grouping                       []int8
	PositiveSign                   []byte
	NegativeSign                   []byte
	InternationalFractionalDigits  uint8
	FractionalDigits               uint8
	PositiveSignPosition           int
	PositiveCurrencySymbolPrecedes bool
	PositiveSeparatedBySpace       int8
	NegativeSignPosition           int
	NegativeCurrencySymbolPrecedes bool
	NegativeSeparatedBySpace       int8
}
