package resources

type Transformer interface {
	Field() []string

	//Transform(item map[string]interface{}) map[string]interface{}
}
