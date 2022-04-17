package coupler

import (
	"reflect"

	"github.com/d34dl0ck/coupler/internal/container"
	"github.com/d34dl0ck/coupler/internal/strategies"
	"github.com/pkg/errors"
)

var (
	c *container.Container
)

type Strategy container.ResolvingStrategy

type Key container.ResolvingKey

type RegistrationOption func(r *Registration)

type ResolveOption func(r *Registration)

type Registration struct {
	Key      Key
	Strategy Strategy
}

func init() {
	c = container.NewContainer()
}

func Register(resolveOption ResolveOption, opts ...RegistrationOption) error {
	r := &Registration{}

	resolveOption(r)

	for _, opt := range opts {
		if opt != nil {
			opt(r)
		}
	}

	key := r.Key

	if r.Strategy == nil {
		return errors.Wrapf(container.ErrStrategyIsEmpty, "no strategy was set for key %s", key)
	}

	if key == nil || key.IsEmpty() {
		r.Key = r.Strategy.ProvideDefaultKey()
	}

	c.Register(r.Key, r.Strategy)
	return nil
}

func WithName(n string) RegistrationOption {
	return func(r *Registration) {
		r.Key = container.NewRawResolvingKey(n)
	}
}

func ByFunc(f interface{}) ResolveOption {
	return func(r *Registration) {
		r.Strategy = strategies.NewFuncStrategy(f)
	}
}

func ByInstance(i interface{}) ResolveOption {
	return func(r *Registration) {
		r.Strategy = strategies.NewInstanceStrategy(i)
	}
}

func ByType[T interface{}]() ResolveOption {
	var def T
	return func(r *Registration) {
		r.Strategy = strategies.NewFieldStrategy(reflect.TypeOf(def))
	}
}

func Resolve[T interface{}]() (T, error) {
	var def T
	desiredType := reflect.TypeOf(def)

	raw, err := c.Resolve(container.NewTypeResolvingKey(desiredType))

	return raw.(T), err
}

func ResolveNamed[T interface{}](name string) (T, error) {
	raw, err := c.Resolve(container.NewRawResolvingKey(name))
	return raw.(T), err
}
