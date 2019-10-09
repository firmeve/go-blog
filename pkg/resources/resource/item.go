package resource

type Item struct {
	*Resource
}

func (i *Item) Resolve() interface{} {
	return i.Transform(i.Source())
}

func NewItem(source interface{}) *Item {
	return &Item{
		Resource: NewResource(source),
	}
}
