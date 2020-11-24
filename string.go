package valkyrie

func IsString(arg interface{}) error {
	_, ok := arg.(string)
	if !ok {
		return errShouldBeString
	}
	return nil
}
