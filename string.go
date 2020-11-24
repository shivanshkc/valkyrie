package valkyrie

type StringValidator func(arg string) error

type String struct {
	Rules []StringValidator
}

func (s *String) Validate(arg interface{}, customErr error) error {
	value, ok := arg.(string)
	if !ok {
		return orErr(customErr, errShouldBeString)
	}

	for _, validator := range s.Rules {
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