package models

import "time"

type Page struct {
	Id      uint64 `json:"id",gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content",gorm:"default null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (m *Page) TableName() string {
	return `pages`
}

func NewPage() *Page {
	return &Page{}
}
