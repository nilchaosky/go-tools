package tools

import "reflect"

func IsStruct(value interface{}) bool {
	v := reflect.ValueOf(value)
	t := v.Type()

	// 判断是否是结构体，排除指针的情况也考虑进去
	if t.Kind() == reflect.Struct {
		return true
	}

	// 如果是指针，判断它指向的是不是结构体
	if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
		return true
	}

	return false
}
