package valkyrie

import (
	"errors"
)

var errShouldBeBool = errors.New("value should be a boolean")
var errShouldBeInt = errors.New("value should be an int")
var errShouldBeFloat = errors.New("value should be a float64")
var errShouldBeString = errors.New("value should be a string")
