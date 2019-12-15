package bob

import "regexp"

// Question represents Bob's answers
type Question struct{}

// IsCriterionMet evaluates a question, such as "How are you?".
func (q Question) IsCriterionMet(remark string) bool {
	r, _ := regexp.Compile(".+\\?$")

	return r.MatchString(remark)
}

// Response returns 'Sure.'
func (q Question) Response() string {
	return "Sure."
}
