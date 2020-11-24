package valkyrie

type Value struct {
	Rules []ValidatorFunc
}

func (v *Value) Validate(arg interface{}, customErr error) error {
	return nil
}
