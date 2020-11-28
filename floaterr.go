package valkyrie

import "fmt"

var errShouldBeFloat = func() error {
	return fmt.Errorf("value should be a float")
}

var errFloatNonZero = func() error {
	return fmt.Errorf("value should be non-zero")
}

var errFloatMin = func(limit float64) error {
	return fmt.Errorf("value should not be smaller than %f", limit)
}

var errFloatMax = func(limit float64) error {
	return fmt.Errorf("value should not be greater than %f", limit)
}
