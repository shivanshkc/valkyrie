package valkyrie

import (
	"reflect"
)

type MapKey struct {
	Name interface{}
	Required bool
	Rule Validatable
}

type Map struct {
	Keys []*MapKey
}

func (m *Map) Validate(arg interface{}, customErr error) error {
	if reflect.TypeOf(arg).Kind() != reflect.Map {
		return orErr(customErr, errShouldBeMap)
	}

	mapValue := reflect.ValueOf(arg)
	for _, key := range m.Keys {
		value := mapValue.MapIndex(reflect.ValueOf(key.Name))

		if value == (reflect.Value{}) {
			if key.Required {
				return orErr(customErr, errKeyMissing(key.Name))
			}
			continue
		}

		err := key.Rule.Validate(value, customErr)
		if err != nil {
			return err
		}
	}
	return nil
}