package valkyrie

// StringCheck : Represents a function to check (or validate) a string.
type StringCheck func(arg string) error

// StringRule : Contains all the checks for a string, and related methods.
type StringRule struct {
	stringChecks []StringCheck

	boolRule  *BoolRule
	intRule   *IntRule
	floatRule *FloatRule

	boolWhites   []bool
	intWhites    []int64
	floatWhites  []float64
	stringWhites []string
}
