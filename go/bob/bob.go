package bob

import (
	"regexp"
	"strings"
)

type bobReaction interface {
	isCriterionMet(remark string) bool
	response() string
}

type yelledQuestion struct{}
type justYell struct{}
type question struct{}
type silence struct{}

func (y yelledQuestion) isCriterionMet(remark string) bool {
	r, _ := regexp.Compile("^[^a-z0-9|:]+\\?$")

	return r.MatchString(remark)
}

func (y yelledQuestion) response() string {
	return "Calm down, I know what I'm doing!"
}

func (y justYell) isCriterionMet(remark string) bool {
	r, _ := regexp.Compile("^[^a-z]+(\\.|\\!|[A-Z])$")

	return r.MatchString(remark)
}

func (y justYell) response() string {
	return "Whoa, chill out!"
}

func (q question) isCriterionMet(remark string) bool {
	r, _ := regexp.Compile(".+\\?$")

	return r.MatchString(remark)
}

func (q question) response() string {
	return "Sure."
}

func (s silence) isCriterionMet(remark string) bool {
	r, _ := regexp.Compile("^\\s+$")

	return remark == "" || r.MatchString(remark)
}

func (s silence) response() string {
	return "Fine. Be that way!"
}

// Hey returns what Bob would answer based on the type of remark
func Hey(remark string) string {
	reactions := [4]bobReaction{yelledQuestion{}, justYell{}, question{}, silence{}}

	for _, reaction := range reactions {
		if reaction.isCriterionMet(strings.Trim(remark, " ")) {
			return reaction.response()
		}
	}

	return "Whatever."
}
