package internal

import (
	"bytes"
)

func SplitMonetaryBytes(data []byte) [][]byte {
	return bytes.Split(data, []byte{10})
}
