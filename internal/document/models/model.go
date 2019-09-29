package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Page struct {
	Id      uint64 `json:"id",gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content",gorm:"default null"`
	//Data Data `gorm:"default null;type:json"`
	Data1 *StringMap `gorm:"default null;type:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type StringMap struct {
	Src   map[string]string
	Valid bool
}

func NewEmptyStringMap() *StringMap {
	return &StringMap{
		Src:   make(map[string]string),
		Valid: true,
	}
}

func NewStringMap(src map[string]string) *StringMap {
	return &StringMap{
		Src:   src,
		Valid: true,
	}
}

func (ls *StringMap) Scan(value interface{}) error {
	if value == nil {
		ls.Src, ls.Valid = make(map[string]string), false
		return nil
	}
	t := make(map[string]string)
	if e := json.Unmarshal(value.([]byte), &t); e != nil {
		return e
	}
	ls.Valid = true
	ls.Src = t
	return nil
}

func (ls *StringMap) Value() (driver.Value, error) {
	if ls == nil {
		return nil, nil
	}
	if !ls.Valid {
		return nil, nil
	}

	b, e := json.Marshal(ls.Src)
	return b, e
}
//type Data struct{
//	value []int
//}
//
//func (d *Data) Scan(src interface{}) error {
//	if src == nil {
//		d.value = make([]int,0)
//		return nil
//	}
//	t := make([]int,0)
//	if e := json.Unmarshal(d.value.([]byte), &t); e != nil {
//		return e
//	}
//	ls.Valid = true
//	ls.Src = t
//	return nil
//}
//
//func (d *Data) Value() (driver.Value, error) {
//	panic("implement me")
//}

//func (ls *StringMap) Scan(value interface{}) error {
//	if value == nil {
//		ls.Src, ls.Valid = make(map[string]string), false
//		return nil
//	}
//	t := make(map[string]string)
//	if e := json.Unmarshal(value.([]byte), &t); e != nil {
//		return e
//	}
//	ls.Valid = true
//	ls.Src = t
//	return nil
//}
//
//func (ls *StringMap) Value() (driver.Value, error) {
//	if ls == nil {
//		return nil, nil
//	}
//	if !ls.Valid {
//		return nil, nil
//	}
//
//	b, e := json.Marshal(ls.Src)
//	return b, e
//}

func (m *Page) TableName() string {
	return `pages`
}

func NewPage() *Page {
	return &Page{}
}
