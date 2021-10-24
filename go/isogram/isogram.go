package isogram

import "strings"

func IsIsogram(input string) bool {
	if input == "" {
		return true
	}
	tag := make(map[rune]bool)
	for _, char := range strings.ToLower(input) {
		if char == '-' || char == ' ' {
			continue
		}
		if _, ok := tag[char]; ok {
			return false
		}
		tag[char] = true
	}
	return true
}
