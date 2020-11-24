package valkyrie

type Value struct {
	Rules []ValidatorFunc
}

func (v *Value) Validate(arg interface{}, customErr error) error {
	for _, rule := range v.Rules {
		if rule == nil {
			continue
		}
		if err := rule(arg); err != nil {
			return orErr(customErr, err)
		}
	}
	return nil
}
