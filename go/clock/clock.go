package clock

import "fmt"

type Clock struct {
	Hour int
	Min  int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Min)
}

func (c Clock) Add(min int) Clock {
	return New(c.Hour, c.Min+min)
}
func (c Clock) Subtract(min int) Clock {
	return New(c.Hour, c.Min-min)
}

func New(hour int, min int) Clock {
	carry := 0 // 分进位
	if min >= 0 {
		carry = min / 60
		min %= 60
	} else {
		carry = min / 60
		min = min % 60
		if min < 0 {
			carry--
			min += 60
		}
	}
	hour += carry
	if hour >= 0 {
		hour %= 24
	} else {
		hour = hour % 24
		if hour < 0 {
			hour += 24
		}
	}
	return Clock{Hour: hour, Min: min}
}
