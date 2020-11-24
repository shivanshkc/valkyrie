package valkyrie

func IsBool(arg interface{}) error {
	_, ok := arg.(bool)
	if !ok {
		return errShouldBeBool
	}
	return nil
}
