package valkyrie

// StringCheck : Represents a function that performs a validation check on a string.
type StringCheck func(arg string) error

// StringRule : Rule interface implementation for a string.
type StringRule struct {
	from   string
	whites []string
	checks []StringCheck
	err    error
}

// Basic Rule Modifiers #############################################

// Allow : Whitelists the provided strings for a rule.
// If the argument is one of the whitelisted values, no checks
// will be performed upon it.
func (s *StringRule) Allow(args ...string) *StringRule {
	s.whites = append(s.whites, args...)
	return s
}

// AddCheck : Adds a custom check function to the rule.
func (s *StringRule) AddCheck(check StringCheck) *StringRule {
	s.checks = append(s.checks, check)
	return s
}

// WithError : Adds a custom error to the rule.
// This custom error (if not nil) will be thrown on every check violation
// instead of the original error.
func (s *StringRule) WithError(err error) *StringRule {
	s.err = err
	return s
}

// Apply : Applies the rule on a given argument.
func (s *StringRule) Apply(arg interface{}) error {
	str, err := toString(arg, s.from)
	if err != nil {
		return orErr(s.err, errString(s.from))
	}
	if s.isWhitelisted(str) {
		return nil
	}

	if err := s.performChecks(str); err != nil {
		return orErr(s.err, err)
	}
	return nil
}

// Constructors for StringRule ######################################

// BoolString : Creates a StringRule which expects the arg to be bool.
// which will be validated after conversion to string.
// Example: true -> "true"
func BoolString() *StringRule {
	return &StringRule{from: boolType}
}

// IntString : Creates a StringRule which expects the arg to be an int64.
// which will be validated after conversion to string.
// Example: 23 -> "23"
func IntString() *StringRule {
	return &StringRule{from: intType}
}

// FloatString : Creates a StringRule which expects the arg to be a float64.
// which will be validated after conversion to string.
// Example: 2.34 -> "2.34"
func FloatString() *StringRule {
	return &StringRule{from: floatType}
}

// PureString : Creates a StringRule which expects the arg to be a string.
func PureString() *StringRule {
	return &StringRule{from: stringType}
}

// Private method of StringRule #####################################

func (s *StringRule) isWhitelisted(value string) bool {
	for _, white := range s.whites {
		if white == value {
			return true
		}
	}
	return false
}

func (s *StringRule) performChecks(arg string) error {
	for _, check := range s.checks {
		if check == nil {
			continue
		}
		if err := check(arg); err != nil {
			return err
		}
	}
	return nil
}

// Utility checks for StringRule ####################################
