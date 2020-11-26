package valkyrie

// MapCheck : Represents a function that performs a validation check on a map.
type MapCheck func(map[string]interface{}) error

// MapRule : A Collection of MapChecks, hence acts as a customizable Rule for a map.
type MapRule struct {
	checks    []MapCheck
	customErr error
}

// Map : An intuitive function to instantiate a MapRule.
func Map(customErr error) *MapRule {
	return &MapRule{
		checks:    []MapCheck{},
		customErr: customErr,
	}
}

// Key : Adds a MapCheck to the MapRule for the specified key.
func (mr *MapRule) Key(name string, required bool, rule Rule) *MapRule {
	check := func(m map[string]interface{}) error {
		value, exists := m[name]
		if !exists && required {
			return errRequiredKeyMissing(name)
		}
		if !exists {
			return nil
		}

		if err := rule.Apply(value); err != nil {
			return err
		}
		return nil
	}

	mr.checks = append(mr.checks, check)
	return mr
}

// Apply : Applies all the checks in the MapRule on the provided args.
func (mr *MapRule) Apply(arg interface{}) error {
	mapValue, ok := arg.(map[string]interface{})
	if !ok {
		return orErr(mr.customErr, errShouldBeMap())
	}

	for _, check := range mr.checks {
		if check == nil {
			continue
		}
		if err := check(mapValue); err != nil {
			return orErr(mr.customErr, err)
		}
	}
	return nil
}
