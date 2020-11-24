package valkyrie

func IsFloat(arg interface{}) error {
	_, ok := arg.(float64)
	if !ok {
		return errShouldBeFloat
	}
	return nil
}
