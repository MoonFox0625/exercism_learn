package pangram

import "strings"

func IsPangram(input string) bool {
	if input == "" {
		return false
	}
	alphabet := make(map[rune]bool)
	var count int
	for _, char := range strings.ToLower(input) {
		if char < 'a' || char > 'z' {
			continue
		}
		if _, ok := alphabet[char]; !ok {
			alphabet[char] = true
			count++
		}
	}
	return count == 26
}
