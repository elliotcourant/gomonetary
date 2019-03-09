package monetary

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

type byteConstants []byte

var (
	Negative = byteConstants("-")
	Decimal  = byteConstants(".")
	Comma    = byteConstants(",")
	Dollar   = byteConstants("$")
)

func sliceEquals(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func floatToInt(f float64) (int64, error) {
	vs := bytes.Split([]byte(fmt.Sprintf("%f", f)), Decimal)
	v1 := []byte(vs[0])
	v2 := []byte(vs[1][:int(math.Min(float64(len(vs[1])), 2))])
	return strconv.ParseInt(fmt.Sprintf("%s%s", v1, v2), 10, 64)
}
