package core

import "reflect"

type DependencyKey interface {
	Value() string
	IsEmpty() bool
}

type stringKey struct {
	value string
}

func (k stringKey) Value() string {
	return k.value
}

func (k stringKey) IsEmpty() bool {
	return k.value == ""
}

func NewTypeDependencyKey(t reflect.Type) stringKey {
	return stringKey{
		value: t.Name(),
	}
}

func NewRawDependencyKey(s string) stringKey {
	return stringKey{
		value: s,
	}
}
