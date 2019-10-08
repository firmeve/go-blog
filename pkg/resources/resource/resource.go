package resource

import (
	"fmt"
	"github.com/blog/pkg/utils"
	_ "github.com/blog/pkg/utils"
	"reflect"
)

var (
	resourcesFields map[reflect.Type]string
)

type resourceFields []string

type Resource struct {
	source       interface{}
	onlyFields   resourceFields
	exceptFields resourceFields
}

func (r *Resource) Only(fields []string) {
	r.onlyFields = fields
}

func (r *Resource) Except(fields []string) {
	r.exceptFields = fields
}

func (r *Resource) Source() interface{} {
	return r.source
}


// Core conversion method
//func (r *Resource) Transform(source interface{}) map[string]interface{} {
//
//	methods := utils.ReflectMethodsName(r.source)
//
//	collection := make(map[string]interface{})
//	for _, field := range r.effectiveFields() {
//		method := utils.StringUcWords([]string{"Get", field, `Field`})
//		fieldKey := utils.StringSnakeCase(field)
//		if utils.SliceStringIn(methods, method) {
//			collection[fieldKey] = utils.ReflectCallMethod(t.source, method)[0].Interface()
//			//collection[fieldKey] = utils.ReflectCallMethod(t, method)[0].Interface()
//		} else {
//			collection[fieldKey] = utils.ReflectFieldValue(t.resource, field).Interface()
//		}
//	}
//
//	return map[string]interface{}{"data":collection}
//}

func (r *Resource) ReflectRelationFields(source interface{}) {
	//fmt.Println(utils.ReflectFieldsName(source))
	reflectType := utils.ReflectTypeIndirect(reflect.TypeOf(source))
	//kind := reflectType.Kind()
	//if  SliceIntIn([]int64{int64(reflect.Array),int64(reflect.Ptr),int64(reflect.Chan),int64(reflect.Map),int64(reflect.Slice)},int64(kind)){
	//	reflectType = reflectType.Elem()
	//}
	//json.Marshal(Apple{"green", 10})
	nums := reflectType.NumField()
	fields := make([]string, 0)
	for i := 0; i < nums; i++ {
		// check struct 字段是否是私有
		fields = append(fields, reflectType.Field(i).Name)
	}
	fmt.Println(fields)
}