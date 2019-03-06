package monetary

type LocaleName string

type LocaleNames []LocaleName

func GetSupportedLocales() LocaleNames {
	return localeNames
}
