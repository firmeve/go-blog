package resource

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ResourceItemExample struct {
	ID       uint   `json:"id",a:"123"`
	Title    string `json:"title"`
	Content  string
	ServerTo    string `resource:"s_t"`
	internal string `json:"internal"`
	E *ItemEmbedded `resource:"es"`
}

type ItemEmbedded struct {
	E1 string `json:"e1"`
	e2 string `json:"e2"`
	ES *itemEmbeddedSecond `resource:"es_internal"`
}

type itemEmbeddedSecond struct {
	ES1 string `json:"es1"`
	es2 string `json:"es2"`
}

func TestResource_ReflectRelationFields(t *testing.T) {
	//resource := &Resource{
	source := &ResourceItemExample{
		1, `title`, `content`, "serverTo",`internal`,
		&ItemEmbedded{
			E1:"abc",
			e2:"def",
			ES:&itemEmbeddedSecond{
				ES1:"ss",
				es2:"es2",
			},
		},
	}
	//}
	item := NewItem(source)
	item.Fields(`ID`, `Title`,`ServerTo`,`E`)
	z , _ := json.Marshal(item.Resolve())
	fmt.Println(string(z))
	//
	//zs := map[string]string{`a`:`a`,`b`:`b`}
	//item2 := NewItem(zs)
	//item2.Fields(`a`, `b`)
	//zsd , _ := json.Marshal(item.Resolve())
	//fmt.Println(string(zsd))

	//bytes,err := json.Marshal(source)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(bytes))
	//resource.ReflectRelationFields(resource.source)
}
