package listops

type IntList []int

// // 待会测试能否换个文件的位置
// type binFunc func(int, int) int
// type predFunc func(int) bool
// type unaryFunc func(int) int

func (myList IntList) Append(other IntList) IntList {
	myList = append(myList, other...)
	return myList
}

func (myList IntList) Concat(others []IntList) IntList {
	for _, other := range others {
		// myList = append(myList, other...)
		myList = myList.Append(other)
	}
	return myList
}

// Filter is given a predicate and a list,
// return the list of all items for which predicate(item) is True)
func (myList IntList) Filter(predicate func(n int) bool) IntList {
	var newList = make(IntList, 0)
	// newList := IntList{}
	// var newList IntList //不能这么初始化
	for _, val := range myList {
		if predicate(val) {
			newList = append(newList, val)
		}
	}
	return newList
}

func (myList IntList) Length() int {
	return len(myList)
}

func (myList IntList) Map(fn func(int) int) IntList {
	for i, v := range myList {
		myList[i] = fn(v)
	}
	return myList
}

func (myList IntList) Reverse() IntList {

	for i, j := 0, len(myList)-1; i < j; i, j = i+1, j-1 {
		myList[i], myList[j] = myList[j], myList[i]
	}

	return myList
}

func (myList IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial
	for _, v := range myList {
		result = fn(result, v)
	}
	return result

}

func (myList IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial
	for i := myList.Length() - 1; i >= 0; i-- {
		result = fn(myList[i], result)
	}
	return result
}
