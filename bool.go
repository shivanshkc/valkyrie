package valkyrie

// BoolCheck : Represents a function to check (or validate) a bool.
type BoolCheck func(arg bool) error

// BoolRule : Contains all the checks for a bool, and related methods.
type BoolRule struct {
	boolChecks []BoolCheck

	boolWhites []bool
}
