package bob

import "regexp"

// YelledQuestion represents Bob's answers
type YelledQuestion struct{}

// IsCriterionMet evaluates if you yell a question at him.
func (y YelledQuestion) IsCriterionMet(remark string) bool {
	r, _ := regexp.Compile("^[^a-z0-9|:]+\\?$")

	return r.MatchString(remark)
}

// Response returns 'Calm down, I know what I'm doing!'
func (y YelledQuestion) Response() string {
	return "Calm down, I know what I'm doing!"
}
