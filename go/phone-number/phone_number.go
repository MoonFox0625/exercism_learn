package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

var errInvalidNumber = errors.New("input number is invalid")

func Number(input string) (string, error) {
	number := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		} else {
			return -1
		}
	}, input)

	if strings.HasPrefix(number, "1") {
		number = number[1:]
	}

	if len(number) != 10 {
		return "", errInvalidNumber
	}
	if !(isN(number[0]) && isN(number[3])) {
		return "", errInvalidNumber
	}

	return number, nil
}
func isN(r uint8) bool {
	if r >= '2' && r <= '9' {
		return true
	}
	return false
}

func AreaCode(input string) (string, error) {
	num, err := Number(input)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

func Format(input string) (string, error) {
	num, err := Number(input)
	if err != nil {
		return "", err
	}
	output := fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:])
	return output, nil
}
