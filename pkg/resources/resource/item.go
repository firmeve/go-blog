package resource

type Item struct {
	*Resource
}

func (i *Item) SetFields(fields ...string) *Item {

	return i
}

func (i *Item) Resolve() interface{} {
	return i.Transform(i.Source())
}

func NewItem(source interface{}, transformer interface{}) *Item {
	return &Item{
		Resource: NewResource(source, transformer),
	}
}
