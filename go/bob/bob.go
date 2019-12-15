package bob

import "strings"

// Hey returns what Bob would answer based on the type of remark
func Hey(remark string) string {
	reactions := [4]Reaction{YelledQuestion{}, JustYell{}, Question{}, Silence{}}

	for _, reaction := range reactions {
		if reaction.IsCriterionMet(strings.Trim(remark, " ")) {
			return reaction.Response()
		}
	}

	return "Whatever."
}
