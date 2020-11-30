package valkyrie

import "strconv"

func orErr(err1 error, err2 error) error {
	if err1 != nil {
		return err1
	}
	return err2
}

func toString(arg interface{}, dataType string) (string, error) {
	switch dataType {
	case boolType:
		boolVal, ok := arg.(bool)
		if !ok {
			return "", errEmpty
		}
		return strconv.FormatBool(boolVal), nil
	case intType:
		intVal, ok := arg.(int64)
		if !ok {
			return "", errEmpty
		}
		return strconv.FormatInt(intVal, 10), nil
	case floatType:
		floatVal, ok := arg.(float64)
		if !ok {
			return "", errEmpty
		}
		return strconv.FormatFloat(floatVal, 'f', -1, 64), nil
	case stringType:
		str, ok := arg.(string)
		if !ok {
			return "", errEmpty
		}
		return str, nil
	default:
		return "", errEmpty
	}
}
