package resource

//type ResourceExample struct {
//	ID      uint   `json:"id",a:"123"`
//	Title   string `json:"title"`
//	Content string
//	internal string `json:"internal"`
//	E *Embedded
//}
//
//type Embedded struct {
//	E1 string `json:"e1"`
//	e2 string `json:"e2"`
//	ES *embeddedSecond
//}
//
//type embeddedSecond struct {
//	ES1 string `json:"es1"`
//	es2 string `json:"es2"`
//}
//
//func TestResource_ReflectRelationFields(t *testing.T) {
//	//resource := &Resource{
//		source := &ResourceExample{
//			1, `title`, `content`,`internal`,
//			&Embedded{
//				E1:"abc",
//				e2:"def",
//				ES:&embeddedSecond{
//					ES1:"ss",
//					es2:"es2",
//				},
//			},
//		}
//	//}
//
//	bytes,err := json.Marshal(source)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(bytes))
//	//resource.ReflectRelationFields(resource.source)
//}
