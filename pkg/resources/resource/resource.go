package resource

import (
	"github.com/blog/pkg/utils"
	"reflect"
	"regexp"
	"strings"
)

type MapCache map[string]map[string]string

var (
	resourcesFields  map[reflect.Type]MapCache
	resourcesMethods map[reflect.Type]MapCache
)

type Resource struct {
	source interface{}
	fields []string
}

func (r *Resource) Fields(fields []string) {
	r.fields = fields
}

func (r *Resource) Source() interface{} {
	return r.source
}

//Core conversion method
func (r *Resource) Transform(source interface{}) map[string]interface{} {
	fields := r.ReflectRelationFields(source)
	methods := r.ReflectRelationMethods(source)
	collection := make(map[string]interface{}, 0)
	for _, field := range r.fields {
		// method 优先
		if v, ok := methods[field]; ok {
			collection[v[`alias`]] = utils.ReflectValueInterface(utils.ReflectCallMethod(source, v[`method`])[0])
		}

		if v, ok := fields[field]; ok {
			if v[`method`] != `` {
				collection[v[`alias`]] = utils.ReflectValueInterface(utils.ReflectCallMethod(source, v[`method`])[0])
			} else {
				collection[v[`alias`]] = utils.ReflectFieldValueInterface(source, field)
			}
		}
	}

	return collection
}

func (r *Resource) ReflectRelationMethods(source interface{}) MapCache {
	reflectType := reflect.TypeOf(source)
	if v, ok := resourcesMethods[reflectType]; ok {
		return v
	}

	methods := make(MapCache, 0)

	for name := range utils.ReflectMethods(source) {
		if regexp.MustCompile("^(.+)Field$").MatchString(name) {
			exceptFieldMethodName := name[0 : len(name)-5]
			methods[exceptFieldMethodName] = map[string]string{`alias`: utils.StringSnakeCase(exceptFieldMethodName), `method`: name}
		}

	}

	resourcesMethods[reflectType] = methods

	return methods
}

func (r *Resource) ReflectRelationFields(source interface{}) MapCache {

	reflectType := reflect.TypeOf(source)
	if v, ok := resourcesFields[reflectType]; ok {
		return v
	}

	fields := make(MapCache, 0)

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

		//if method == `` {
		//	method = utils.StringUcWords([]string{name, `Field`})
		//}

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
