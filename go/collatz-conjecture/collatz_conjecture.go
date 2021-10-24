package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, errors.New("input should be positive")
	}
	var steps int
	for {
		if input == 1 {
			break
		}
		if input%2 == 0 {
			input = input / 2
			steps++
		} else {
			input = input*3 + 1
			steps++
		}
	}

	return steps, nil

}
