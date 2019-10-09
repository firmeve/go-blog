package transfromers

import (
	"github.com/blog/internal/document/models"
	"github.com/blog/internal/pkg/transform"
)

type Page struct {
	*transform.BaseTransform
}

func (p *Page) GetData1Field() interface{} {
	return p.Resource().(*models.Page).Data1.Src
}

func (t *Page) GetTitleField() string {
	return t.Resource().(*models.Page).Title
}

func NewPage(page *models.Page) *Page {
	transform := &Page{
		BaseTransform:transform.NewTransform(page),
	}
	transform.BaseTransform.SetSource(transform)
	return transform
}