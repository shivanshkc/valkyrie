package valkyrie

import (
	"errors"
	"fmt"
)

var errShouldBeBool = errors.New("value should be a boolean")
var errShouldBeInt = errors.New("value should be an int")
var errShouldBeFloat = errors.New("value should be a float64")
var errShouldBeString = errors.New("value should be a string")

var errShouldBeSlice = errors.New("value should be a slice")
var errSliceTooFewItems = errors.New("slice contains too few arguments")

var errShouldBeMap = errors.New("value should be a map")

func errKeyMissing(keyName interface{}) error {
	return fmt.Errorf("%v key is missing", keyName)
}