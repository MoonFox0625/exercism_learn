package scrabble

import "strings"

// var letterValue = map[rune]int{
// 	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
// 	'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
// 	'D': 2, 'G': 2,
// 	'B': 3, 'C': 3, 'M': 3, 'P': 3,
// 	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
// 	'K': 5,
// 	'J': 8, 'X': 8,
// 	'Q': 10, 'Z': 10,
// }

var letterValue = make(map[rune]int)

func init() {
	levels := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}
	for str, val := range levels {
		for _, char := range str {
			letterValue[char] = val
		}
	}
}
func Score(input string) int {
	if input == "" {
		return 0
	}
	input = strings.ToUpper(input)
	score := 0
	for _, val := range input {
		score += letterValue[val]
	}
	return score

}
