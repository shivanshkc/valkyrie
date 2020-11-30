package valkyrie

import "fmt"

var (
	errMap = func() error { return fmt.Errorf("value should follow: type map[string]interface{}") }
)
