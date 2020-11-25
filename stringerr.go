package valkyrie

import (
	"fmt"
	"regexp"
)

var errShouldBeString = func() error {
	return fmt.Errorf("value should be string")
}

var errStringNonEmpty = func() error {
	return fmt.Errorf("string should not be empty")
}

var errStringMaxLength = func(length int) error {
	return fmt.Errorf("string length should not be greater than %d", length)
}

var errStringMinLength = func(length int) error {
	return fmt.Errorf("string length should not be less than %d", length)
}

var errStringPattern = func(pattern regexp.Regexp) error {
	return fmt.Errorf("string should satisfy regex: %s", pattern.String())
}
