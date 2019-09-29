package transfromers

import (
	"github.com/blog/internal/document/models"
	"github.com/blog/internal/pkg/transform"
)

type Page struct {
	*transform.BaseTransform
}

func NewPage(page *models.Page) *Page {
	return &Page{
		BaseTransform:transform.NewTransform(page),
	}
}