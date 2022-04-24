package coupler

import (
	"reflect"

	"github.com/d34dl0ck/coupler/internal/container"
	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/pkg/errors"
)

var (
	c                           core.Container
	ErrRegistration             = errors.New("cannot register dependency")
	ErrDependenciesInconsistent = errors.New("some of registered dependencies cannot be resolved cause of missing dependency")
)

func init() {
	c = container.NewContainer()
}

func Register(resolveOption ResolveOption, opts ...RegistrationOption) error {
	r := &Registration{}

	err := resolveOption(r)

	if err != nil {
		return err
	}

	for _, opt := range opts {
		if opt != nil {
			err := opt(r)
			if err != nil {
				return err
			}
		}
	}

	key := r.Key

	if r.Strategy == nil {
		return errors.Wrapf(core.ErrStrategyIsEmpty, "no strategy was set for key %s", key)
	}

	if key == nil || key.IsEmpty() {
		r.Key = r.Strategy.ProvideDefaultKey()
	}

	c.Register(r.Key, r.Strategy)
	return nil
}

func Resolve[T interface{}]() (T, error) {
	var def T

	desiredType := reflect.TypeOf(def)

	if desiredType == nil {
		desiredType = reflect.TypeOf((*T)(nil)).Elem()
	}

	raw, err := c.Resolve(core.NewTypeDependencyKey(desiredType))

	if err != nil {
		return def, err
	}

	return raw.(T), nil
}

func Check() error {
	err := c.Check()

	if err != nil {
		return ErrDependenciesInconsistent
	}

	return nil
}

func ResolveNamed[T interface{}](name string) (T, error) {
	raw, err := c.Resolve(core.NewRawDependencyKey(name))
	return raw.(T), err
}
