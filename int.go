package valkyrie

// IntCheck : Represents a function to check (or validate) an int.
type IntCheck func(arg int64) error

// IntRule : Contains all the checks for an int, and related methods.
type IntRule struct {
	intChecks []IntCheck

	floatRule  *FloatRule
	stringRule *StringRule

	intWhites    []int64
	floatWhites  []float64
	stringWhites []string
}
