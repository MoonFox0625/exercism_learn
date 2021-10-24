// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) (output []string) {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	if len(rhyme) == 0 {
		return nil
	}
	prev := rhyme[0]
	for i := 1; i < len(rhyme); i++ {
		recent := rhyme[i]
		sentence := fmt.Sprintf("For want of a %s the %s was lost.",
			prev, recent)
		output = append(output, sentence)
		prev = recent

	}
	last := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	output = append(output, last)

	return output
}
