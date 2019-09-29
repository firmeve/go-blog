package transform

import (
	"github.com/blog/internal/document/models"
	"github.com/blog/pkg/utils"
)

type fields []string

type BaseTransformer interface {
	BaseTransform()
}

type BaseTransform struct {
	onlyFields   fields
	exceptFields fields
	resource     interface{}
}

func (t *BaseTransform) Only(field ...string) *BaseTransform {
	t.onlyFields = utils.SliceStringUnique(append(t.onlyFields, field...))
	return t
}

func (t *BaseTransform) Except(field ...string) *BaseTransform {
	t.exceptFields = utils.SliceStringUnique(append(t.exceptFields, field...))
	return t
}

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

func (t *BaseTransform) Transform() map[string]interface{} {
	methods := utils.ReflectMethodsName(t)
	collection := make(map[string]interface{})
	for _, field := range t.effectiveFields() {
		method := utils.StringUcWords([]string{"Get", field, `Field`})
		fieldKey := utils.StringSnakeCase(field)
		if utils.SliceStringIn(methods, method) {
			collection[fieldKey] = utils.ReflectCallMethod(t, method)[0].Interface()
		} else {
			collection[fieldKey] = utils.ReflectFieldValue(t.resource, field).Interface()
		}
	}

	return collection
}

func (t *BaseTransform) GetTitleField() string {
	return t.resource.(*models.Page).Title
}

func NewTransform(resource interface{}) *BaseTransform {
	return &BaseTransform{
		resource:     resource,
		onlyFields:   make(fields, 0),
		exceptFields: make(fields, 0),
	}
}
