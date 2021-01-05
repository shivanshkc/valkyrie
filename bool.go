package valkyrie

// BoolCheck : Represents a function that performs a validation check on a bool.
type BoolCheck func(arg bool) error

// BoolRule : Rule interface implementation for a bool.
type BoolRule struct {
	// base : base is the name of the type from which the bool value will be inferred.
	base string
	// whites : the list of whitelisted values for this rule.
	whites []interface{}
	// checks : the list of checks to be performed as part of this rule.
	checks []BoolCheck
	// err : the error to be thrown if the rule fails.
	err error
}

// BoolRule PRIMARY PUBLIC METHODS ##################################

// Allow : Whitelists the provided values for a rule.
// If the argument is one of the whitelisted values, no checks
// will be performed upon it.
func (b *BoolRule) Allow(args ...interface{}) *BoolRule {
	b.whites = append(b.whites, args...)
	return b
}

// AddCheck : Adds a custom check function to the rule.
func (b *BoolRule) AddCheck(check BoolCheck) *BoolRule {
	b.checks = append(b.checks, check)
	return b
}

// WithError : Adds a custom error to the rule.
// This custom error (if not nil) will be thrown on every check violation
// instead of the original error.
func (b *BoolRule) WithError(err error) *BoolRule {
	b.err = err
	return b
}

// Apply : Applies the rule on a given argument.
func (b *BoolRule) Apply(arg interface{}) error {
	if b.isWhitelisted(arg) {
		return nil
	}
	boolVal, err := toBool(arg, b.base)
	if err != nil {
		return orErr(b.err, errBool(b.base))
	}

	if err := b.performChecks(boolVal); err != nil {
		return orErr(b.err, err)
	}
	return nil
}

// BoolRule CONSTRUCTORS ############################################

// PureBool : Creates a BoolRule which expects the arg to be a bool.
func PureBool() *BoolRule {
	return &BoolRule{base: boolType}
}

// BoolRule PRIVATE METHODS #########################################

func (b *BoolRule) isWhitelisted(value interface{}) bool {
	for _, white := range b.whites {
		if white == value {
			return true
		}
	}
	return false
}

func (b *BoolRule) performChecks(arg bool) error {
	for _, check := range b.checks {
		if check == nil {
			continue
		}
		if err := check(arg); err != nil {
			return err
		}
	}
	return nil
}

// BoolRule UTILITY PUBLIC METHODS  #################################
