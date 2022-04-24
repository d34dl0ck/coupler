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

	expected := strategies.TestStruct{
		SomeString: strategies.ExpectedTestString,
	}
	f := func(s string) strategies.TestStruct {
		return strategies.TestStruct{
			SomeString: s,
		}
	}
	strategy, err := strategies.NewFuncStrategy(f)
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeDependencyKey(reflect.TypeOf(strategies.ExpectedTestString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(strategies.ExpectedTestString, nil)

	actual, err := strategy.Resolve(resolverMock)
	require.NoError(t, err, "error was not expected")
	require.Equal(t, expected, actual, "resolved instance mismatch")
}

func TestReturnFuncArgDependencyResolveError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := func(s string) strategies.TestStruct {
		return strategies.TestStruct{
			SomeString: s,
		}
	}
	strategy, err := strategies.NewFuncStrategy(f)
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeDependencyKey(reflect.TypeOf(strategies.ExpectedTestString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(nil, strategies.ErrExpected)

	_, err = strategy.Resolve(resolverMock)

	require.ErrorIs(t, err, strategies.ErrExpected, "error was expected")
}

func TestFuncStrategyDefaultKey(t *testing.T) {
	t.Parallel()

	input := strategies.TestStruct{
		SomeString: strategies.ExpectedTestString,
	}
	f := func(s string) strategies.TestStruct {
		return strategies.TestStruct{
			SomeString: s,
		}
	}
	expected := core.NewTypeDependencyKey(reflect.TypeOf(input))
	strategy, err := strategies.NewFuncStrategy(f)
	require.NoError(t, err, "err was not expected")

	actual := strategy.ProvideDefaultKey()

	require.Equal(t, expected, actual, "key mismatch")
}

func TestErrorNilFunc(t *testing.T) {
	t.Parallel()

	_, err := strategies.NewFuncStrategy(nil)
	require.ErrorIs(t, err, strategies.ErrNilInput, "error mismatch")
}

func TestErrorFuncNilDependency(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := func(s string) strategies.TestStruct {
		return strategies.TestStruct{
			SomeString: s,
		}
	}
	strategy, err := strategies.NewFuncStrategy(f)
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeDependencyKey(reflect.TypeOf(strategies.ExpectedTestString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(nil, nil)

	_, err = strategy.Resolve(resolverMock)
	require.ErrorIs(t, err, strategies.ErrNilDependency, "error mismatch")
}
