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

var errStringShouldBeUUIDv4 = func() error {
	return fmt.Errorf("string should be a valid UUIDv4 string")
}

var errStringPattern = func(pattern *regexp.Regexp) error {
	return fmt.Errorf("string should satisfy regex: %s", pattern.String())
}

var errStringBool = func() error {
	return fmt.Errorf("string should be parse-able to bool")
}

var errStringInt = func() error {
	return fmt.Errorf("string should be parse-able to int")
}

var errStringFloat = func() error {
	return fmt.Errorf("string should be parse-able to float")
}
