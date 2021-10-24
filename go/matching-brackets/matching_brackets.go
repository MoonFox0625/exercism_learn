package brackets

type Stack []rune

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return 0
	}
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

var bracketMap = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

func Bracket(input string) bool {
	stack := Stack{}
	for _, char := range input {
		if char == '{' || char == '[' || char == '(' {
			stack.Push(char)
		}
		if char == '}' || char == ']' || char == ')' {
			if stack.IsEmpty() {
				return false
			}
			value := stack.Pop()
			if value != bracketMap[char] {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
