package valkyrie

import (
	"errors"
	"fmt"
)

var (
	errEmpty = errors.New("")
	errBlind = errors.New("blind validation")

	errBool = func(t string) error { return fmt.Errorf("value should follow: type %s && convertible to bool", t) }

	errInt64     = func(t string) error { return fmt.Errorf("value should follow: type %s && convertible to int64", t) }
	errIntGTE    = func(value int64) error { return fmt.Errorf("value should follow: type int64 && >= %d", value) }
	errIntLTE    = func(value int64) error { return fmt.Errorf("value should follow: type int64 && <= %d", value) }
	errIntGT     = func(value int64) error { return fmt.Errorf("value should follow: type int64 && > %d", value) }
	errIntLT     = func(value int64) error { return fmt.Errorf("value should follow: type int64 && < %d", value) }
	errIntExcept = func(value int64) error { return fmt.Errorf("value should follow: type int64 && != %d", value) }

	errFloat64     = func(t string) error { return fmt.Errorf("value should follow: type %s && convertible to float64", t) }
	errFloatGTE    = func(value float64) error { return fmt.Errorf("value should follow: type float64 && >= %f", value) }
	errFloatLTE    = func(value float64) error { return fmt.Errorf("value should follow: type float64 && <= %f", value) }
	errFloatGT     = func(value float64) error { return fmt.Errorf("value should follow: type float64 && > %f", value) }
	errFloatLT     = func(value float64) error { return fmt.Errorf("value should follow: type float64 && < %f", value) }
	errFloatExcept = func(value float64) error { return fmt.Errorf("value should follow: type float64 && != %f", value) }

	errString        = func(t string) error { return fmt.Errorf("value should follow: type %s && convertible to string", t) }
	errStringLenGTE  = func(value int64) error { return fmt.Errorf("value should follow: type: string && length >= %d", value) }
	errStringLenLTE  = func(value int64) error { return fmt.Errorf("value should follow: type: string && length <= %d", value) }
	errStringLenGT   = func(value int64) error { return fmt.Errorf("value should follow: type: string && length > %d", value) }
	errStringLenLT   = func(value int64) error { return fmt.Errorf("value should follow: type: string && length < %d", value) }
	errStringPattern = func(value string) error { return fmt.Errorf("value should follow: type string && pattern: %s", value) }
	errStringUUIDv4  = func() error { return fmt.Errorf("value should follow: type string && valid UUIDv4") }
	errStringExcept  = func(value string) error { return fmt.Errorf("value should follow: type string && != %s", value) }

	errMap           = func() error { return fmt.Errorf("value should follow: type map[string]interface{}") }
	errMapKeyMissing = func(name string) error { return fmt.Errorf("required key '%s' is missing", name) }
)
