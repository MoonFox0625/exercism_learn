package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed > 11 {
		return 0
	} else if speed > 8 {
		return 0.77
	} else if speed > 4 {
		return 0.9
	} else if speed > 0 {
		return 1
	}
	return 0
}

const producePerHour = 221

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return producePerHour * float64(speed) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	produce := CalculateProductionRatePerHour(speed)
	if produce > limit {
		return limit
	}
	return produce
}
