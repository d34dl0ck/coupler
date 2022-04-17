package strategies

import (
	"reflect"

	"github.com/d34dl0ck/coupler/internal/container"
)

type InstanceStrategy struct {
	instance interface{}
}

func NewInstanceStrategy(instance interface{}) container.ResolvingStrategy {
	return InstanceStrategy{
		instance: instance,
	}
}

func (s InstanceStrategy) Resolve(r container.Registrations) (interface{}, error) {
	return s.instance, nil
}

func (s InstanceStrategy) ProvideDefaultKey() container.ResolvingKey {
	return container.NewTypeResolvingKey(reflect.TypeOf(s.instance))
}
