package valkyrie

import "fmt"

var (
	errString = func(dt string) error { return fmt.Errorf("value should follow: type %s && convertible to string", dt) }
)
