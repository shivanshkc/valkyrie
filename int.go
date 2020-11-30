package valkyrie

// IntCheck : Represents a function that performs a validation check on an int64.
type IntCheck func(arg int64) error

// IntRule : Rule interface implementation for an int64.
type IntRule struct {
	from   string
	whites []int64
	checks []IntCheck
	err    error
}

// IntRule PRIMARY PUBLIC METHODS ###################################

// Allow : Whitelists the provided int(s) for a rule.
// If the argument is one of the whitelisted values, no checks
// will be performed upon it.
func (i *IntRule) Allow(args ...int64) *IntRule {
	i.whites = append(i.whites, args...)
	return i
}

// AddCheck : Adds a custom check function to the rule.
func (i *IntRule) AddCheck(check IntCheck) *IntRule {
	i.checks = append(i.checks, check)
	return i
}

// WithError : Adds a custom error to the rule.
// This custom error (if not nil) will be thrown on every check violation
// instead of the original error.
func (i *IntRule) WithError(err error) *IntRule {
	i.err = err
	return i
}

// Apply : Applies the rule on a given argument.
func (i *IntRule) Apply(arg interface{}) error {
	intVal, err := toInt64(arg, i.from)
	if err != nil {
		return orErr(i.err, errInt64(i.from))
	}
	if i.isWhitelisted(intVal) {
		return nil
	}

	if err := i.performChecks(intVal); err != nil {
		return orErr(i.err, err)
	}
	return nil
}

// IntRule CONSTRUCTORS #############################################

// FloatInt : Creates an IntRule which expects the arg to be a float64.
// which will be validated after conversion to int64.
// Example: 23.12 -> 23
func FloatInt() *IntRule {
	return &IntRule{from: floatType}
}

// StringInt : Creates an IntRule which expects the arg to be a string.
// which will be validated after conversion to int64.
// Example: "23" -> 23, note that "23.34" or any float will throw an error.
func StringInt() *IntRule {
	return &IntRule{from: stringType}
}

// PureInt : Creates an IntRule which expects the arg to be an int64.
func PureInt() *IntRule {
	return &IntRule{from: intType}
}

// IntRule PRIVATE METHODS ##########################################

func (i *IntRule) isWhitelisted(value int64) bool {
	for _, white := range i.whites {
		if white == value {
			return true
		}
	}
	return false
}

func (i *IntRule) performChecks(arg int64) error {
	for _, check := range i.checks {
		if check == nil {
			continue
		}
		if err := check(arg); err != nil {
			return err
		}
	}
	return nil
}

// IntRule UTILITY PUBLIC METHODS  ##################################
