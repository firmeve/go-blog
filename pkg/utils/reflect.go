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
