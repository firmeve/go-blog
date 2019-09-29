package utils

import (
	"reflect"
)

// Reflect get methods name
// But only support public method
func ReflectMethodsName(object interface{}) []string {
	nums := reflect.TypeOf(object).NumMethod()
	methods := make([]string, 0)
	for i := 0; i < nums; i++ {
		methods = append(methods, reflect.TypeOf(object).Method(i).Name)
	}

	return methods
}

func ReflectFieldValue(object interface{}, name string) reflect.Value {
	return reflect.Indirect(reflect.ValueOf(object)).FieldByName(name)
}

func ReflectCallMethod(object interface{}, name string, params ...interface{}) []reflect.Value {
	newParams := make([]reflect.Value, 0)
	if len(params) > 0 {
		for _, param := range params {
			newParams = append(newParams, reflect.ValueOf(param))
		}
	}

	return reflect.ValueOf(object).MethodByName(name).Call(newParams)
}

func ReflectFieldsName(object interface{}) []string {
	reflectType := reflect.TypeOf(object)
	kind := reflectType.Kind()
	if  SliceIntIn([]int64{int64(reflect.Array),int64(reflect.Ptr),int64(reflect.Chan),int64(reflect.Map),int64(reflect.Slice)},int64(kind)){
		reflectType = reflectType.Elem()
	}
	nums := reflectType.NumField()
	fields := make([]string, 0)
	for i := 0; i < nums; i++ {
		fields = append(fields, reflectType.Field(i).Name)
	}

	return fields
}
