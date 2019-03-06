package main

import (
	"bytes"
	"fmt"
	"github.com/elliotcourant/gomonetary"
	"github.com/kataras/golog"
	"io/ioutil"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

var (
	localeMonetaryCommand = `LC_MONETARY=%s locale -k LC_MONETARY`
)

func main() {
	golog.Infof("generating locale metadata")
	meta := GetMonetaryMetadata()
	result := GenerateGo(meta)
	if err := ioutil.WriteFile("lc_generated.go", result, 0644); err != nil {
		panic(err)
	}
	if err := exec.Command("go", "fmt").Run(); err != nil {
		panic(err)
	}
	golog.Infof("generated metadata for %d locales", len(meta))
}

var (
	validKeys = map[string]interface{}{
		"int_curr_symbol":    nil,
		"currency_symbol":    nil,
		"mon_decimal_point":  nil,
		"mon_thousands_sep":  nil,
		"mon_grouping":       nil,
		"positive_sign":      nil,
		"negative_sign":      nil,
		"int_frac_digits":    nil,
		"frac_digits":        nil,
		"p_cs_precedes":      nil,
		"p_sep_by_space":     nil,
		"n_cs_precedes":      nil,
		"n_sep_by_space":     nil,
		"p_sign_posn":        nil,
		"n_sign_posn":        nil,
		"crncystr":           nil,
		"int_p_cs_precedes":  nil,
		"int_p_sep_by_space": nil,
		"int_n_cs_precedes":  nil,
		"int_n_sep_by_space": nil,
		"int_p_sign_posn":    nil,
		"int_n_sign_posn":    nil,
	}
)

func GetMonetaryMetadata() map[string]map[string]interface{} {
	if !monetary.LocaleSupported() {
		panic("locales cannot be generated on this platform. locales must be generated on linux or darwin.")
	}
	list, err := exec.Command("locale", "-a").Output()
	if err != nil {
		panic(err)
	}

	allItems := bytes.Split(list, []byte{10})
	items := make([]string, 0)
	for _, item := range allItems {
		if len(item) == 0 {
			continue
		}
		items = append(items, string(item))
	}
	sort.Strings(items)

	metaMap := map[string]map[string]interface{}{}
	for _, item := range items {
		result, err := exec.Command("/bin/sh", "-c", fmt.Sprintf(localeMonetaryCommand, item)).CombinedOutput()
		if err != nil {
			golog.Fatalf("could not generate metadata for locale [%s]: %s", item, err.Error())
		}
		metaMap[item] = map[string]interface{}{}
		fieldList := bytes.Split(result, []byte{10})
		for _, field := range fieldList {
			if len(field) == 0 {
				continue
			}
			kv := bytes.Split(field, []byte("="))
			key := string(kv[0])
			if _, ok := validKeys[key]; !ok {
				continue
			}
			val := kv[1]
			if len(val) == 0 {
				metaMap[item][key] = nil
				continue
			}
			isNumber := true
			if val[0] == byte('"') {
				isNumber = false
				val = val[1 : len(val)-1]
			}
			if key == "mon_grouping" {
				groups := bytes.Split(val, []byte(";"))
				groupVal := make([]int8, len(groups))
				for i, group := range groups {
					if g, err := strconv.ParseInt(string(group), 10, 8); err != nil {
						golog.Errorf("%s", string(result))
						panic(err)
					} else {
						groupVal[i] = int8(g)
					}
				}
				metaMap[item][key] = groupVal
			} else {
				if !isNumber {
					metaMap[item][key] = bytes.TrimSpace(val)
				} else {
					if v, err := strconv.ParseInt(string(val), 10, 32); err != nil {
						golog.Errorf("%s", string(result))
						panic(err)
					} else {
						metaMap[item][key] = int(v)
					}
				}
			}
		}
	}

	return metaMap
}

func GenerateGo(meta map[string]map[string]interface{}) []byte {
	data := `package monetary
var (
	localeNames = LocaleNames{
		%s
	}
	metaData = map[string]localeInfo{
		%s
	}
)`
	kvFormatter := func(key, value string) string {
		return fmt.Sprintf(`"%s": %s,`, key, value)
	}
	stringFormatter := func(str interface{}) string {
		if str == nil {
			str = ""
		}
		return fmt.Sprintf(`"%s"`, str)
	}
	bytesFormatter := func(bytes interface{}) string {
		if bytes == nil {
			return "nil"
		}
		str := fmt.Sprintf("%+v", bytes.([]byte))
		return fmt.Sprintf("[]byte{%s}", strings.Join(strings.Split(str[1:len(str)-1], " "), ", "))
	}
	tinyIntArrayFormatter := func(groups interface{}) string {
		if groups == nil {
			return "nil"
		}
		str := fmt.Sprintf("%+v", groups.([]int8))
		return fmt.Sprintf("[]int8{%s}", strings.Join(strings.Split(str[1:len(str)-1], " "), ", "))
	}
	localeInfoFormatter := func(data map[string]interface{}) string {
		return fmt.Sprintf(`{
			InternationalCurrencySymbol: %s,
			CurrencySymbol: %s,
			DecimalPoint: %s,
			ThousandsSeparator: %s,
			Grouping: %s,
			PositiveSign: %s,
			NegativeSign: %s,
			FractionalDigits: %d,
		}`,
			stringFormatter(data["int_curr_symbol"]),
			bytesFormatter(data["currency_symbol"]),
			bytesFormatter(data["mon_decimal_point"]),
			bytesFormatter(data["mon_thousands_sep"]),
			tinyIntArrayFormatter(data["mon_grouping"]),
			bytesFormatter(data["positive_sign"]),
			bytesFormatter(data["negative_sign"]),
			data["frac_digits"],
		)
	}
	body := ""
	names := ""
	for locale, data := range meta {
		names += fmt.Sprintf(`"%s",
`, locale)
		result := localeInfoFormatter(data)
		body += kvFormatter(locale, result)
	}

	data = fmt.Sprintf(data, names, body)

	return []byte(data)
}
