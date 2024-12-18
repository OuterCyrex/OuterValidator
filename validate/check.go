package validate

import (
	"errors"
	"reflect"
)

// Check Provides a way for client to check if an object meets the constraints
// It will return a bool and an error, for client to check which constraint is not met
func (v *Validator) Check(Struct interface{}) (bool, error) {
	Fields, Values := reflect.TypeOf(Struct), reflect.ValueOf(Struct)
	for i := 0; i < Fields.NumField(); i++ {
		field := Fields.Field(i)
		if _, ok := v.cs[field.Name]; !ok {
			return false, errors.New("[Validator]: field " + field.Name + " is not Contained into Validator")
		}
		if v.cs[field.Name].required && !checkRequired(Values.FieldByName(field.Name)) {
			return false, errors.New("[Validator]: field " + field.Name + " is required")
		}
		if v.cs[field.Name].max.ok && !checkMax(v.cs[field.Name].max.num, Values.FieldByName(field.Name)) {
			return false, errors.New("[Validator]: the length of field " + field.Name + " is greater then max")
		}
		if v.cs[field.Name].min.ok && !checkMin(v.cs[field.Name].min.num, Values.FieldByName(field.Name)) {
			return false, errors.New("[Validator]: the length of field " + field.Name + " is smaller then min")
		}
	}
	return true, nil
}

func checkRequired(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return !rv.IsZero()
	case reflect.Uintptr, reflect.Map, reflect.Slice, reflect.Pointer, reflect.UnsafePointer:
		return !rv.IsNil()
	case reflect.String:
		return !(rv.Len() == 0)
	default:
		return false
	}
}

func checkMin(min int, rv reflect.Value) bool {
	if rv.Len() < min {
		return false
	}
	return true
}

func checkMax(max int, rv reflect.Value) bool {
	if rv.Len() > max {
		return false
	}
	return true
}
