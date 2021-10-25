package series

func All(n int, s string) []string {
	size := len(s)
	var output []string
	var leftIndex = 0
	if n > size {
		return nil
	} else if n == size {
		return []string{s}
	} else {
		for GetRightIndex(leftIndex, n) <= size {
			rightIndex := GetRightIndex(leftIndex, n)
			if leftIndex == rightIndex {
				output = append(output, string(s[leftIndex]))
			} else {
				output = append(output, s[leftIndex:rightIndex])
			}
			leftIndex++
		}
	}
	return output
}
func GetRightIndex(leftIndex, size int) int {
	// if size == 1 {
	// 	return leftIndex
	// }
	return leftIndex + size
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
func First(n int, s string) (first string, ok bool) {
	size := len(s)
	if n > size || n < 1 {
		return
	} else {
		first = s[:n]
		ok = true
	}
	return
}
