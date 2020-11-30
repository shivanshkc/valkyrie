package valkyrie

// Rule : Representation of a basic validation rule.
type Rule interface {
	Apply(arg interface{}) error
}
