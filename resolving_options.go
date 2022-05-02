package coupler

import (
	"reflect"

	"github.com/d34dl0ck/coupler/internal/strategies"
	"github.com/pkg/errors"
)

func ByFunc(f interface{}) ResolveOption {
	return func(r *Registration) error {
		s, err := strategies.NewFuncStrategy(f)
		if err != nil {
			return errors.Wrap(ErrRegistration, err.Error())
		}
		r.Strategy = s
		return nil
	}
}

func ByInstance(i interface{}) ResolveOption {
	return func(r *Registration) error {
		s, err := strategies.NewInstanceStrategy(i)
		if err != nil {
			return errors.Wrap(ErrRegistration, err.Error())
		}
		r.Strategy = s
		return nil
	}
}

func ByType[T interface{}]() ResolveOption {
	var def T
	return func(r *Registration) error {
		s, err := strategies.NewFieldStrategy(reflect.TypeOf(def))
		if err != nil {
			return errors.Wrap(ErrRegistration, err.Error())
		}
		r.Strategy = s
		return nil
	}
}
