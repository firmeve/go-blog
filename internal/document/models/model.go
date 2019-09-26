package models

import (
	"github.com/jinzhu/gorm"
)

type Page struct {
	gorm.Model
	ID      uint64 `json:"id",gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content",gorm:"default null"`
}

func (m *Page) TableName() string {
	return `pages`
}

func NewPage() *Page {
	return &Page{}
}
