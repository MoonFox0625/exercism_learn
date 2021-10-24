// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral 等边
	Iso             // isosceles 等腰
	Sca             // scalene  不等边三角形
)

func IsTriangle(a, b, c float64) bool {
	var condition1, condition2 bool
	if a+b > c && a+c > b && b+c > a {
		condition1 = true
	}
	// to check for degenerate triangles.
	if IsPositive(a, b, c) && !IsInf(a, b, c) && (a+b == c || a+c == b || b+c == a) {
		condition2 = true
	}
	return condition1 || condition2
}
func IsInf(a, b, c float64) bool {
	// return math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
	// to avoid check everyone
	return math.IsInf(a+b+c, 0)
}
func IsPositive(a, b, c float64) bool {
	return a > 0 && b > 0 && c > 0
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind
	if !IsTriangle(a, b, c) {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
