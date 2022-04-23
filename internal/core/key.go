package core

import "reflect"

type ResolvingKey interface {
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

func NewTypeResolvingKey(t reflect.Type) stringKey {
	return stringKey{
		value: t.Name(),
	}
}

func NewRawResolvingKey(s string) stringKey {
	return stringKey{
		value: s,
	}
}
