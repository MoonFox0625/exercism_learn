package encode

import (
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}
	var output string
	prev := input[0]
	count := 1
	for i := 1; i < len(input); i++ {
		recent := input[i]
		if recent == prev {
			count++
		} else {
			if count == 1 {
				output += string(prev)
			} else {
				output += strconv.Itoa(count) + string(prev)
				count = 1
			}
			prev = recent
		}
	}
	if count == 1 {
		output += string(prev)
	} else {
		output += strconv.Itoa(count) + string(prev)
		count = 1
	}
	return output
}

func RunLengthDecode(input string) string {
	var output string
	if input == "" {
		return ""
	}
	count := 1
	preDigit := false
	for _, char := range input {
		if unicode.IsLetter(char) || char == ' ' {
			if count == 1 {
				output += string(char)
			} else {
				for count > 0 {
					output += string(char)
					count--
				}
				//output += strconv.Itoa(count) + string(char)
				count = 1
				preDigit = false
			}
		} else if unicode.IsDigit(char) {
			if !preDigit {
				count = int(char - '0')
				preDigit = true
			} else {
				count = count*10 + int(char-'0')
			}
		}
	}
	return output
}
