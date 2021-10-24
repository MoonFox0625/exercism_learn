package darts

import "math"

const (
	outerRadius  = 10
	middleRadius = 5
	innerRadius  = 1
)

func Score(x, y float64) (score int) {
	radius := math.Sqrt(x*x + y*y)
	if radius > outerRadius {
		score = 0
	} else if radius > middleRadius {
		score = 1
	} else if radius > innerRadius {
		score = 5
	} else {
		score = 10
	}
	return score
}
