package valkyrie

type ValidatorFunc func(arg interface{}) error

type Validatable interface {
	Validate(arg interface{}, customErr error) error
}
