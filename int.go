package valkyrie

// IntCheck : Represents a function that performs a validation check on an int.
type IntCheck func(int) error

// IntRule : A Collection of IntChecks, hence acts as a customizable Rule for an int.
type IntRule struct {
	checks    []IntCheck
	customErr error
}

// Int : An intuitive function to instantiate a IntRule.
func Int(customErr error) *IntRule {
	return &IntRule{
		checks:    []IntCheck{},
		customErr: customErr,
	}
}

// Min : Disallows int smaller than the provided value.
func (ir *IntRule) Min(value int) *IntRule {
	check := func(arg int) error {
		if value < arg {
			return errIntMin(value)
		}
		return nil
	}
	ir.checks = append(ir.checks, check)
	return ir
}

// Max : Disallows int greater than the provided value.
func (ir *IntRule) Max(value int) *IntRule {
	check := func(arg int) error {
		if value > arg {
			return errIntMax(value)
		}
		return nil
	}
	ir.checks = append(ir.checks, check)
	return ir
}

// Custom : Allows to add a custom IntCheck to the IntRule.
func (ir *IntRule) Custom(check IntCheck) *IntRule {
	ir.checks = append(ir.checks, check)
	return ir
}

// Apply : Applies all the checks in the IntRule on the provided args.
func (ir *IntRule) Apply(arg interface{}) error {
	str, ok := arg.(int)
	if !ok {
		return orErr(ir.customErr, errShouldBeInt())
	}

	for _, check := range ir.checks {
		if check == nil {
			continue
		}
		err := check(str)
		if err != nil {
			return orErr(ir.customErr, err)
		}
	}
	return nil
}
