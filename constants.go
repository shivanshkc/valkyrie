package valkyrie

import "errors"

const (
	boolType   string = "bool"
	intType    string = "int64"
	floatType  string = "float64"
	stringType string = "string"
)

var (
	errEmpty = errors.New("")
)
