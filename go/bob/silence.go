package bob

import "regexp"

// Silence represents Bob's answers
type Silence struct{}

// IsCriterionMet evaluates if you address him without actually saying anything.
func (s Silence) IsCriterionMet(remark string) bool {
	r, _ := regexp.Compile("^\\s+$")

	return remark == "" || r.MatchString(remark)
}

// Response returns 'Fine. Be that way!'
func (s Silence) Response() string {
	return "Fine. Be that way!"
}
