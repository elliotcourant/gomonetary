package monetary

var (
	localeNames = LocaleNames{
		"it_CH.ISO8859-1",
		"zh_HK.Big5HKSCS",
		"POSIX",
		"af_ZA",
		"af_ZA.ISO8859-1",
		"hr_HR.ISO8859-2",
		"en_GB.ISO8859-1",
		"fr_BE.ISO8859-1",
		"ca_ES",
		"da_DK",
		"de_AT.ISO8859-15",
		"en_CA.UTF-8",
		"de_DE.ISO8859-15",
		"eu_ES.UTF-8",
		"is_IS",
		"ru_RU",
		"it_IT.ISO8859-1",
		"pt_PT",
		"sk_SK",
		"en_GB",
		"eu_ES.ISO8859-1",
		"hy_AM.ARMSCII-8",
		"is_IS.ISO8859-15",
		"it_IT.UTF-8",
		"ja_JP",
		"sv_SE.UTF-8",
		"cs_CZ.ISO8859-2",
		"en_US.ISO8859-15",
		"es_ES",
		"hy_AM.UTF-8",
		"it_CH",
		"tr_TR.ISO8859-9",
		"zh_CN.UTF-8",
		"en_NZ.ISO8859-15",
		"en_US.ISO8859-1",
		"et_EE.ISO8859-15",
		"is_IS.ISO8859-1",
		"kk_KZ",
		"pl_PL.UTF-8",
		"zh_CN.GB2312",
		"C",
		"en_CA.ISO8859-1",
		"fi_FI",
		"fr_FR.ISO8859-1",
		"hu_HU.UTF-8",
		"is_IS.UTF-8",
		"pt_BR",
		"pt_BR.UTF-8",
		"de_AT",
		"en_AU",
		"es_ES.UTF-8",
		"he_IL",
		"zh_TW",
		"de_AT.UTF-8",
		"et_EE",
		"fr_CA.UTF-8",
		"tr_TR",
		"ro_RO",
		"ro_RO.ISO8859-2",
		"uk_UA",
		"uk_UA.ISO8859-5",
		"ca_ES.ISO8859-1",
		"en_NZ.UTF-8",
		"fr_CH.UTF-8",
		"hy_AM",
		"zh_CN.eucCN",
		"de_DE.UTF-8",
		"et_EE.UTF-8",
		"en_US.US-ASCII",
		"fr_FR.ISO8859-15",
		"it_IT",
		"sr_YU.ISO8859-5",
		"de_DE",
		"el_GR.ISO8859-7",
		"en_CA",
		"en_NZ.ISO8859-1",
		"sv_SE",
		"de_CH.UTF-8",
		"ja_JP.SJIS",
		"pl_PL",
		"sl_SI.ISO8859-2",
		"af_ZA.UTF-8",
		"da_DK.ISO8859-15",
		"de_AT.ISO8859-1",
		"de_CH.ISO8859-1",
		"zh_CN.GB18030",
		"ru_RU.UTF-8",
		"de_DE.ISO8859-1",
		"en_US",
		"eu_ES",
		"ru_RU.CP1251",
		"ko_KR.eucKR",
		"pt_BR.ISO8859-1",
		"pt_PT.ISO8859-1",
		"zh_HK",
		"cs_CZ",
		"it_CH.UTF-8",
		"kk_KZ.PT154",
		"ko_KR.UTF-8",
		"es_ES.ISO8859-15",
		"sl_SI.UTF-8",
		"sr_YU.UTF-8",
		"uk_UA.UTF-8",
		"am_ET",
		"be_BY.CP1251",
		"no_NO.ISO8859-15",
		"pt_PT.ISO8859-15",
		"it_CH.ISO8859-15",
		"nl_BE",
		"nl_BE.UTF-8",
		"zh_TW.UTF-8",
		"be_BY.ISO8859-5",
		"fr_BE.ISO8859-15",
		"no_NO.ISO8859-1",
		"ru_RU.KOI8-R",
		"no_NO.UTF-8",
		"sk_SK.ISO8859-2",
		"sr_YU",
		"en_IE.UTF-8",
		"en_NZ.US-ASCII",
		"hr_HR.UTF-8",
		"ja_JP.eucJP",
		"en_AU.UTF-8",
		"pl_PL.ISO8859-2",
		"zh_CN.GBK",
		"af_ZA.ISO8859-15",
		"am_ET.UTF-8",
		"be_BY.UTF-8",
		"ca_ES.UTF-8",
		"nl_NL.ISO8859-1",
		"ru_RU.ISO8859-5",
		"sv_SE.ISO8859-1",
		"bg_BG",
		"eu_ES.ISO8859-15",
		"fr_CH.ISO8859-15",
		"ko_KR.CP949",
		"en_IE",
		"en_NZ",
		"he_IL.UTF-8",
		"nl_NL.ISO8859-15",
		"en_AU.ISO8859-15",
		"en_AU.US-ASCII",
		"en_GB.ISO8859-15",
		"en_GB.US-ASCII",
		"be_BY",
		"ca_ES.ISO8859-15",
		"da_DK.UTF-8",
		"de_CH.ISO8859-15",
		"fr_CH",
		"hr_HR",
		"ro_RO.UTF-8",
		"sk_SK.UTF-8",
		"en_GB.UTF-8",
		"fi_FI.ISO8859-1",
		"fi_FI.UTF-8",
		"fr_BE.UTF-8",
		"zh_TW.Big5",
		"be_BY.CP1131",
		"es_ES.ISO8859-1",
		"lt_LT",
		"pt_PT.UTF-8",
		"fr_BE",
		"fr_FR",
		"ko_KR",
		"lt_LT.UTF-8",
		"bg_BG.UTF-8",
		"cs_CZ.UTF-8",
		"da_DK.ISO8859-1",
		"el_GR",
		"nl_NL",
		"sv_SE.ISO8859-15",
		"nl_BE.ISO8859-1",
		"zh_HK.UTF-8",
		"en_CA.US-ASCII",
		"hu_HU.ISO8859-2",
		"ja_JP.UTF-8",
		"lt_LT.ISO8859-13",
		"nl_BE.ISO8859-15",
		"nl_NL.UTF-8",
		"ru_RU.CP866",
		"sl_SI",
		"bg_BG.CP1251",
		"de_CH",
		"fi_FI.ISO8859-15",
		"fr_FR.UTF-8",
		"zh_CN",
		"tr_TR.UTF-8",
		"uk_UA.KOI8-U",
		"en_AU.ISO8859-1",
		"fr_CA.ISO8859-15",
		"hu_HU",
		"kk_KZ.UTF-8",
		"en_CA.ISO8859-15",
		"fr_CH.ISO8859-1",
		"no_NO",
		"lt_LT.ISO8859-4",
		"sr_YU.ISO8859-2",
		"en_US.UTF-8",
		"fr_CA",
		"hi_IN.ISCII-DEV",
		"it_IT.ISO8859-15",
		"el_GR.UTF-8",
		"fr_CA.ISO8859-1",
	}
	metaData = map[string]localeInfo{
		"it_CH.ISO8859-1": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "zh_HK.Big5HKSCS": {
			InternationalCurrencySymbol: "HKD",
			CurrencySymbol:              []byte{72, 75, 36},
		}, "POSIX": {
			InternationalCurrencySymbol: "",
			CurrencySymbol:              nil,
		}, "af_ZA": {
			InternationalCurrencySymbol: "ZAR",
			CurrencySymbol:              []byte{82},
		}, "af_ZA.ISO8859-1": {
			InternationalCurrencySymbol: "ZAR",
			CurrencySymbol:              []byte{82},
		}, "hr_HR.ISO8859-2": {
			InternationalCurrencySymbol: "HRK",
			CurrencySymbol:              []byte{75, 110},
		}, "en_GB.ISO8859-1": {
			InternationalCurrencySymbol: "GBP",
			CurrencySymbol:              []byte{163},
		}, "fr_BE.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "ca_ES": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "da_DK": {
			InternationalCurrencySymbol: "DKK",
			CurrencySymbol:              []byte{107, 114},
		}, "de_AT.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "en_CA.UTF-8": {
			InternationalCurrencySymbol: "CAD",
			CurrencySymbol:              []byte{36},
		}, "de_DE.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "eu_ES.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "is_IS": {
			InternationalCurrencySymbol: "ISK",
			CurrencySymbol:              []byte{107, 114},
		}, "ru_RU": {
			InternationalCurrencySymbol: "RUB",
			CurrencySymbol:              []byte{209, 128, 209, 131, 208, 177, 46},
		}, "it_IT.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "pt_PT": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "sk_SK": {
			InternationalCurrencySymbol: "SKK",
			CurrencySymbol:              []byte{83, 107},
		}, "en_GB": {
			InternationalCurrencySymbol: "GBP",
			CurrencySymbol:              []byte{194, 163},
		}, "eu_ES.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "hy_AM.ARMSCII-8": {
			InternationalCurrencySymbol: "AMD",
			CurrencySymbol:              []byte{184, 240},
		}, "is_IS.ISO8859-15": {
			InternationalCurrencySymbol: "ISK",
			CurrencySymbol:              []byte{107, 114},
		}, "it_IT.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "ja_JP": {
			InternationalCurrencySymbol: "JPY",
			CurrencySymbol:              []byte{194, 165},
		}, "sv_SE.UTF-8": {
			InternationalCurrencySymbol: "SEK",
			CurrencySymbol:              []byte{107, 114},
		}, "cs_CZ.ISO8859-2": {
			InternationalCurrencySymbol: "CZK",
			CurrencySymbol:              []byte{75, 232},
		}, "en_US.ISO8859-15": {
			InternationalCurrencySymbol: "USD",
			CurrencySymbol:              []byte{36},
		}, "es_ES": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "hy_AM.UTF-8": {
			InternationalCurrencySymbol: "AMD",
			CurrencySymbol:              []byte{212, 180, 213, 144},
		}, "it_CH": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "tr_TR.ISO8859-9": {
			InternationalCurrencySymbol: "TRL",
			CurrencySymbol:              []byte{76},
		}, "zh_CN.UTF-8": {
			InternationalCurrencySymbol: "CNY",
			CurrencySymbol:              []byte{239, 191, 165},
		}, "en_NZ.ISO8859-15": {
			InternationalCurrencySymbol: "NZD",
			CurrencySymbol:              []byte{36},
		}, "en_US.ISO8859-1": {
			InternationalCurrencySymbol: "USD",
			CurrencySymbol:              []byte{36},
		}, "et_EE.ISO8859-15": {
			InternationalCurrencySymbol: "EEK",
			CurrencySymbol:              []byte{107, 114},
		}, "is_IS.ISO8859-1": {
			InternationalCurrencySymbol: "ISK",
			CurrencySymbol:              []byte{107, 114},
		}, "kk_KZ": {
			InternationalCurrencySymbol: "KZT",
			CurrencySymbol:              []byte{209, 130, 208, 179, 46},
		}, "pl_PL.UTF-8": {
			InternationalCurrencySymbol: "PLN",
			CurrencySymbol:              []byte{122, 197, 130},
		}, "zh_CN.GB2312": {
			InternationalCurrencySymbol: "CNY",
			CurrencySymbol:              []byte{163, 164},
		}, "C": {
			InternationalCurrencySymbol: "",
			CurrencySymbol:              nil,
		}, "en_CA.ISO8859-1": {
			InternationalCurrencySymbol: "CAD",
			CurrencySymbol:              []byte{36},
		}, "fi_FI": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "fr_FR.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "hu_HU.UTF-8": {
			InternationalCurrencySymbol: "HUF",
			CurrencySymbol:              []byte{70, 116},
		}, "is_IS.UTF-8": {
			InternationalCurrencySymbol: "ISK",
			CurrencySymbol:              []byte{107, 114},
		}, "pt_BR": {
			InternationalCurrencySymbol: "BRL",
			CurrencySymbol:              []byte{82, 36},
		}, "pt_BR.UTF-8": {
			InternationalCurrencySymbol: "BRL",
			CurrencySymbol:              []byte{82, 36},
		}, "de_AT": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "en_AU": {
			InternationalCurrencySymbol: "AUD",
			CurrencySymbol:              []byte{36},
		}, "es_ES.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "he_IL": {
			InternationalCurrencySymbol: "ILS",
			CurrencySymbol:              []byte{215, 169, 215, 151},
		}, "zh_TW": {
			InternationalCurrencySymbol: "TWD",
			CurrencySymbol:              []byte{78, 84, 36},
		}, "de_AT.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "et_EE": {
			InternationalCurrencySymbol: "EEK",
			CurrencySymbol:              []byte{107, 114},
		}, "fr_CA.UTF-8": {
			InternationalCurrencySymbol: "CAD",
			CurrencySymbol:              []byte{36},
		}, "tr_TR": {
			InternationalCurrencySymbol: "TRL",
			CurrencySymbol:              []byte{76},
		}, "ro_RO": {
			InternationalCurrencySymbol: "ROL",
			CurrencySymbol:              []byte{76, 101, 105},
		}, "ro_RO.ISO8859-2": {
			InternationalCurrencySymbol: "ROL",
			CurrencySymbol:              []byte{76, 101, 105},
		}, "uk_UA": {
			InternationalCurrencySymbol: "UAH",
			CurrencySymbol:              []byte{208, 179, 209, 128, 208, 189, 46},
		}, "uk_UA.ISO8859-5": {
			InternationalCurrencySymbol: "UAH",
			CurrencySymbol:              []byte{211, 224, 221, 46},
		}, "ca_ES.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "en_NZ.UTF-8": {
			InternationalCurrencySymbol: "NZD",
			CurrencySymbol:              []byte{36},
		}, "fr_CH.UTF-8": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "hy_AM": {
			InternationalCurrencySymbol: "AMD",
			CurrencySymbol:              []byte{212, 180, 213, 144},
		}, "zh_CN.eucCN": {
			InternationalCurrencySymbol: "CNY",
			CurrencySymbol:              []byte{163, 164},
		}, "de_DE.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "et_EE.UTF-8": {
			InternationalCurrencySymbol: "EEK",
			CurrencySymbol:              []byte{107, 114},
		}, "en_US.US-ASCII": {
			InternationalCurrencySymbol: "USD",
			CurrencySymbol:              []byte{36},
		}, "fr_FR.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "it_IT": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "sr_YU.ISO8859-5": {
			InternationalCurrencySymbol: "YUD",
			CurrencySymbol:              []byte{212, 216, 221},
		}, "de_DE": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "el_GR.ISO8859-7": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "en_CA": {
			InternationalCurrencySymbol: "CAD",
			CurrencySymbol:              []byte{36},
		}, "en_NZ.ISO8859-1": {
			InternationalCurrencySymbol: "NZD",
			CurrencySymbol:              []byte{36},
		}, "sv_SE": {
			InternationalCurrencySymbol: "SEK",
			CurrencySymbol:              []byte{107, 114},
		}, "de_CH.UTF-8": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "ja_JP.SJIS": {
			InternationalCurrencySymbol: "JPY",
			CurrencySymbol:              []byte{92},
		}, "pl_PL": {
			InternationalCurrencySymbol: "PLN",
			CurrencySymbol:              []byte{122, 197, 130},
		}, "sl_SI.ISO8859-2": {
			InternationalCurrencySymbol: "SIT",
			CurrencySymbol:              []byte{83, 73, 84},
		}, "af_ZA.UTF-8": {
			InternationalCurrencySymbol: "ZAR",
			CurrencySymbol:              []byte{82},
		}, "da_DK.ISO8859-15": {
			InternationalCurrencySymbol: "DKK",
			CurrencySymbol:              []byte{107, 114},
		}, "de_AT.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "de_CH.ISO8859-1": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "zh_CN.GB18030": {
			InternationalCurrencySymbol: "CNY",
			CurrencySymbol:              []byte{163, 164},
		}, "ru_RU.UTF-8": {
			InternationalCurrencySymbol: "RUB",
			CurrencySymbol:              []byte{209, 128, 209, 131, 208, 177, 46},
		}, "de_DE.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "en_US": {
			InternationalCurrencySymbol: "USD",
			CurrencySymbol:              []byte{36},
		}, "eu_ES": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "ru_RU.CP1251": {
			InternationalCurrencySymbol: "RUB",
			CurrencySymbol:              []byte{240, 243, 225, 46},
		}, "ko_KR.eucKR": {
			InternationalCurrencySymbol: "KRW",
			CurrencySymbol:              []byte{92},
		}, "pt_BR.ISO8859-1": {
			InternationalCurrencySymbol: "BRL",
			CurrencySymbol:              []byte{82, 36},
		}, "pt_PT.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "zh_HK": {
			InternationalCurrencySymbol: "HKD",
			CurrencySymbol:              []byte{72, 75, 36},
		}, "cs_CZ": {
			InternationalCurrencySymbol: "CZK",
			CurrencySymbol:              []byte{75, 196, 141},
		}, "it_CH.UTF-8": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "kk_KZ.PT154": {
			InternationalCurrencySymbol: "KZT",
			CurrencySymbol:              []byte{242, 227, 46},
		}, "ko_KR.UTF-8": {
			InternationalCurrencySymbol: "KRW",
			CurrencySymbol:              []byte{226, 130, 169},
		}, "es_ES.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "sl_SI.UTF-8": {
			InternationalCurrencySymbol: "SIT",
			CurrencySymbol:              []byte{83, 73, 84},
		}, "sr_YU.UTF-8": {
			InternationalCurrencySymbol: "YUD",
			CurrencySymbol:              []byte{100, 105, 110},
		}, "uk_UA.UTF-8": {
			InternationalCurrencySymbol: "UAH",
			CurrencySymbol:              []byte{208, 179, 209, 128, 208, 189, 46},
		}, "am_ET": {
			InternationalCurrencySymbol: "ETB",
			CurrencySymbol:              []byte{36},
		}, "be_BY.CP1251": {
			InternationalCurrencySymbol: "BYR",
			CurrencySymbol:              []byte{240, 243, 225, 46},
		}, "no_NO.ISO8859-15": {
			InternationalCurrencySymbol: "NOK",
			CurrencySymbol:              []byte{107, 114},
		}, "pt_PT.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "it_CH.ISO8859-15": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "nl_BE": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "nl_BE.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "zh_TW.UTF-8": {
			InternationalCurrencySymbol: "TWD",
			CurrencySymbol:              []byte{78, 84, 36},
		}, "be_BY.ISO8859-5": {
			InternationalCurrencySymbol: "BYR",
			CurrencySymbol:              []byte{224, 227, 209, 46},
		}, "fr_BE.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "no_NO.ISO8859-1": {
			InternationalCurrencySymbol: "NOK",
			CurrencySymbol:              []byte{107, 114},
		}, "ru_RU.KOI8-R": {
			InternationalCurrencySymbol: "RUB",
			CurrencySymbol:              []byte{210, 213, 194, 46},
		}, "no_NO.UTF-8": {
			InternationalCurrencySymbol: "NOK",
			CurrencySymbol:              []byte{107, 114},
		}, "sk_SK.ISO8859-2": {
			InternationalCurrencySymbol: "SKK",
			CurrencySymbol:              []byte{83, 107},
		}, "sr_YU": {
			InternationalCurrencySymbol: "YUD",
			CurrencySymbol:              []byte{100, 105, 110},
		}, "en_IE.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{226, 130, 172},
		}, "en_NZ.US-ASCII": {
			InternationalCurrencySymbol: "NZD",
			CurrencySymbol:              []byte{36},
		}, "hr_HR.UTF-8": {
			InternationalCurrencySymbol: "HRK",
			CurrencySymbol:              []byte{75, 110},
		}, "ja_JP.eucJP": {
			InternationalCurrencySymbol: "JPY",
			CurrencySymbol:              []byte{92},
		}, "en_AU.UTF-8": {
			InternationalCurrencySymbol: "AUD",
			CurrencySymbol:              []byte{36},
		}, "pl_PL.ISO8859-2": {
			InternationalCurrencySymbol: "PLN",
			CurrencySymbol:              []byte{122, 179},
		}, "zh_CN.GBK": {
			InternationalCurrencySymbol: "CNY",
			CurrencySymbol:              []byte{163, 164},
		}, "af_ZA.ISO8859-15": {
			InternationalCurrencySymbol: "ZAR",
			CurrencySymbol:              []byte{82},
		}, "am_ET.UTF-8": {
			InternationalCurrencySymbol: "ETB",
			CurrencySymbol:              []byte{36},
		}, "be_BY.UTF-8": {
			InternationalCurrencySymbol: "BYR",
			CurrencySymbol:              []byte{209, 128, 209, 131, 208, 177, 46},
		}, "ca_ES.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "nl_NL.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "ru_RU.ISO8859-5": {
			InternationalCurrencySymbol: "RUB",
			CurrencySymbol:              []byte{224, 227, 209, 46},
		}, "sv_SE.ISO8859-1": {
			InternationalCurrencySymbol: "SEK",
			CurrencySymbol:              []byte{107, 114},
		}, "bg_BG": {
			InternationalCurrencySymbol: "BGN",
			CurrencySymbol:              []byte{208, 187, 208, 178, 46},
		}, "eu_ES.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "fr_CH.ISO8859-15": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "ko_KR.CP949": {
			InternationalCurrencySymbol: "KRW",
			CurrencySymbol:              []byte{92},
		}, "en_IE": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{226, 130, 172},
		}, "en_NZ": {
			InternationalCurrencySymbol: "NZD",
			CurrencySymbol:              []byte{36},
		}, "he_IL.UTF-8": {
			InternationalCurrencySymbol: "ILS",
			CurrencySymbol:              []byte{215, 169, 215, 151},
		}, "nl_NL.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "en_AU.ISO8859-15": {
			InternationalCurrencySymbol: "AUD",
			CurrencySymbol:              []byte{36},
		}, "en_AU.US-ASCII": {
			InternationalCurrencySymbol: "AUD",
			CurrencySymbol:              []byte{36},
		}, "en_GB.ISO8859-15": {
			InternationalCurrencySymbol: "GBP",
			CurrencySymbol:              []byte{163},
		}, "en_GB.US-ASCII": {
			InternationalCurrencySymbol: "GBP",
			CurrencySymbol:              []byte{163},
		}, "be_BY": {
			InternationalCurrencySymbol: "BYR",
			CurrencySymbol:              []byte{209, 128, 209, 131, 208, 177, 46},
		}, "ca_ES.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "da_DK.UTF-8": {
			InternationalCurrencySymbol: "DKK",
			CurrencySymbol:              []byte{107, 114},
		}, "de_CH.ISO8859-15": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "fr_CH": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "hr_HR": {
			InternationalCurrencySymbol: "HRK",
			CurrencySymbol:              []byte{75, 110},
		}, "ro_RO.UTF-8": {
			InternationalCurrencySymbol: "ROL",
			CurrencySymbol:              []byte{76, 101, 105},
		}, "sk_SK.UTF-8": {
			InternationalCurrencySymbol: "SKK",
			CurrencySymbol:              []byte{83, 107},
		}, "en_GB.UTF-8": {
			InternationalCurrencySymbol: "GBP",
			CurrencySymbol:              []byte{194, 163},
		}, "fi_FI.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "fi_FI.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "fr_BE.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "zh_TW.Big5": {
			InternationalCurrencySymbol: "TWD",
			CurrencySymbol:              []byte{78, 84, 36},
		}, "be_BY.CP1131": {
			InternationalCurrencySymbol: "BYR",
			CurrencySymbol:              []byte{224, 227, 161, 46},
		}, "es_ES.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "lt_LT": {
			InternationalCurrencySymbol: "LTL",
			CurrencySymbol:              []byte{76, 116},
		}, "pt_PT.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "fr_BE": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "fr_FR": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "ko_KR": {
			InternationalCurrencySymbol: "KRW",
			CurrencySymbol:              []byte{226, 130, 169},
		}, "lt_LT.UTF-8": {
			InternationalCurrencySymbol: "LTL",
			CurrencySymbol:              []byte{76, 116},
		}, "bg_BG.UTF-8": {
			InternationalCurrencySymbol: "BGN",
			CurrencySymbol:              []byte{208, 187, 208, 178, 46},
		}, "cs_CZ.UTF-8": {
			InternationalCurrencySymbol: "CZK",
			CurrencySymbol:              []byte{75, 196, 141},
		}, "da_DK.ISO8859-1": {
			InternationalCurrencySymbol: "DKK",
			CurrencySymbol:              []byte{107, 114},
		}, "el_GR": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "nl_NL": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "sv_SE.ISO8859-15": {
			InternationalCurrencySymbol: "SEK",
			CurrencySymbol:              []byte{107, 114},
		}, "nl_BE.ISO8859-1": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "zh_HK.UTF-8": {
			InternationalCurrencySymbol: "HKD",
			CurrencySymbol:              []byte{72, 75, 36},
		}, "en_CA.US-ASCII": {
			InternationalCurrencySymbol: "CAD",
			CurrencySymbol:              []byte{36},
		}, "hu_HU.ISO8859-2": {
			InternationalCurrencySymbol: "HUF",
			CurrencySymbol:              []byte{70, 116},
		}, "ja_JP.UTF-8": {
			InternationalCurrencySymbol: "JPY",
			CurrencySymbol:              []byte{194, 165},
		}, "lt_LT.ISO8859-13": {
			InternationalCurrencySymbol: "LTL",
			CurrencySymbol:              []byte{76, 116},
		}, "nl_BE.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "nl_NL.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "ru_RU.CP866": {
			InternationalCurrencySymbol: "RUB",
			CurrencySymbol:              []byte{224, 227, 161, 46},
		}, "sl_SI": {
			InternationalCurrencySymbol: "SIT",
			CurrencySymbol:              []byte{83, 73, 84},
		}, "bg_BG.CP1251": {
			InternationalCurrencySymbol: "BGN",
			CurrencySymbol:              []byte{235, 226, 46},
		}, "de_CH": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "fi_FI.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "fr_FR.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "zh_CN": {
			InternationalCurrencySymbol: "CNY",
			CurrencySymbol:              []byte{239, 191, 165},
		}, "tr_TR.UTF-8": {
			InternationalCurrencySymbol: "TRL",
			CurrencySymbol:              []byte{76},
		}, "uk_UA.KOI8-U": {
			InternationalCurrencySymbol: "UAH",
			CurrencySymbol:              []byte{199, 210, 206, 46},
		}, "en_AU.ISO8859-1": {
			InternationalCurrencySymbol: "AUD",
			CurrencySymbol:              []byte{36},
		}, "fr_CA.ISO8859-15": {
			InternationalCurrencySymbol: "CAD",
			CurrencySymbol:              []byte{36},
		}, "hu_HU": {
			InternationalCurrencySymbol: "HUF",
			CurrencySymbol:              []byte{70, 116},
		}, "kk_KZ.UTF-8": {
			InternationalCurrencySymbol: "KZT",
			CurrencySymbol:              []byte{209, 130, 208, 179, 46},
		}, "en_CA.ISO8859-15": {
			InternationalCurrencySymbol: "CAD",
			CurrencySymbol:              []byte{36},
		}, "fr_CH.ISO8859-1": {
			InternationalCurrencySymbol: "CHF",
			CurrencySymbol:              []byte{70, 114, 46},
		}, "no_NO": {
			InternationalCurrencySymbol: "NOK",
			CurrencySymbol:              []byte{107, 114},
		}, "lt_LT.ISO8859-4": {
			InternationalCurrencySymbol: "LTL",
			CurrencySymbol:              []byte{76, 116},
		}, "sr_YU.ISO8859-2": {
			InternationalCurrencySymbol: "YUD",
			CurrencySymbol:              []byte{100, 105, 110},
		}, "en_US.UTF-8": {
			InternationalCurrencySymbol: "USD",
			CurrencySymbol:              []byte{36},
		}, "fr_CA": {
			InternationalCurrencySymbol: "CAD",
			CurrencySymbol:              []byte{36},
		}, "hi_IN.ISCII-DEV": {
			InternationalCurrencySymbol: "INR",
			CurrencySymbol:              []byte{207, 222, 200, 172},
		}, "it_IT.ISO8859-15": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "el_GR.UTF-8": {
			InternationalCurrencySymbol: "EUR",
			CurrencySymbol:              []byte{69, 117},
		}, "fr_CA.ISO8859-1": {
			InternationalCurrencySymbol: "CAD",
			CurrencySymbol:              []byte{36},
		},
	}
)
