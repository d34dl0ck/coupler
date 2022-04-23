package strategies

import (
	"reflect"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/pkg/errors"
)

var (
	ErrNilType = errors.New("input is nil")
)

type FieldStrategy struct {
	t reflect.Type
}

func NewFieldStrategy(t reflect.Type) (core.ResolvingStrategy, error) {
	if t == nil {
		return nil, ErrNilType
	}

	return FieldStrategy{
		t: t,
	}, nil
}

func (s FieldStrategy) Resolve(r core.Resolver) (interface{}, error) {
	result := reflect.New(s.t)
	for i := 0; i < result.Elem().NumField(); i++ {
		field := result.Elem().Field(i).Type()
		key := core.NewTypeResolvingKey(field)

		instance, err := r.Resolve(key)

		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve key %s", key)
		}

		result.Elem().Field(i).Set(reflect.ValueOf(instance))
	}

	return result.Elem().Interface(), nil
}

func (s FieldStrategy) ProvideDefaultKey() core.ResolvingKey {
	return core.NewTypeResolvingKey(s.t)
}
