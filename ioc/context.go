package ioc

import "go/types"

type Feature interface{}

type Context interface {
	Only(theType types.Type) Feature

	Find(types types.Type) []Feature

	OnlyOf(types types.Type, name string) Feature
}

type Graph interface {
	Provide(object ...*Object) error

	Populate() error
}

type Injected interface {
	Name() string

	Done() error
}

type InjectGroup interface {
	Injected

	Components() []*Object
}

type Object struct {
	Value Feature
}
