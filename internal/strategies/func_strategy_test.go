package strategies_test

import (
	"reflect"
	"testing"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/d34dl0ck/coupler/internal/core/testdata"
	"github.com/d34dl0ck/coupler/internal/strategies"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestResolveByFunc(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := TestStruct{
		SomeString: expectedString,
	}
	f := func(s string) TestStruct {
		return TestStruct{
			SomeString: s,
		}
	}
	strategy, err := strategies.NewFuncStrategy(f)
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeResolvingKey(reflect.TypeOf(expectedString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(expectedString, nil)

	actual, err := strategy.Resolve(resolverMock)
	require.NoError(t, err, "error was not expected")
	require.Equal(t, expected, actual, "resolved instance mismatch")
}

func TestReturnFuncArgDependencyResolveError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := func(s string) TestStruct {
		return TestStruct{
			SomeString: s,
		}
	}
	strategy, err := strategies.NewFuncStrategy(f)
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeResolvingKey(reflect.TypeOf(expectedString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(nil, ErrExpected)

	_, err = strategy.Resolve(resolverMock)

	require.ErrorIs(t, err, ErrExpected, "error was expected")
}

func TestFuncStrategyDefaultKey(t *testing.T) {
	t.Parallel()

	input := TestStruct{
		SomeString: expectedString,
	}
	f := func(s string) TestStruct {
		return TestStruct{
			SomeString: s,
		}
	}
	expected := core.NewTypeResolvingKey(reflect.TypeOf(input))
	strategy, err := strategies.NewFuncStrategy(f)
	require.NoError(t, err, "err was not expected")

	actual := strategy.ProvideDefaultKey()

	require.Equal(t, expected, actual, "key mismatch")
}

func TestErrorNilFunc(t *testing.T) {
	t.Parallel()

	_, err := strategies.NewFuncStrategy(nil)
	require.ErrorIs(t, err, strategies.ErrNilType, "error mismatch")
}
