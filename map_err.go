package valkyrie

import "fmt"

var (
	errMap           = func() error { return fmt.Errorf("value should follow: type map[string]interface{}") }
	errMapKeyMissing = func(name string) error { return fmt.Errorf("required key '%s' is missing", name) }
)