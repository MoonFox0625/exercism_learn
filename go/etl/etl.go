package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	var output = make(map[string]int)
	for val, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = val
		}
	}
	return output
}
