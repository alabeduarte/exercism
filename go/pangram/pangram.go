package pangram

import "strings"

const testVersion = 1

func ToMap(array []string) map[string]string {
	mapOfElements := make(map[string]string)

	for _, elem := range array {
		mapOfElements[elem] = elem
	}

	return mapOfElements
}

func IsPangram(phrase string) bool {
	alphabet := ToMap(strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", ""))

	phrase = strings.ToUpper(strings.Replace(phrase, " ", "", -1))
	phraseArray := strings.Split(phrase, "")

	charactersFound := make(map[string]string)

	for _, elem := range phraseArray {
		_, exist := alphabet[elem]

		if exist {
			charactersFound[elem] = elem
		}
	}

	if len(charactersFound) == len(alphabet) {
		return true
	}

	return false
}
