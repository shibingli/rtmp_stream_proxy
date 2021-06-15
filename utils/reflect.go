package utils

import (
	"fmt"
	"reflect"
)

//Contains 判断 source 是否包含在target中，target支持类型:Array/Slice/Map
func Contains(source interface{}, target interface{}) (bool, error) {
	t := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < t.Len(); i++ {
			if t.Index(i).Interface() == source {
				return true, nil
			}
		}
	case reflect.Map:
		if t.MapIndex(reflect.ValueOf(source)).IsValid() {
			return true, nil
		}
	}

	return false, fmt.Errorf("%s", "Invalid array")
}
