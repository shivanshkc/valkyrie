package valkyrie

// BoolCheck : Represents a function that performs a validation check on a bool.
type BoolCheck func(bool) error

// BoolRule : A Collection of BoolChecks, hence acts as a customizable Rule for a bool.
type BoolRule struct {
	checks    []BoolCheck
	customErr error
}

// Bool : An intuitive function to instantiate a BoolRule.
func Bool(customErr error) *BoolRule {
	return &BoolRule{
		checks:    []BoolCheck{},
		customErr: customErr,
	}
}

// True : Enforces the bool to be true.
func (br *BoolRule) True() *BoolRule {
	check := func(arg bool) error {
		if !arg {
			return errShouldBeTrue()
		}
		return nil
	}

	br.checks = append(br.checks, check)
	return br
}

// False : Enforces the bool to be false.
func (br *BoolRule) False() *BoolRule {
	check := func(arg bool) error {
		if arg {
			return errShouldBeFalse()
		}
		return nil
	}

	br.checks = append(br.checks, check)
	return br
}

// Custom : Allows to add a custom BoolCheck to the BoolRule.
func (br *BoolRule) Custom(check BoolCheck) *BoolRule {
	br.checks = append(br.checks, check)
	return br
}

// Apply : Applies all the checks in the BoolRule on the provided args.
func (br *BoolRule) Apply(arg interface{}) error {
	str, ok := arg.(bool)
	if !ok {
		return orErr(br.customErr, errShouldBeBool())
	}

	for _, check := range br.checks {
		if check == nil {
			continue
		}
		err := check(str)
		if err != nil {
			return orErr(br.customErr, err)
		}
	}
	return nil
}
