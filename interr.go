package valkyrie

import "fmt"

var errShouldBeInt = func() error {
	return fmt.Errorf("value should be an int")
}

var errIntMin = func(limit int) error {
	return fmt.Errorf("value should not be smaller than %d", limit)
}

var errIntMax = func(limit int) error {
	return fmt.Errorf("value should not be greater than %d", limit)
}
