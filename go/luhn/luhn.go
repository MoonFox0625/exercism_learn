package luhn

import (
	"strings"
	"unicode"
)

func Valid(input string) bool {
	var sum int
	var tag bool = false
	input = strings.Join(strings.Fields(input), "")
	if len(input) < 2 {
		return false
	}
	for i := len(input) - 1; i >= 0; i = i - 1 {
		var char = rune(input[i])
		if !unicode.IsDigit(char) {
			return false
		}
		num := int(char - '0')
		if tag {
			if num*2 > 9 {
				num = num*2 - 9
			} else {
				num = num * 2
			}
			sum += num
		} else {
			sum += num
		}
		tag = !tag
	}

	return sum%10 == 0
}
