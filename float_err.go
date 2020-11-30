package valkyrie

import "fmt"

var (
	errFloat64 = func(dt string) error { return fmt.Errorf("value should follow: type %s && convertible to float64", dt) }
)
