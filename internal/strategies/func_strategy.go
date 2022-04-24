package strategies

import (
	"reflect"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/pkg/errors"
)

type FuncStrategy struct {
	function interface{}
}

func NewFuncStrategy(function interface{}) (core.ResolvingStrategy, error) {
	if function == nil {
		return nil, ErrNilInput
	}

	return FuncStrategy{
		function: function,
	}, nil
}

func (s FuncStrategy) Resolve(r core.Resolver) (interface{}, error) {
	fType := reflect.TypeOf(s.function)
	args := []reflect.Value{}
	for i := 0; i < fType.NumIn(); i++ {
		arg := fType.In(i)
		key := core.NewTypeDependencyKey(arg)
		instance, err := r.Resolve(key)

		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve dependency %s", key)
		}

		if instance == nil {
			return nil, ErrNilDependency
		}

		args = append(args, reflect.ValueOf(instance))
	}

	result := reflect.ValueOf(s.function).Call(args)
	return result[0].Interface(), nil
}

func (s FuncStrategy) ProvideDefaultKey() core.DependencyKey {
	fType := reflect.TypeOf(s.function)
	resultType := fType.Out(0)

	return core.NewTypeDependencyKey(resultType)
}
