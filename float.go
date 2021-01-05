package valkyrie

// FloatCheck : Represents a function that performs a validation check on a float64.
type FloatCheck func(arg float64) error

// FloatRule : Rule interface implementation for a float64.
type FloatRule struct {
	// base : base is the name of the type from which the float value will be inferred.
	base string
	// whites : the list of whitelisted values for this rule.
	whites []interface{}
	// checks : the list of checks to be performed as part of this rule.
	checks []FloatCheck
	// err : the error to be thrown if the rule fails.
	err error
}

// FloatRule PRIMARY PUBLIC METHODS #################################

// Allow : Whitelists the provided values for a rule.
// If the argument is one of the whitelisted values, no checks
// will be performed upon it.
func (f *FloatRule) Allow(args ...interface{}) *FloatRule {
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
	if f.isWhitelisted(arg) {
		return nil
	}
	floatVal, err := toFloat64(arg, f.base)
	if err != nil {
		return orErr(f.err, errFloat64(f.base))
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
	return &FloatRule{base: intType}
}

// StringFloat : Creates an FloatRule which expects the arg to be a string.
// which will be validated after conversion to float64.
// Example: "23.23" -> 23.23
func StringFloat() *FloatRule {
	return &FloatRule{base: stringType}
}

// PureFloat : Creates an FloatRule which expects the arg to be a float64.
func PureFloat() *FloatRule {
	return &FloatRule{base: floatType}
}

// FloatRule PRIVATE METHODS ########################################

func (f *FloatRule) isWhitelisted(value interface{}) bool {
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

// Except : Invalidates if arg == provided value
func (f *FloatRule) Except(value float64) *FloatRule {
	f.AddCheck(func(arg float64) error {
		if arg == value {
			return errFloatExcept(value)
		}
		return nil
	})
	return f
}
