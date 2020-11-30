package valkyrie

import "regexp"

// StringCheck : Represents a function that performs a validation check on a string.
type StringCheck func(arg string) error

// StringRule : Rule interface implementation for a string.
type StringRule struct {
	from   string
	whites []interface{}
	checks []StringCheck
	err    error
}

// StringRule PRIMARY PUBLIC METHODS ################################

// Allow : Whitelists the provided values for a rule.
// If the argument is one of the whitelisted values, no checks
// will be performed upon it.
func (s *StringRule) Allow(args ...interface{}) *StringRule {
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
	if s.isWhitelisted(arg) {
		return nil
	}
	str, err := toString(arg, s.from)
	if err != nil {
		return orErr(s.err, errString(s.from))
	}

	if err := s.performChecks(str); err != nil {
		return orErr(s.err, err)
	}
	return nil
}

// StringRule CONSTRUCTORS ##########################################

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

// StringRule PRIVATE METHODS #######################################

func (s *StringRule) isWhitelisted(value interface{}) bool {
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

// StringRule UTILITY PUBLIC METHODS  ###############################

// LenGTE : Adds a '>=' check on the string length.
func (s *StringRule) LenGTE(value int64) *StringRule {
	s.AddCheck(func(arg string) error {
		if len(arg) < int(value) {
			return errStringLenGTE(value)
		}
		return nil
	})
	return s
}

// LenLTE : Adds a '<=' check on the string length.
func (s *StringRule) LenLTE(value int64) *StringRule {
	s.AddCheck(func(arg string) error {
		if len(arg) > int(value) {
			return errStringLenLTE(value)
		}
		return nil
	})
	return s
}

// LenGT : Adds a '>' check on the string length.
func (s *StringRule) LenGT(value int64) *StringRule {
	s.AddCheck(func(arg string) error {
		if len(arg) <= int(value) {
			return errStringLenGT(value)
		}
		return nil
	})
	return s
}

// LenLT : Adds a '<' check on the string length.
func (s *StringRule) LenLT(value int64) *StringRule {
	s.AddCheck(func(arg string) error {
		if len(arg) >= int(value) {
			return errStringLenLT(value)
		}
		return nil
	})
	return s
}

// Pattern : Adds a regex check to the string.
func (s *StringRule) Pattern(reg *regexp.Regexp) *StringRule {
	s.AddCheck(func(arg string) error {
		matches := reg.MatchString(arg)
		if !matches {
			return errStringPattern(reg.String())
		}
		return nil
	})
	return s
}

// UUIDv4 : Adds a UUIDv4 check on the string.
func (s *StringRule) UUIDv4() *StringRule {
	s.AddCheck(func(arg string) error {
		matches := regexp.MustCompile(uuidRegex).MatchString(arg)
		if !matches {
			return errStringUUIDv4()
		}
		return nil
	})
	return s
}

// Blind : Invalidates everything except the whitelisted (allowed) values.
func (s *StringRule) Blind() *StringRule {
	s.AddCheck(func(arg string) error {
		return errBlind
	})
	return s
}
