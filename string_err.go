package valkyrie

import "fmt"

var (
	errString        = func(dt string) error { return fmt.Errorf("value should follow: type %s && convertible to string", dt) }
	errStringLenGTE  = func(value int64) error { return fmt.Errorf("value should follow: type: string && length >= %d", value) }
	errStringLenLTE  = func(value int64) error { return fmt.Errorf("value should follow: type: string && length <= %d", value) }
	errStringLenGT   = func(value int64) error { return fmt.Errorf("value should follow: type: string && length > %d", value) }
	errStringLenLT   = func(value int64) error { return fmt.Errorf("value should follow: type: string && length < %d", value) }
	errStringPattern = func(value string) error { return fmt.Errorf("value should follow: type string && pattern: %s", value) }
)
