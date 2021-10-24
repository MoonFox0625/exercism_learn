package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(input string) string {
	input = strings.ToLower(input)
	words := strings.FieldsFunc(input, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsSymbol(r) || unicode.IsPunct(r)
	})
	str := strings.Join(words, "")
	size := len(str)
	r, c := GetRectRC(size)
	columns := make([]string, c)
	for i, v := range str {
		columns[i%c] += string(v)
	}
	for i, column := range columns {
		if len(column) < r {
			columns[i] += strings.Repeat(" ", r-len(columns[i]))
		}
	}
	output := strings.Join(columns, " ")
	return output
}
func GetRectRC(size int) (r, c int) { // r<=c c-r<=1
	r = int(math.Sqrt(float64(size)))
	for {
		if r*r >= size {
			c = r
			return
		}
		if r*(r+1) >= size {
			c = r + 1
			return
		}
		r++
	}
}
