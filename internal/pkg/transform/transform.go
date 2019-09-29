package transform

import (
	"github.com/blog/pkg/utils"
)

type fields []string

type BaseTransformer interface {
	Transform()
}

type BaseTransform struct {
	onlyFields   fields
	exceptFields fields
	resource     interface{}
	source       interface{}
}

func (t *BaseTransform) Resource() interface{} {
	return t.resource
}

func (t *BaseTransform) Only(field ...string) *BaseTransform {
	t.onlyFields = utils.SliceStringUnique(append(t.onlyFields, field...))
	return t
}

func (t *BaseTransform) Except(field ...string) *BaseTransform {
	t.exceptFields = utils.SliceStringUnique(append(t.exceptFields, field...))
	return t
}

// Get effective fields
func (t *BaseTransform) effectiveFields() fields {
	if len(t.onlyFields) > 0 {
		return t.onlyFields
	}

	fields := utils.ReflectFieldsName(t.resource)
	if len(t.exceptFields) > 0 {
		return utils.SliceStringExcept(fields, t.exceptFields)
	}

	return fields
}

// Set transfer source
// fix When child inherits, current is not a child
func (t *BaseTransform) SetSource(source interface{}) {
	t.source = source
}

// Core conversion method
func (t *BaseTransform) Transform() map[string]interface{} {
	//@todo 这里有问题，子级继承时,t并不是子级
	//methods := utils.ReflectMethodsName(t)

	methods := utils.ReflectMethodsName(t.source)

	collection := make(map[string]interface{})
	for _, field := range t.effectiveFields() {
		method := utils.StringUcWords([]string{"Get", field, `Field`})
		fieldKey := utils.StringSnakeCase(field)
		if utils.SliceStringIn(methods, method) {
			collection[fieldKey] = utils.ReflectCallMethod(t.source, method)[0].Interface()
			//collection[fieldKey] = utils.ReflectCallMethod(t, method)[0].Interface()
		} else {
			collection[fieldKey] = utils.ReflectFieldValue(t.resource, field).Interface()
		}
	}

	return map[string]interface{}{"data":collection}
}

// Create a new transform
func NewTransform(resource interface{}) *BaseTransform {
	return &BaseTransform{
		resource:     resource,
		onlyFields:   make(fields, 0),
		exceptFields: make(fields, 0),
	}
}
