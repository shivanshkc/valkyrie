package valkyrie

import "fmt"

var errShouldBeMap = func() error {
	return fmt.Errorf("value should be a map")
}

var errRequiredKeyMissing = func(keyName string) error {
	return fmt.Errorf("required key '%s' missing", keyName)
}
