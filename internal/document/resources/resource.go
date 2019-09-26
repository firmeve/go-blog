package resources

import (
	"fmt"
	"github.com/blog/internal/document/models"
	"github.com/blog/pkg/utils"
)

type Fields []string

type PageResource struct {
	//Id           uint64 `json:"id"`
	//Title        string `json:"title"`
	//Content      string `json:"content"`
	//CreatedAt    string `json:"created_at"`
	//UpdatedAt    string `json:"updated_at"`
	onlyFields   Fields
	exceptFields Fields
	resource     interface{}
}

func (r *PageResource) Only(field ...string) *PageResource {
	r.onlyFields = append(r.onlyFields, field...)
	r.onlyFields = utils.SliceStringUnqiue(r.onlyFields)
	return r
}

func (r *PageResource) effectiveFields() Fields {
	if len(r.onlyFields) > 0 {
		return r.onlyFields
	}

	return r.exceptFields
}

func (r *PageResource) Resource() map[string]interface{} {
	//fmt.Println(reflect.ValueOf(r).Elem().Kind().String())

	methods := utils.ReflectMethodsName(r)
	result := make(map[string]interface{})
	for _, field := range r.effectiveFields() {
		if utils.InSlice(methods.([]string),field) {
			
		}
		if method, ok := methods[field]; ok {

		}
	}
	fmt.Println()
	z := map[string]interface{}{
		"x": 1,
		"z": "z",
	}
	return z
}

func (r *PageResource) GetTitleField() string {
	return r.resource.(*models.Page).Title
}

func NewPageResource(page *models.Page) *PageResource {
	return &PageResource{
		resource:     page,
		onlyFields:   make(Fields, 0),
		exceptFields: make(Fields, 0),
	}
}

//func PageResource(page *models.Page) {
//
//}
