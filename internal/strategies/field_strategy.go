package strategies

import (
	"reflect"

	"github.com/pkg/errors"

	"github.com/d34dl0ck/coupler/internal/container"
)

type FieldStrategy struct {
	t reflect.Type
}

func NewFieldStrategy(t reflect.Type) container.ResolvingStrategy {
	return FieldStrategy{
		t: t,
	}
}

func (s FieldStrategy) Resolve(r container.Registrations) (interface{}, error) {
	result := reflect.New(s.t)
	for i := 0; i < result.Elem().NumField(); i++ {
		field := result.Elem().Field(i).Type()
		key := container.NewTypeResolvingKey(field)
		strategy, hasValue := r[key]

		if hasValue {
			instance, err := strategy.Resolve(r)

			if err != nil {
				return nil, errors.Wrapf(container.ErrResolveFailed, "failed to resolve key %s", key)
			}

			result.Elem().Field(i).Set(reflect.ValueOf(instance))
		} else {
			return nil, errors.Wrapf(container.ErrDependencyNotRegistered, "dependency with key %s was not registered", key)
		}
	}

	return result.Elem().Interface(), nil
}

func (s FieldStrategy) ProvideDefaultKey() container.ResolvingKey {
	return container.NewTypeResolvingKey(s.t)
}
