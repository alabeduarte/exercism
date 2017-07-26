package pangram

import "strings"

const testVersion = 1

func IsPangram(phrase string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	phrase = strings.ToLower(phrase)

	for _, character := range alphabet {
		if !strings.Contains(phrase, string(character)) {
			return false
		}
	}

	return true
}
