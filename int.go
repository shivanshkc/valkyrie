package valkyrie

type IntValidator func(arg int) error

type Int struct {
	Rules []IntValidator
}

func (i *Int) Validate(arg interface{}, customErr error) error {
	value, ok := arg.(int)
	if !ok {
		return orErr(customErr, errShouldBeInt)
	}

	for _, validator := range i.Rules {
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