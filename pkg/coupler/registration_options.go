package coupler

import (
	"reflect"

	"github.com/d34dl0ck/coupler/internal/core"
)

func WithName(n string) RegistrationOption {
	return func(r *Registration) error {
		r.Key = core.NewRawDependencyKey(n)
		return nil
	}
}

func AsImplementationOf[T interface{}]() RegistrationOption {
	return func(r *Registration) error {
		t := reflect.TypeOf((*T)(nil))
		elemType := t.Elem()
		r.Key = core.NewTypeDependencyKey(elemType)
		return nil
	}
}
