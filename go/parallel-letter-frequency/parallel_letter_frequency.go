package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
func ConcurrentFrequency(inputs []string) FreqMap {
	outputMap := FreqMap{}
	c := make(chan FreqMap)
	for _, input := range inputs {
		go func(input string) {
			c <- Frequency(input)
		}(input)
	}
	for _ = range inputs {
		m := <-c
		for key, value := range m {
			outputMap[key] += value
		}
	}
	return outputMap
}
