package utils

import (
	"reflect"
)

// Reflect get methods name
// But only support public method
func ReflectMethodsName(object interface{}) []string {
	nums := ReflectTypeIndirect(reflect.TypeOf(object)).NumMethod()
	methods := make([]string, 0)
	for i := 0; i < nums; i++ {
		methods = append(methods, reflect.TypeOf(object).Method(i).Name)
	}

	return methods
}

func ReflectMethods(object interface{}) map[string]reflect.Method {
	nums := ReflectTypeIndirect(reflect.TypeOf(object)).NumMethod()
	methods := make(map[string]reflect.Method, 0)
	for i := 0; i < nums; i++ {
		method := reflect.TypeOf(object).Method(i)
		methods[method.Name] = method
	}

	return methods
}

func ReflectFieldValue(object interface{}, name string) reflect.Value {
	return reflect.Indirect(reflect.ValueOf(object)).FieldByName(name)
}

func ReflectFieldValueInterface(object interface{}, name string) interface{} {
	return ReflectValueInterface(reflect.Indirect(reflect.ValueOf(object)).FieldByName(name))
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
	reflectType := ReflectTypeIndirect(reflect.TypeOf(object))
	nums := reflectType.NumField()
	fields := make([]string, 0)
	for i := 0; i < nums; i++ {
		fields = append(fields, reflectType.Field(i).Name)
	}

	return fields
}

func ReflectStructFields(object interface{}, onlyPublic bool) map[string]map[string]interface{} {
	reflectType := ReflectTypeIndirect(reflect.TypeOf(object))
	reflectValue := reflect.Indirect(reflect.ValueOf(object))
	nums := reflectType.NumField()
	fields := make(map[string]map[string]interface{}, 0)

	for i := 0; i < nums; i++ {
		typeField := reflectType.Field(i)
		valueField := reflectValue.Field(i)
		fieldName := typeField.Name
		fieldTag := typeField.Tag

		if onlyPublic && reflectValue.Field(i).CanSet() {
			fields[fieldName] = map[string]interface{}{`tag`:fieldTag,`value`:valueField}
		} else if !onlyPublic {
			fields[fieldName] = map[string]interface{}{`tag`:fieldTag,`value`:valueField}
		}
	}

	return fields
}

func ReflectStructFieldsTag(object interface{}) map[string]reflect.StructTag {
	reflectType := ReflectTypeIndirect(reflect.TypeOf(object))
	nums := reflectType.NumField()
	fields := make(map[string]reflect.StructTag, 0)

	for i := 0; i < nums; i++ {
		typeField := reflectType.Field(i)
		fields[typeField.Name] = typeField.Tag
	}

	return fields
}

//func ReflectPublicFieldsValue(object interface{}) map[string]reflect.Value {
//	reflectValue := reflect.ValueOf(object)
//	reflectType := reflect.TypeOf(object)
//
//	nums := reflectValue.NumField()
//	fields := make(map[string]reflect.Value, 0)
//	for i := 0; i < nums; i++ {
//		if reflectValue.Field(i).CanSet() {
//			fields[reflectType.Field(i).Name] = reflectValue.Field(i)
//		}
//	}
//
//	return fields
//}

//func ReflectStructFields(object interface{}, onlyPublic bool) map[reflect.StructField]reflect.Value {
//	reflectValue := reflect.Indirect(reflect.ValueOf(object))
//	reflectType := ReflectTypeIndirect(reflect.TypeOf(object))
//
//	nums := reflectValue.NumField()
//	fields := make(map[reflect.StructField]reflect.Value, 0)
//	for i := 0; i < nums; i++ {
//		valueField := reflect.Indirect(reflectValue.Field(i))
//		typeField := reflectType.Field(i)
//		if valueField.CanSet() && onlyPublic {
//			fields[typeField] = valueField
//			continue
//		}
//
//		fields[typeField] = valueField
//	}
//
//	return fields
//}

func ReflectTypeIndirect(reflectType reflect.Type) reflect.Type {
	if SliceUintIn(
		[]uint{uint(reflect.Array), uint(reflect.Ptr), uint(reflect.Chan), uint(reflect.Map), uint(reflect.Slice)},
		uint(reflectType.Kind()),
	) {
		reflectType = reflectType.Elem()
	}

	return reflectType
}

func ReflectIsTypeKind(object interface{}, kind reflect.Kind) bool {
	return reflect.TypeOf(object).Kind() == kind
}

func ReflectIsValueKind(object interface{}, kind reflect.Kind) bool {
	return reflect.ValueOf(object).Kind() == kind
}

func ReflectValueInterface(value reflect.Value) interface{} {
	if value.CanAddr() {
		return value.Addr().Interface()
	}

	return value.Interface()
}
