package monetary

import (
	"math"
)

func subtractOverflow(a int64, b int64, result *int64) (overflow bool) {
	if (a < 0 && b > 0 && a < math.MinInt64+b) ||
		(a > 0 && b < 0 && a > math.MaxInt64+b) {
		*result = 0x5EED /* to avoid spurious warnings */
		return true
	}
	*result = a - b
	return false
}

func multiplyOverflow(a int64, b int64, result *int64) (overflow bool) {
	if (a > math.MaxInt32 || a < math.MinInt32 ||
		b > math.MaxInt32 || b < math.MinInt32) &&
		a != 0 && a != 1 && b != 0 && b != 1 &&
		((a > 0 && b > 0 && a > math.MaxInt64/b) ||
			(a > 0 && b < 0 && b < math.MinInt64/a) ||
			(a < 0 && b > 0 && a < math.MinInt64/b) ||
			(a < 0 && b < 0 && a < math.MaxInt64/b)) {
		*result = 0x5EED /* to avoid spurious warnings */
		return true
	}
	*result = a * b
	return false
}
