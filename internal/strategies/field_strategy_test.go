package strategies_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/d34dl0ck/coupler/internal/core/testdata"
	"github.com/d34dl0ck/coupler/internal/strategies"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	expectedString = "some_expected"
)

var (
	ErrExpected = errors.New("some_expected_error")
)

type TestStruct struct {
	SomeString string
}

func TestResolveByFields(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := TestStruct{
		SomeString: expectedString,
	}
	strategy, err := strategies.NewFieldStrategy(reflect.TypeOf(expected))
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeResolvingKey(reflect.TypeOf(expectedString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(expectedString, nil)

	actual, err := strategy.Resolve(resolverMock)
	require.NoError(t, err, "error was not expected")
	require.Equal(t, expected, actual, "resolved instance mismatch")
}

func TestReturnFieldDependencyResolveError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := TestStruct{
		SomeString: expectedString,
	}
	strategy, err := strategies.NewFieldStrategy(reflect.TypeOf(expected))
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeResolvingKey(reflect.TypeOf(expectedString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(nil, ErrExpected)

	_, err = strategy.Resolve(resolverMock)

	require.ErrorIs(t, err, ErrExpected, "error was expected")
}

func TestFieldStrategyDefaultKey(t *testing.T) {
	t.Parallel()

	input := TestStruct{
		SomeString: expectedString,
	}
	expected := core.NewTypeResolvingKey(reflect.TypeOf(input))
	strategy, err := strategies.NewFieldStrategy(reflect.TypeOf(input))
	require.NoError(t, err, "err was not expected")

	actual := strategy.ProvideDefaultKey()

	require.Equal(t, expected, actual, "key mismatch")
}

func TestErrorNilType(t *testing.T) {
	t.Parallel()

	_, err := strategies.NewFieldStrategy(nil)
	require.ErrorIs(t, err, strategies.ErrNilType, "error mismatch")
}
