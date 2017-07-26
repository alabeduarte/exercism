package acronym

import "strings"

const testVersion = 3

func Abbreviate(phrase string) string {
	phraseWithoutSlashes := strings.Replace(phrase, "-", " ", -1)
	acronym := ""

	for _, elem := range strings.Split(phraseWithoutSlashes, " ") {
		acronym += strings.ToUpper(string(elem[0]))
	}

	return acronym
}
