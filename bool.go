package valkyrie

type BoolValidator func(arg bool) error

type Bool struct {
	Rules []BoolValidator
}

func (b *Bool) Validate(arg interface{}, customErr error) error {
	value, ok := arg.(bool)
	if !ok {
		return orErr(customErr, errShouldBeBool)
	}

	for _, validator := range b.Rules {
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