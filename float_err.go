package valkyrie

import "fmt"

var (
	errFloat64  = func(dt string) error { return fmt.Errorf("value should follow: type %s && convertible to float64", dt) }
	errFloatGTE = func(value float64) error { return fmt.Errorf("value should follow: type float64 && >= %f", value) }
	errFloatLTE = func(value float64) error { return fmt.Errorf("value should follow: type float64 && <= %f", value) }
	errFloatGT  = func(value float64) error { return fmt.Errorf("value should follow: type float64 && > %f", value) }
	errFloatLT  = func(value float64) error { return fmt.Errorf("value should follow: type float64 && < %f", value) }
)
