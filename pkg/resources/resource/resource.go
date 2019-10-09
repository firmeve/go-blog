package resource

import (
	"github.com/blog/pkg/utils"
	"reflect"
	"strings"
)

var (
	resourcesFields map[reflect.Type]map[string]map[string]string
)

const (
	Only   = iota
	Except = iota
)

type Resource struct {
	source       interface{}
	fields       []string
	status       int
	onlyFields   resourceFields
	exceptFields resourceFields
}

func (r *Resource) Only(fields []string) {
	r.fields = fields
	r.status = Only
}

func (r *Resource) Except(fields []string) {
	r.fields = fields
	r.status = Except
}

func (r *Resource) Source() interface{} {
	return r.source
}

// @todo effectiveFields  Transform
// @todo 只写到这，需要在transform里面取出有有效交集或差集字段
// @todo 然后用call调用字段值或方法

// Get effective fields
func (r *Resource) effectiveFields() fields {
	if len(r.onlyFields) > 0 {
		return r.onlyFields
	}

	fields := utils.ReflectFieldsName(r.resource)
	if len(r.exceptFields) > 0 {
		return utils.SliceStringExcept(fields, t.exceptFields)
	}

	return fields
}

//Core conversion method
func (r *Resource) Transform(source interface{}) map[string]interface{} {

	for name, field := range r.ReflectRelationFields(source) {
		if utils.SliceStringIn(r.fields, name) && r.status == Except {
			continue
		}

		if utils.SliceStringIn(methods, method) {
			collection[fieldKey] = utils.ReflectCallMethod(t.source, method)[0].Interface()
			//collection[fieldKey] = utils.ReflectCallMethod(t, method)[0].Interface()
		} else {
			collection[fieldKey] = utils.ReflectFieldValue(t.resource, field).Interface()
		}
	}

	methods := utils.ReflectMethodsName(r.source)

	collection := make(map[string]interface{})
	for _, field := range r.effectiveFields() {
		method := utils.StringUcWords([]string{"Get", field, `Field`})
		fieldKey := utils.StringSnakeCase(field)
		if utils.SliceStringIn(methods, method) {
			collection[fieldKey] = utils.ReflectCallMethod(t.source, method)[0].Interface()
			//collection[fieldKey] = utils.ReflectCallMethod(t, method)[0].Interface()
		} else {
			collection[fieldKey] = utils.ReflectFieldValue(t.resource, field).Interface()
		}
	}

	return map[string]interface{}{"data": collection}
}

type testResource struct {
	ID uint `resource:"id,method"`
}

func (r *Resource) ReflectRelationFields(source interface{}) map[string]map[string]string {

	reflectType := reflect.TypeOf(source)
	if v, ok := resourcesFields[reflectType]; ok {
		return v
	}

	fields := make(map[string]map[string]string, 0)

	for name, tag := range utils.ReflectStructFieldsTag(source) {
		var alias, method string

		if tag.Get(`resource`) != `` {
			tagNames := strings.Split(tag.Get(`resource`), `,`)
			alias = tagNames[0]
			if len(tagNames) >= 2 {
				method = tagNames[1]
			}
		} else { //method
			alias = utils.StringSnakeCase(name)
		}

		if method == `` {
			method = utils.StringUcWords([]string{name, `Field`})
		}

		fields[name] = map[string]string{`alias`: alias, `method`: method,}
	}

	resourcesFields[reflectType] = fields

	return fields
}

//func (r *Resource) ReflectRelationFields(source interface{}) map[string]interface{} {
//
//	methods := utils.ReflectMethodsName(source)
//	fields := make(map[string]interface{})
//
//	for name, value := range utils.ReflectStructFields(source, true) {
//		tag := value[`tag`].(reflect.StructTag)
//		callValue := value[`value`].(reflect.Value)
//		var alias string
//		if tag.Get(`resource`) != `` {
//			tagNames := strings.Split(tag.Get(`resource`), `,`)
//			alias = tagNames[0]
//			if len(tagNames) == 2 && utils.SliceStringIn(methods, tagNames[1]) {
//				callValue = utils.ReflectCallMethod(source, tagNames[1])[0]
//			}
//		} else { //method
//			alias = utils.StringSnakeCase(name)
//			method := utils.StringUcWords([]string{name, `Field`})
//			if utils.SliceStringIn(methods, method) {
//				callValue = utils.ReflectCallMethod(source, method)[0]
//			}
//		}
//
//		fields[alias] = utils.ReflectValueInterface(callValue)
//	}
//
//	return fields
//}
