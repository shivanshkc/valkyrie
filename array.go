package valkyrie

import "reflect"

type Slice struct {
	Rules []Validatable
}

func (s *Slice) Validate(arg interface{}, customErr error) error {
	if reflect.TypeOf(arg).Kind() != reflect.Slice {
		return orErr(customErr, errShouldBeSlice)
	}

	slice := reflect.ValueOf(arg)
	if slice.Len() < len(s.Rules) {
		return orErr(customErr, errSliceTooFewItems)
	}

	for ind := 0; ind < slice.Len(); ind++ {
		validator := s.Rules[ind]
		err := validator.Validate(slice.Index(ind), customErr)
		if err != nil {
			return err
		}
	}
	return nil
}
