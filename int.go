package valkyrie

func IsInt(arg interface{}) error {
	_, ok := arg.(int)
	if !ok {
		return errShouldBeInt
	}
	return nil
}
