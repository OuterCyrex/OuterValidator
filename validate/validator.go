package validate

import (
	"fmt"
	"reflect"
)

// Constraints defines the types of Constraints for a number of fields
// Each Type of Constraint has its types that can be adapted.
type Constraints struct {

	// t stands for the type of the field that owns this constraint
	t reflect.Type

	// Required means that the Value of this Field should not be None
	required bool

	// min and max means that if the Type of this Field is String,
	// the length of the Value should between the Min and Max.
	min struct {
		ok  bool
		num int
	}

	max struct {
		ok  bool
		num int
	}
}

// Validator contains two hashmaps, one ties the field names to their Constraints
// as the other one ties names to their types.
type Validator struct {
	cs map[string]*Constraints
}

// NewValidator returns an initial Validator Object
func NewValidator(Struct interface{}) Validator {
	Fields := reflect.TypeOf(Struct)
	if Fields.Kind() != reflect.Struct {
		panic("Generate New Validator Failed: Unsupported Data Types")
	}
	v := Validator{
		make(map[string]*Constraints, Fields.NumField()),
	}
	for i := 0; i < Fields.NumField(); i++ {
		if Fields.Field(i).Type.Kind() == reflect.Struct {
			panic("struct field is not allowed in validator")
		}
		v.cs[Fields.Field(i).Name] = &Constraints{
			t: Fields.Field(i).Type,
		}
	}
	return v
}

// Set returns the Constraints struct for client to Add
// Constraints
func (v *Validator) Set(field string) *Constraints {
	if _, ok := v.cs[field]; ok {
		return v.cs[field]
	}
	panic(fmt.Sprintf("No such Field: %s", field))
	return nil
}

// Required constrains that the field must not be empty
func (c *Constraints) Required() *Constraints {
	c.required = true
	return c
}

// Min constrains that the length of string must be greater than min
func (c *Constraints) Min(min int) *Constraints {
	if c.t.Kind() != reflect.String {
		panic("Type must be string when Using Min or Max")
		return c
	}
	if min < 0 {
		panic("Min must be Positive")
		return c
	}
	if c.max.ok && c.max.num < min {
		panic(fmt.Sprintf("Min must be Smaller than %d", c.max.num))
		return c
	}
	c.min = struct {
		ok  bool
		num int
	}{ok: true, num: min}
	return c
}

// Max constrains that the length of string must be smaller than max
func (c *Constraints) Max(max int) *Constraints {
	if c.t.Kind() != reflect.String {
		panic("Type must be string when Using Min or Max")
		return c
	}
	if max < 0 {
		panic("Max must be Positive")
		return c
	}
	if c.min.ok && c.min.num > max {
		panic(fmt.Sprintf("Max must be Greater than %d", c.min.num))
		return c
	}
	c.max = struct {
		ok  bool
		num int
	}{ok: true, num: max}
	return c
}
