package valkyrie

type FloatValidator func(arg float64) error

type Float struct {
	Rules []FloatValidator
}

func (f *Float) Validate(arg interface{}, customErr error) error {
	value, ok := arg.(float64)
	if !ok {
		return orErr(customErr, errShouldBeFloat)
	}

	for _, validator := range f.Rules {
		if validator == nil {
			continue
		}
		err := validator(value)
		if err == nil {
			continue
		}
		return orErr(customErr, err)
	}

	return nil
}
