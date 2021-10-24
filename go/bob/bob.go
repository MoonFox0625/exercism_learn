// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	lastChar := remark[len(remark)-1]
	// 保证remark 不为空 且为字母
	if remark == strings.ToUpper(remark) && remark != strings.ToLower(remark) {
		if lastChar == '?' {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}
	if lastChar == '?' {
		return "Sure."
	}
	return "Whatever."

}
func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
