package accumulate

func Accumulate(collection []string, operation func(string) string) []string {
	for i, str := range collection {
		collection[i] = operation(str)
	}
	return collection
}
