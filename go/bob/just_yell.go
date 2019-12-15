package bob

import "regexp"

// JustYell represents Bob's answers
type JustYell struct{}

// IsCriterionMet evaluates if you YELL AT HIM (in all capitals).
func (j JustYell) IsCriterionMet(remark string) bool {
	r, _ := regexp.Compile("^[^a-z]+(\\.|\\!|[A-Z])$")

	return r.MatchString(remark)
}

// Response returns 'Whoa, chill out!'
func (j JustYell) Response() string {
	return "Whoa, chill out!"
}
