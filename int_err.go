package valkyrie

import "fmt"

var (
	errInt64 = func(dt string) error { return fmt.Errorf("value should follow: type %s && convertible to int64", dt) }
)
