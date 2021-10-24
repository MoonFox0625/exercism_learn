package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) (output []string) {
	subjectSort := GetSortString(subject)
	for _, candidate := range candidates {
		if len(subjectSort) != len(candidate) ||
			strings.EqualFold(subject, candidate) {
			continue
		}
		//if strings.ToLower(subject) == strings.ToLower(candidate) {
		//	continue
		//}
		if subjectSort == GetSortString(candidate) {
			output = append(output, candidate)
		}
	}
	return output
}

func GetSortString(input string) string {
	runeInput := []rune(strings.ToLower(input))
	sort.Slice(runeInput, func(i, j int) bool {
		return runeInput[i] < runeInput[j]
	})

	return string(runeInput)
}
