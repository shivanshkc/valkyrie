package valkyrie

import "regexp"

// StringCheck : Represents a function that performs a validation check on a string.
type StringCheck func(string) error

// StringRule : A Collection of StringChecks, hence acts as a customizable Rule for a string.
type StringRule struct {
	checks    []StringCheck
	customErr error
}

// String : An intuitive function to instantiate a StringRule.
func String(customErr error) *StringRule {
	return &StringRule{
		checks:    []StringCheck{},
		customErr: customErr,
	}
}

// NonEmpty : Disallows strings of length 0.
func (sr *StringRule) NonEmpty() *StringRule {
	sr.checks = append(sr.checks, func(s string) error {
		if len(s) == 0 {
			return errStringNonEmpty()
		}
		return nil
	})
	return sr
}

// MaxLength : Disallows strings with length greater than the provided length.
func (sr *StringRule) MaxLength(length int) *StringRule {
	sr.checks = append(sr.checks, func(s string) error {
		if len(s) > length {
			return errStringMaxLength(length)
		}
		return nil
	})
	return sr
}

// MinLength : Disallows strings with length smaller than the provided length.
func (sr *StringRule) MinLength(length int) *StringRule {
	sr.checks = append(sr.checks, func(s string) error {
		if len(s) < length {
			return errStringMinLength(length)
		}
		return nil
	})
	return sr
}

// UUIDv4 : Disallows strings that are not UUIDv4.
func (sr *StringRule) UUIDv4() *StringRule {
	sr.checks = append(sr.checks, func(s string) error {
		if !regexp.MustCompile(uuidRegex).MatchString(s) {
			return errStringShouldBeUUIDv4()
		}
		return nil
	})
	return sr
}

// Regex : Checks the given string against the provided regex.
func (sr *StringRule) Regex(pattern *regexp.Regexp) *StringRule {
	sr.checks = append(sr.checks, func(s string) error {
		matches := pattern.MatchString(s)
		if !matches {
			return errStringPattern(pattern)
		}
		return nil
	})
	return sr
}

// Custom : Allows to add a custom StringCheck to the StringRule.
func (sr *StringRule) Custom(check StringCheck) *StringRule {
	sr.checks = append(sr.checks, check)
	return sr
}

// Apply : Applies all the checks in the StringRule on the provided args.
func (sr *StringRule) Apply(arg interface{}) error {
	str, ok := arg.(string)
	if !ok {
		return orErr(sr.customErr, errShouldBeString())
	}

	for _, check := range sr.checks {
		if check == nil {
			continue
		}
		err := check(str)
		if err != nil {
			return orErr(sr.customErr, err)
		}
	}
	return nil
}
