package strtools

import (
	"reflect"
)

// IsString check if a value is a string
func IsString(v interface{}) bool {
	if reflect.TypeOf(v).String() == "string" {
		return true
	} else {
		return false
	}
}
