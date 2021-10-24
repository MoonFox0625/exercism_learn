package romannumerals

import "errors"

type Roman struct {
	num int
	str string
}

var romanTable = []Roman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", errors.New("invalid num")
	}
	var output string
	for _, roman := range romanTable {
		if input == 0 {
			break
		}
		var count int
		count, input = Divmod(input, roman)
		for count > 0 {
			output += roman.str
			count--
		}
	}
	return output, nil

}

// Divmod is 返回值是商和余数，参数分别代表分子和分母
func Divmod(numerator int, roman Roman) (quotient, remainder int) {
	denominator := roman.num
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}
