package datastruct

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type AnyType any

// ToRightValue return code-generated right value of the input (right side of the equals sign).
// For example, if the input is :
// 1. int = 17 -> 17
// 2. string = "Hello" -> "Hello"
// since in code, you can do:
// variable := <the right value>

func ToRightValue(anyValue AnyType) string {
	baseType := reflect.TypeOf(anyValue).Kind()
	value := reflect.ValueOf(anyValue)
	switch baseType {
	case reflect.String:
		realValue := value.String()
		return fmt.Sprintf("%q", realValue) // Quote it
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		realValue := value.Int()
		return strconv.FormatInt(realValue, 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		realValue := value.Uint()
		return strconv.FormatUint(realValue, 10)
	case reflect.Float64:
		realValue := value.Float()
		return strconv.FormatFloat(realValue, 'f', -1, 64)
	case reflect.Float32:
		realValue := value.Float()
		return strconv.FormatFloat(realValue, 'f', -1, 32)
	case reflect.Bool:
		realValue := value.Bool()
		return strconv.FormatBool(realValue)
	case reflect.Complex64:
		realValue := value.Complex()
		return strconv.FormatComplex(realValue, 'f', -1, 64)
	case reflect.Complex128:
		realValue := value.Complex()
		return strconv.FormatComplex(realValue, 'f', -1, 128)

		// From here -> recursive values.
	case reflect.Map:
		resultString := make([]string, 0)
		resultString = append(resultString, fmt.Sprintf("map[%v]%v{", value.Type().Key(), value.Type().Elem()))

		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key() // Key is not recursive
			v := iter.Value()

			valueRightKey := ToRightValue(k.Interface())
			valueRightVal := ToRightValue(v.Interface())

			resultString = append(resultString, fmt.Sprintf("%v: %v,", valueRightKey, valueRightVal))
		}
		resultString = append(resultString, "}") // Close
		return strings.Join(resultString, "\n")

	case reflect.Slice, reflect.Array:
		resultString := make([]string, 0)
		resultString = append(resultString, fmt.Sprintf("%v{", value.Type()))

		l := value.Len()
		for i := 0; i < l; i++ {
			v := value.Index(i)
			valueRightVal := ToRightValue(v.Interface())
			resultString = append(resultString, fmt.Sprintf("%v,", valueRightVal))
		}
		resultString = append(resultString, "}") // Close
		return strings.Join(resultString, "\n")

	case reflect.Struct:
		resultString := make([]string, 0)
		resultString = append(resultString, fmt.Sprintf("%v{", value.Type()))

		l := value.NumField()
		to := reflect.TypeOf(value.Interface())
		valueOf := reflect.ValueOf(value.Interface())
		for i := 0; i < l; i++ {
			valueRightKey := to.Field(i).Name
			firstLetter := valueRightKey[0:1]
			if strings.ToUpper(firstLetter) != firstLetter { // Validate that first letter is upper case (Public)
				continue // Neglect privates
			}
			valueRightVal := ToRightValue(valueOf.Field(i).Interface())
			resultString = append(resultString, fmt.Sprintf("%v: %v,", valueRightKey, valueRightVal))
		}
		resultString = append(resultString, "}") // Close
		return strings.Join(resultString, "\n")

	default:
		return fmt.Sprintf("Unable to parse. Type: %v, name: %v", baseType, reflect.TypeOf(value).Name()) // Quote it

	}

}
