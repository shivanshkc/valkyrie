package valkyrie

// MapCheck : Represents a function that performs a validation check on a map[string]interface{}.
type MapCheck func(map[string]interface{}) error

// MapRule : Rule interface implementation for a map[string]interface{}.
type MapRule struct {
	checks []MapCheck
	err    error
}

// MapRule PRIMARY PUBLIC METHODS ###################################

// AddCheck : Adds a custom check function to the rule.
func (m *MapRule) AddCheck(check MapCheck) *MapRule {
	m.checks = append(m.checks, check)
	return m
}

// WithError : Adds a custom error to the rule.
// This custom error (if not nil) will be thrown on every check violation
// instead of the original error.
func (m *MapRule) WithError(err error) *MapRule {
	m.err = err
	return m
}

// Apply : Applies the rule on a given argument.
func (m *MapRule) Apply(arg interface{}) error {
	mapVal, ok := arg.(map[string]interface{})
	if !ok {
		return orErr(m.err, errMap())
	}

	if err := m.performChecks(mapVal); err != nil {
		return orErr(m.err, err)
	}
	return nil
}

// MapRule CONSTRUCTORS #############################################

// PureMap : Creates an MapRule which expects the arg to be a map[string]interface{}.
func PureMap() *MapRule {
	return &MapRule{}
}

// MapRule PRIVATE METHODS ##########################################

func (m *MapRule) performChecks(arg map[string]interface{}) error {
	for _, check := range m.checks {
		if check == nil {
			continue
		}
		if err := check(arg); err != nil {
			return err
		}
	}
	return nil
}

// MapRule UTILITY PUBLIC METHODS  ##################################
