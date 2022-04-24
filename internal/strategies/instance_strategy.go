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
		return nil, ErrNilInput
	}

	return InstanceStrategy{
		instance: instance,
	}, nil
}

func (s InstanceStrategy) Resolve(_ core.Resolver) (interface{}, error) {
	return s.instance, nil
}

func (s InstanceStrategy) ProvideDefaultKey() core.DependencyKey {
	return core.NewTypeDependencyKey(reflect.TypeOf(s.instance))
}
