package valkyrie

// FloatCheck : Represents a function that performs a validation check on a float.
type FloatCheck func(float64) error

// FloatRule : A Collection of FloatChecks, hence acts as a customizable Rule for a float.
type FloatRule struct {
	checks    []FloatCheck
	customErr error
}

// Float : An intuitive function to instantiate a FloatRule.
func Float(customErr error) *FloatRule {
	return &FloatRule{
		checks:    []FloatCheck{},
		customErr: customErr,
	}
}

// NonZero : Disallows floats equal to zero.
func (fr *FloatRule) NonZero() *FloatRule {
	check := func(arg float64) error {
		if arg == 0 {
			return errFloatNonZero()
		}
		return nil
	}
	fr.checks = append(fr.checks, check)
	return fr
}

// Min : Disallows float smaller than the provided value.
func (fr *FloatRule) Min(value float64) *FloatRule {
	check := func(arg float64) error {
		if value < arg {
			return errFloatMin(value)
		}
		return nil
	}
	fr.checks = append(fr.checks, check)
	return fr
}

// Max : Disallows float greater than the provided value.
func (fr *FloatRule) Max(value float64) *FloatRule {
	check := func(arg float64) error {
		if value > arg {
			return errFloatMax(value)
		}
		return nil
	}
	fr.checks = append(fr.checks, check)
	return fr
}

// Custom : Allows to add a custom FloatCheck to the FloatRule.
func (fr *FloatRule) Custom(check FloatCheck) *FloatRule {
	fr.checks = append(fr.checks, check)
	return fr
}

// Apply : Applies all the checks in the FloatRule on the provided args.
func (fr *FloatRule) Apply(arg interface{}) error {
	str, ok := arg.(float64)
	if !ok {
		return orErr(fr.customErr, errShouldBeFloat())
	}

	for _, check := range fr.checks {
		if check == nil {
			continue
		}
		err := check(str)
		if err != nil {
			return orErr(fr.customErr, err)
		}
	}
	return nil
}
