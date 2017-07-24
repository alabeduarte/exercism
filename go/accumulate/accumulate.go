package accumulate

const testVersion = 1

func Accumulate(collection []string, operation func(string) string) []string {
	accumulatedCollection := make([]string, len(collection), cap(collection))

	for i, elem := range collection {
		accumulatedCollection[i] = operation(elem)
	}

	return accumulatedCollection
}
