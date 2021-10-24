package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	input = strings.ToLower(input)
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r != '\'' && (unicode.IsSpace(r) ||
			unicode.IsPunct(r) || unicode.IsSymbol(r))
	})
	output := make(Frequency)
	for _, word := range words {
		if word[0] == '\'' && word[len(word)-1] == '\'' {
			word = word[1 : len(word)-1]
		}
		output[word]++
	}
	return output
}

//func WordCount(input string) Frequency {
//	input = ReplaceMark(input)
//	words := strings.Fields(input)
//	output := make(Frequency)
//	for _, word := range words {
//		word = strings.ToLower(word)
//		tmpWord := make([]rune, 0)
//		for i, char := range word {
//			if char == '\'' && (i > 0 && i < len(word)-1) {
//				tmpWord = append(tmpWord, char)
//			} else if unicode.IsLetter(char) || unicode.IsDigit(char) {
//				tmpWord = append(tmpWord, char)
//			}
//		}
//		newWord := string(tmpWord)
//		if newWord != "" {
//
//			output[newWord]++
//		}
//	}
//	return output
//}
//
//func ReplaceMark(input string) string {
//	const mark = "?!.;,*"
//	for _, char := range mark {
//		input = strings.ReplaceAll(input, string(char), " ")
//	}
//	return input
//}
