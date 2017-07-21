package hamming

import (
	"errors"
	"strings"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings must have the same size")
	}

	if a == b {
		return 0, nil
	}

	arrayA := strings.SplitAfter(a, "")
	arrayB := strings.SplitAfter(b, "")

	counter := 0
	for i := range arrayA {
		if arrayA[i] != arrayB[i] {
			counter++
		}
	}

	return counter, nil
}
