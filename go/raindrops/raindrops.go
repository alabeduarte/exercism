package raindrops

import "strconv"

const testVersion = 3

func Convert(number int) string {
	const pling, plang, plong = 3, 5, 7
	result := ""

	factors := []int{pling, plang, plong}
	mapOfFactors := map[int]string{
		pling: "Pling",
		plang: "Plang",
		plong: "Plong",
	}

	for _, elem := range factors {
		if number%elem == 0 {
			result += mapOfFactors[elem]
		}
	}

	if result == "" {
		return strconv.Itoa(number)
	}

	return result
}
