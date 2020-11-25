package valkyrie

type Rule interface {
	Apply(arg interface{}) error
}
