package bob

// Reaction represents the different reactions that Bob can have
// taking into consideration his criteria
type Reaction interface {
	IsCriterionMet(remark string) bool
	Response() string
}
