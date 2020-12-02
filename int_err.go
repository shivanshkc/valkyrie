package valkyrie

import "fmt"

var (
	errInt64     = func(dt string) error { return fmt.Errorf("value should follow: type %s && convertible to int64", dt) }
	errIntGTE    = func(value int64) error { return fmt.Errorf("value should follow: type int64 && >= %d", value) }
	errIntLTE    = func(value int64) error { return fmt.Errorf("value should follow: type int64 && <= %d", value) }
	errIntGT     = func(value int64) error { return fmt.Errorf("value should follow: type int64 && > %d", value) }
	errIntLT     = func(value int64) error { return fmt.Errorf("value should follow: type int64 && < %d", value) }
	errIntExcept = func(value int64) error { return fmt.Errorf("value should follow: type int64 && != %d", value) }
)
