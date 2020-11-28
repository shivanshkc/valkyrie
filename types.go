package valkyrie

// Rule : The Rule interface, implemented by StringRule and MapRule etc.
type Rule interface {
	Apply(arg interface{}) error
}
