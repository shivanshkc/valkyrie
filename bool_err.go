package valkyrie

import "fmt"

var (
	errBool = func(dt string) error { return fmt.Errorf("value should follow: type %s && convertible to bool", dt) }
)
