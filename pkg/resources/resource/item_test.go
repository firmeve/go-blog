package resource

import (
	"fmt"
	"testing"
)

type ResourceItemExample struct {
	ID      uint   `json:"id",a:"123"`
	Title   string `json:"title"`
	Content string
	internal string `json:"internal"`
	//E *ItemEmbedded
}

type ItemEmbedded struct {
	E1 string `json:"e1"`
	e2 string `json:"e2"`
	ES *itemEmbeddedSecond
}

type itemEmbeddedSecond struct {
	ES1 string `json:"es1"`
	es2 string `json:"es2"`
}

func TestResource_ReflectRelationFields(t *testing.T) {
	//resource := &Resource{
		source := &ResourceItemExample{
			1, `title`, `content`,`internal`,
			//&ItemEmbedded{
			//	E1:"abc",
			//	e2:"def",
			//	ES:&itemEmbeddedSecond{
			//		ES1:"ss",
			//		es2:"es2",
			//	},
			//},
		}
	//}

	fmt.Printf("%#v",NewItem(source).Fields(`ID`,`Title`).Resolve())
	//bytes,err := json.Marshal(source)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(bytes))
	//resource.ReflectRelationFields(resource.source)
}