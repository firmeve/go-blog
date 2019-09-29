package utils

import "reflect"

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
