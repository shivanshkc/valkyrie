package valkyrie

type Validatable interface {
	Validate(arg interface{}, customErr error) error
}

