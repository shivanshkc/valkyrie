package valkyrie

// FloatCheck : Represents a function to check (or validate) a float.
type FloatCheck func(arg float64) error

// FloatRule : Contains all the checks for a float, and related methods.
type FloatRule struct {
	floatChecks []FloatCheck

	intRule    *IntRule
	stringRule *StringRule

	intWhites    []int64
	floatWhites  []float64
	stringWhites []string
}
