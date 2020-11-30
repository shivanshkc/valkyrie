package valkyrie

// FloatCheck : Represents a function that performs a validation check on a float64.
type FloatCheck func(arg float64) error

// FloatRule : Rule interface implementation for a float64.
type FloatRule struct {
	from   string
	whites []float64
	checks []FloatCheck
	err    error
}

// FloatRule PRIMARY PUBLIC METHODS #################################

// Allow : Whitelists the provided float(s) for a rule.
// If the argument is one of the whitelisted values, no checks
// will be performed upon it.
func (f *FloatRule) Allow(args ...float64) *FloatRule {
	f.whites = append(f.whites, args...)
	return f
}

// AddCheck : Adds a custom check function to the rule.
func (f *FloatRule) AddCheck(check FloatCheck) *FloatRule {
	f.checks = append(f.checks, check)
	return f
}

// WithError : Adds a custom error to the rule.
// This custom error (if not nil) will be thrown on every check violation
// instead of the original error.
func (f *FloatRule) WithError(err error) *FloatRule {
	f.err = err
	return f
}

// Apply : Applies the rule on a given argument.
func (f *FloatRule) Apply(arg interface{}) error {
	floatVal, err := toFloat64(arg, f.from)
	if err != nil {
		return orErr(f.err, errFloat64(f.from))
	}
	if f.isWhitelisted(floatVal) {
		return nil
	}

	if err := f.performChecks(floatVal); err != nil {
		return orErr(f.err, err)
	}
	return nil
}

// FloatRule CONSTRUCTORS #############################################

// IntFloat : Creates a FloatRule which expects the arg to be an int64.
// which will be validated after conversion to float64.
// Example: 23 -> 23.00
func IntFloat() *FloatRule {
	return &FloatRule{from: intType}
}

// StringFloat : Creates an FloatRule which expects the arg to be a string.
// which will be validated after conversion to float64.
// Example: "23.23" -> 23.23
func StringFloat() *FloatRule {
	return &FloatRule{from: stringType}
}

// PureFloat : Creates an FloatRule which expects the arg to be a float64.
func PureFloat() *FloatRule {
	return &FloatRule{from: floatType}
}

// FloatRule PRIVATE METHODS ########################################

func (f *FloatRule) isWhitelisted(value float64) bool {
	for _, white := range f.whites {
		if white == value {
			return true
		}
	}
	return false
}

func (f *FloatRule) performChecks(arg float64) error {
	for _, check := range f.checks {
		if check == nil {
			continue
		}
		if err := check(arg); err != nil {
			return err
		}
	}
	return nil
}

// FloatRule UTILITY PUBLIC METHODS  ################################

// GTE : Adds a '>=' check to the rule.
func (f *FloatRule) GTE(value float64) *FloatRule {
	f.AddCheck(func(arg float64) error {
		if arg < value {
			return errFloatGTE(value)
		}
		return nil
	})
	return f
}

// LTE : Adds a '<=' check to the rule.
func (f *FloatRule) LTE(value float64) *FloatRule {
	f.AddCheck(func(arg float64) error {
		if arg > value {
			return errFloatLTE(value)
		}
		return nil
	})
	return f
}

// GT : Adds a '>' check to the rule.
func (f *FloatRule) GT(value float64) *FloatRule {
	f.AddCheck(func(arg float64) error {
		if arg <= value {
			return errFloatGT(value)
		}
		return nil
	})
	return f
}

// LT : Adds a '<' check to the rule.
func (f *FloatRule) LT(value float64) *FloatRule {
	f.AddCheck(func(arg float64) error {
		if arg >= value {
			return errFloatLT(value)
		}
		return nil
	})
	return f
}
