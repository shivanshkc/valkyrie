package valkyrie

func orErr(err1 error, err2 error) error {
	if err1 != nil {
		return err1
	}
	return err2
}
