package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		var tag bool
		for _, divisor := range divisors {
			if divisor == 0 {
				continue
			}
			if i%divisor == 0 {
				tag = true
				break
			}
		}
		if tag {
			sum += i
		}
	}
	return sum
}
