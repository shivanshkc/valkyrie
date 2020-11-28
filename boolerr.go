package valkyrie

import "fmt"

var errShouldBeBool = func() error {
	return fmt.Errorf("value should be a bool")
}

var errShouldBeTrue = func() error {
	return fmt.Errorf("value should be true")
}

var errShouldBeFalse = func() error {
	return fmt.Errorf("value should be false")
}
