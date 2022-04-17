package strategies

import (
	"reflect"

	"github.com/pkg/errors"

	"github.com/d34dl0ck/coupler/internal/container"
)

type FuncStrategy struct {
	function interface{}
}

func NewFuncStrategy(function interface{}) container.ResolvingStrategy {
	return FuncStrategy{
		function: function,
	}
}

func (s FuncStrategy) Resolve(r container.Registrations) (interface{}, error) {
	fType := reflect.TypeOf(s.function)
	args := []reflect.Value{}
	for i := 0; i < fType.NumIn(); i++ {
		arg := fType.In(i)
		key := container.NewTypeResolvingKey(arg)
		strategy, hasValue := r[key]

		if hasValue {
			instance, err := strategy.Resolve(r)

			if err != nil {
				return nil, errors.Wrapf(container.ErrResolveFailed, "failed to resolve key %s", key)
			}

			args = append(args, reflect.ValueOf(instance))
		} else {
			return nil, errors.Wrapf(container.ErrDependencyNotRegistered, "dependency with key %s was not registered", key)
		}
	}

	result := reflect.ValueOf(s.function).Call(args)
	return result[0].Interface(), nil
}

func (s FuncStrategy) ProvideDefaultKey() container.ResolvingKey {
	fType := reflect.TypeOf(s.function)
	resultType := fType.Out(0)

	return container.NewTypeResolvingKey(resultType)
}
