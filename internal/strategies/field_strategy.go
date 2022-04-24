package strategies

import (
	"reflect"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/pkg/errors"
)

type FieldStrategy struct {
	t reflect.Type
}

func NewFieldStrategy(t reflect.Type) (core.ResolvingStrategy, error) {
	if t == nil {
		return nil, ErrNilInput
	}

	empty := reflect.New(t)
	for i := 0; i < empty.Elem().NumField(); i++ {
		if !empty.Elem().Field(i).CanSet() {
			return nil, ErrUnexportedFieldDetected
		}
	}

	return FieldStrategy{
		t: t,
	}, nil
}

func (s FieldStrategy) Resolve(r core.Resolver) (interface{}, error) {
	result := reflect.New(s.t)
	for i := 0; i < result.Elem().NumField(); i++ {
		field := result.Elem().Field(i)

		if !field.CanSet() {
			return nil, ErrUnexportedFieldDetected
		}

		fieldType := field.Type()
		key := core.NewTypeDependencyKey(fieldType)

		instance, err := r.Resolve(key)

		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve key %s", key)
		}

		if instance == nil {
			return nil, ErrNilDependency
		}

		field.Set(reflect.ValueOf(instance))
	}

	return result.Elem().Interface(), nil
}

func (s FieldStrategy) ProvideDefaultKey() core.DependencyKey {
	return core.NewTypeDependencyKey(s.t)
}
