package strain

type Ints []int
type Lists [][]int
type Strings []string

func (list Ints) Keep(pred func(int) bool) Ints {
	if list == nil {
		return nil
	}
	output := make(Ints, 0)
	for _, v := range list {
		if pred(v) {
			output = append(output, v)
		}
	}
	return output
}
func (list Ints) Discard(pred func(int) bool) Ints {
	if list == nil {
		return nil
	}
	output := make(Ints, 0)

	for _, v := range list {
		if !pred(v) {
			output = append(output, v)
		}
	}
	return output
}

func (lists Lists) Keep(pred func([]int) bool) Lists {
	if lists == nil {
		return nil
	}
	output := make(Lists, 0)
	for _, list := range lists {
		if pred(list) {
			output = append(output, list)
		}
	}
	return output
}

func (strs Strings) Keep(pred func(string) bool) Strings {
	if strs == nil {
		return nil
	}
	output := make(Strings, 0)
	for _, str := range strs {
		if pred(str) {
			output = append(output, str)
		}
	}
	return output
}
