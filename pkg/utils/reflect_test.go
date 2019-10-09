package utils

import (
	"fmt"
	"testing"
)

type testResourceExample struct {
	ID      uint   `json:"id",a:"123"`
	Title   string `json:"title"`
	Content string
	internal string `json:"internal"`
	E *testEmbedded
}

type testEmbedded struct {
	E1 string `json:"e1"`
	e2 string `json:"e2"`
	ES *testEmbeddedSecond
}

type testEmbeddedSecond struct {
	ES1 string `json:"es1"`
	es2 string `json:"es2"`
}

func TestReflectFields(t *testing.T) {
	source := &testResourceExample{
		1, `title`, `content`,`internal`,
		&testEmbedded{
			E1:"abc",
			e2:"def",
			ES:&testEmbeddedSecond{
				ES1:"ss",
				es2:"es2",
			},
		},
	}

	//source2 := resource.TestResourceExample2{
	//	ID:1,Title: `title`, Content:`content`,
	//}

	//b,_ := json.Marshal(source2)
	//fmt.Println(string(b))
	//source2 := map[string]interface{}{"a":"a","b":1,"c":true}
	//fmt.Printf("%#v",reflect.TypeOf(source).Elem().Field(0))
	//fmt.Println(ReflectStructFields(source,true))
	//fmt.Println(ReflectStructFields(source,true))
	for name,_ := range ReflectStructFields(source,true) {
		//println(name,)
		//fmt.Printf("%s==%s==%v\n",name,v[`tag`].(reflect.StructTag).Get(`json`),ReflectValueInterface(v[`value`].(reflect.Value)))
		fmt.Printf("%s\n",name)
	}
}