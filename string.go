package valkyrie

import "regexp"

type StringCheck func(string) error

type StringRule struct {
	checks    []StringCheck
	customErr error
}

func String(customErr error) *StringRule {
	return &StringRule{
		checks:    []StringCheck{},
		customErr: customErr,
	}
}

func (sr *StringRule) NonEmpty() *StringRule {
	sr.checks = append(sr.checks, func(s string) error {
		if len(s) == 0 {
			return errStringNonEmpty()
		}
		return nil
	})
	return sr
}

func (sr *StringRule) MaxLength(length int) *StringRule {
	sr.checks = append(sr.checks, func(s string) error {
		if len(s) > length {
			return errStringMaxLength(length)
		}
		return nil
	})
	return sr
}

func (sr *StringRule) MinLength(length int) *StringRule {
	sr.checks = append(sr.checks, func(s string) error {
		if len(s) < length {
			return errStringMinLength(length)
		}
		return nil
	})
	return sr
}

func (sr *StringRule) Regex(pattern regexp.Regexp) *StringRule {
	sr.checks = append(sr.checks, func(s string) error {
		matches := pattern.MatchString(s)
		if !matches {
			return errStringPattern(pattern)
		}
		return nil
	})
	return sr
}

func (sr *StringRule) Custom(check StringCheck) *StringRule {
	sr.checks = append(sr.checks, check)
	return sr
}

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
