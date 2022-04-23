package strategies

import (
	"reflect"

	"github.com/d34dl0ck/coupler/internal/core"
)

type InstanceStrategy struct {
	instance interface{}
}

func NewInstanceStrategy(instance interface{}) (core.ResolvingStrategy, error) {
	if instance == nil {
		return nil, ErrNilType
	}

	return InstanceStrategy{
		instance: instance,
	}, nil
}

func (s InstanceStrategy) Resolve(_ core.Resolver) (interface{}, error) {
	return s.instance, nil
}

func (s InstanceStrategy) ProvideDefaultKey() core.ResolvingKey {
	return core.NewTypeResolvingKey(reflect.TypeOf(s.instance))
}
