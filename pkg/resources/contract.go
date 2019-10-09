package resources


type Resolver interface {
	Resolve() interface{}
}

type