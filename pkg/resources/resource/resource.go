package resource

import (
	"github.com/blog/pkg/resources"
	"github.com/blog/pkg/utils"
	"reflect"
	"regexp"
	"strings"
)

type MapCache map[string]map[string]string

type TransformerFunc func(source interface{}) map[string]interface{}

var (
	resourcesFields  = make(map[reflect.Type]MapCache, 0)
	resourcesMethods = make(map[reflect.Type]MapCache, 0)
)

type Resolver interface {
	Resolve() map[string]interface{}
}

type Resource struct {
	source      interface{}
	fields      []string
	transformer interface{}
}

func (r *Resource) SetTransformer(transformer interface{}) *Resource {
	r.transformer = transformer
	return r
}

func (r *Resource) SetSource(source interface{}) *Resource {
	r.transformer = source
	return r
}

func (r *Resource) SetFields(fields ...string) *Resource {
	r.fields = fields
	return r
}

func (r *Resource) Source() interface{} {
	return r.source
}

//Core conversion method
func (r *Resource) Transform(source interface{}) map[string]interface{} {

	if fn, ok := r.transformer.(TransformerFunc); ok {
		return fn(source)
	} else if transformer,ok := r.transformer.(resources.Transformer); ok {
		fields := r.ReflectRelationFields(source)
		methods := r.ReflectRelationMethods(transformer)
		collection := make(map[string]interface{}, 0)
		for _, field := range transformer.Field() {
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

		//transformer.Transform()
		return collection
	}
	//if utils.ReflectIsTypeKind(r.transformer, reflect.Func) {
	//
	//} else {
	//
	//}

	//fields := r.ReflectRelationFields(source)
	//methods := r.ReflectRelationMethods(r.transformer)
	//collection := make(map[string]interface{}, 0)
	//for _, field := range r.fields {
	//	// method 优先
	//	if v, ok := methods[field]; ok {
	//		collection[v[`alias`]] = utils.ReflectValueInterface(utils.ReflectCallMethod(source, v[`method`])[0])
	//	}
	//
	//	if v, ok := fields[field]; ok {
	//		if v[`method`] != `` {
	//			collection[v[`alias`]] = utils.ReflectValueInterface(utils.ReflectCallMethod(source, v[`method`])[0])
	//		} else {
	//			collection[v[`alias`]] = utils.ReflectFieldValueInterface(source, field)
	//		}
	//	}
	//}
	//
	//return collection
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
	//fmt.Println(fields)
	resourcesFields[reflectType] = fields

	return fields
}

func NewResource(source interface{},transformer interface{}) *Resource {
	return &Resource{
		source: source,
		transformer:transformer,
	}
}
//func NewResource(source interface{}, fields ...string) *Resource {
//	return &Resource{
//		source: source,
//		fields: fields,
//	}
//}

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
