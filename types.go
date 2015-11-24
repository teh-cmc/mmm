package mmm

import (
	"fmt"
	"reflect"
)

// -----------------------------------------------------------------------------

// Type represents the underlying type of an interface.
type Type int

const (
	// TypeInvalid is an illegal type.
	TypeInvalid Type = iota
	// TypeNumeric is any of bool/int/uint/float/complex and their variants.
	TypeNumeric Type = iota
	// TypeArray is an array of any underlying type.
	TypeArray Type = iota
	// TypeStruct is any struct.
	TypeStruct Type = iota
	// TypeUnsafePointer is any pointer from the unsafe package.
	TypeUnsafePointer Type = iota
)

// TypeOf returns the underlying type of an interface.
func TypeOf(v interface{}) (Type, error) {
	t := reflect.TypeOf(v)
	if t == nil {
		return TypeInvalid, Error(fmt.Sprintf("unsuppported type: %#v", v))
	}
	k := t.Kind()

	switch k {
	case reflect.Bool:
		return TypeNumeric, nil
	case reflect.Int:
		return TypeNumeric, nil
	case reflect.Int8:
		return TypeNumeric, nil
	case reflect.Int16:
		return TypeNumeric, nil
	case reflect.Int32:
		return TypeNumeric, nil
	case reflect.Int64:
		return TypeNumeric, nil
	case reflect.Uint:
		return TypeNumeric, nil
	case reflect.Uint8:
		return TypeNumeric, nil
	case reflect.Uint16:
		return TypeNumeric, nil
	case reflect.Uint32:
		return TypeNumeric, nil
	case reflect.Uint64:
		return TypeNumeric, nil
	case reflect.Uintptr:
		return TypeNumeric, nil
	case reflect.Float32:
		return TypeNumeric, nil
	case reflect.Float64:
		return TypeNumeric, nil
	case reflect.Complex64:
		return TypeNumeric, nil
	case reflect.Complex128:
		return TypeNumeric, nil
	case reflect.Array:
		return TypeArray, nil
	case reflect.Struct:
		return TypeStruct, nil
	case reflect.UnsafePointer:
		return TypeUnsafePointer, nil
	default:
		return TypeInvalid, Error(fmt.Sprintf("unsuppported type: %#v", k.String()))
	}
}

// TypeCheck recursively checks the underlying types of `v` and returns an error
// if one or more of those types are illegal.
func TypeCheck(v interface{}) error {
	t, err := TypeOf(v)
	if err != nil {
		return nil
	}

	rv := reflect.ValueOf(v)

	if t == TypeArray && rv.Len() > 0 {
		return TypeCheck(rv.Index(0).Interface())
	}

	if t == TypeStruct {
		fields := rv.NumField()
		for i := 0; i < fields; i++ {
			if err := TypeCheck(rv.Field(i).Interface()); err != nil {
				return err
			}
		}
	}

	return nil
}
